package payment

import (
	"errors"

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

func (ps *PaymentSystem) AddTransaction(fromID, toID string, amount float64) {
	ps.TransactionQueue = append(ps.TransactionQueue, Transaction{
		FromID: fromID,
		ToID:   toID,
		Amount: amount})
}

func (ps *PaymentSystem) ProcessingTransactions() error {
	check := func(ok bool) error {
		if !ok {
			return errors.New("пользователь с таким ID не найден")
		}
		return nil
	}

	for _, transaction := range ps.TransactionQueue {
		receiver, ok := ps.Users[transaction.ToID]
		if err := check(ok); err != nil {
			return err
		}

		sender, ok := ps.Users[transaction.FromID]
		if err := check(ok); err != nil {
			return err
		}

		err := sender.Withdraw(transaction.Amount)
		if err != nil {
			return err
		}

		receiver.Deposit(transaction.Amount)
	}

	return nil
}
