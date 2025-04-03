package actioninfo

import (
	"fmt"
)

// создайте интерфейс DataParser
type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	// fmt.Println(dataset, dp)
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			fmt.Println(err)
			continue
		}

		actionInfo, err := dp.ActionInfo()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(actionInfo)
		// fmt.Println(v)
	}
}
