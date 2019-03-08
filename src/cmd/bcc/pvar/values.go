package pvar

import (
	"errors"
	"strings"
)

// SplitValues splitValues for split values to []string
func splitValues(values string) (splitValues []string, err error) {

	// step 1. split string with "\\", and concat it to splitValues's tail item at last
	splitSlash := strings.Split(values, "\\\\")
	for _, item := range splitSlash {

		// step 2. split splitSlash item with "\,", and concat "," to slashComma's tail item at last
		splitSlashComma := strings.Split(item, "\\,")
		var slashComma []string
		for _, item = range splitSlashComma {

			// step 3. split splitSlashComma item with ",", and append items to comma
			temp3 := strings.Split(item, ",")
			var comma []string
			for _, item = range temp3 {
				if strings.Contains(item, "\\") {
					return splitValues, errors.New("slash number failed")
				}

				comma = append(comma, item)
			}

			if len(slashComma) == 0 {
				slashComma = append(slashComma, comma...)
			} else {
				last := slashComma[len(slashComma)-1]
				slashComma = slashComma[:len(slashComma)-1]
				slashComma = append(slashComma, last+","+comma[0])
				comma = comma[1:]
				slashComma = append(slashComma, comma...)
			}
		}

		if len(splitValues) == 0 {
			splitValues = append(splitValues, slashComma...)
		} else {
			last := splitValues[len(splitValues)-1]
			splitValues = splitValues[:len(splitValues)-1]
			splitValues = append(splitValues, last+"\\\\"+slashComma[0])
			slashComma = slashComma[1:]
			splitValues = append(splitValues, slashComma...)
		}
	}

	return
}
