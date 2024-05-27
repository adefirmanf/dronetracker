package estate

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/SawitProRecruitment/UserService/repository/estate"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestRetrieveEstate(t *testing.T) {
	tests := []struct {
		name      string
		id        string
		returns   *estate.Estate
		returnErr error
		want      *estate.Estate
		wantErr   error
	}{
		{
			name: "should return estate",
			id:   "1",
			returns: &estate.Estate{
				ID:        "1",
				Width:     10,
				Length:    10,
				CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			returnErr: nil,
			want: &estate.Estate{
				ID:        "1",
				Width:     10,
				Length:    10,
				CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
			},
			wantErr: nil,
		},
		{
			name:      "should return error",
			id:        "1",
			returns:   nil,
			returnErr: errors.New("db: no rows in result set"),
			want:      nil,
			wantErr:   errors.New("db: no rows in result set"),
		},
	}

	for _, v := range tests {
		mockRepo := estate.NewMockRepositoryInterface(gomock.NewController(t))
		mockRepo.EXPECT().FindByID(context.Background(), v.id).Return(v.returns, nil)

		estateService := NewEstateService(mockRepo)

		estate, err := estateService.RetrieveEstate(context.Background(), v.id)
		if err != nil {
			assert.Equal(t, v.wantErr, err)
		}
		assert.Equal(t, estate, v.returns)
	}
}

func TestCreateEstate(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			width  int
			length int
		}
		returns   string
		returnErr error
		want      string
		wantErr   error
	}{
		{
			name: "should create new estate",
			args: struct {
				width  int
				length int
			}{10, 20},
			returns:   "1",
			returnErr: nil,
			want:      "1",
			wantErr:   nil,
		},
	}

	for _, v := range tests {
		mockRepo := estate.NewMockRepositoryInterface(gomock.NewController(t))
		mockRepo.EXPECT().Insert(context.Background(), &estate.Estate{Width: v.args.width, Length: v.args.length}).Return(v.returns, nil)

		estateService := NewEstateService(mockRepo)

		estate, err := estateService.CreateNewEstate(context.Background(), v.args.width, v.args.length)
		if err != nil {
			assert.Equal(t, v.wantErr, err)
		}
		assert.Equal(t, estate, v.returns)
	}
}
