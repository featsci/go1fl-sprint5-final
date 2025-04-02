package daysteps

import (
	"time"

	"github.com/featsci/go1fl-sprint5-final/internal/personaldata"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {

}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {

}
