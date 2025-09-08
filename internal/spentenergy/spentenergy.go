package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("invalid steps value: %d (must be positive)", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("invalid weight value: %.2f (must be positive)", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("invalid height value: %.2f (must be positive)", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("invalid duration value: %v (must be positive)", duration)
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	calories := (weight * meanSpeed * durationInMinutes) / minInH
	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("invalid steps value: %d (must be positive)", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("invalid weight value: %.2f (must be positive)", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("invalid height value: %.2f (must be positive)", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("invalid duration value: %v (must be positive)", duration)
	}
	meanSpeed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	return (weight * meanSpeed * durationInMinutes) / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps == 0 || duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := float64(height) * stepLengthCoefficient
	distanceM := float64(steps) * stepLength
	return distanceM / mInKm
}
