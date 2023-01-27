package BackEnd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func QueryExtras() [][]string {

	fmt.Printf("\n\n Would you like any extras added to your order? \n   Items in stock : \n")
	extrasOrdered := [][]string{}
	extras := GetExtras()
	//displaying items available
	if len(extras) > 0 {
		for i, thisSlice := range extras {
			fmt.Printf(" %d. %s , £%s per item \n", i+1, thisSlice[0], thisSlice[1])
		}
	} else {
		fmt.Println("Error loading extras, returning with none ordered")
		return [][]string{{}}
	}
	//while, the user wants to order more items
	moreOrders := true
	scanner := bufio.NewScanner(os.Stdin)

	for moreOrders {
		fmt.Printf("Would you like to add any additional items to your order? [Y/N] :")
		thisItem := []string{}

		//get their decision to order more or not
		scanner.Scan()
		textIn := scanner.Text()

		//check user input
		if strings.ToUpper(textIn) == "Y" || strings.ToUpper(textIn) == "YES" {
			//if they want to order more items, query the item
			fmt.Println("\nPlease enter the full name of the item as it appears, or '#' Hashtag, and then the number in the list")
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
					fmt.Printf("Inspecting item number %s\n", itemInString[1:])
					//in : #2
					inString, err := strconv.ParseInt(itemInString[1:], 10, 64)
					fmt.Println(inString)
					//out : 2

					if err == nil {
						//int accepted
						thisItem = extras[(inString - 1)]
						fmt.Printf("Selected Item '%s' Successsfully \n", thisItem[0])
						validExtraName = true
					} else {
						fmt.Println("An error occurred. Attempt to recognise #x notation failed. Try again")
					}
				} else {

					//if they are (potentially) entering the full name...
					//loop over possible items
					for i, item := range extras {
						fmt.Println(i, item)
						//check if the item name, and input are identical
						if strings.EqualFold(itemInString, item[0]) {
							//if so, set the item being ordered to this item, tell the user, set flag
							thisItem = item
							fmt.Printf("Selected %s\n", thisItem[0])
							validExtraName = true
						}
					}
					if !validExtraName {
						fmt.Println("Sorry, that was not recognised. Please make sure you enter the name exactly as it appears and try again...")
					}
				}
			}
			//item found, thisItem[0]&[1] are filled with the selected item info.
			//now get the quantity to order with getQuantity(item)

			//var thisQuant []string = []string{fmt.Sprintf("%f")}
			fmt.Println(thisItem, thisItem[0])
			//append the result of getQuantity
			thisItem = append(thisItem, fmt.Sprintf("%.0f", getQuantity(thisItem[0])))
			extrasOrdered = append(extrasOrdered, thisItem)
			moreOrders = true

			//if the user does not want to order other items
		} else if strings.ToUpper(textIn) == "N" || strings.ToUpper(textIn) == "NO" {

			//stop looping
			moreOrders = false
		} else if strings.ToUpper(textIn) == "I" || strings.ToUpper(textIn) == "Item" || strings.ToUpper(textIn) == "Items" {
			displayExtras(extras)
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
		{"Overall set", "10.00"},
	}
}

func displayExtras(extras [][]string) {
	if len(extras) > 0 {
		for i, thisSlice := range extras {
			fmt.Printf(" %d. %s , £%s per item \n", i+1, thisSlice[0], thisSlice[1])
		}
	}
}

func getQuantity(thisItemName string) float64 {
	validInput := false
	scanner := bufio.NewScanner(os.Stdin)
	var numberToOrder float64
	for !validInput {
		fmt.Printf("How many '%ss' would you like to order? : ", thisItemName)
		scanner.Scan()
		inString := scanner.Text()
		numberOrdered, err := strconv.ParseFloat(inString, 64)
		if err != nil {
			fmt.Println("Sorry, that wasn't a valid number to order. Try again...")
		} else {
			numberToOrder = numberOrdered
			fmt.Printf("Adding %.0f x %ss to your order....\n \n", numberToOrder, thisItemName)
			validInput = true
		}
	}
	return numberToOrder
}
