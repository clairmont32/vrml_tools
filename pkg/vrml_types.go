package vrml

import (
	"time"
)

// Donations contains recent and alltime stats
type Donations struct {
	Recent  DonationData  `json:"recent"`
	Alltime DonorsAlltime `json:"alltime"`
}

type DonationData []struct {
	CreatedAt  time.Time `json:"created_at"`
	Currency   string    `json:"currency"`
	Amount     string    `json:"amount"`
	AmountText string    `json:"amount_text"`
	Name       string    `json:"name"`
	Message    string    `json:"message"`
}

type DonorsAlltime []struct {
	Name string `json:"name"`
	Rank string `json:"rank"`
}

// Sponsors holds the response fields for performing a GET against /Sponsors
type Sponsors []struct {
	Name        string `json:"sponsorName"`
	Url         string `json:"sponsorUrl"`
	Logo        string `json:"sponsorLogo"`
	Information string `json:"information"`
	Sort        int    `json:"sort"`
}
