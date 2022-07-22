package handler

import (
	"github.com/fede/golang_api/internal/platform/mocks"
	"github.com/fede/golang_api/internal/platform/service"
	"github.com/fede/golang_api/internal/platform/storage/repository"
	"github.com/fede/golang_api/kit"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

var (
	bodyAssignTripToDriverOK         = []byte(`{"driver_file": 654321}`)
	bodyAssignTripToDriverNoOk       = []byte(`{"driver_file": 123456}`)
	bodyAssignTripToDriverNotExistOk = []byte(`{"driver_file": 123455646466}`)
	bodyAssignTripToDriverBadRequest = []byte(`{"driver_file": "s"}`)
)

func newTripHandler() *TripControllers {
	db := mocks.MockDB()
	repoTrip := repository.NewTripRepository(db)
	repoDriver := repository.NewDriverRepository(db)
	serviceDriver := service.NewDriverService(repoDriver)
	serviceTrip := service.NewTripService(repoTrip)
	handler := NewTripController(serviceTrip, serviceDriver)

	return handler
}

func TestCloseTripToDriverOK(t *testing.T) {
	handler := newTripHandler()
	response, contextTest := kit.GetTestConfig(http.MethodPost, "/close/driver?id=1", nil, nil)

	handler.CloseTripToDriver(contextTest)

	assert.Nil(t, contextTest.Errors)
	assert.Equal(t, http.StatusNoContent, response.Code)
}

func TestCloseTripToDriverPanicWithOutID(t *testing.T) {
	handler := newTripHandler()
	_, contextTest := kit.GetTestConfig(http.MethodPost, "/close/driver", nil, nil)

	assert.Panics(t, func() {
		handler.CloseTripToDriver(contextTest)
	})
}

func TestCloseTripToDriverPanicWithUserNotFound(t *testing.T) {
	handler := newTripHandler()
	_, contextTest := kit.GetTestConfig(http.MethodPost, "/close/driver?id=55", nil, nil)

	assert.Panics(t, func() {
		handler.CloseTripToDriver(contextTest)
	})
}

func TestCloseTripToDriverPanicWithUserWithOutTripOpen(t *testing.T) {
	handler := newTripHandler()
	_, contextTest := kit.GetTestConfig(http.MethodPost, "/close/driver?id=2", nil, nil)

	assert.Panics(t, func() {
		handler.CloseTripToDriver(contextTest)
	})
}

func TestAssignTripToDriverOk(t *testing.T) {
	handler := newTripHandler()
	response, contextTest := kit.GetTestConfig(http.MethodPost, "/assign/driver", bodyAssignTripToDriverOK, nil)

	handler.AssignTripToDriver(contextTest)

	assert.Nil(t, contextTest.Errors)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestAssignTripToDriverBadRequest(t *testing.T) {
	handler := newTripHandler()
	_, contextTest := kit.GetTestConfig(http.MethodPost, "/assign/driver", bodyAssignTripToDriverBadRequest, nil)

	assert.Panics(t, func() {
		handler.AssignTripToDriver(contextTest)
	})

}

func TestAssignTripToDriverNotExist(t *testing.T) {
	handler := newTripHandler()
	_, contextTest := kit.GetTestConfig(http.MethodPost, "/assign/driver", bodyAssignTripToDriverNotExistOk, nil)

	assert.Panics(t, func() {
		handler.AssignTripToDriver(contextTest)
	})
}

func TestAssignTripToDriverWitTripOpen(t *testing.T) {
	handler := newTripHandler()
	_, contextTest := kit.GetTestConfig(http.MethodPost, "/assign/driver", bodyAssignTripToDriverNoOk, nil)

	assert.Panics(t, func() {
		handler.AssignTripToDriver(contextTest)
	})
}
