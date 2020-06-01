package scan

import (
	"errors"
	"fmt"
	"github.com/bcbchain/toolbox/bcscan/conf"
	"github.com/bcbchain/toolbox/bcscan/db"
	"github.com/bcbchain/toolbox/bcscan/log"
	"github.com/bcbchain/toolbox/bcscan/query"
	"math/rand"
	"sort"
	"time"
)

type Controller struct {
	Config *conf.BcScanConfig

	URLs           []string
	FromHeight     int64
	ToHeight       int64
	FilterTypeList []string

	CurrentScanHeight int64

	ScanResultChan chan *Result
	ScanResultList []*Result

	ScannerMap map[ScannerID]*Scanner
}

func NewController(c *conf.BcScanConfig, from, to int64, filterList []string) *Controller {
	return &Controller{
		Config:         c,
		URLs:           c.URLs,
		FromHeight:     from,
		ToHeight:       to,
		FilterTypeList: filterList,
		ScanResultChan: make(chan *Result),
		ScannerMap:     map[ScannerID]*Scanner{},
	}
}

func (c *Controller) Start() error {
	log.GetLogger().Info("Scan controller start", "controller", *c)

	if err := c.checkHeight(); err != nil {
		return err
	}

	if db.GetFirstHeight() == 0 && c.FromHeight != 0 {
		log.GetLogger().Info("Scan controller", "SetFirstHeight", c.FromHeight)
		db.SetFirstHeight(c.FromHeight)
	}

	if c.FromHeight == 0 && db.GetLastHeight() == 0 {
		// todo query genesis info and save to db.
		c.FromHeight = 1
	}

	if c.FromHeight == 0 {
		c.CurrentScanHeight = db.GetLastHeight()
		c.FromHeight = c.CurrentScanHeight + 1
	}

	log.GetLogger().Info("Scan controller runScanners", "controller", *c)
	c.runScanners()

	return c.collectScanResult()

}

func (c *Controller) runScanners() {
	var (
		toHeight int64
		err      error
	)

	if c.ToHeight == 0 {
		toHeight, err = query.GetCurrentHeight(c.Config.URLs[0])
		if err != nil {
			panic(err)
		}
		log.GetLogger().Info("runScanners", "query.GetCurrentHeight", toHeight)
	} else {
		toHeight = c.ToHeight
	}

	for i := 0; i < c.Config.ConcurrentNumberOfOneNode; i++ {
		for _, url := range c.URLs {
			scanner := &Scanner{
				ID:             calcScannerID(url, i),
				URL:            url,
				BeginHeight:    0,
				EndHeight:      0,
				FilterTypeList: c.FilterTypeList,
				signalChan:     make(chan bool, 1),
				scanResultChan: c.ScanResultChan,
			}
			scanner.BeginHeight, scanner.EndHeight = c.calcScannerHeight(toHeight)
			if scanner.EndHeight == 0 {
				return
			}

			c.ScannerMap[scanner.ID] = scanner
			go scanner.run()
			scanner.signalChan <- true
			log.GetLogger().Debug("Scanner", "ID", scanner.ID, "scanner", *scanner)
		}
	}
}

func (c *Controller) calcScannerHeight(toHeight int64) (bHeight, eHeight int64) {
	if toHeight <= c.CurrentScanHeight {
		return
	}

	// 暂定每个 scanner 一次查询十个区块
	if toHeight > c.CurrentScanHeight+10 {
		bHeight = c.CurrentScanHeight + 1
		eHeight = c.CurrentScanHeight + 10
		c.CurrentScanHeight = eHeight
	} else {
		bHeight = c.CurrentScanHeight + 1
		eHeight = toHeight
		c.CurrentScanHeight = toHeight
	}

	return
}

func (c *Controller) collectScanResult() error {
	log.GetLogger().Info("collectScanResult")
	var timer *time.Timer
	timer = time.NewTimer(time.Second * 30)

	for {
		timer.Reset(time.Second * 30)
		if len(c.ScannerMap) == 0 {
			return nil
		}

		select {
		case result := <-c.ScanResultChan:
			c.processResult(result)

		case <-timer.C:
			// 如果三十秒内所有 scanner 都没有返回结果，所有节点都超时，程序退出。
			return errors.New("timeout when query block chain info by all urls")
		}
	}
}

func (c *Controller) processResult(result *Result) {
	// add scan result to cacheList and sort
	lastHeight := db.GetLastHeight()
	log.GetLogger().Debug("processResult", "result", *result, "lastHeight", lastHeight)
	if result.Err == nil {
		if result.EndHeight > lastHeight {
			c.ScanResultList = append(c.ScanResultList, result)
			sort.Slice(c.ScanResultList, func(i, j int) bool {
				return c.ScanResultList[i].BeginHeight < c.ScanResultList[j].BeginHeight
			})
		}
	}

	c.processScanner(result.ScannerID, result.Err != nil)
	c.processResultList()
}

func (c *Controller) processScanner(id ScannerID, hasErr bool) {
	log.GetLogger().Debug("processScanner", "ID", id, "hasErr", hasErr)
	if hasErr {
		c.processResultErr(id)
		return
	}
	if c.ToHeight == 0 {
		var toHeight int64
		for i := 0; i < 3; i++ {
			h, err := query.GetCurrentHeight(c.Config.URLs[0])
			if err != nil {
				continue
			}
			toHeight = h
			break
		}

		if c.CurrentScanHeight+10 < toHeight {
			c.notifyScanner(id, toHeight)
		} else {
			c.deleteScanner(id)
		}
	} else {
		if c.CurrentScanHeight < c.ToHeight {
			c.notifyScanner(id, c.ToHeight)
		} else {
			c.deleteScanner(id)
		}
	}
}

func (c *Controller) processResultErr(id ScannerID) {
	log.GetLogger().Debug("processResultErr", "ID", id)
	oldScanner := c.ScannerMap[id]
	url := id.getURL()

	newURLs := make([]string, 0)
	for _, v := range c.URLs {
		if v != url {
			newURLs = append(newURLs, v)
		}
	}

	c.URLs = newURLs

	rand.Seed(time.Now().Unix())
	randIndex := rand.Int63n(int64(len(c.URLs)))
	newURL := c.URLs[randIndex]

	maxIndex := 0
	for k := range c.ScannerMap {
		if k.getURL() == newURL {
			if k.getIndex() > maxIndex {
				maxIndex = k.getIndex()
			}
		}
	}

	newID := calcScannerID(newURL, maxIndex+1)

	scanner := &Scanner{
		ID:             newID,
		URL:            newURL,
		BeginHeight:    oldScanner.BeginHeight,
		EndHeight:      oldScanner.EndHeight,
		FilterTypeList: c.FilterTypeList,
		signalChan:     make(chan bool, 1),
		scanResultChan: c.ScanResultChan,
	}
	c.ScannerMap[scanner.ID] = scanner
	go scanner.run()
	scanner.signalChan <- true
	log.GetLogger().Debug("processResultErr", "newScanner", *scanner)

	c.deleteScanner(id)
}

func (c *Controller) notifyScanner(id ScannerID, toHeight int64) {
	scanner, ok := c.ScannerMap[id]
	if !ok {
		return
	}
	bh, eh := c.calcScannerHeight(toHeight)
	if eh == 0 {
		c.deleteScanner(id)
		return
	}
	scanner.BeginHeight = bh
	scanner.EndHeight = eh
	c.ScannerMap[id] = scanner
	scanner.signalChan <- true
}

func (c *Controller) deleteScanner(id ScannerID) {
	log.GetLogger().Debug("deleteScanner", "ID", id)
	scanner := c.ScannerMap[id]
	delete(c.ScannerMap, id)
	scanner.signalChan <- false
}

func (c *Controller) processResultList() {
	log.GetLogger().Debug("processResultList")
	if len(c.ScanResultList) == 0 {
		log.GetLogger().Debug("processResultList", "ScanResultList length", 0)
		return
	}

	lastHeight := db.GetLastHeight()
	index := 0
	flag := false

	for i, r := range c.ScanResultList {
		if r.EndHeight < lastHeight+1 {
			continue
		}
		if r.BeginHeight == lastHeight+1 {
			db.SetLastHeight(r.EndHeight)
			flag = true
			lastHeight = r.EndHeight
		} else {
			index = i
			break
		}
	}

	if len(c.ScanResultList) > 1 {
		c.ScanResultList = c.ScanResultList[index:]
	} else if flag {
		c.ScanResultList = c.ScanResultList[len(c.ScanResultList):]
	}
}

func (c *Controller) checkHeight() error {
	if c.FromHeight < 0 {
		return errors.New("from must greater than zero")
	}

	if c.ToHeight < 0 {
		return errors.New("to must greater than zero")
	}

	if c.FromHeight > c.ToHeight {
		return errors.New("to must greater than from")
	}

	if len(c.Config.URLs) == 0 {
		return errors.New("URLs can not be empty")
	}

	if c.FromHeight == 0 && c.ToHeight == 0 {
		return nil
	}

	// check height by block chain height
	h, e := query.GetCurrentHeight(c.Config.URLs[0])
	if e != nil {
		return fmt.Errorf("can not get current height from %s, error: %s", c.Config.URLs[0], e.Error())
	}

	if c.FromHeight != 0 && c.FromHeight > h {
		return errors.New("from height can not greater than current height")
	}

	if c.ToHeight != 0 && c.ToHeight > h {
		return errors.New("to height can not greater than current height")
	}

	// check height by local first height and last height
	fh := db.GetFirstHeight()
	if c.FromHeight != 0 && c.FromHeight < fh {
		return fmt.Errorf("local first height = %d, from height must greater than it", fh)
	}

	lh := db.GetLastHeight()
	if c.FromHeight != 0 && c.FromHeight > lh {
		return fmt.Errorf("local last height = %d, from height must NOT greater than it", lh)
	}

	if c.ToHeight != 0 && c.ToHeight <= lh {
		return fmt.Errorf("local last height = %d, to height must greater than it", lh)
	}

	return nil
}
