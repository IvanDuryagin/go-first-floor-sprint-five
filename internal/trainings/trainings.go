package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	var datastringParse []string
	var current []rune
	for _, char := range datastring {
		if char == ',' {
			datastringParse = append(datastringParse, string(current))
			current = []rune{}
		} else {
			current = append(current, char)
		}
	}
	datastringParse = append(datastringParse, string(current))

	if len(datastringParse) != 3 {
		return errors.New("invalid data format")
	}

	steps, err := strconv.Atoi(datastringParse[0])
	if err != nil {
		return errors.New("invalid steps format, must be integer")
	}

	t.Steps = steps

	t.TrainingType = datastringParse[1]

	duration, err := time.ParseDuration(datastringParse[2])
	if err != nil {
		return errors.New("invalid duration format: must be a valid time.Duration (e.g., '1h30m')")
	}

	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var spentCalories float64
	var err error

	switch t.TrainingType {
	case "Ходьба":
		spentCalories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Бег":
		spentCalories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		return "", errors.New("unknown type of training")
	}

	if err != nil {
		return "", err
	}

	durationInHours := t.Duration.Hours()

	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f", t.TrainingType, durationInHours, distance, meanSpeed, spentCalories)

	return result, nil
}
