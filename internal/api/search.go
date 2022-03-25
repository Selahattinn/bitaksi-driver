package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Selahattinn/bitaksi-driver/internal/api/response"
	"github.com/Selahattinn/bitaksi-driver/internal/model"
)

// GetAllDrivers godoc
// @Summary      Searches all drivers
// @Description  Searches all drivers and finds the suitable drivers
// @Tags         Search
// @Accept       json
// @Produce      json
// @Param 		 Flag header string true "Flag of the Match Service"
// @Param        Search   body      model.Search  true  "An Search object with json format"
// @Success      200  {object}  []model.SearchResult "Search result info with json format"
// @Failure      401  "UnAuthorized"
// @Failure      404  "Not found"
// @Router       /search [post]
func (a *API) Search(w http.ResponseWriter, r *http.Request) {
	var search model.Search
	err := json.NewDecoder(r.Body).Decode(&search)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error parsing Search info: %v", err), http.StatusNotFound, "")
		return
	}

	err = search.Validate()
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error validating Search info: %v", err), http.StatusNotFound, "")
		return
	}
	drivers, err := a.Service.GetSearchService().FindSuitableDrivers(&search)
	if err != nil {
		response.Errorf(w, r, fmt.Errorf("error getting Search: %v", err), http.StatusNotFound, "")
		return
	}
	response.Write(w, r, drivers)
}
