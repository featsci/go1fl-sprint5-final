package trainings

import (
	"time"

	"github.com/featsci/go1fl-sprint5-final/internal/personaldata"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(datastring string) (err error) {

}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {

}
