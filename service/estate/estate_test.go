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
		name    string
		id      string
		mock    func(mock *estate.MockRepositoryInterfaceMockRecorder)
		want    *estate.Estate
		wantErr error
	}{
		{
			name: "should return estate",
			id:   "1",
			mock: func(stub *estate.MockRepositoryInterfaceMockRecorder) {
				stub.FindByID(context.Background(), "1").Return(&estate.Estate{
					ID:        "1",
					Width:     10,
					Length:    10,
					CreatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				}, nil)
			},
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
			name: "should return error",
			id:   "1",
			mock: func(stub *estate.MockRepositoryInterfaceMockRecorder) {
				stub.FindByID(context.Background(), "1").Return(nil, errors.New("db: no rows in result set"))
			},
			want:    nil,
			wantErr: errors.New("db: no rows in result set"),
		},
	}

	for _, tt := range tests {
		mockRepo := estate.NewMockRepositoryInterface(gomock.NewController(t))
		tt.mock(mockRepo.EXPECT())

		estateService := NewEstateService(mockRepo)

		estate, err := estateService.RetrieveEstate(context.Background(), tt.id)
		if err != nil {
			assert.Equal(t, tt.wantErr, err)
		}
		assert.Equal(t, estate, tt.want)
	}
}

func TestCreateEstate(t *testing.T) {
	tests := []struct {
		name string
		args struct {
			width  int
			length int
		}
		mock    func(m *estate.MockRepositoryInterfaceMockRecorder)
		want    string
		wantErr error
	}{
		{
			name: "should create new estate",
			args: struct {
				width  int
				length int
			}{10, 20},
			mock: func(stub *estate.MockRepositoryInterfaceMockRecorder) {
				stub.Insert(context.Background(), &estate.Estate{Width: 10, Length: 20}).Return("1", nil)
			},
			want:    "1",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		mockRepo := estate.NewMockRepositoryInterface(gomock.NewController(t))
		tt.mock(mockRepo.EXPECT())

		estateService := NewEstateService(mockRepo)

		estate, err := estateService.CreateNewEstate(context.Background(), tt.args.width, tt.args.length)
		if err != nil {
			assert.Equal(t, tt.wantErr, err)
		}
		assert.Equal(t, tt.want, estate)
	}
}
