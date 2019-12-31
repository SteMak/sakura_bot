package imagework

import (
	"fmt"
	"image"
	"path/filepath"
	"strings"

	"github.com/SteMak/sakura_bot/utils/strcont"

	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/segment"
	"github.com/anthonynsimon/bild/transform"
	"github.com/otiai10/gosseract"
)

// ConvertImage crop and convert image to left only captcha
func ConvertImage(messageTimeStamp string) {

	img, err := imgio.Open(filepath.Join("sakura_images", messageTimeStamp+".png"))
	if err != nil {
		fmt.Println("ERROR opening file", err.Error())
		return
	}

	img = transform.Crop(img, image.Rect(coordOfCode(img)))
	img = effect.Invert(img)
	img = effect.Grayscale(img)
	img = segment.Threshold(img, 128)

	err = imgio.Save(filepath.Join("sakura_images", "conv-"+messageTimeStamp+".jpg"), img, imgio.JPEGEncoder(100))
	if err != nil {
		fmt.Println("ERROR saving converted file", err.Error())
		return
	}
}

// ParseImage find codes in image
func ParseImage(messageTimeStamp string) (string, string) {

	var err error

	client := gosseract.NewClient()
	err = client.SetImage(filepath.Join("sakura_images", "conv-"+messageTimeStamp+".jpg"))
	if err != nil {
		fmt.Println("ERROR opening converted file", err.Error())
	}

	text, err := client.Text()
	if err != nil {
		fmt.Println("ERROR reading text", err.Error())
	}

	client.Close()

	codes := strings.Split(strcont.ReplaceBadSymbols(strcont.ClearStrange(text)), "/")
	if len(codes) != 2 {
		return text, ""
	}

	return codes[0], codes[1]
}
