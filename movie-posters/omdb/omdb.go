package omdb

const (
	omdb_api_key = "db9fc3f8"
	OMDBURL      = "http://www.omdbapi.com/"
)

type Movie struct {
	Title  string `json:"Title"`
	Poster string `json:"Poster"`
}
