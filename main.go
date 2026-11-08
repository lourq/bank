package main

import (
	"fmt"

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

	john.Deposit(50.1)
	err := john.Withdraw(50.1)
	checkErr(err)
	john.ShowBalance()

	jack.Deposit(00.1)
	err = jack.Withdraw(10.1)
	checkErr(err)
	err = jack.Withdraw(100.1)
	checkErr(err)
	jack.ShowBalance()

	err = bob.Withdraw(100)
	checkErr(err)
	bob.Deposit(100)
	bob.ShowBalance()

	// Part 2
}
