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
	fmt.Printf("На счет пользователя %s зачислено %.2f\n", u.Name, sum)
}

func (u *User) Withdraw(sum float64) error {
	if u.Balance < sum {
		return errors.New("No enough funds on the account")
	}

	u.Balance -= sum
	fmt.Printf("Со счета пользователя %s списано %.2f\n", u.Name, sum)

	return nil
}

func (u *User) ShowBalance() {
	fmt.Printf("Имя : %s | Баланс: %.2f\n", u.Name, u.Balance)
}
