// 03_atm_errors.go

package main
import "fmt"

type BankAccount struct {
	balance int
}

func (account *BankAccount) Deposit(amount int) {
	if (amount <= 0) {
		panic(fmt.Errorf("Error: cannot deposit a negative amount! KRW %d \n", amount))
	}
	account.balance += amount
	fmt.Printf("Deposit successful! Current balance: KRW %d \n\n", account.balance)
}

func (account *BankAccount) Withdraw(amount int) {
	if (amount <= 0) {
		panic(fmt.Errorf("Error: cannot withdrwaw a negative amount! KRW %d \n", amount))
	}

	if (amount > account.balance) {
		panic(fmt.Errorf("Error: Insufficient balance!\nCurrent balance: KRW %d, Requested amount: KRW %d \n", account.balance, amount))
	}
	
	account.balance -= amount
	fmt.Printf("Withdrawal successful! Current balance: KRW %d \n\n", account.balance)
}


func main() {
    account := BankAccount{balance: 100} // initial balance
    var choice string

    for {
        fmt.Print("Deposit (1), Withdraw (2), Exit (Others): ")
        fmt.Scanln(&choice)

        switch choice {
        case "1":
            var amount int
            fmt.Print("Enter the amount to deposit: ")
            fmt.Scanln(&amount)

            // try to deposit -> if an error occurs, handles the error
            func() {
                defer func() {
                    if r := recover(); r != nil {
                        fmt.Println(r)  // print error message
                    }
                }()
                account.Deposit(amount)
            }()

        case "2":
            var amount int
            fmt.Print("Enter the amount to withdraw: ")
            fmt.Scanln(&amount)

            // try to withdraw -> if an error occurs, handles the error
            func() {
                defer func() {
                    if r := recover(); r != nil {
                        fmt.Println(r)  // print error message
                    }
                }()
                account.Withdraw(amount)
            }()

        default:
            fmt.Println(">> Exiting the program!")
            return
        }
    }
}