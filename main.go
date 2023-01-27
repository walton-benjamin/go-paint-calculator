package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	b "example.com/go-paint-calculator/BackEnd"
)

/*
TODO:
prints an automated invoice/receipt.
accept multiple rooms, add all to receipts
*/

func main() {
	fmt.Printf("\nWelcome to the Wall Painting Calculator \n\n")
	//run first room
	//runProgram()
	fmt.Println(b.QueryExtras())

	//does the user want to add a second room?

}

func runProgram() {

	//get the number of walls to paint
	//get the total surface area of wall to be painted

	fmt.Printf("Please enter a name for this room : ")
	/*
		scanner = bufio.NewScanner(os.stdin)
		scanner.Scan()
		input := scanner.Text()
	*/
	var roomWalls = []float64{0.0}
	for roomWalls[0] == 0.0 {
		roomWalls = b.WallAreaCalculator()
	}

	wallAreasToPaint := wallAreaTotal(roomWalls)

	//fmt.Println(wallAreasToPaint)

	fmt.Printf("\nTotal : %.2f m^2 of Wall to paint.\n \n \n", wallAreasToPaint)

	//decide on the type of paint.

	paintOut := b.DecidePaint()
	paint, customPaint := paintOut[0], paintOut[1]

	fmt.Printf("Paint selected is : %s\n", paint)
	if customPaint != "none" {
		fmt.Printf(" \n  with custom colour being : %s \n", customPaint)
	}

	//find the number of tins of paint needed for the walls
	//calculate cost.
	paintToOrder := b.CalculateTinsNeeded(wallAreasToPaint, paint)
	tinsNeeded, costOfTins := paintToOrder[0], paintToOrder[1]
	fmt.Printf("\n\n Tins Needed : %.0f,Total Price of Paint : £%.2f \n", tinsNeeded, costOfTins)

	//query extras

	extras := b.QueryExtras()

	//roomWalls - []float64 of all walls surface areas
	//paint - colour selected,
	//     optional customPaint for a custom colour description
	//thisPaintMetaData, the can-specific details of the chosen paint
	//paintToOrder - []float64 of the tins of paint needed, and cost of those tins.
	//extras - [][]string of optional extras added {product, quantity, price per}
	thisPaintMetaData := b.GetPaintsAvailable()[paint]

	printReceipt(roomWalls, paint, customPaint, thisPaintMetaData, paintToOrder, extras)

}
func wallAreaTotal(walls []float64) float64 {
	totalArea := 0.0
	for _, wall := range walls {
		totalArea += wall
	}
	return float64(totalArea)
}

func pagifyString(inputWord string) string {
	//Will output an input string, with trailing spaces to ensure a target width
	inputLength := len(inputWord)
	var targetLength int = 59
	for inputLength < targetLength {
		t := inputWord + " "
		inputWord = t
		inputLength = len(inputWord)
	}

	if strings.Contains(inputWord, "£") {
		numberOfDecimals := strings.Count(inputWord, "£")
		for i := 0; i < numberOfDecimals; i++ {
			t := inputWord + " "
			inputWord = t
		}
	}
	inputWord = fmt.Sprintf("%s|", inputWord)
	return inputWord
}
func pagifyStringRight(inputWord string) string {
	//Will output an input string, with trailing spaces to ensure a target width
	inputLength := len(inputWord)
	var targetLength int = 57

	for inputLength < targetLength {
		t := " " + inputWord
		inputWord = t
		inputLength = len(inputWord)
	}

	if strings.Contains(inputWord, "£") {
		numberOfDecimals := strings.Count(inputWord, "£")
		for i := 0; i < numberOfDecimals; i++ {
			t := " " + inputWord
			inputWord = t
		}
	}
	inputWord = fmt.Sprintf("|%s |", inputWord)
	return inputWord
}

func printReceipt(roomWalls []float64, chosenPaint string, customPaint string, paintMetaData []float64, tinsNeeded []float64, extras [][]string) {
	//receipt width : 60 character including borders
	//paintMetaData {Volume in Litres, Price per tin(£), Surface Area Covered by one Tin}
	subTotal := tinsNeeded[1]
	fmt.Printf(" ++------------------------------------------------------++ \n")
	fmt.Printf("++                                                        ++\n")
	fmt.Printf("|                     Paint Calculator                     |\n")
	fmt.Printf("|                                                          |\n")
	fmt.Printf("+----------------------------------------------------------+\n")
	fmt.Printf("|                                                          |\n")
	fmt.Printf("|   Rooms entered:                                         |\n")
	fmt.Printf("|                                                          |\n")
	for i := 0; i < len(roomWalls); i++ {
		thisWall := fmt.Sprintf("|     Wall %d : %.2fm^2 ", i+1, roomWalls[i])
		thisWallString := pagifyString(thisWall)
		fmt.Println(thisWallString)
	}
	fmt.Printf("|                                                          |\n")
	totalAreaString := fmt.Sprintf("|   Total Surface Area to paint : %.2fm^2 ", wallAreaTotal(roomWalls))
	fmt.Println(pagifyString(totalAreaString))
	fmt.Printf("|                                                          |\n")
	fmt.Printf("+----------------------------------------------------------+\n")
	fmt.Printf("|                                                          |\n")
	//paint selected
	if customPaint != "none" {
		customPaintReceiptString := fmt.Sprintf("|   Paint Colour Selected: %s ", chosenPaint)
		customPaintReceiptString = pagifyString(customPaintReceiptString)
		customPaintReceiptString_custom := fmt.Sprintf("|      with custom colour being : '%s'", customPaint)
		customPaintReceiptString_custom = pagifyString(customPaintReceiptString_custom)
		fmt.Println(customPaintReceiptString)
		fmt.Println(customPaintReceiptString_custom)
	} else {
		customPaintReceiptString := fmt.Sprintf("|   Paint Colour Selected: %s ", chosenPaint)
		//customPaintReceiptString = pagifyString(customPaintReceiptString, 59)
		fmt.Println(pagifyString(customPaintReceiptString))
	}
	fmt.Printf("|                                                          |\n")

	// this type of paint comes in X volume tins, costing Y each.
	fmt.Println(pagifyString(fmt.Sprintf("|   This type of paint comes in %.2fL tins,", paintMetaData[0])))
	fmt.Println(pagifyString(fmt.Sprintf("|      Costing £%.2f each.", paintMetaData[1])))
	fmt.Printf("|                                                          |\n")

	//Each tin lists it will cover Z m^2
	fmt.Println(pagifyString(fmt.Sprintf("|   This tin lists it will cover %.2fm^2 per tin ", paintMetaData[2])))
	// So, to cover totalSA, you will need A tins
	fmt.Println(pagifyString(fmt.Sprintf("|   So, to cover %.2fm^2 of wall, you will need... ", wallAreaTotal(roomWalls))))
	fmt.Printf("|                                                          |\n")

	// A tins = £xxx.xx
	fmt.Println(pagifyString(fmt.Sprintf("|   %s Tins of %s paint ", fmt.Sprintf("%.0f", math.Ceil(tinsNeeded[0])), chosenPaint)))
	fmt.Printf("|                                                          |\n")
	fmt.Println(pagifyStringRight(fmt.Sprintf("   Costing = £%.2f   ", tinsNeeded[1])))
	fmt.Printf("|                                                          |\n")
	fmt.Printf("+----------------------------------------------------------+\n")
	fmt.Printf("|                                                          |\n")

	//extras :
	fmt.Printf("|   Extras :                                               |\n")
	//product, quantity, price per
	if len(extras) != 0 {
		for _, thisSlice := range extras {
			//fmt.Println(thisSlice)
			fmt.Println(pagifyString(fmt.Sprintf("|   %s x %s        ", thisSlice[1], thisSlice[0])))
			thisItemCost, _ := strconv.ParseFloat(thisSlice[2], 64)
			thisQuantity, _ := strconv.ParseFloat(thisSlice[1], 64)
			thisTotal := thisItemCost * thisQuantity
			fmt.Println(pagifyString(fmt.Sprintf("|       %.0f * %.2f ", thisQuantity, thisItemCost)))
			fmt.Println(pagifyStringRight(fmt.Sprintf("  = £ %.2f   ", thisTotal)))
			subTotal += thisTotal
			fmt.Printf("|                                                          |\n")
		}
	} else {
		fmt.Printf("|   	None                                               |\n")
	}
	fmt.Printf("|                                                          |\n")
	fmt.Printf("+----------------------------------------------------------+\n")
	fmt.Printf("|                                                          |\n")
	fmt.Printf("|   Total :                                                |\n")
	fmt.Println(pagifyStringRight(fmt.Sprintf("       £%.2f   ", subTotal)))
	fmt.Printf("|                                                          |\n")
	fmt.Printf("|                                                          |\n")
	fmt.Printf("|                Thank you for your order!                 |\n")
	fmt.Printf("|                                                          |\n")
	fmt.Printf("++                                                        ++\n")
	fmt.Printf(" ++------------------------------------------------------++ \n")
}
