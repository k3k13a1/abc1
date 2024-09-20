package main

import "math"

// Структура для хранения информации о грузе
type Cargo struct {
	Weight      float64 // вес груза
	Floor       int     // этаж
	HasElevator bool    // наличие работающего лифта
}

// Метод для расчёта базовой стоимости в зависимости от массы
func (c *Cargo) BaseCost() float64 {
	switch {
	case c.Weight <= 50:
		return 300
	case c.Weight <= 100:
		return 1000
	case c.Weight <= 300:
		return 2000
	default:
		return 0 // Логика для грузов свыше 300 кг
	}
}

// Тесты
// Метод для расчёта дополнительной стоимости за подъём вручную
func (c *Cargo) ManualLiftCost() float64 {
	if c.HasElevator {
		return 0 // Если есть лифт, доплата не требуется
	}

	additionalCostPerFloor := 300.0
	weightMultiplier := math.Ceil(c.Weight / 100.0)
	floorsToClimb := float64(c.Floor - 1)
	return additionalCostPerFloor * weightMultiplier * floorsToClimb
}

// Метод для расчёта общей стоимости
func (c *Cargo) TotalCost() float64 {
	return c.BaseCost() + c.ManualLiftCost()
}
