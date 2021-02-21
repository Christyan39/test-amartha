package database

import (
	"fmt"
	"reflect"
	"regexp"
	"test/amartha/usecase/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestDB_GetShortenByCode(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockDb := sqlx.NewDb(db, "sqlmock")
	defer func() {
		db.Close()
	}()

	query := `
	SELECT 
		code,
		url,
		created_at,
		last_seen_at, 
		count
	FROM 
		shorten
	WHERE
		code = ?
`

	column := []string{
		"code",
		"url",
		"created_at",
		"last_seen_at",
		"count"}

	type fields struct {
		Master *sqlx.DB
	}
	type args struct {
		code string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantShorten *model.ShortlnRequest
		patch       func()
	}{
		{
			name:        "Err Query",
			fields:      fields{},
			args:        args{},
			wantShorten: nil,
			patch: func() {
				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("").WillReturnError(fmt.Errorf("Error"))
			},
		},
		{
			name:        "Err Scan",
			fields:      fields{},
			args:        args{},
			wantShorten: nil,
			patch: func() {
				results := sqlmock.NewRows(column)
				results.AddRow(
					"",
					"",
					time.Time{},
					time.Time{},
					"",
				)
				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("").WillReturnRows(results)
			},
		},
		{
			name:        "Success Query",
			fields:      fields{},
			args:        args{},
			wantShorten: &model.ShortlnRequest{},
			patch: func() {
				results := sqlmock.NewRows(column)
				results.AddRow(
					"",
					"",
					time.Time{},
					time.Time{},
					0,
				)
				mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs("").WillReturnRows(results)
			},
		},
	}
	for _, tt := range tests {
		tt.patch()
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				Master: mockDb,
			}
			if gotShorten := db.GetShortenByCode(tt.args.code); !reflect.DeepEqual(gotShorten, tt.wantShorten) {
				t.Errorf("DB.GetShortenByCode() = %v, want %v", gotShorten, tt.wantShorten)
			}
		})
	}
}

func TestDB_CreateShortenCode(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockDb := sqlx.NewDb(db, "sqlmock")
	defer func() {
		db.Close()
	}()

	type fields struct {
		Master *sqlx.DB
	}
	type args struct {
		shorten *model.ShortlnRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		patch   func()
	}{
		{
			name:   "Err Query",
			fields: fields{},
			args: args{
				shorten: &model.ShortlnRequest{},
			},
			wantErr: true,
			patch: func() {
				mock.ExpectExec("INSERT").
					WithArgs(
						"",
						"",
					).WillReturnError(fmt.Errorf("Some Error"))
			},
		},
		{
			name:   "Success Query",
			fields: fields{},
			args: args{
				shorten: &model.ShortlnRequest{},
			},
			wantErr: false,
			patch: func() {
				mock.ExpectExec("INSERT").
					WithArgs(
						"",
						"",
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
	}
	for _, tt := range tests {
		tt.patch()
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				Master: mockDb,
			}
			if err := db.CreateShortenCode(tt.args.shorten); (err != nil) != tt.wantErr {
				t.Errorf("DB.CreateShortenCode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDB_CountVisitingURL(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockDb := sqlx.NewDb(db, "sqlmock")
	defer func() {
		db.Close()
	}()
	type fields struct {
		Master *sqlx.DB
	}
	type args struct {
		code string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		patch   func()
	}{
		{
			name:    "err query",
			fields:  fields{},
			args:    args{},
			wantErr: true,
			patch: func() {
				mock.ExpectExec("UPDATE").
					WithArgs(
						"",
					).WillReturnError(fmt.Errorf("Some Error"))
			},
		},
		{
			name:    "success query",
			fields:  fields{},
			args:    args{},
			wantErr: false,
			patch: func() {
				mock.ExpectExec("UPDATE").
					WithArgs(
						"",
					).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
	}
	for _, tt := range tests {
		tt.patch()
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				Master: mockDb,
			}
			if err := db.CountVisitingURL(tt.args.code); (err != nil) != tt.wantErr {
				t.Errorf("DB.CountVisitingURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
