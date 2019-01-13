package domain

type Feed struct {
	Zone       string `json:"zone"`
	Sitecode   string `json:"sitecode"`
	Id         string `json:"id"`
	Feedfilter string `json:"feedfilter"`
	Feedlink   string `json:"feedlink"`
	Feedtype   string `json:"feedtype"`
}

type Feeds []Feed

