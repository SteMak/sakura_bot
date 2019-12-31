package image_work

import (
	"fmt"
	"path/filepath"
	"strings"
	"image"

	"github.com/SteMak/sakura_bot/utils/string_control"

	"github.com/otiai10/gosseract"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/anthonynsimon/bild/segment"
)

func ConvertImage(messageTimeStamp string) {

	img, err := imgio.Open(filepath.Join("sakura_images", messageTimeStamp + ".png"))
	if err != nil {
		fmt.Println("ERROR opening file", err.Error())
		return
	}

	img = transform.Crop(img, image.Rect(CoordOfCode(img)))
	img = effect.Invert(img)
	img = effect.Grayscale(img)
	img = segment.Threshold(img, 128)

	err = imgio.Save(filepath.Join("sakura_images", "conv-" + messageTimeStamp + ".jpg"), img, imgio.JPEGEncoder(100))
	if err != nil {
		fmt.Println("ERROR saving converted file", err.Error())
		return
	}
}

func ParseImage(messageTimeStamp string) (string, string) {

	var err error = nil

	client := gosseract.NewClient()
	err = client.SetImage(filepath.Join("sakura_images", "conv-" + messageTimeStamp + ".jpg"))
	if (err != nil) {
		fmt.Println("ERROR opening converted file", err.Error())
	}

	text, err := client.Text()
	if (err != nil)	{
		fmt.Println("ERROR reading text", err.Error())
	}

	client.Close()

	codes := strings.Split(string_control.ReplaceBadSymbols(string_control.ClearStrange(text)), "/")
	if len(codes) != 2 {
		return text, ""
	}

	return codes[0], codes[1]
}
