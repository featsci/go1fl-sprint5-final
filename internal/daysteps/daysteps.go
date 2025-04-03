package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/featsci/go1fl-sprint5-final/internal/personaldata"
	"github.com/featsci/go1fl-sprint5-final/internal/spentenergy"
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
	// ваш код ниже
	var s []string = strings.Split(datastring, ",")

	if len(s) != 2 {
		return errors.New("datastring != 2")
	}

	countSteps, err := strconv.Atoi(s[0])
	if err != nil {
		return err
	}

	timeDuration, err := time.ParseDuration(s[1])
	if err != nil {
		return err
	}
	ds.Steps = countSteps
	ds.Duration = timeDuration
	return nil
	// return err
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {
	// return "---", errors.New("Errrrr")
	// countSteps, timeDuration, err := parsePackage(data)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return ""
	// }

	if ds.Duration <= 0 {
		return "", errors.New("ds.Duration > 0")
	}
	calories := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	//spentcalories.WalkingSpentCalories(countSteps, weight, height, timeDuration)

	distanceSteps := spentenergy.Distance(ds.Steps)

	s := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distanceSteps, calories)

	return s, nil
}
