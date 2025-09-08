package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("invalid data format")
	}

	rawSteps := parts[0]
	trimmedSteps := strings.TrimSpace(rawSteps)
	if rawSteps != trimmedSteps {
		return errors.New("error: invalid spaces in the number of steps")
	}

	steps, err := strconv.Atoi(trimmedSteps)
	if err != nil {
		return fmt.Errorf("error converting number of steps:number of steps: %w", err)
	}
	if steps <= 0 {
		return errors.New("error: invalid step count value")
	}
	t.Steps = steps

	t.TrainingType = strings.TrimSpace(parts[1])
	if t.TrainingType == "" {
		return errors.New("error: empty training type")
	}

	rawDuration := parts[2]
	trimmedDuration := strings.TrimSpace(rawDuration)
	if rawDuration != trimmedDuration {
		return errors.New("error: invalid spaces in duration")
	}

	duration, err := time.ParseDuration(trimmedDuration)
	if err != nil {
		return fmt.Errorf("invalid time data: %w", err)
	}
	if duration <= 0 {
		return errors.New("error: invalid duration")
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

	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType, durationInHours, distance, meanSpeed, spentCalories)

	return result, nil
}
