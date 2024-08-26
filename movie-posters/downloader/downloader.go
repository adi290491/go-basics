package downloader

import (
	"fmt"
	"image/jpeg"
	"log"
	"movie-posters/omdb"
	"net/http"
	"os"
)

func DownloadPoster(movie *omdb.Movie) (bool, error) {

	filename, err := os.Create(movie.Title + ".jpg")

	if err != nil {
		log.Fatalf("")
	}

	defer filename.Close()

	resp, err := http.Get(movie.Poster)

	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("bad status: %s", resp.Status)
	}

	image, err := jpeg.Decode(resp.Body)

	if err != nil {
		return false, err
	}

	if err = jpeg.Encode(filename, image, nil); err != nil {
		return false, err
	}

	return true, nil
}
