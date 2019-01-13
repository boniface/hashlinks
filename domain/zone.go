package domain

type Zone struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Zonestatus string `json:"zonestatus"`
	Logo   string `json:"logo"`
}

type Zones []Zone
