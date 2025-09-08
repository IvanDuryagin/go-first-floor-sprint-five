package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("invalid data format")
	}

	rawSteps := parts[0]
	trimmedSteps := strings.TrimSpace(rawSteps)

	if rawSteps != trimmedSteps {
		return errors.New("error: invalid spaces in the number of steps")
	}

	steps, err := strconv.Atoi(trimmedSteps)
	if err != nil {
		return errors.New("step count conversion error")
	}
	if steps <= 0 {
		return errors.New("step count must be positive")
	}
	ds.Steps = steps

	durationStr := strings.TrimSpace(parts[1])
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return errors.New("invalid time format")
	}
	if duration <= 0 {
		return errors.New("duration must be positive")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	spentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", errors.New("invalid data")
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, spentCalories)

	return result, nil
}
