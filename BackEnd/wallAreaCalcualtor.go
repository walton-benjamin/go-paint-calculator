package BackEnd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func WallAreaCalculator() []float64 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("An error occurred....Restarting...\n \n \n")
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	totalWallsValidFlag := false
	var numWalls int64

	if !totalWallsValidFlag {

		fmt.Println("How many walls would you like to paint? ")
		fmt.Println("If you know the total surface area of all walls, Please Enter '1'. ")
		fmt.Print("Please enter the number of walls you would like to paint : ")
		scanner.Scan()
		wallsIn := scanner.Text()
		n, allWallsErr := strconv.ParseInt(wallsIn, 10, 64)
		numWalls = n

		if allWallsErr != nil {
			panic(fmt.Sprintf("'%d' was not recognised as an integer. Try again....\n", n))
		} else {
			//fmt.Printf("will work with %d walls", numWalls)
			totalWallsValidFlag = true
		}
	}
	//numWalls, int64 of number of walls to paint

	//ask if they want to use SA for each wall, or dimensions of each wall
	fmt.Printf("\nHow would you like to enter your wall sizes? By Height & Width, or total Surface Area?\n")
	fmt.Printf("Please enter 'D' to use Height and Width\n")
	fmt.Printf("Please enter 'S' to use Total Surface Area\n")
	enteranceMethodSelected := false
	var walls []float64
	for !enteranceMethodSelected {
		fmt.Printf("\n Enter here : ")
		scanner.Scan()
		enteranceMethodIn := scanner.Text()
		switch strings.ToLower(enteranceMethodIn) {
		case "d":
			walls = enterWallDimentions(numWalls)
			enteranceMethodSelected = true
		case "s":
			walls = enterWallSurfaceArea(numWalls)
			enteranceMethodSelected = true
		case "x":
			os.Exit(0)
		default:
			fmt.Printf("Invalid input, try again, or enter 'x' to cancel")
		}
	}
	return walls
}

func enterWallSurfaceArea(numWalls int64) []float64 {
	scanner := bufio.NewScanner(os.Stdin)
	var walls []float64 // size of walls

	fmt.Printf("\nNow entering data for %d wall surface areas....\n ", numWalls)
	i := int64(1)
	for i <= numWalls {
		fmt.Printf("Please enter Surface Area of wall number %d  you would like to paint (Meters^2) : ", i)
		scanner.Scan()
		t := scanner.Text()
		thisWall, wallErr := strconv.ParseFloat(t, 64)

		if wallErr != nil {
			fmt.Printf("'%f' is not recognised as a number.", thisWall)
		} else {
			walls = append(walls, thisWall)
			i++
		}
	}
	return walls
}

func enterWallDimentions(numWalls int64) []float64 {
	scanner := bufio.NewScanner(os.Stdin)
	var walls []float64 // size of walls

	fmt.Printf("\nNow entering data for %d wall dimensions.... \n", numWalls)
	fmt.Printf("Please note, we calculate surface areas for each wall entered automatically\n")

	i := int64(1)
	for i <= numWalls {
		fmt.Printf("\nPlease enter Height of wall number %d you would like to paint (Meters) : ", i)
		scanner.Scan()
		t := scanner.Text()
		thisWallHeight, wallHeightErr := strconv.ParseFloat(t, 64)

		fmt.Printf("Please enter Width of wall number %d you would like to paint (Meters) : ", i)
		scanner.Scan()
		t = scanner.Text()
		thisWallWidth, wallWidthErr := strconv.ParseFloat(t, 64)

		if wallHeightErr != nil || wallWidthErr != nil {
			fmt.Printf("'%f' is not recognised as a number.", thisWallHeight)
		} else {
			thisSA := thisWallHeight * thisWallWidth
			walls = append(walls, thisSA)
			i++
		}
	}
	return walls
}
