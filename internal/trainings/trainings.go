package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	_ "time"

	"github.com/featsci/go1fl-sprint5-final/internal/personaldata"
	_ "github.com/featsci/go1fl-sprint5-final/internal/personaldata"
	"github.com/featsci/go1fl-sprint5-final/internal/spentenergy"
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
	// ваш код ниже
	var s []string = strings.Split(datastring, ",")

	if len(s) != 3 {
		return errors.New("datastring != 3")
	}

	if s[1] != "Бег" && s[1] != "Ходьба" {
		return errors.New("not match with move type")
	}
	t.TrainingType = s[1]

	countSteps, err := strconv.Atoi(s[0])
	if err != nil {
		return err
	}
	t.Steps = countSteps

	timeDuration, err := time.ParseDuration(s[2])
	if err != nil {
		return err
	}
	t.Duration = timeDuration
	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	// ваш код ниже
	var (
		d,
		m,
		w float64
	)
	d = spentenergy.Distance(t.Steps)
	m = spentenergy.MeanSpeed(t.Steps, t.Duration)
	switch t.TrainingType {
	case "Ходьба":
		w = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		w = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Duration)
	default:
		return "неизвестный тип тренировки", errors.New("unknown training type")
	}

	s := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", t.TrainingType, t.Duration.Hours(), d, m, w)

	return s, nil

	// return "", nil
}
