package payment

import (
	"errors"
	"fmt"
	"sync"

	u "github.com/lourq/bank/user"
)

type Transaction struct {
	FromID, ToID string
	Amount       float64
}

type PaymentSystem struct {
	Users            map[string]*u.User
	TransactionQueue []Transaction
}

func (ps PaymentSystem) AddUser(user *u.User) {
	ps.Users[user.ID] = user
}

func (ps *PaymentSystem) AddTransaction(t Transaction) {
	ps.TransactionQueue = append(ps.TransactionQueue, t)
}

func (ps *PaymentSystem) ProcessingTransactions(t Transaction) error {
	check := func(ok bool) error {
		if !ok {
			return errors.New("User not found")
		}
		return nil
	}

	receiver, ok := ps.Users[t.ToID]
	if err := check(ok); err != nil {
		return err
	}

	sender, ok := ps.Users[t.FromID]
	if err := check(ok); err != nil {
		return err
	}

	err := sender.Withdraw(t.Amount)
	if err != nil {
		return err
	}

	receiver.Deposit(t.Amount)

	fmt.Printf("Транзакция успешно выполнена: %s -> %s, сумма: %.2f\n", t.FromID, t.ToID, t.Amount)
	return nil
}

func (ps *PaymentSystem) Worker(ch <-chan Transaction, wg *sync.WaitGroup) {
	defer wg.Done()

	for t := range ch {
		err := ps.ProcessingTransactions(t)
		if err != nil {
			fmt.Println("Ошибка при обработке транзакции:", err)
		}
	}
}
