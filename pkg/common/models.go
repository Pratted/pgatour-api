// Package common provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package common

// Course defines model for Course.
type Course struct {
	City         string `json:"city"`
	Country      string `json:"country"`
	CourseName   string `json:"courseName"`
	CourseNumber string `json:"courseNumber"`
	IsHost       bool   `json:"isHost"`
	State        string `json:"state"`
}

// Leaderboard defines model for Leaderboard.
type Leaderboard struct {
	CurrentRoundNumber int    `json:"currentRoundNumber"`
	RoundStatus        string `json:"roundStatus"`
	Year               string `json:"year"`
}

// PgaTourScheduleV2 defines model for PgaTourScheduleV2.
type PgaTourScheduleV2 struct {
	TourSchedules []Schedule `json:"tourSchedules"`
	Year          string     `json:"year"`
}

// Schedule defines model for Schedule.
type Schedule struct {
	TourCode    string                `json:"tourCode"`
	Tournaments []ScheduledTournament `json:"tournaments"`
}

// ScheduledTournament defines model for ScheduledTournament.
type ScheduledTournament struct {
	Courses        []Course       `json:"courses"`
	TournamentDate TournamentDate `json:"tournamentDate"`
	TournamentId   string         `json:"tournamentId"`
	TournamentName TournamentName `json:"tournamentName"`
}

// Tournament defines model for Tournament.
type Tournament struct {
	Courses        []Course         `json:"courses"`
	Leaderboard    Leaderboard      `json:"leaderboard"`
	Status         TournamentStatus `json:"status"`
	TournamentDate TournamentDate   `json:"tournamentDate"`
	TournamentId   string           `json:"tournamentId"`
	TournamentName TournamentName   `json:"tournamentName"`
}

// TournamentDate defines model for TournamentDate.
type TournamentDate struct {
	EndDate    string `json:"endDate"`
	StartDate  string `json:"startDate"`
	WeekNumber string `json:"weekNumber"`
}

// TournamentName defines model for TournamentName.
type TournamentName struct {
	Long     string `json:"long"`
	Medium   string `json:"medium"`
	Official string `json:"official"`
	Short    string `json:"short"`
}

// TournamentStatus defines model for TournamentStatus.
type TournamentStatus struct {
	CourseNumber string `json:"courseNumber"`
	CurrentRound string `json:"currentRound"`
	LastUpdated  string `json:"lastUpdated"`
	RoundStatus  string `json:"roundStatus"`
	SeasonYear   string `json:"seasonYear"`
	TotalRounds  string `json:"totalRounds"`
	TourCode     string `json:"tourCode"`
	TournamentId string `json:"tournamentId"`
}

