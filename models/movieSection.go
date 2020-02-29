package models

import (
	"cinema-admin/db"
	"cinema-admin/utils"
	"errors"
	"log"
	"time"
)

//MovieSessionInTheater ...
type MovieSessionInTheater struct {
	ID        string     `json:"id,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-"`
	StartTime *time.Time `json:"start_time,omitempty"`
	EndTime   *time.Time `json:"end_time,omitempty"`
	MovieID   *string    `json:"movie_id,omitempty"`
	TheaterID *string    `json:"theater_id,omitempty"`
}

// GetMovieSessionsByTheaterIDAndMovieID ...
func GetMovieSessionsByTheaterIDAndMovieID(theaterID,
	movieID, date string) ([]MovieSessionInTheater, error) {
	dbConn := db.GetDB()

	movieSessions := []MovieSessionInTheater{}

	dateTime, err := utils.ConvertStringToTime(date)

	if err != nil {
		go utils.LogErrToFile(err.Error())
		return movieSessions, errors.New("Server is busy")
	}

	dateAfterTime := dateTime.Add(time.Hour * 24)

	res := dbConn.Where("theater_id = ? and movie_id = ?",
		theaterID, movieID).
		Where("start_time >= ? and end_time <= ?",
			dateTime, dateAfterTime).
		Find(&movieSessions)

	log.Println("movieSessions:", movieSessions)
	if res.Error != nil {
		go utils.LogErrToFile(res.Error.Error())
		return movieSessions, errors.New("Data or data type is invalid")
	}
	return movieSessions, nil
}

// GetMovieSessionsByTheaterID ...
func GetMovieSessionsByTheaterID(theaterID,
	date string) ([]MovieSessionInTheater, error) {
	dbConn := db.GetDB()
	movieSessions := []MovieSessionInTheater{}

	dateTime, err := utils.ConvertStringToTime(date)

	if err != nil {
		go utils.LogErrToFile(err.Error())
		return movieSessions, errors.New("Server is busy")
	}

	dateAfterTime := dateTime.Add(time.Hour * 24)

	res := dbConn.Where("theater_id = ?", theaterID).
		Where("start_time >= ? and end_time <= ?",
			dateTime, dateAfterTime).
		Find(&movieSessions)
	if res.Error != nil {
		go utils.LogErrToFile(res.Error.Error())
		return movieSessions, errors.New("Data or data type is invalid")
	}
	return movieSessions, nil
}
