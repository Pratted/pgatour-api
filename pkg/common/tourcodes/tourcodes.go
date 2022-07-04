package tourcodes

const (
	PgaTour             = "r"
	KornFerryTour       = "h"
	PgaTourChampions    = "s"
	PgaTourCanada       = "c"
	PgaTourLatinAmerica = "m"
)

type TourCode string

var (
	TourCodes = map[TourCode]string{
		PgaTour:             "PGA Tour",
		KornFerryTour:       "Korn Ferry Tour",
		PgaTourChampions:    "PGA Tour Champions",
		PgaTourCanada:       "PGA Tour Canada",
		PgaTourLatinAmerica: "PGA Tour Latino America",
	}
)

func IsValidTourCode(tourCode TourCode) bool {
	_, exists := TourCodes[tourCode]
	return exists
}

func IsValidTourCodeStr(tourCode string) bool {
	return IsValidTourCode(TourCode(tourCode))
}

func FromString(tourCode string) TourCode {
	return TourCode(tourCode)
}
