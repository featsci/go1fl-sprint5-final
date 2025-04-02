package actioninfo

import (
	...
)

// создайте интерфейс DataParser
type DataParser interface {
	Parse()
	ActionInfo()
}
// создайте функцию Info()
func Info(dataset []string, dp DataParser) {

}
