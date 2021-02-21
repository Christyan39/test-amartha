package database

import (
	"fmt"
	"reflect"
	"test/amartha/config"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func TestInit(t *testing.T) {
	mockErr := fmt.Errorf("Some Error")

	tempHandleSQLXConnect := handleSQLXConnect
	defer func() {
		handleSQLXConnect = tempHandleSQLXConnect
	}()

	type args struct {
		cfg config.Database
	}
	tests := []struct {
		name    string
		args    args
		want    *DB
		wantErr bool
		patch   func()
	}{
		{
			name: "when_ErrConnectDB_ExpectError",
			args: args{
				cfg: config.Database{},
			},
			want:    nil,
			wantErr: true,
			patch: func() {
				handleSQLXConnect = func(driverName, dataSourceName string) (*sqlx.DB, error) {
					return nil, mockErr
				}
			},
		},
		{
			name: "when_SuccessConnectDB_ExpectSuccess",
			args: args{
				cfg: config.Database{},
			},
			want: &DB{
				Master: &sqlx.DB{},
			},
			wantErr: false,
			patch: func() {
				handleSQLXConnect = func(driverName, dataSourceName string) (*sqlx.DB, error) {
					return &sqlx.DB{}, nil
				}
			},
		},
	}
	for _, tt := range tests {
		tt.patch()
		t.Run(tt.name, func(t *testing.T) {
			got, err := Init(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
