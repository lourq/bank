package user

import (
	"errors"
	"fmt"
)

type User struct {
	ID, Name string
	Balance  float32
}

func (u *User) Deposit(sum float32) {
	u.Balance += sum
}

func (u *User) Withdraw(sum float32) error {
	if u.Balance < sum {
		return errors.New("You don't have enough money")
	}
	u.Balance -= sum

	return nil
}

func (u User) ShowBalance() {
	fmt.Printf("Name : %s | Balance : %.2f\n", u.Name, u.Balance)
}
