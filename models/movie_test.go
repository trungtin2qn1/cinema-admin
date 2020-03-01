package models

import (
	"cinema-admin/db"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetMovieByID(t *testing.T) {

	dbConn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbConn.Close()

	db.Init(dbConn)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "movies" 
		WHERE "movies"."deleted_at" IS NULL AND ((id = $1))`)).
		WithArgs("1").WillReturnRows(sqlmock.NewRows([]string{
		"id", "name", "image", "trailer", "duration",
		"rating", "views", "type", "manual_point", "algorithm_point",
	}).AddRow("1", "name", "image", "trailer", 20,
		8.5, 25, 1, 6, 1))

	movie := Movie{
		ID:             "1",
		Name:           "name",
		Image:          "image",
		Trailer:        "trailer",
		Duration:       20,
		Rating:         8.5,
		Views:          25,
		Type:           1,
		ManualPoint:    6,
		AlgorithmPoint: 1,
	}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    Movie
		wantErr bool
	}{
		{
			name: "Valid Input",
			args: args{
				id: "1",
			},
			want:    movie,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMovieByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMovieByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMovieByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
