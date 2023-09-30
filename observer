package main

import (
	"fmt"
)

type Observer interface {
	Update(paymentMethod string)
}

type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	Notify(paymentMethod string)
}

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

type PaymentSubject struct {
	observers     []Observer
	currentMethod PaymentMethod
}

func (ps *PaymentSubject) Register(observer Observer) {
	ps.observers = append(ps.observers, observer)
}

func (ps *PaymentSubject) Unregister(observer Observer) {
	for i, obs := range ps.observers {
		if obs == observer {
			ps.observers = append(ps.observers[:i], ps.observers[i+1:]...)
			break
		}
	}
}

func (ps *PaymentSubject) Notify(paymentMethod string) {
	for _, observer := range ps.observers {
		observer.Update(paymentMethod)
	}
}

func (ps *PaymentSubject) SetPaymentMethod(paymentMethod PaymentMethod) {
	ps.currentMethod = paymentMethod
	ps.Notify(paymentMethod.Pay(0))
}

func main() {
	creditCard := &CreditCardPayment{}
	payPal := &PayPalPayment{}
	cash := &CashPayment{}

	paymentSubject := &PaymentSubject{}

	observer1 := &PaymentObserver{name: "Observer 1"}
	observer2 := &PaymentObserver{name: "Observer 2"}

	paymentSubject.Register(observer1)
	paymentSubject.Register(observer2)

	paymentSubject.SetPaymentMethod(creditCard)

	paymentSubject.SetPaymentMethod(payPal)

	paymentSubject.Unregister(observer1)

	paymentSubject.SetPaymentMethod(cash)
}

type PaymentObserver struct {
	name string
}

func (o *PaymentObserver) Update(paymentMethod string) {
	fmt.Printf("%s received an update: Payment method changed to %s\n", o.name, paymentMethod)
}
