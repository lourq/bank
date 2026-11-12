package main

import (
	"fmt"

	"github.com/lourq/bank/payment"
	"github.com/lourq/bank/user"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// Part 1
	john := &user.User{ID: "1", Name: "John", Balance: 100.0}
	jack := &user.User{ID: "2", Name: "Jack", Balance: 10.0}
	bob := &user.User{ID: "3", Name: "Bob", Balance: 0.0}

	john.ShowBalance()
	jack.ShowBalance()
	bob.ShowBalance()

	fmt.Println("")
	fmt.Println("")

	john.Deposit(50.1)
	john.ShowBalance()
	err := john.Withdraw(50.1)
	checkErr(err)
	john.ShowBalance()

	jack.Deposit(00.1)
	jack.ShowBalance()
	err = jack.Withdraw(10.1)
	checkErr(err)
	jack.ShowBalance()
	err = jack.Withdraw(100.1)
	checkErr(err)
	jack.ShowBalance()

	err = bob.Withdraw(100)
	bob.ShowBalance()
	checkErr(err)
	bob.Deposit(100)
	bob.ShowBalance()

	john.Deposit(50000)
	bob.Deposit(50000)
	jack.Deposit(50000)

	// Part 2
	fmt.Println()
	fmt.Println()
	fmt.Println()

	ps := &payment.PaymentSystem{
		Users:            map[string]*user.User{},
		TransactionQueue: []payment.Transaction{},
	}

	ps.AddUser(john)
	ps.AddUser(jack)
	ps.AddUser(bob)

	ps.AddTransaction(payment.Transaction{
		FromID: john.ID,
		ToID:   bob.ID,
		Amount: 50})
	ps.AddTransaction(payment.Transaction{
		FromID: john.ID,
		ToID:   jack.ID,
		Amount: 20})
	ps.AddTransaction(payment.Transaction{
		FromID: john.ID,
		ToID:   jack.ID,
		Amount: 10})

	for _, t := range ps.TransactionQueue {
		ps.ProcessingTransactions(t)
	}

	john.ShowBalance()
	bob.ShowBalance()
	jack.ShowBalance()
}
