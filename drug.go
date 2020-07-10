package pilltime

import (
	"math"
	"time"
)

type Dose struct {
	Dosage        Dosage
	Chemical      *Chemical
	ScheduledTime time.Time
	// ScheduledTime = TakenTime until interactive reminder timer functional
	TakenTime time.Time
}

type Dosage struct {
	Qty   float64
	Units string
}

type Chemical struct {
	Name        string
	HalfLife    time.Duration
	ClearedTime time.Duration
	//EmptyStomachOK string //yes/no/doesn't matter
	// Don't need to implement until scheduling system differentiates between
	//  meal times and nonmeal times
}

func (chem *Chemical) WithClearedTime() {
	if chem.ClearedTime == 0 {
		chem.ClearedTime = 10 * chem.HalfLife
	}
}

func CumulativeDosage(doses []Dose, point time.Time) {
	accumulated := float64(0)
	chem := doses[0].Chemical
	chem.WithClearedTime()
	for _, dose := range doses {
		if point.After(dose.ScheduledTime) {
			if point.Add(-1 * point.Sub(dose.ScheduledTime)).Before(dose.ScheduledTime.Add(chem.ClearedTime)) {
				accumulated += dose.Left(point)
			}
		}
	}
}

func (d Dose) Left(now time.Time) float64 {
	return d.Dosage.Qty * math.Exp2(-d.Halflives(now))
}

func (d Dose) Halflives(now time.Time) float64 {
	if d.TakenTime.IsZero() {
		d.TakenTime = d.ScheduledTime
	}
	elapsed := now.Sub(d.TakenTime)
	return (elapsed.Hours() / d.Chemical.HalfLife.Hours())
}

func (d Dose) ZeroTime() time.Time {
	// when give dose is no longer in system
	return time.Now()
}
