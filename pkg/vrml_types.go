package vrml

import (
	"time"
)

// Donations contains recent and alltime stats
type Donations struct {
	Recent  DonationData
	Alltime DonorsAlltime
}

type DonationData struct {
	Data []struct {
		CreatedAt  string `json:"created_at"`
		Currency   string `json:"currency"`
		Amount     string `json:"amount"`
		AmountText string `json:"amount_text"`
		Name       string `json:"name"`
		Message    string `json:"message"`
	}
}

type DonorsAlltime struct {
	Donors []struct {
		Name string `json:"name"`
		Rank int    `json:"rank"`
	} `json:"donors"`
}

// Sponsors holds the response fields for performing a GET against /Sponsors
type Sponsors []struct {
	Name        string `json:"sponsorName"`
	Url         string `json:"sponsorUrl"`
	Logo        string `json:"sponsorLogo"`
	Information string `json:"information"`
	Sort        int    `json:"sort"`
}

// MatchesUpcoming Gets the upcoming Matches. URL Params: ?game= Onward, EchoArena, Pavlov, Snapshot
//If the Authorization header with the Token is set, additional information will be retrieved.
type MatchesUpcoming []struct {
	MatchInfo          MatchesUpcomingProduction
	Week               uint16 `json:"week"`
	IsScheduled        bool   `json:"isScheduled"`
	IsSpecificDivision bool   `json:"isSpecificDivision"`
	IsChallenge        bool   `json:"isChallenge"`
	IsCup              bool   `json:"isCup"`
	VodUrl             string `json:"vodUrl"`
}

// MatchesUpcomingProduction Gets the upcoming Match information geared for production.
// GET {game}/Matches/UpcomingProduction/{matchID}
type MatchesUpcomingProduction struct {
	MatchID          string    `json:"matchID"`
	HomeTeam         struct{}  `json:"homeTeam"`
	AwayTeam         struct{}  `json:"awayTeam"`
	CastingInfo      struct{}  `json:"castingInfo"`
	HomeBetCount     uint8     `json:"homeBetCount"`
	AwayBetCount     uint8     `json:"awayBetCount"`
	DateScheduledUTC time.Time `json:"dateScheduledUTC"`
}

// CastingInfo The casting information for the Match
type CastingInfo struct {
	ChannelID         string `json:"channelID"`
	ChannelType       string `json:"channelType"`
	ChannelUrl        string `json:"channelUrl"`
	CasterId          string `json:"casterId"`
	Caster            string `json:"caster"`
	CoCaster          string `json:"coCaster"`
	Cameraman         string `json:"cameraman"`
	PostGameInterview string `json:"postGameInterview"`
}

type TeamObject struct {
	TeamID          string `json:"teamID"`
	TeamName        string `json:"teamName"`
	TeamLogo        string `json:"teamLogo"`
	RegionID        string `json:"regionID"`
	RegionName      string `json:"regionName"`
	DivisionName    string `json:"divisionName"`
	SubmittedScores bool   `json:"submittedScores"`
}
