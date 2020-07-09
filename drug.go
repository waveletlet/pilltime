package pilltime

import (
	"time"
)

type Dose struct {
	Dosage        Dosage
	ScheduledTime time.Time
	// ScheduledTime = TakenTime until interactive reminder timer functional
	TakenTime time.Time
	ZeroTime  time.Time
}

type Dosage struct {
	Qty      float64
	QtyUnits string
	Chemical Chemical
}

type Chemical struct {
	HalfLife       time.Duration
	EmptyStomachOK string //yes/no/doesn't matter
	// Don't need to implement until scheduling system differentiates between
	//  meal times and nonmeal times
}
