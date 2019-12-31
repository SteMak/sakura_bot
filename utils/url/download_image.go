package url

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// GetImageByURL get image by URL and save it
func GetImageByURL(fileURL string, name string) {

	err := DownloadFile(filepath.Join("sakura_images", name+".png"), fileURL)
	if err != nil {
		fmt.Println("ERROR getting image by url:", err.Error())
	}
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
