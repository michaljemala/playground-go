package pkg

import "time"

type Advertiser struct {
	ID     string
	Name   string
	Active bool
}

type Campaign struct {
	ID           string
	Name         string
	ActiveFrom   time.Time
	ActiveTo     time.Time
	AdvertiserID string
}

type Flight struct {
	ID         string
	Name       string
	ActiveFrom time.Time
	ActiveTo   time.Time
	CampaignID string
}

type Coupon struct {
	ID         string
	Name       string
	ActiveFrom time.Time
	ActiveTo   time.Time
	FlightID   string
}
