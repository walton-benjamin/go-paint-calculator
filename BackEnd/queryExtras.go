package BackEnd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func QueryExtras() [][]string {

	fmt.Printf("\n\n Would you like any extras added to your order? Items in stock : ")
	extrasOrdered := [][]string{{"0"}}
	extras := GetExtras()
	//displaying items available
	for i, thisSlice := range extras {
		fmt.Printf(" %d. %s , £%s per item \n", i+1, thisSlice[0], thisSlice[2])
	}

	//while, the user wants to order more items
	moreOrders := true
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Would you like to add any items to your order? [Y/N] :")
	for moreOrders {
		thisItem := []string{}
		//get their decision to order more or not
		scanner.Scan()
		textIn := scanner.Text()
		//check user input
		if strings.ToUpper(textIn) == "Y" || strings.ToUpper(textIn) == "YES" {
			//if they want to order more items, query the item
			fmt.Println("Please enter the full name of the item as it appears, or '#' Hashtag, and then the number in the list")
			fmt.Println("If this was a mistake, type 'x', 'c' or 'cancel'.")

			//loop until a valid extra item name is entered
			validExtraName := false
			for !validExtraName {
				//ask for a new attempt to enter a correct name
				fmt.Println("Which item would you like to order? : ")
				scanner.Scan()
				itemInString := scanner.Text()
				//check if item name entered is found in the available products, through one of two ways
				if strings.Contains(itemInString, "#") {
					//if they are using #x notation correctly...
					fmt.Printf("Inspecting item number %s", itemInString[1:len(itemInString)])
				} else {
					//if they are (potentially) entering the full name...
					for i, item := range extras {
						fmt.Println(i, item)
					}
				}
			}
			//TC Which items
			quantity := getQuantity(thisItem[0])

			thisItem[2] = quantity
			extrasOrdered = append(extrasOrdered, []string{itemOrdered, priceEach, quantity})

		} else if strings.ToUpper(textIn) == "N" || strings.ToUpper(textIn) == "NO" {
			//if they don't want to order more items
			//stop looping
			moreOrders = false
		} else {
			//unaccepted inputs
			fmt.Printf("Sorry, that was not recognised as 'Y' or 'N'. Try again : ")
		}
	}

	return extrasOrdered
}
func GetExtras() [][]string {
	return [][]string{
		{"2 inch bristles Paintbrush", "3.99"},
		{"4 inch bristles Paintbrush", "5.99"},
		{"Paint Roller 6 inch brush", "5.99"},
		{"Paint Roller Long Handle", "8.99"},
		{"Paint Roller Short Handle", "4.99"},
		{"Paint tray", "1.99"},
		{"Overall", "10.00"},
	}
}

func getQuantity(thisItem string) float64 {
	validInput := false
	scanner := bufio.NewScanner(os.Stdin)
	var numberToOrder float64
	for validInput {
		fmt.Printf("How many '%ss' would you like to order? : ")
		scanner.Scan()
		inString := scanner.Text()
		numberOrdered, err := strconv.ParseFloat(inString, 64)
		if err != nil {
			fmt.Println("Sorry, that wasn't a valid number to order. Try again...")
		} else {
			numberToOrder = numberOrdered
			validInput = true
		}
	}
	return numberToOrder
}
