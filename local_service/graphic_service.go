package localservices

import (
	"DMood/domain"
	"fmt"
	"image/color"
	"strconv"
	"time"
	"github.com/golang/freetype/truetype"
	"github.com/shomali11/gridder"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)



func WhereBlock(yesterday, today, tomorrow int) int {

	return 0
}


func Getpng(slice []domain.Mood)error  {
	Rows := 6
	Collumns := len(slice)-1
	if Collumns < 7 {
		Collumns = 7
		dayCount:=1
		for len(slice) != 7 {
			slice = append(slice, domain.Mood{UserId: "", Date: time.Now().Add(time.Hour*time.Duration(24*dayCount)), MoodRating: "1", Description: "", DayIdea: ""})
			dayCount++
		}
	}

	imageConfig := gridder.ImageConfig{
		Width:  100 * Collumns,
		Height: Rows * 170,
		Name:   "C:\\Users\\Serj\\go\\src\\DMood\\myMood.png",
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
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	fontFace := truetype.NewFace(font, &truetype.Options{Size: 10, DPI: 150})
	IdeaFont:= truetype.NewFace(font, &truetype.Options{Size: 8, DPI: 150})
	lineConfig := gridder.PathConfig{Dashes: 15, StrokeWidth: 5}
	circleConfig := gridder.CircleConfig{Color: color.Gray{}, Radius: 15}

	grid.DrawString(0, 0, slice[0].Date.Format("2006-01-02"), fontFace, gridder.StringConfig{Rotate: 45})
	SetRating(slice, grid, lineConfig, circleConfig, IdeaFont)
	SetDate(slice, grid, fontFace)
	 if err=grid.SavePNG();err!=nil{
		 return err
	 }
	 return nil	
}

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
		grid.DrawString(rating, column, slice[column].DayIdea, fontFace, gridder.StringConfig{Rotate: 0})
		log.Info(fmt.Sprintf("||i:%d||rating:%d||rating+1:%d,", column, rating, rating1))
		grid.DrawCircle(6-rating1, column+1, circleConfig)
	}
}