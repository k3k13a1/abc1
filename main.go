package main

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
