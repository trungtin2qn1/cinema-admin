package models

import (
	"cinema-admin/db"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetTheaterByID(t *testing.T) {

	dbConn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbConn.Close()

	db.Init(dbConn)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "theaters" 
		WHERE "theaters"."deleted_at" IS NULL AND ((id = $1))`)).
		WithArgs("12").WillReturnRows(sqlmock.NewRows([]string{
		"id", "name",
		"description", "state", "city", "district", "ward", "street",
	}).AddRow("12", "name", "description",
		"state", "city", "district", "ward", "street"))

	wantTheater := Theater{
		ID:          "12",
		Name:        "name",
		Description: "description",
		State:       "state",
		City:        "city",
		District:    "district",
		Ward:        "ward",
		Street:      "street",
	}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    Theater
		wantErr bool
	}{
		{
			name: "Valid Input",
			args: args{
				id: "12",
			},
			want:    wantTheater,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTheaterByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTheaterByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTheaterByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
