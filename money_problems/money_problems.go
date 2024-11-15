package money_problems

import (
	"errors"
	"fmt"
)

type Bitcoin int

// implement the stdlib Stringer interface
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// note the pointer here is by convention as receiving a copy of w would work for reading the value
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

var ErrNotEnoughJuice = errors.New("not enough juice in the weasel")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrNotEnoughJuice
	}

	w.balance -= amount

	return nil
}
