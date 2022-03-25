package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Selahattinn/bitaksi-driver/internal/api/response"
	"github.com/Selahattinn/bitaksi-driver/internal/model"
	"github.com/gorilla/mux"
)

func (a *API) AddDriver(w http.ResponseWriter, r *http.Request) {

	var driver model.Driver
	err := json.NewDecoder(r.Body).Decode(&driver)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error parsing Driver info: %v", err), http.StatusNotFound, "")
		return
	}
	err = driver.Validate()
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error validating Driver info: %v", err), http.StatusNotFound, "")
		return
	}
	id, err := a.Service.GetDriverService().CreateDriver(&driver)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error creating Driver: %v", err), http.StatusNotFound, "")
		return
	}
	response.Write(w, r, id)
}

func (a *API) FindDriver(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	// convert id to int64
	idInt64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error parsing Driver id: %v", err), http.StatusNotFound, "")
		return
	}
	driver, err := a.Service.GetDriverService().GetDriver(idInt64)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting Driver: %v", err), http.StatusNotFound, "")
		return
	}
	response.Write(w, r, driver)
}

func (a *API) UpdateDriver(w http.ResponseWriter, r *http.Request) {
	var driver model.Driver
	err := json.NewDecoder(r.Body).Decode(&driver)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error parsing Driver info: %v", err), http.StatusNotFound, "")
		return
	}
	err = driver.Validate()
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error validating Driver info: %v", err), http.StatusNotFound, "")
		return
	}
	id, err := a.Service.GetDriverService().UpdateDriver(&driver)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error updating Driver: %v", err), http.StatusNotFound, "")
		return
	}
	response.Write(w, r, id)
}

func (a *API) DeleteDriver(w http.ResponseWriter, r *http.Request) {
	var driver model.Driver
	err := json.NewDecoder(r.Body).Decode(&driver)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error parsing Driver info: %v", err), http.StatusNotFound, "")
		return
	}
	err = driver.Validate()
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error validating Driver info: %v", err), http.StatusNotFound, "")
		return
	}
	id, err := a.Service.GetDriverService().DeleteDriver(&driver)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error updating Driver: %v", err), http.StatusNotFound, "")
		return
	}
	response.Write(w, r, id)
}

func (a *API) GetAllDrivers(w http.ResponseWriter, r *http.Request) {
	drivers, err := a.Service.GetDriverService().GetAllDrivers()
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting Drivers: %v", err), http.StatusNotFound, "")
		return
	}
	response.Write(w, r, drivers)
}
