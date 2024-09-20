package main

import (
	"fmt"
	"math"
)

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
// Добавление проверки для этажей
func (c *Cargo) ManualLiftCost() float64 {
	if c.HasElevator || c.Floor <= 1 {
		return 0 // Нет доплаты за лифт или первый этаж
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

// Добавление метода валидации для проверки ввода
func (c *Cargo) Validate() error {
	if c.Weight <= 0 {
		return fmt.Errorf("Вес груза должен быть больше 0")
	}
	if c.Floor < 1 {
		return fmt.Errorf("Этаж должен быть не меньше 1")
	}
	return nil
}

func main() {
	var weight float64
	var floor int
	var hasElevator bool

	fmt.Println("Введите вес груза (в кг):")
	fmt.Scan(&weight)

	fmt.Println("Введите этаж:")
	fmt.Scan(&floor)

	fmt.Println("Есть ли лифт? (1 - да, 0 - нет):")
	var elevatorInput int
	fmt.Scan(&elevatorInput)
	hasElevator = elevatorInput == 1

	cargo := Cargo{
		Weight:      weight,
		Floor:       floor,
		HasElevator: hasElevator,
	}

	// Валидация
	if err := cargo.Validate(); err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	totalCost := cargo.TotalCost()
	fmt.Printf("Общая стоимость подъёма груза: %.2f руб.\n", totalCost)
}
