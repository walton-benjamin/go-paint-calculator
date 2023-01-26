package BackEnd

import (
	"fmt"
	"math"
)

func CalculateTinsNeeded(area float64, paint string) []float64 {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("An error occurred.......")
		}
	}()

	paints := GetPaintsAvailable()
	paintMetaData := paints[paint]
	var areaPerTinOfThisPaint float64 = paintMetaData[2]
	tinsNeeded := math.Ceil(area / areaPerTinOfThisPaint) //int
	costOfTins := tinsNeeded * paintMetaData[1]           //float64

	fmt.Printf("\nCovering %.2fm^2 of wall. \nEach tin of the selected paint will cover %.2fm^2 of wall, therefore...", area, areaPerTinOfThisPaint)

	return []float64{tinsNeeded, costOfTins}
}
