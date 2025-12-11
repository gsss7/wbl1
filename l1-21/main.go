package main

import "fmt"

type LegacyPaymentSystem struct{}

func (l *LegacyPaymentSystem) ProcessPaymentLegacy(amountCents int) bool {
	fmt.Printf("Legacy: Processing payment of %d cents\n", amountCents)
	return true
}

type ModernPaymentGateway struct{}

func (m *ModernPaymentGateway) Authorize(amountDollars float64) string {
	fmt.Printf("Modern: Authorizing payment of $%.2f\n", amountDollars)
	return fmt.Sprintf("auth_%f", amountDollars)
}

func (m *ModernPaymentGateway) Capture(transactionID string) bool {
	fmt.Printf("Modern: Capturing transaction %s\n", transactionID)
	return true
}

type PaymentProcessor interface {
	ProcessPayment(amount float64, currency string) (bool, error)
}

// Адаптер для старой системы
type LegacyPaymentAdapter struct {
	legacySystem *LegacyPaymentSystem
}

func NewLegacyPaymentAdapter(system *LegacyPaymentSystem) *LegacyPaymentAdapter {
	return &LegacyPaymentAdapter{legacySystem: system}
}

func (a *LegacyPaymentAdapter) ProcessPayment(amount float64, currency string) (bool, error) {
	amountCents := int(amount * 100)
	success := a.legacySystem.ProcessPaymentLegacy(amountCents)

	if !success {
		return false, fmt.Errorf("legacy payment failed")
	}
	return true, nil
}

// Адаптер для новой системы
type ModernPaymentAdapter struct {
	modernGateway *ModernPaymentGateway
}

func NewModernPaymentAdapter(gateway *ModernPaymentGateway) *ModernPaymentAdapter {
	return &ModernPaymentAdapter{modernGateway: gateway}
}

func (a *ModernPaymentAdapter) ProcessPayment(amount float64, currency string) (bool, error) {
	if currency != "USD" {
		return false, fmt.Errorf("modern gateway only supports USD")
	}

	transactionID := a.modernGateway.Authorize(amount)
	success := a.modernGateway.Capture(transactionID)

	if !success {
		return false, fmt.Errorf("modern payment capture failed")
	}
	return true, nil
}

func main() {
	legacySystem := &LegacyPaymentSystem{}
	modernGateway := &ModernPaymentGateway{}
	legacyAdapter := NewLegacyPaymentAdapter(legacySystem)
	modernAdapter := NewModernPaymentAdapter(modernGateway)
	processors := []PaymentProcessor{legacyAdapter, modernAdapter}

	fmt.Println("=== Processing payments through adapters ===")

	for i, processor := range processors {
		success, err := processor.ProcessPayment(99.99, "USD")
		if err != nil {
			fmt.Printf("Processor %d: Error: %v\n", i+1, err)
		} else {
			fmt.Printf("Processor %d: Payment successful: %v\n", i+1, success)
		}
		fmt.Println()
	}
}

/*
	Когда использовать:
	интеграция legacy кода, когда нужно использовать старый класс в новой системе
	создание адаптеров для мок-объектов
	когда переходите с одной системы на другую

	ПЛЮСЫ:
		Принцип единой ответственности: отделяет преобразование интерфейсов от основной логики
		Позволяет работать с несовместимыми интерфейсами
		Позволяет подменять реальные объекты адаптерами для тестов

	МИНУСЫ:
		Добавляет дополнительные классы/структуры
		Небольшой overhead из-за дополнительного слоя

	Реальные примеры:
		В каких-либо задачах никогда с этим не сталкивался. Чисто теоретически адаптеры применяют между
		разными СУБД и единым интерфейсом в ORM. В брокерах сообщений, ввод в Kafka.
*/
