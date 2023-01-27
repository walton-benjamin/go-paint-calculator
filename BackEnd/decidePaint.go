package BackEnd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func DecidePaint() []string {
	paints := GetPaintsAvailable()
	fmt.Println("Please select a colour from the following options : \n ")
	maxLengthColourName := getLongestColourName()
	for paintColour, PaintData := range paints {
		thisPaintColourDisplay := pagifyPaintColour(paintColour, maxLengthColourName)
		fmt.Printf("Colour : %s  ; Price per litre £%.2f\n", thisPaintColourDisplay, (PaintData[2] / PaintData[0]))
	}
	//assume paint is not correct
	paintSelected, customColour := "nul", ""
	var pOut []string
	//while no valid paint is entered
	for paintSelected == "nul" {
		//attempt to pick a paint
		pOut = pickPaint(paints)
		paintSelected, customColour = pOut[0], pOut[1]
	}

	return []string{paintSelected, customColour}
}

func pagifyPaintColour(thisColour string, maxColourNameLength int) string {
	//hoping to make all the colours display the same length on the screen
	colourLength := len(thisColour)
	for colourLength < maxColourNameLength {
		t := thisColour + " "
		thisColour = t
		colourLength = len(thisColour)
	}
	return thisColour
}

func pickPaint(paintsAvailable map[string][]float64) []string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("\n Select a colour : ")
	scanner.Scan()
	inputPaint := scanner.Text()
	paintFound := false
	selectedPaint, customColour := "", "none"
	for paintColour := range paintsAvailable {
		if strings.EqualFold(inputPaint, paintColour) {
			//if paint is found in the map,
			selectedPaint = paintColour
			if selectedPaint == "Custom" {
				fmt.Printf("\nPlease describe the colour of paint you would like : ")
				scanner.Scan()
				customColour = scanner.Text()
			}
			paintFound = true
		}
	}

	if !paintFound {
		fmt.Printf("\nThat colour wasn't found. Please select one from the list...\n \n")
		return []string{"nul", "nul"}
	} else {
		return []string{selectedPaint, customColour}
	}
}

func getLongestColourName() int {
	paints := GetPaintsAvailable()
	longestNameLength := 0
	for thisPaintColour := range paints {
		if len(thisPaintColour) > longestNameLength {
			thisLength := len(thisPaintColour)
			longestNameLength = thisLength
		}
	}
	return longestNameLength
}

func GetPaintsAvailable() map[string][]float64 {

	paints := map[string][]float64{
		//"Name": {Volume in Litres, Price per tin(£), Surface Area Covered by one Tin}
		"Red":         {1.00, 7.99, 6.10},
		"Navy Blue":   {1.00, 8.99, 6.00},
		"Purple":      {1.00, 7.99, 5.50},
		"Cream":       {7.50, 50.00, 45.00},
		"White":       {1.00, 7.00, 6.00},
		"Taupe":       {1.00, 7.50, 6.15},
		"Light Brown": {1.00, 7.25, 6.00},
		"Custom":      {1.00, 10.50, 6.00},
	}
	return paints
}
