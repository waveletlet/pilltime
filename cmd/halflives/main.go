package main

import (
	"fmt"
	"time"

	"gitlab.com/waveletlet/pilltime"
)

// Contrived usage example:
// how much caffeine from this morning is left by bedtime?
var (
	drug     = "caffeine"
	qty      = float64(300) // Super Mega Coffee
	units    = "mg"
	halflife = 6 * time.Hour
	ago      = 16 * time.Hour // You know, one of those normal 16 hour days
)

// Contrived usage example:
// how much drug is left after 2 halflives?
//var (
//	drug     = "drug"
//	qty      = float64(20)
//	units    = "mg"
//	halflife = 10 * time.Hour
//	ago      = 2 * halflife
//)

func main() {
	pill := pilltime.Dose{
		Dosage: pilltime.Dosage{
			Qty:   qty,
			Units: units,
		},
		Chemical: &pilltime.Chemical{
			Name:     drug,
			HalfLife: halflife,
		},
	}
	pill.ScheduledTime = time.Now().Add(-ago)
	fmt.Printf("%v %s left of %s.\n", pill.Left(time.Now()), pill.Dosage.Units, pill.Chemical.Name)
	fmt.Printf("%v half lives since taking it.\n", pill.Halflives(time.Now()))
}
