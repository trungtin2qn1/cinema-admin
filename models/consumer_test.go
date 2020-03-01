package models

import (
	"cinema-admin/db"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetConsumerByID(t *testing.T) {

	dbConn, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer dbConn.Close()

	db.Init(dbConn)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "consumers" 
		WHERE "consumers"."deleted_at" IS NULL AND ((id = $1))`)).
		WithArgs("1").WillReturnRows(sqlmock.NewRows([]string{
		"id", "email", "name", "phone", "address",
	}).AddRow("1", "email", "name", "phone", "address"))

	consumer := Consumer{
		ID:      "1",
		Email:   "email",
		Name:    "name",
		Phone:   "phone",
		Address: "address",
	}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    Consumer
		wantErr bool
	}{
		{
			name: "Valid Input",
			args: args{
				id: "1",
			},
			want:    consumer,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetConsumerByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConsumerByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConsumerByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
