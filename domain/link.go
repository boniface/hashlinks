package domain

import "time"

type Link struct {
	Zone          string `json:"zone"`
	Datepublished time.Time `json:"datepublished"`
	Linkhash      string `json:"linkhash"`
	Linkurl       string `json:"linkurl"`
	Linksite      string `json:"linksite"`
	Linktitle     string `json:"linksitle"`
	Linktype      string `json:"linktype"`
	Linksitecode  string `json:"linksitecode"`
}


type Links []Link
