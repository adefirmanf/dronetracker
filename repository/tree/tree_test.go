package tree

import (
	"context"
	"errors"
	"regexp"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/SawitProRecruitment/UserService/repository"
	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInsertTree(t *testing.T) {
	id := uuid.New().String()
	returnedUUID := uuid.New().String()
	tests := []struct {
		name       string
		args       Tree
		beforeTest func(sqlmock.Sqlmock)
		want       string
		wantErr    error
	}{
		{
			name: "success insert tree",
			args: Tree{
				EstateID:    id,
				XCoordinate: 10,
				YCoordinate: 10,
				Height:      5,
			},
			beforeTest: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					regexp.QuoteMeta(
						"INSERT INTO tree (estate_id, x_coordinate, y_coordinate, height) VALUES ($1, $2, $3, $4) RETURNING id")).
					WithArgs(id, 10, 10, 5).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(returnedUUID))
			},
			want:    returnedUUID,
			wantErr: nil,
		},
		{
			name: "fail insert tree",
			args: Tree{
				EstateID:    id,
				XCoordinate: 10,
				YCoordinate: 10,
				Height:      5,
			},
			beforeTest: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					regexp.QuoteMeta(
						"INSERT INTO tree (estate_id, x_coordinate, y_coordinate, height) VALUES ($1, $2, $3, $4) RETURNING id")).
					WithArgs(id, 10, 10, 5).
					WillReturnError(errors.New("some error"))
			},
			want:    "",
			wantErr: errors.New("some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			r := NewRepository(&repository.Repository{Db: db})
			tt.beforeTest(mock)
			got, err := r.Insert(context.Background(), &tt.args)
			assert.Equal(t, got, tt.want)
			if err != nil {
				assert.Equal(t, err, tt.wantErr)
			}

		})
	}
}

func TestGetStats(t *testing.T) {
	id := uuid.New().String()
	tests := []struct {
		name       string
		args       string
		beforeTest func(sqlmock.Sqlmock)
		want       *TreeStats
		wantErr    error
	}{
		{
			name: "success get stats",
			args: id,
			beforeTest: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					regexp.QuoteMeta(
						"SELECT median(height), count(id), min(height), max(height) FROM tree WHERE estate_id = $1")).
					WithArgs(id).
					WillReturnRows(sqlmock.NewRows([]string{"median", "count", "min", "max"}).
						AddRow(5, 1, 5, 5))
			},
			want: &TreeStats{
				Count:  1,
				Max:    5,
				Min:    5,
				Median: 5,
			},
			wantErr: nil,
		},
		{
			name: "fail get stats",
			args: id,
			beforeTest: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					regexp.QuoteMeta(
						"SELECT median(height), count(id), min(height), max(height) FROM tree WHERE estate_id = $1")).
					WithArgs(id).
					WillReturnError(errors.New("some error"))
			},
			want:    nil,
			wantErr: errors.New("some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			r := NewRepository(&repository.Repository{Db: db})
			tt.beforeTest(mock)
			got, err := r.GetStats(context.Background(), tt.args)

			assert.Equal(t, got, tt.want)
			if err != nil {
				assert.Equal(t, err, tt.wantErr)
			}
		})
	}
}

func TestFindByEstateID(t *testing.T) {
	id := uuid.New().String()
	tests := []struct {
		name       string
		args       string
		beforeTest func(sqlmock.Sqlmock)
		want       []*Tree
		wantErr    error
	}{
		{
			name: "success find by estate id",
			args: id,
			beforeTest: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					regexp.QuoteMeta(
						"SELECT id, estate_id, x_coordinate, y_coordinate, height, created_at, updated_at FROM tree WHERE estate_id = $1")).
					WithArgs(id).
					WillReturnRows(sqlmock.NewRows([]string{"id", "estate_id", "x_coordinate", "y_coordinate", "height", "created_at", "updated_at"}).
						AddRow("1", id, 10, 10, 5, time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)).
						AddRow("2", id, 10, 10, 5, time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)))
			},
			want: []*Tree{
				{
					ID:          "1",
					EstateID:    id,
					XCoordinate: 10,
					YCoordinate: 10,
					Height:      5,
					CreatedAt:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				},
				{
					ID:          "2",
					EstateID:    id,
					XCoordinate: 10,
					YCoordinate: 10,
					Height:      5,
					CreatedAt:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: nil,
		},
		{
			name: "fail find by estate id",
			args: id,
			beforeTest: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					regexp.QuoteMeta(
						"SELECT id, estate_id, x_coordinate, y_coordinate, height, created_at, updated_at FROM tree WHERE estate_id = $1")).
					WithArgs(id).
					WillReturnError(errors.New("some error"))
			},
			want:    nil,
			wantErr: errors.New("some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			r := NewRepository(&repository.Repository{Db: db})
			tt.beforeTest(mock)
			got, err := r.FindByEstateID(context.Background(), tt.args)

			assert.Equal(t, got, tt.want)
			if err != nil {
				assert.Equal(t, err, tt.wantErr)
			}
		})
	}
}

func TestIsExistInTree(t *testing.T) {
	id := uuid.New().String()
	tests := []struct {
		name string
		args struct {
			estateID    string
			xCoordinate int
			yCoordinate int
		}
		beforeTest func(sqlmock.Sqlmock)
		want       bool
		wantErr    bool
	}{
		{
			name: "success find by estate id",
			args: struct {
				estateID    string
				xCoordinate int
				yCoordinate int
			}{
				estateID:    id,
				xCoordinate: 10,
				yCoordinate: 10,
			},
			beforeTest: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					regexp.QuoteMeta(
						"SELECT EXISTS(SELECT 1 FROM tree WHERE estate_id = $1 AND x_coordinate = $2 AND y_coordinate = $3)")).
					WithArgs(id, 10, 10).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "fail find by estate id",
			args: struct {
				estateID    string
				xCoordinate int
				yCoordinate int
			}{
				estateID:    id,
				xCoordinate: 10,
				yCoordinate: 10,
			},
			beforeTest: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery(
					regexp.QuoteMeta(
						"SELECT EXISTS(SELECT 1 FROM tree WHERE estate_id = $1 AND x_coordinate = $2 AND y_coordinate = $3)")).
					WithArgs(id, 10, 10).
					WillReturnError(errors.New("some error"))
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			r := NewRepository(&repository.Repository{Db: db})
			tt.beforeTest(mock)
			got, err := r.IsExistInEstate(context.Background(), tt.args.estateID, tt.args.xCoordinate, tt.args.yCoordinate)

			if (err != nil) != tt.wantErr {
				t.Errorf("IsExistInTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsExistInTree() got = %v, want %v", got, tt.want)
			}
		})
	}
}
