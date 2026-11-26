package main

import "fmt"

func main() {
	// Исходные данные
	isBlocked := false
	balance := 1500.0
	amount := 200
	level := "gold"
	points := 950

	// Проверяем, заблокирована ли карта
	if isBlocked {
		fmt.Println("Card is blocked")
		return // Завершаем выполнение, если карта заблокирована
	}

	// Проверяем, достаточно ли средств на счёте
	// Преобразуем amount из int в float64, чтобы типы совпадали
	if float64(amount) > balance {
		fmt.Println("Insufficient funds")
	} else {
		fmt.Println("Payment approved")
	}

	// Проверяем уровень пользователя и начисляем кешбэк
	if level == "gold" {
		fmt.Println("Cashback: 5%")
	} else {
		fmt.Println("Cashback: 1%")
	}

	// Определяем статус по количеству баллов
	// Проверяем от большего порога к меньшему
	switch {
	case points >= 900:
		fmt.Println("Status: Platinum")
	case points >= 500:
		fmt.Println("Status: Gold")
	case points >= 100:
		fmt.Println("Status: Silver")
	default:
		fmt.Println("Status: Bronze")
	}
}