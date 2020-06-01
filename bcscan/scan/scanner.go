package scan

import (
	"errors"
	rpcclient "github.com/bcbchain/bclib/rpc/lib/client"
	"github.com/bcbchain/tendermint/types"
	"github.com/bcbchain/toolbox/bcscan/db"
	"github.com/bcbchain/toolbox/bcscan/log"
	"github.com/bcbchain/toolbox/bcscan/query"
	"github.com/bcbchain/toolbox/bcscan/utils"
)

type Scanner struct {
	ID             ScannerID
	URL            string
	BeginHeight    int64
	EndHeight      int64
	FilterTypeList []string

	signalChan     chan bool
	scanResultChan chan *Result

	client *rpcclient.JSONRPCClient
}

func (s *Scanner) run() {
	log.GetLogger().Info("scanner run", "scanner", *s)
	s.client = query.GetClient(s.URL, true)
	for {
		flag := <-s.signalChan
		if !flag {
			return
		}

		s.scan()
	}
}

func (s *Scanner) scan() {
	log.GetLogger().Info("scanner scan")
	var (
		hasHeaderType, hasTxType bool
	)
	for _, v := range s.FilterTypeList {
		if v == utils.FilterHeader {
			hasHeaderType = true
		}
		if v == utils.FilterTx {
			hasTxType = true
		}
	}

	if hasHeaderType && hasTxType {
		s.scanHeaderAndTx()
	} else if hasHeaderType {
		s.scanHeader()
	} else if hasTxType {
		s.scanTx()
	} else {
		s.sendScanResult(errors.New("invalid filter type"))
	}
}

func (s *Scanner) scanHeaderAndTx() {
	log.GetLogger().Info("scanHeaderAndTx", "begin height", s.BeginHeight, "end height", s.EndHeight)
	for h := s.BeginHeight; h <= s.EndHeight; h++ {
		if err := s.queryHeader(h); err != nil {
			s.sendScanResult(err)
			log.GetLogger().Error("queryHeader", "error", err)
			return
		}
		if err := s.queryTx(h); err != nil {
			s.sendScanResult(err)
			log.GetLogger().Error("queryTx", "error", err)
			return
		}
	}
	s.sendScanResult(nil)
}

func (s *Scanner) scanHeader() {
	log.GetLogger().Info("scanHeader", "begin height", s.BeginHeight, "end height", s.EndHeight)
	for h := s.BeginHeight; h <= s.EndHeight; h++ {
		if err := s.queryHeader(h); err != nil {
			s.sendScanResult(err)
			return
		}
	}
	s.sendScanResult(nil)
}

func (s *Scanner) scanTx() {
	log.GetLogger().Info("scanTx", "begin height", s.BeginHeight, "end height", s.EndHeight)
	for h := s.BeginHeight; h <= s.EndHeight; h++ {
		if err := s.queryHeader(h); err != nil {
			s.sendScanResult(err)
			return
		}
	}
	s.sendScanResult(nil)
}

func (s *Scanner) queryHeader(h int64) error {
	if db.HasHeader(h) {
		return nil
	}

	var (
		header *types.Header
		err    error
	)

	for i := 0; i <= 5; i++ {
		err = nil
		header, err = query.GetHeader(s.client, h)
		if err != nil {
			log.GetLogger().Error("query.GetHeader", "error", err)
			continue
		}
		log.GetLogger().Debug("query.GetHeader", "header", *header)
		db.SetHeader(h, header)
		return nil
	}
	return err
}

func (s *Scanner) queryTx(h int64) error {
	if db.HasTxWithHeight(h) {
		return nil
	}

	var (
		txList []string
		err    error
	)

	for i := 0; i <= 5; i++ {
		err = nil
		txList, err = query.GetTx(s.client, h)
		if err != nil {
			log.GetLogger().Error("query.GetTx", "error", err)
			continue
		}
		for i, v := range txList {
			log.GetLogger().Debug("query.GetTx", "height", h, "index", i, "tx", v)
			db.SetTx(h, int64(i), v)
		}
		db.SetTxWithHeightOK(h)
		return nil
	}
	return err
}

func (s *Scanner) sendScanResult(err error) {
	s.scanResultChan <- &Result{
		ScannerID:   s.ID,
		BeginHeight: s.BeginHeight,
		EndHeight:   s.EndHeight,
		Err:         err,
	}
}
