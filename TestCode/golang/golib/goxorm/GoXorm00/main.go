package main

import (
	"fmt"
)

const prompt = `Please enter number of operation
1. Create new account
2. Show detail of account
3. Deposit
4. Withdrow
5. Make transfer
6. List account by Id
7. List account by balance
8. Delete account
9. Exit`

func main() {
	fmt.Println("Welcome bank of xorm!")

	for {
		fmt.Println(prompt)

		var num int
		fmt.Scanf("%d\n", &num)
		switch num {
		case 1:
			fmt.Println("Please enter <name> <balance>:")
			var name string
			var balance float64
			fmt.Scanf("%s %f\n", &name, &balance)
			if err := NewAccount(name, balance); err != nil {
				fmt.Println(err)
			}
		}
	}
}
