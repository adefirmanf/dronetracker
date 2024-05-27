package tree

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/SawitProRecruitment/UserService/repository/estate"
	mockEstate "github.com/SawitProRecruitment/UserService/repository/estate"
	"github.com/SawitProRecruitment/UserService/repository/tree"
	mockTree "github.com/SawitProRecruitment/UserService/repository/tree"
	"github.com/SawitProRecruitment/UserService/types"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateNewTree(t *testing.T) {
	estateID := uuid.New().String()
	treeID := uuid.New().String()

	tests := []struct {
		name    string
		payload struct {
			estateID    string
			xCoordinate int
			yCoordinate int
			height      int
		}
		mockTree   func(expect *mockTree.MockRepositoryInterfaceMockRecorder)
		mockEstate func(expect *mockEstate.MockRepositoryInterfaceMockRecorder)
		wantErr    error
		want       string
	}{
		{
			name: "success create new tree",
			payload: struct {
				estateID    string
				xCoordinate int
				yCoordinate int
				height      int
			}{
				estateID:    estateID,
				xCoordinate: 10,
				yCoordinate: 10,
				height:      5,
			},
			mockTree: func(stub *mockTree.MockRepositoryInterfaceMockRecorder) {
				stub.Insert(context.TODO(), &tree.Tree{EstateID: estateID, XCoordinate: 10, YCoordinate: 10, Height: 5, CreatedAt: time.Time{}, UpdatedAt: time.Time{}}).Return(treeID, nil)
				stub.IsExistInEstate(context.TODO(), estateID, 10, 10).Return(false, nil)
			},
			mockEstate: func(stub *mockEstate.MockRepositoryInterfaceMockRecorder) {
				stub.FindByID(context.TODO(), estateID).Return(&estate.Estate{ID: estateID, Width: 10, Length: 10, CreatedAt: time.Time{}, UpdatedAt: time.Time{}}, nil)
			},
			wantErr: nil,
			want:    treeID,
		},
		{
			name: "fail create new tree (estate not found)",
			payload: struct {
				estateID    string
				xCoordinate int
				yCoordinate int
				height      int
			}{
				estateID:    estateID,
				xCoordinate: 10,
				yCoordinate: 10,
				height:      5,
			},
			mockTree: func(stub *mockTree.MockRepositoryInterfaceMockRecorder) {},
			mockEstate: func(stub *mockEstate.MockRepositoryInterfaceMockRecorder) {
				stub.FindByID(context.TODO(), estateID).Return(nil, errors.New("sql: no rows in result set"))
			},
			wantErr: types.ErrorEstateNotFound,
			want:    "",
		},
		{
			name: "fail create new tree (tree out of bound)",
			payload: struct {
				estateID    string
				xCoordinate int
				yCoordinate int
				height      int
			}{
				estateID:    estateID,
				xCoordinate: 11,
				yCoordinate: 10,
				height:      5,
			},
			mockTree: func(stub *mockTree.MockRepositoryInterfaceMockRecorder) {},
			mockEstate: func(stub *mockEstate.MockRepositoryInterfaceMockRecorder) {
				stub.FindByID(context.TODO(), estateID).Return(&estate.Estate{ID: estateID, Width: 10, Length: 10, CreatedAt: time.Time{}, UpdatedAt: time.Time{}}, nil)
			},
			wantErr: types.ErorrTreeOutOfBound,
			want:    "",
		},
		{
			name: "fail create new tree (tree already planted)",
			payload: struct {
				estateID    string
				xCoordinate int
				yCoordinate int
				height      int
			}{
				estateID:    estateID,
				xCoordinate: 10,
				yCoordinate: 10,
				height:      5,
			},
			mockTree: func(stub *mockTree.MockRepositoryInterfaceMockRecorder) {
				stub.IsExistInEstate(context.TODO(), estateID, 10, 10).Return(true, nil)
			},
			mockEstate: func(stub *mockEstate.MockRepositoryInterfaceMockRecorder) {
				stub.FindByID(context.TODO(), estateID).Return(&estate.Estate{ID: estateID, Width: 10, Length: 10, CreatedAt: time.Time{}, UpdatedAt: time.Time{}}, nil)
			},
			wantErr: types.ErrorTreeAlreadyPlanted,
			want:    "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTreeRepo := tree.NewMockRepositoryInterface(gomock.NewController(t))
			mockEstateRepo := estate.NewMockRepositoryInterface(gomock.NewController(t))

			tt.mockTree(mockTreeRepo.EXPECT())
			tt.mockEstate(mockEstateRepo.EXPECT())

			treeService := NewTreeService(mockTreeRepo, mockEstateRepo)
			tree, err := treeService.CreateNewTree(context.TODO(), tt.payload.estateID, tt.payload.xCoordinate, tt.payload.yCoordinate, tt.payload.height)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			}
			assert.Equal(t, tt.want, tree)
		},
		)
	}
}

func TestRetrievesTreeByEstateID(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		mock    func(mock *mockTree.MockRepositoryInterfaceMockRecorder)
		want    []*tree.Tree
		wantErr error
	}{
		{
			name: "should return tree",
			id:   "1",
			mock: func(stub *mockTree.MockRepositoryInterfaceMockRecorder) {
				stub.FindByEstateID(context.TODO(), "1").Return([]*tree.Tree{
					{
						ID:          "1",
						EstateID:    "1",
						XCoordinate: 10,
						YCoordinate: 10,
						Height:      10,
						CreatedAt:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
						UpdatedAt:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
					},
				}, nil)
			},
			want: []*tree.Tree{
				{
					ID:          "1",
					EstateID:    "1",
					XCoordinate: 10,
					YCoordinate: 10,
					Height:      10,
					CreatedAt:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTreeRepo := tree.NewMockRepositoryInterface(gomock.NewController(t))
			tt.mock(mockTreeRepo.EXPECT())
			treeService := NewTreeService(mockTreeRepo, nil)
			tree, err := treeService.RetrievesByEstateID(context.TODO(), tt.id)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			}
			assert.Equal(t, tt.want, tree)
		},
		)
	}
}

func TestGetStats(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		mock    func(mock *mockTree.MockRepositoryInterfaceMockRecorder)
		want    *tree.TreeStats
		wantErr error
	}{
		{
			name: "should return stats",
			id:   "1",
			mock: func(stub *mockTree.MockRepositoryInterfaceMockRecorder) {
				stub.GetStats(context.TODO(), "1").Return(&tree.TreeStats{
					Count:  1,
					Max:    5,
					Min:    5,
					Median: 5,
				}, nil)
			},
			want: &tree.TreeStats{
				Count:  1,
				Max:    5,
				Min:    5,
				Median: 5,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTreeRepo := tree.NewMockRepositoryInterface(gomock.NewController(t))
			tt.mock(mockTreeRepo.EXPECT())

			treeService := NewTreeService(mockTreeRepo, nil)
			stats, err := treeService.GetStats(context.TODO(), tt.id)
			if err != nil {
				assert.Equal(t, tt.wantErr, err)
			}
			assert.Equal(t, tt.want, stats)
		},
		)
	}
}
