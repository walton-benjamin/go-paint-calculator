package BackEnd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func QueryExtras() [][]string {

	fmt.Printf("\n\n Would you like any extras added to your order? Items in stock : ")
	extrasOrdered := {{"0"}}
	extras := [][]string{
		{"2 inch bristles Paintbrush", "3.99"},
		{"4 inch bristles Paintbrush", "5.99"},
		{"Paint Roller 6 inch brush", "5.99"},
		{"Paint Roller Long Handle", "8.99"},
		{"Paint Roller Short Handle", "4.99"},
		{"Paint tray", "1.99"},
		{"Overall", "10.00"},
	}

	for i, thisSlice := range extras {
		fmt.Printf(" %d. %s , £%s per item \n", i+1 , thisSlice[0], thisSlice[2])
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
			itemInString := scanner.Text()
			//check if item is found in the available products
			if strings.Contains(itemInString, "#"){
				
			}else{

			}


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

	return extrasOrdered
}

func getQuantity(thisItem string) int {
	validInput := false
	var numberToOrder int
	for validInput {
		fmt.Printf("How many '%ss' would you like to order? : ")
		scanner.Scan()
		inString := scanner.Text()
		numberOrdered, err := stsrings.ParseInt(inString,10, 64)
		if err != nil{
			fmt.Println("Sorry, that wasn't a valid number to order. Try again...")
		} else {
			numberToOrder = numberOrdered
			validInput = true
		}
	}
	return numberOrdered
}
