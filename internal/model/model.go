package model

type LocationResponse struct {
	Count           int
	NextPageURL     string
	PreviousPageURL string
	Results         []LocationResults
}

type LocationResults struct {
	Name string
	URL  string
}
