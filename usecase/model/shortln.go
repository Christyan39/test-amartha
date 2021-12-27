package model

import "time"

type ShortlnRequest struct {
	Url       string `json:"url"`
	Shortcode string `json:"shortcode"`
	Count     int64
	CreatedAt time.Time
	LastSeen  time.Time
}

type ShortlnResponse struct {
	Shortcode string `json:"shortcode"`
}

type ShortlnStatus struct {
	Count     int64     `json:"redirectCount"`
	CreatedAt time.Time `json:"startDate"`
	LastSeen  time.Time `json:"lastSeenDate"`
}
