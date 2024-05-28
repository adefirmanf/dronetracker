package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/service/drone"
	"github.com/SawitProRecruitment/UserService/service/estate"
	"github.com/SawitProRecruitment/UserService/service/tree"
	"github.com/SawitProRecruitment/UserService/types"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	typeestate "github.com/SawitProRecruitment/UserService/repository/estate"
	typeetree "github.com/SawitProRecruitment/UserService/repository/tree"
)

func TestPostEstate(t *testing.T) {
	tests := []struct {
		name           string
		mock           func(m *estate.MockServiceMockRecorder)
		requestBody    map[string]interface{}
		wantResponse   string
		wantHTTPStatus int
	}{
		{
			name: "should create new estate",
			mock: func(stub *estate.MockServiceMockRecorder) {
				stub.CreateNewEstate(context.Background(), 10, 10).Return("1", nil)
			},
			requestBody:    map[string]interface{}{"length": 10, "width": 10},
			wantResponse:   `{"id":"1"}`,
			wantHTTPStatus: http.StatusOK,
		},
		{
			name: "missing field",
			mock: func(m *estate.MockServiceMockRecorder) {},
			requestBody: map[string]interface{}{
				"length": 10,
			},
			wantResponse:   `{"error":"Width must be greater than or equal to 1 "}`,
			wantHTTPStatus: http.StatusBadRequest,
		},
		{
			name: "invalid value",
			mock: func(m *estate.MockServiceMockRecorder) {},
			requestBody: map[string]interface{}{
				"length": "10",
				"width":  "ABC",
			},
			wantResponse:   `{"error":"failed to bind request"}`,
			wantHTTPStatus: http.StatusBadRequest,
		},
		{
			name: "value reached the limit",
			mock: func(m *estate.MockServiceMockRecorder) {},
			requestBody: map[string]interface{}{
				"length": 10000000000,
				"width":  10000000000,
			},
			wantResponse:   `{"error":"Length must be less than or equal to 1000000 \nWidth must be less than or equal to 1000000 "}`,
			wantHTTPStatus: http.StatusBadRequest,
		},
		{
			name: "internal server error",
			mock: func(stub *estate.MockServiceMockRecorder) {
				stub.CreateNewEstate(context.Background(), 10, 10).Return("", errors.New("internal server error"))
			},
			requestBody:    map[string]interface{}{"length": 10, "width": 10},
			wantResponse:   `{"error":"internal server error"}`,
			wantHTTPStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		estateService := estate.NewMockService(gomock.NewController(t))
		s := Server{
			Validator:     NewValidator(),
			EstateService: estateService,
		}
		tt.mock(estateService.EXPECT())

		e := echo.New()
		bodyRequest, _ := json.Marshal(tt.requestBody)
		req := httptest.NewRequest(http.MethodPost, "/estate", bytes.NewBuffer(bodyRequest))
		req.Header.Add("Content-Type", "application/json")

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		s.PostEstate(c)
		assert.JSONEq(t, tt.wantResponse, rec.Body.String())
		assert.Equal(t, tt.wantHTTPStatus, rec.Code)
	}
}

func TestPostEstateEstateIdTree(t *testing.T) {
	failUUID := uuid.New().String()
	successUUID := uuid.New().String()

	tests := []struct {
		name           string
		mock           func(stub *tree.MockServiceMockRecorder)
		requestBody    map[string]interface{}
		requestParam   string
		wantResponse   string
		wantHTTPStatus int
	}{
		{
			name: "should create new tree",
			mock: func(stub *tree.MockServiceMockRecorder) {
				stub.CreateNewTree(context.Background(), successUUID, 10, 10, 10).Return("1", nil)
			},
			requestParam:   successUUID,
			requestBody:    map[string]interface{}{"x": 10, "y": 10, "height": 10},
			wantResponse:   `{"id":"1"}`,
			wantHTTPStatus: http.StatusOK,
		},
		{
			name:           "invalid value",
			mock:           func(stub *tree.MockServiceMockRecorder) {},
			requestBody:    map[string]interface{}{"x": 10, "y": "ABC", "height": 10},
			requestParam:   failUUID,
			wantResponse:   `{"error":"failed to bind request"}`,
			wantHTTPStatus: http.StatusBadRequest,
		},
		{
			name:           "missing field",
			mock:           func(stub *tree.MockServiceMockRecorder) {},
			requestParam:   failUUID,
			requestBody:    map[string]interface{}{"x": 10, "y": 20},
			wantResponse:   `{"error":"Height must be greater than or equal to 1 "}`,
			wantHTTPStatus: http.StatusBadRequest,
		},
		{
			name:           "should return an error when UUID is invalid",
			mock:           func(stub *tree.MockServiceMockRecorder) {},
			requestParam:   "1212-12001-3121",
			requestBody:    map[string]interface{}{"x": 10, "y": 10, "height": 10},
			wantResponse:   `{"error":"EstateID must be a valid UUID"}`,
			wantHTTPStatus: http.StatusBadRequest,
		},
		{
			name: "estate id is not found",
			mock: func(stub *tree.MockServiceMockRecorder) {
				stub.CreateNewTree(context.Background(), failUUID, 10, 10, 10).Return("", types.ErrorEstateNotFound)
			},
			requestParam:   failUUID,
			requestBody:    map[string]interface{}{"x": 10, "y": 10, "height": 10},
			wantResponse:   `{"error":"estate not found"}`,
			wantHTTPStatus: http.StatusNotFound,
		},
		{
			name: "tree is out ouf bound",
			mock: func(stub *tree.MockServiceMockRecorder) {
				stub.CreateNewTree(context.Background(), failUUID, 190, 120, 10).Return("", types.ErorrTreeOutOfBound)
			},
			requestParam:   failUUID,
			requestBody:    map[string]interface{}{"x": 190, "y": 120, "height": 10},
			wantResponse:   `{"error":"tree out of bound"}`,
			wantHTTPStatus: http.StatusBadRequest,
		},
		{
			name: "error is already planted",
			mock: func(stub *tree.MockServiceMockRecorder) {
				stub.CreateNewTree(context.Background(), failUUID, 10, 10, 10).Return("", types.ErrorTreeAlreadyPlanted)
			},
			requestParam:   failUUID,
			requestBody:    map[string]interface{}{"x": 10, "y": 10, "height": 10},
			wantResponse:   `{"error": "tree already planted in same coordinate"}`,
			wantHTTPStatus: http.StatusBadRequest,
		},

		{
			name: "internal server error",
			mock: func(stub *tree.MockServiceMockRecorder) {
				stub.CreateNewTree(context.Background(), failUUID, 10, 120, 10).Return("", errors.New("unknown error"))
			},
			requestParam:   failUUID,
			requestBody:    map[string]interface{}{"x": 10, "y": 120, "height": 10},
			wantResponse:   `{"error":"internal server error"}`,
			wantHTTPStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		treeService := tree.NewMockService(gomock.NewController(t))

		s := Server{
			Validator:   NewValidator(),
			TreeService: treeService,
		}
		tt.mock(treeService.EXPECT())

		e := echo.New()
		bodyRequest, _ := json.Marshal(tt.requestBody)
		req := httptest.NewRequest(http.MethodPost, "/estate/"+tt.requestParam+"/tree", bytes.NewBuffer(bodyRequest))
		req.Header.Add("Content-Type", "application/json")

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)
		s.PostEstateEstateIdTree(c, tt.requestParam)
		assert.JSONEq(t, tt.wantResponse, rec.Body.String())
		assert.Equal(t, tt.wantHTTPStatus, rec.Code)
	}
}

func TestGetEstateEstateIDDronePlan(t *testing.T) {
	successUUID := uuid.New().String()
	failedUUID := uuid.New().String()

	tests := []struct {
		name           string
		mockEstate     func(stub *estate.MockServiceMockRecorder)
		mockTree       func(stub *tree.MockServiceMockRecorder)
		mockDrone      func(stub *drone.MockServiceMockRecorder)
		requestParam   string
		wantResponse   string
		wantHTTPStatus int
	}{
		{
			name: "should return a plan",
			mockEstate: func(stub *estate.MockServiceMockRecorder) {
				stub.RetrieveEstate(context.Background(), successUUID).Return(&typeestate.Estate{
					ID:     successUUID,
					Width:  10,
					Length: 10,
				}, nil)
			},
			mockTree: func(stub *tree.MockServiceMockRecorder) {
				stub.RetrievesByEstateID(context.Background(), successUUID).Return([]*typeetree.Tree{
					{
						XCoordinate: 1,
						YCoordinate: 1,
					},
					{
						XCoordinate: 2,
						YCoordinate: 1,
					},
				}, nil)
			},
			mockDrone: func(stub *drone.MockServiceMockRecorder) {
				stub.GetDronePlane(&typeestate.Estate{
					ID:     successUUID,
					Width:  10,
					Length: 10,
				}, []*typeetree.Tree{
					{
						XCoordinate: 1,
						YCoordinate: 1,
					},
					{
						XCoordinate: 2,
						YCoordinate: 1,
					},
				}, 0).Return(120, drone.Coordinate{
					X: 1,
					Y: 1})
			},
			requestParam:   successUUID,
			wantResponse:   `{"distance":120}`,
			wantHTTPStatus: 200,
		},
		{
			name:           "should return an error when UUID is invalid",
			mockEstate:     func(stub *estate.MockServiceMockRecorder) {},
			mockTree:       func(stub *tree.MockServiceMockRecorder) {},
			mockDrone:      func(stub *drone.MockServiceMockRecorder) {},
			requestParam:   "1212-12001-3121",
			wantResponse:   `{"error":"EstateID must be a valid UUID"}`,
			wantHTTPStatus: http.StatusBadRequest,
		},
		{
			name: "internal server error when retrieve estate",
			mockEstate: func(stub *estate.MockServiceMockRecorder) {
				stub.RetrieveEstate(context.Background(), failedUUID).Return(nil, errors.New("unknown error"))
			},
			mockTree:       func(stub *tree.MockServiceMockRecorder) {},
			mockDrone:      func(stub *drone.MockServiceMockRecorder) {},
			requestParam:   failedUUID,
			wantResponse:   `{"error":"internal server error"}`,
			wantHTTPStatus: http.StatusInternalServerError,
		},
		{
			name: "internal server error when retrieve trees",
			mockEstate: func(stub *estate.MockServiceMockRecorder) {
				stub.RetrieveEstate(context.Background(), failedUUID).Return(&typeestate.Estate{}, nil)
			},
			mockTree: func(stub *tree.MockServiceMockRecorder) {
				stub.RetrievesByEstateID(context.Background(), failedUUID).Return(nil, errors.New("unknown error"))
			},
			mockDrone:      func(stub *drone.MockServiceMockRecorder) {},
			requestParam:   failedUUID,
			wantResponse:   `{"error":"internal server error"}`,
			wantHTTPStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		estateService := estate.NewMockService(gomock.NewController(t))
		treeService := tree.NewMockService(gomock.NewController(t))
		droneService := drone.NewMockService(gomock.NewController(t))

		s := Server{
			Validator:     NewValidator(),
			EstateService: estateService,
			TreeService:   treeService,
			DroneService:  droneService,
		}
		tt.mockEstate(estateService.EXPECT())
		tt.mockTree(treeService.EXPECT())
		tt.mockDrone(droneService.EXPECT())

		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/estate/"+tt.requestParam+"/drone-plan", nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		s.GetEstateEstateIdDronePlan(c, tt.requestParam, generated.GetEstateEstateIdDronePlanParams{})

		assert.JSONEq(t, tt.wantResponse, rec.Body.String())
		assert.Equal(t, tt.wantHTTPStatus, rec.Code)
	}
}

func TestGetEstateEstateIdStats(t *testing.T) {
	successUUID := uuid.New().String()

	tests := []struct {
		name           string
		mockTree       func(stub *tree.MockServiceMockRecorder)
		requestParam   string
		wantResponse   string
		wantHTTPStatus int
	}{
		{
			name: "should return a stats",
			mockTree: func(stub *tree.MockServiceMockRecorder) {
				stub.GetStats(context.Background(), successUUID).Return(&typeetree.TreeStats{
					Count:  10,
					Max:    5,
					Min:    1,
					Median: 4,
				}, nil)
			},
			requestParam:   successUUID,
			wantResponse:   `{"count":10,"max":5,"min":1,"median":4}`,
			wantHTTPStatus: 200,
		},
	}

	for _, tt := range tests {
		treeService := tree.NewMockService(gomock.NewController(t))

		s := Server{
			Validator:   NewValidator(),
			TreeService: treeService,
		}
		tt.mockTree(treeService.EXPECT())

		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/estate/"+tt.requestParam+"/stats", nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		s.GetEstateEstateIdStats(c, tt.requestParam)

		assert.JSONEq(t, tt.wantResponse, rec.Body.String())
		assert.Equal(t, tt.wantHTTPStatus, rec.Code)
	}
}
