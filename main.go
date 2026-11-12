package main

import (
	"fmt"
	"sync"

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

	// Part 2
	fmt.Println()
	fmt.Println()

	john.Deposit(50000)
	bob.Deposit(50000)
	jack.Deposit(50000)

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
		err := ps.ProcessingTransactions(t)
		checkErr(err)
	}

	john.ShowBalance()
	bob.ShowBalance()
	jack.ShowBalance()

	// Part 3
	fmt.Println()
	fmt.Println("Обработка новых транзакций через канал:")

	ps.TransactionQueue = nil
	transactions := []payment.Transaction{
		{FromID: john.ID, ToID: bob.ID, Amount: 100},
		{FromID: bob.ID, ToID: jack.ID, Amount: 40},
		{FromID: jack.ID, ToID: john.ID, Amount: 15},
		{FromID: john.ID, ToID: jack.ID, Amount: 200},
		{FromID: jack.ID, ToID: bob.ID, Amount: 75},
		{FromID: bob.ID, ToID: john.ID, Amount: 30},
		{FromID: "unknown", ToID: john.ID, Amount: 10},
		{FromID: john.ID, ToID: "unknown", Amount: 10},
		{FromID: john.ID, ToID: bob.ID, Amount: 1_000_000},
		{FromID: john.ID, ToID: bob.ID, Amount: 5},
	}
	for _, t := range transactions {
		ps.AddTransaction(t)
	}

	var wg sync.WaitGroup
	ch := make(chan payment.Transaction, len(ps.TransactionQueue))

	wg.Add(1)
	go ps.Worker(ch, &wg)

	for _, t := range ps.TransactionQueue {
		ch <- t
	}

	close(ch)
	wg.Wait()

	john.ShowBalance()
	bob.ShowBalance()
	jack.ShowBalance()
}
