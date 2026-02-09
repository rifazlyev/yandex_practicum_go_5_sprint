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

func (ds *DaySteps) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 2 {
		return errors.New("invalid string format")
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("invalid steps format: %w", err)
	}
	if steps <= 0 {
		return errors.New("steps count must be positive")
	}

	duration, err := time.ParseDuration(data[1])
	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}
	if duration <= 0 {
		return errors.New("duration must be positive")
	}

	ds.Steps = steps
	ds.Duration = duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories), nil
}
