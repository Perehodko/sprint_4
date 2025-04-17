package daysteps

import (
	"time"
	"fmt"
	"strconv"
	"strings"	
	"go1fl-4-sprint-final/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

//"678,0h50m"
func parsePackage(data string) (int, time.Duration, error) {
	var (
		step int 
		err error
		walkDuration time.Duration
	)

	if data == "" {
		return 0, 0, fmt.Errorf("empty input")
	}

	slices := strings.Split(data, ",")
	if len(slices) != 2 {
		return 0, 0, fmt.Errorf("invalid input")
	}

	if slices[0] == "" {
		return 0, 0, fmt.Errorf("empty step count")
	}

	if step, err = strconv.Atoi(slices[0]); err != nil {
		return 0, 0, fmt.Errorf("invalid step count: %w", err)
	}

	if step <= 0 {
		return 0, 0, fmt.Errorf("invalid steps")
	}

	if slices[1] == "" {
		return 0, 0, fmt.Errorf("empty duration")
	}

	if walkDuration, err = time.ParseDuration(slices[1]); err != nil {
		return 0, 0, fmt.Errorf("invalid duration: %w", err)
	}

	if walkDuration <= 0 {
		return 0, 0, fmt.Errorf("invalid duration")
	}

	return step, walkDuration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, walkDuration, err := parsePackage(data)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	if steps < 1 {
		return ""
	}

	distanseMeters := float64(steps) * float64(stepLength)
	distanseKm := distanseMeters / mInKm
	calories := spentcalories.WalkingSpentCalories(steps, weight, height, walkDuration)
	dayActionInfo := fmt.Sprintf("Количество шагов: %v.\nДистанция составила %v км.\nВы сожгли %v ккал.", steps, distanseKm, calories) 

	return dayActionInfo
}
