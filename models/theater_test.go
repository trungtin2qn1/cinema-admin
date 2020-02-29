package models

// func recordStats(db *sql.DB, userID, productID int64) (err error) {
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return
// 	}

// 	defer func() {
// 		switch err {
// 		case nil:
// 			err = tx.Commit()
// 		default:
// 			tx.Rollback()
// 		}
// 	}()

// 	if _, err = tx.Exec("UPDATE products SET views = views + 1"); err != nil {
// 		return
// 	}
// 	if _, err = tx.Exec("INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)", userID, productID); err != nil {
// 		return
// 	}
// 	return
// }

// func TestGetTheaterByID(t *testing.T) {

// 	db, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
// 	}
// 	defer db.Close()

// 	mock.ExpectBegin()
// 	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectCommit()

// 	// now we execute our method
// 	if err = recordStats(db, 2, 3); err != nil {
// 		t.Errorf("error was not expected while updating stats: %s", err)
// 	}

// 	// we make sure that all expectations were met
// 	if err := mock.ExpectationsWereMet(); err != nil {
// 		t.Errorf("there were unfulfilled expectations: %s", err)
// 	}

// 	type args struct {
// 		id string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    Theater
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.

// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := GetTheaterByID(tt.args.id)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetTheaterByID() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GetTheaterByID() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
