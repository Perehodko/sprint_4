package daysteps

import (
	"time"
	"fmt"
	"strconv"
	"strings"	

)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

//"678,0h50m"
func parsePackage(data string) (int, time.Duration, error) {
	if data == "" {
		return 0, 0, fmt.Errorf("empty input")
	}

	slices := strings.Split(data, ",")
	if len(slices) != 2 {
		return 0, 0, fmt.Errorf("invalid input")
	}

	stepCount := slices[0]
	if stepCount == "" {
		return 0, 0, fmt.Errorf("empty step count")
	}

	if steps, err := strconv.Atoi(stepCount); err != nil {
		return 0, 0, fmt.Errorf("invalid step count: %w", err)
	}

	if steps <= 0 {
		return 0, 0, fmt.Errorf("invalid steps")
	}

	duration := slices[1]
	if duration == "" {
		return 0, 0, fmt.Errorf("empty duration")
	}

	if walkDuration, err := time.ParseDuration(duration); err != nil {
		return 0, 0, fmt.Errorf("invalid duration: %w", err)
	}

	if walkDuration <= 0 {
		return 0, 0, fmt.Errorf("invalid duration")
	}

	return steps, walkDuration, nil
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
