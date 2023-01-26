package BackEnd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func QueryExtras() [][]string {

	fmt.Printf("\n\n Would you like any extras added to your order? Items in stock : ")
	extras := [][]string{
		{"Paintbrush 2 inch", "2", "3.99"},
		{"Paint tray", "1", "1.99"},
	}

	for _, thisSlice := range extras {
		fmt.Printf(" 1. %s , £%s per item \n", thisSlice[0], thisSlice[2])
	}

	moreOrders := true
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Would you like to add any items to your order? [Y/N] :")
	for moreOrders {
		scanner.Scan()
		textIn := scanner.Text()
		if strings.ToUpper(textIn) == "Y" || strings.ToUpper(textIn) == "YES" {

			fmt.Println("Please enter the full name of the item, or '#' Hashtag, and then the number in the list")
			fmt.Println("If this was a mistake, type 'x', 'c' or 'cancel'.")
			fmt.Println("Which item would you like to order? : ")
			scanner.Scan()
			thisItem := scanner.Text()

			//TC Which item
			quantity := getQuantity(thisItem)
			// test comment

		} else if strings.ToUpper(textIn) == "N" || strings.ToUpper(textIn) == "No" {
			moreOrders = false
		} else {
			fmt.Printf("Sorry, that was not recognised as 'Y' or 'N'. Try again : ")
		}
		fmt.Printf("")
	}

	return [][]string{{"0"}}
}

func getQuantity(thisItem string) int {
	validInput := false
	for validInput {
		fmt.Printf("How many '%s' would you like to order? : ")
	}
	return 0
}
