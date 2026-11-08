package user

import (
	"errors"
	"fmt"
)

type User struct {
	ID, Name string
	Balance  float64
}

func (u *User) Deposit(sum float64) {
	u.Balance += sum
}

func (u *User) Withdraw(sum float64) error {
	if u.Balance < sum {
		return errors.New("You don't have enough money")
	}
	u.Balance -= sum

	return nil
}

func (u User) ShowBalance() {
	fmt.Printf("Name : %s | Balance : %.2f\n", u.Name, u.Balance)
}
