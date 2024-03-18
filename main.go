package main

import (
	"bufio"
	"fmt"
	"homework1662/Bank"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var client = Bank.NewBankAccount(0)

func init() {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; i < 10; i++ {
		go depositGoRoutine(client, r)
	}

	time.Sleep(time.Second)
	for i := 0; i < 5; i++ {
		go withdrawGoRoutine(client, r)
	}

}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter command (balance, deposit, withdrawal, exit): ")
		scanner.Scan()
		command := scanner.Text()

		switch command {
		case "balance":
			handlerPrint(client)
		case "deposit":
			handlerDeposit(client, scanner)
		case "withdrawal":
			handlerWithdraw(client, scanner)
		case "exit":
			handlerExit()
			return
		default:
			handlerUnknown()
		}
	}
}

func sleep() {
	time.Sleep(time.Duration(rand.Intn(500)+500) * time.Millisecond)
}

func depositGoRoutine(client *Bank.BankAccount, r *rand.Rand) {
	for {
		sleep()
		client.Deposit(r.Intn(10) + 1)
	}
}

func withdrawGoRoutine(client *Bank.BankAccount, r *rand.Rand) {
	for {
		sleep()
		err := client.Withdraw(r.Intn(5) + 1)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
func handlerPrint(client *Bank.BankAccount) {
	fmt.Println("Your balance is", client.Balance())
}

func handlerDeposit(client *Bank.BankAccount, scanner *bufio.Scanner) {
	fmt.Print("Enter deposit amount: ")
	scanner.Scan()
	amount, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid amount:")
		return
	}
	client.Deposit(amount)
	fmt.Println("Deposited", amount, "into the account. Current balance:", client.Balance())
}

func handlerWithdraw(client *Bank.BankAccount, scanner *bufio.Scanner) {
	fmt.Print("Enter withdrawal amount: ")
	scanner.Scan()
	amount, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid amount:")
		return
	}
	err = client.Withdraw(amount)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Withdrew", amount, "from the account. Current balance:", client.Balance())
}

func handlerExit() {
	fmt.Println("Exiting the program.")
}

func handlerUnknown() {
	fmt.Println("Unsupported command. You can use commands: balance, deposit, withdrawal, exit")
}
