package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SawitProRecruitment/UserService/service/estate"
	"github.com/SawitProRecruitment/UserService/service/tree"
	"github.com/SawitProRecruitment/UserService/types"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
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
			wantResponse:   `{"error":"bad type value"}`,
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
