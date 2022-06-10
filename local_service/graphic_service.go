package local_service

import (
	"DMood/domain"
	"errors"
	"github.com/golang/freetype/truetype"
	"github.com/shomali11/gridder"
	log "github.com/sirupsen/logrus"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"image/color"
	"strconv"
	"time"
)

const RowsNum =6


func setCollumnsNumber(slice []domain.Mood)(int,[]domain.Mood) {
	Collumns := len(slice) - 1
	if Collumns < 7 {
		lastDay := slice[Collumns].Date
		Collumns = 7
		dayCount := 1
		for len(slice) != 7 {
			slice = append(slice, domain.Mood{UserId: slice[0].UserId, Date: lastDay.Add(time.Hour * time.Duration(24*dayCount)), MoodRating: "1", Description: "", DayIdea: ""})
			dayCount++
		}
	}
	return Collumns,slice
}

func Getpng(slice []domain.Mood, path string)error  {
	if len(slice)==0{
		log.Error("Getpng: empty slice")
		return errors.New("Getpng: empty slice")
	}

	Rows := RowsNum
	Collumns,slice:=setCollumnsNumber(slice)
	imageConfig := gridder.ImageConfig{
		Width:  100 * Collumns,
		Height: Rows * 170,
		Name:   path,
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
	SetIdea(slice, grid,IdeaFont)
	 if err=grid.SavePNG();err!=nil{
	 	log.Error(err)
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
	first,_:=strconv.Atoi(slice[0].MoodRating)
	grid.DrawCircle(6-first, 0, circleConfig)
	for column := range slice {
		if column == len(slice)-1 {
			break
		}
		rating, _ := strconv.Atoi(slice[column].MoodRating)
		nextRating, _ := strconv.Atoi(slice[column+1].MoodRating)
		grid.DrawPath(6-rating, column, 6-nextRating, column+1, lineConfig)
		grid.DrawCircle(6-nextRating, column+1, circleConfig)
	}


}

func SetIdea(slice []domain.Mood, grid *gridder.Gridder, fontFace font.Face) {
	for column := range slice {
		if column == len(slice)-1 {
			break
		}
		rating, _ := strconv.Atoi(slice[column].MoodRating)
		if rating==3{
			if nextRating, _ := strconv.Atoi(slice[column+1].MoodRating); nextRating>=3{
				rating=5
			}else {
				rating=1
			}
		}
		grid.DrawString(rating, column, slice[column].DayIdea, fontFace, gridder.StringConfig{Rotate: 25})
	}
}