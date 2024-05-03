package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

// Technically you do not need to change Balance to use a pointer receiver as taking a copy of the balance is fine.
// However, by convention you should keep your method receiver types the same for consistency.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
	// In fact, here we returning struct pointer: (*w).balance
	// https://go.dev/ref/spec#Method_values
}

// When calling func (w Wallet) Deposit(amount float64) the w is a copy of whatever we called the method from.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("Cannot withdraw, insufficient funds")
	}

	w.balance -= amount
	return nil
}

// This interface is defined in the fmt package and lets you define how your type is printed
// when used with the %s format string in prints.
func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}
