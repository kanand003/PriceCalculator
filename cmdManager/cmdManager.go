package cmdmanager

import "fmt"

type CommandManager struct {
}

func (cmd CommandManager) ReadLines() ([]string, error) {
	fmt.Println("Please Enter the Prices...")
	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func (cmd CommandManager) WriteResult(data interface{}) error {
	fmt.Println("Result:")
	fmt.Println(data)
	return nil
}

func NewCommandManager() *CommandManager {
	return &CommandManager{}
}
