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

// AddDriver godoc
// @Summary      Adds a driver
// @Description  Gets driver information with post request and saves it to database
// @Tags         Driver
// @Accept       json
// @Produce      json
// @Param        Driver   body      model.Driver  true  "An Driver object with json format"
// @Success      200  {object}  int "Driver id which is added"
// @Failure      404  "Not found"
// @Router       /driver [post]
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

// FindDriver godoc
// @Summary      Gets a driver
// @Description  Gets driver's id from url and returns related driver
// @Tags         Driver
// @Accept       json
// @Produce      json
// @Param        id   path    int  true  "ID of the Driver"
// @Success      200  {object}  model.Driver "Driver info with json format"
// @Failure      404  "Not found"
// @Router       /driver/{id} [get]
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

// UpdateDriver godoc
// @Summary      Update a driver
// @Description  Gets driver's information from body and update this driver by id
// @Tags         Driver
// @Accept       json
// @Produce      json
// @Param        Driver   body      model.Driver  true  "An Driver object with json format"
// @Success      200  {object}  model.Driver "Return updated Driver info"
// @Failure      404  "Not found"
// @Router       /driver [put]
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
	d, err := a.Service.GetDriverService().UpdateDriver(&driver)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error updating Driver: %v", err), http.StatusNotFound, "")
		return
	}
	response.Write(w, r, d)
}

// DeleteDriver godoc
// @Summary      Delete a driver
// @Description  Gets driver's information from body and delete this driver by id
// @Tags         Driver
// @Accept       json
// @Produce      json
// @Param        Driver   body      model.Driver  true  "An Driver object with json format"
// @Success      200  {object}  int "return 1 if  success else returns 0"
// @Failure      404  "Not found"
// @Router       /driver [delete]
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

// GetAllDrivers godoc
// @Summary      Gets all drivers
// @Description  Gets all information of drivers
// @Tags         Driver
// @Accept       json
// @Produce      json
// @Success      200  {object}  []model.Driver "Driver info with json format"
// @Failure      404  "Not found"
// @Router       /driver [get]
func (a *API) GetAllDrivers(w http.ResponseWriter, r *http.Request) {
	drivers, err := a.Service.GetDriverService().GetAllDrivers()
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting Drivers: %v", err), http.StatusNotFound, "")
		return
	}
	response.Write(w, r, drivers)
}
