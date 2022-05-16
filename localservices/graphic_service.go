package localservices

import (
	"DMood/domain"
	"bytes"
	"encoding/base64"
	"fmt"
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/shomali11/gridder"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"

	//	"math/rand"
	"os"
	//	"strconv"
)

func SetDate(slice []domain.Mood, grid *gridder.Gridder, fontFace font.Face) {
	for column := range slice {
		grid.PaintCell(0, column, color.White)
		grid.DrawString(0, column, slice[column].Date.Format("2006-01-02"), fontFace, gridder.StringConfig{Rotate: 45})
	}
}

func SetRating(slice []domain.Mood, grid *gridder.Gridder, lineConfig gridder.PathConfig, circleConfig gridder.CircleConfig, fontFace font.Face) {
	for column := range slice {
		if column == len(slice)-2 {
			break
		}
		rating, _ := strconv.Atoi(slice[column].MoodRating)
		rating1, _ := strconv.Atoi(slice[column+1].MoodRating)
		grid.DrawPath(6-rating, column, 6-rating1, column+1, lineConfig)
		grid.DrawString(rating, column, slice[column].DayIdea, fontFace, gridder.StringConfig{Rotate: 45})
		log.Info(fmt.Sprintf("||i:%d||rating:%d||rating+1:%d,", column, rating, rating1))
		grid.DrawCircle(6-rating1, column+1, circleConfig)
	}

}

func WhereBlock(yesterday, today, tomorrow int) int {

	return 0
}
func Getpng(slice []domain.Mood) {
	GRIDEEER(slice)
}

func GRIDEEER(slice []domain.Mood) {

	Rows := 6
	Collumns := len(slice) - 1
	if Collumns < 7 {
		Collumns = 7
	}
	for len(slice) != 7 {
		slice = append(slice, domain.Mood{UserId: "", Date: time.Now(), MoodRating: "1", Description: "", DayIdea: ""})
	}
	imageConfig := gridder.ImageConfig{
		Width:  100 * Collumns,
		Height: Rows * 170,
		Name:   "C:/Users/Serj/go/src/DMood/localservices/myMood.png",
	}
	gridConfig := gridder.GridConfig{
		Rows:              Rows,
		Columns:           Collumns,
		MarginWidth:       10,
		LineStrokeWidth:   2,
		BorderStrokeWidth: 5,
		LineColor:         color.Black,
	}
	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}
	log.Error("SOZDAL GRID")
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	fontFace := truetype.NewFace(font, &truetype.Options{Size: 10, DPI: 150})
	lineConfig := gridder.PathConfig{Dashes: 15, StrokeWidth: 5}
	circleConfig := gridder.CircleConfig{Color: color.Gray{}, Radius: 15}
	log.Error("Sha SetRating")

	SetRating(slice, grid, lineConfig, circleConfig, fontFace)
	SetDate(slice, grid, fontFace)
	grid.SavePNG()
}

func PAINt() {
	imageConfig := gridder.ImageConfig{
		Width:  6000,
		Height: 1500,
		Name:   "example12.png",
	}
	gridConfig := gridder.GridConfig{
		Rows:              5,
		Columns:           31,
		LineStrokeWidth:   4,
		BorderStrokeWidth: 4,
		LineColor:         color.Gray{},
		BorderColor:       color.Gray{},
		BackgroundColor:   color.NRGBA{R: 220, G: 220, B: 220, A: 255},
	}

	grid, err := gridder.New(imageConfig, gridConfig)
	if err != nil {
		log.Fatal(err)
	}

	//blue := color.RGBA{B: 128, A: 255}

	// create a random chart
	for col := 0; col < gridConfig.Columns; col++ {
		height := rand.Intn(gridConfig.Rows - 1)
		for topRow := 0; topRow < height; topRow++ {
			grid.DrawCircle(gridConfig.Rows-topRow, col, gridder.CircleConfig{Radius: 70, Color: color.Black, StrokeWidth: 15, Stroke: true})
		}
	}

	// encode image as byte string
	bImage := new(bytes.Buffer)
	grid.EncodePNG(bImage)

	// convert to base64 string to support storing into database
	imageString := base64.StdEncoding.EncodeToString(bImage.Bytes())

	// convert back from string and save as binary image
	bDecodedImage, err := base64.StdEncoding.DecodeString(imageString)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(imageConfig.Name, bDecodedImage, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
