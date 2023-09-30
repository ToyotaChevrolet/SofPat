package main

import "fmt"

type PaymentMethod interface {
	Pay(amount float64) string
}

type CreditCardPayment struct{}

func (cc *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using Credit Card", amount)
}

type PayPalPayment struct{}

func (pp *PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using PayPal", amount)
}

type CashPayment struct{}

func (cp *CashPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f in Cash", amount)
}

type PaymentProcessor struct {
	PaymentMethod PaymentMethod
}

func (pp *PaymentProcessor) ProcessPayment(amount float64) string {
	return pp.PaymentMethod.Pay(amount)
}

func main() {
	creditCard := &CreditCardPayment{}
	payPal := &PayPalPayment{}
	cash := &CashPayment{}

	processor := &PaymentProcessor{PaymentMethod: creditCard}

	fmt.Println(processor.ProcessPayment(190.0))
	fmt.Println(processor.ProcessPayment(120.0))

	processor.PaymentMethod = payPal
	fmt.Println(processor.ProcessPayment(100.0))

	processor.PaymentMethod = cash
	fmt.Println(processor.ProcessPayment(30.0))
}
