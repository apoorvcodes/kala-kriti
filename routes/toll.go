package routes

import (
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"strings"

	"github.com/gominima/minima"
)


type Toll struct {
	Weight string
	IsOverweight bool
	Timestamp string
	Base64 string
	TollName string
	NumberPlate string
}

func CreateNewToll(w string, b bool, t string, img string, n string) *Toll {
	return &Toll{
		Weight: w,
		IsOverweight: b,
		Timestamp: t,
		Base64: img,
		TollName: n,
	}
}



func TollRouter() *minima.Router {
	return minima.NewRouter().Get("/toll/:name/:bool/:read/:ocr", SaveTollHandler())
}


func BaseToImg(data string) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
    m, formatString, err := image.Decode(reader)
    if err != nil {
        log.Fatal(err)
    }
    bounds := m.Bounds()
    fmt.Println(bounds, formatString)

    //Encode from image format to writer
    pngFilename := "test.png"
    f, err := os.OpenFile(pngFilename, os.O_WRONLY|os.O_CREATE, 0777)
    if err != nil {
        log.Fatal(err)
        return
    }

    err = png.Encode(f, m)
    if err != nil {
        log.Fatal(err)
        return
    }
    fmt.Println("Png file", pngFilename, "created")

}




func deleteImg(name string) {

}

func OcrImg(name string) string {
	return ""
}