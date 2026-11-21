package spentenergy

import (
	"errors"
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
		return 0, errors.New("количество шагов должно быть больше 0")
	}
	if weight <= 0 || height <= 0 {
		return 0, errors.New("вес и рост должны быть больше 0")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность тренировки должна быть больше 0")
	}
	speed := MeanSpeed(steps, height, duration)
	durationInMinutes := duration.Minutes()
	spentCalories := (weight * speed * durationInMinutes) / minInH

	return spentCalories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть больше 0")
	}
	if weight <= 0 || height <= 0 {
		return 0, errors.New("вес и рост должны быть больше 0")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность тренировки должна быть больше 0")
	}
	speed := MeanSpeed(steps, height, duration)

	durationInMinutes := duration.Minutes()
	spentCalories := (weight * speed * durationInMinutes) / minInH

	return spentCalories, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 || duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)

	return distance / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	stepLength := height * stepLengthCoefficient
	distanceInMeters := float64(steps) * stepLength

	return distanceInMeters / mInKm
}
