package spentcalories

import (
	"time"
	"strconv"
	"strings"
	"fmt"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

//"3456,Ходьба,3h00m"
func parseTraining(data string) (int, string, time.Duration, error) {
	var duration time.Duration

	parts := strings.Split(data, ",")

	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("invalid input")
	}

	stepsStr := parts[0]
	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid steps")
	}

	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("invalid steps")
	}

	kindActivity := parts[1]

	durationStr := parts[2]
	duration, err = time.ParseDuration(durationStr)
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid duration")
	}

	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("invalid duration")
	}

	return steps, kindActivity, duration, nil
}

func distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distance := float64(steps) * stepLength

	return distance/mInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}
	avrSpeed := distance(steps, height) / duration.Hours()

	return avrSpeed
}

//"3456,Ходьба,3h00m"
func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, kindActivity, duration, err := parseTraining(data)
	if err != nil {
		return "", err
	}

	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return "", fmt.Errorf("invalid input")		
	}

	switch kindActivity {
	case "Бег":
		kacl, err := RunningSpentCalories(steps, weight, height, duration)
		avgSpeed := meanSpeed(steps, height, duration)
		distance := distance(steps, height)
		if err != nil {
			return "", err
		}
		result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %0.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			kindActivity, duration.Hours(), distance, avgSpeed, kacl)
		return result, nil
	case "Ходьба":
		kacl, err := WalkingSpentCalories(steps, weight, height, duration)
		avgSpeed := meanSpeed(steps, height, duration)
		distance := distance(steps, height)
		if err != nil {
			return "", err		
		}
		result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %0.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		kindActivity, duration.Hours(), distance, avgSpeed, kacl)
		return result, nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {	
		return 0, fmt.Errorf("invalid input")
	}
	avgSpeed := meanSpeed(steps, height, duration)
	if avgSpeed == 0 {
		return 0, fmt.Errorf("invalid input")
	}
	runningCalories := (duration.Minutes() * weight * avgSpeed )/ minInH 

	return runningCalories, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {	
		return 0, fmt.Errorf("invalid input")
	}

	avgSpeed := meanSpeed(steps, height, duration)
	if avgSpeed <= 0 {
		return 0, fmt.Errorf("invalid input")
	}

	walkingCalories := ((duration.Minutes() * weight * avgSpeed ) / minInH) * walkingCaloriesCoefficient

	return walkingCalories, nil
}
