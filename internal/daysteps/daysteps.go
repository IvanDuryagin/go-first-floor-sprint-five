package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
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

	if len(datastringParse) != 2 {
		return errors.New("invalid data format")
	}

	steps, err := strconv.Atoi(datastringParse[0])
	if err != nil {
		return errors.New("invalid steps format, must be integer")
	}

	ds.Steps = steps

	duration, err := time.ParseDuration(datastringParse[1])
	if err != nil {
		return errors.New("invalid duration format: must be a valid time.Duration (e.g., '1h30m')")
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
