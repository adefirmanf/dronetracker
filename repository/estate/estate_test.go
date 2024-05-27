package estate

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/stretchr/testify/assert"
)

func TestFindByID(t *testing.T) {
	tests := []struct {
		name       string
		args       string
		beforeTest func(sqlmock.Sqlmock)
		want       *Estate
		wantErr    error
	}{
		{
			name: "success find by id",
			args: "1",
			beforeTest: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					regexp.QuoteMeta(
						"SELECT id, width, length, created_at, updated_at FROM estate WHERE id = $1")).
					WithArgs("1").
					WillReturnRows(sqlmock.NewRows([]string{"id", "width", "length", "created_at", "updated_at"}).
						AddRow("1", 10, 10, time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)))
			},
			want: &Estate{
				ID:        "1",
				Width:     10,
				Length:    10,
				CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name: "fail find by id",
			args: "1",
			beforeTest: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					regexp.QuoteMeta(
						"SELECT id, width, length, created_at, updated_at FROM estate WHERE id = $1")).
					WithArgs("1").
					WillReturnError(errors.New("sql: no rows in result set"))
			},
			want:    nil,
			wantErr: errors.New("sql: no rows in result set"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			tt.beforeTest(mock)
			r := NewRepository(&repository.Repository{Db: db})

			got, err := r.FindByID(context.Background(), tt.args)
			assert.Equal(t, got, tt.want)
			if err != nil {
				assert.Equal(t, err, tt.wantErr)
			}
		})
	}
}

func TestInsert(t *testing.T) {

	tests := []struct {
		name       string
		args       *Estate
		beforeTest func(sqlmock.Sqlmock)
		want       string
		wantErr    error
	}{
		{
			name: "success insert",
			args: &Estate{
				Width:  10,
				Length: 10,
			},
			beforeTest: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					regexp.QuoteMeta(
						"INSERT INTO estate (width, length) VALUES ($1, $2) RETURNING id")).
					WithArgs(10, 10).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
			},
			want:    "1",
			wantErr: nil,
		},
		{
			name: "fail insert",
			args: &Estate{
				Width:  10,
				Length: 10,
			},
			beforeTest: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					regexp.QuoteMeta(
						"INSERT INTO estate (width, length) VALUES ($1, $2) RETURNING id")).
					WithArgs(10, 10).
					WillReturnError(errors.New("sql: some errors"))
			},
			want:    "",
			wantErr: errors.New("sql: some errors"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()
			tt.beforeTest(mock)
			r := NewRepository(&repository.Repository{Db: db})

			got, err := r.Insert(context.Background(), tt.args)
			assert.Equal(t, got, tt.want)
			if err != nil {
				assert.Equal(t, err, tt.wantErr)
			}
		})
	}
}
