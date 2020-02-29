package models

import (
	"cinema-admin/db"
	"cinema-admin/utils"
	"errors"
	"log"
	"time"
)

//Movie ...
type Movie struct {
	ID             string     `json:"id,omitempty"`
	CreatedAt      time.Time  `json:"created_at,omitempty"`
	UpdatedAt      time.Time  `json:"-"`
	DeletedAt      *time.Time `json:"-"`
	Name           string     `json:"name,omitempty"`
	Description    string     `json:"description,omitempty"`
	Image          string     `json:"image,omitempty"`
	Trailer        string     `json:"trailer,omitempty"`
	StartedAt      *time.Time `json:"started_at,omitempty"`
	Duration       int        `json:"duration,omitempty"`
	Rating         float32    `json:"rating,omitempty"`
	Views          int        `json:"views,omitempty"`
	Type           int        `json:"type,omitempty"`
	ManualPoint    int        `json:"manual_point,omitempty"`
	AlgorithmPoint int        `json:"algorithm_point,omitempty"`
}

// GetMovieByID ...
func GetMovieByID(id string) (Movie, error) {
	dbConn := db.GetDB()

	movie := Movie{}
	res := dbConn.Where("id = ?", id).Find(&movie)
	if res.Error != nil {
		return movie, errors.New("Data or data type is invalid")
	}
	return movie, nil
}

// GetMoviesByTheaterID ...
func GetMoviesByTheaterID(theaterID, date string) ([]Movie, error) {
	dbConn := db.GetDB()

	movieSessions := []MovieSessionInTheater{}
	movies := []Movie{}

	dateTime, err := utils.ConvertStringToTime(date)

	if err != nil {
		log.Println("err in time:", err)
		return movies, errors.New("Server is busy")
	}

	dateAfterTime := dateTime.Add(time.Hour * 24)

	res := dbConn.Where("theater_id = ?", theaterID).
		Select("DISTINCT(movie_id)").
		Where("start_time >= ? and end_time <= ?",
			dateTime, dateAfterTime).
		Find(&movieSessions)

	log.Println("movieSessions:", movieSessions)

	if res.Error != nil {
		return movies, errors.New("Data or data type is invalid")
	}

	for _, movieSession := range movieSessions {
		if movieSession.MovieID == nil {
			continue
		}
		movie, err := GetMovieByID(*movieSession.MovieID)
		if err != nil {
			continue
		}
		movies = append(movies, movie)
	}

	return movies, nil
}
