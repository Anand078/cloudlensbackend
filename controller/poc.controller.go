package controller

import (
	"backend/dbconfig"
	"backend/logging"
	"backend/models"
	repo "backend/service"
	poc "backend/service/poc"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Poc struct {
	repo repo.PocRepo
}

func NewPocHandler(db *dbconfig.DB) *Poc {
	return &Poc{
		repo: poc.NewPocRepo(db.SQL),
	}
}

// swagger:route GET /poc Pocs getPocs
// Get Pocs list
// responses:
//   200: jsonResponse
//
// swagger:response jsonResponse

// Get Pocs list
func (e *Poc) GetPocsList(w http.ResponseWriter, r *http.Request) {
	res, err := e.repo.FetchPocs(r.Context())
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		logging.Logger.Errorf(err.Error())
		return
	}
	// On success
	respondwithJSON(w, 200, res)
}

// swagger:route GET /poc/{id} Pocs getPocById
// Get Poc by id
// responses:
//   200: jsonResponse
//
// swagger:response jsonResponse

// Get Poc by id
func (e *Poc) GetPocById(w http.ResponseWriter, r *http.Request) {
	// Covert id from str to int64
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad request")
		logging.Logger.Errorf(err.Error())
		return
	}
	res, err := e.repo.GetPocByID(r.Context(), id)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		respondWithError(w, http.StatusNotFound, "Not Found")
		return
	}
	// On success
	respondwithJSON(w, 200, res)
}

// swagger:route POST /poc Pocs createPoc
// Create new poc
// responses:
//   200: jsonResponse
//
// swagger:response jsonResponse

// Create Poc
func (e *Poc) CreatePoc(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logging.Logger.Errorf(err.Error())
		return
	}

	if r.Body == nil {
		logging.Logger.Infof("empty body")
	}
	// Create a variable to hold the JSON data
	req := models.Poc{}

	// Decode the JSON data from the request body
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logging.Logger.Errorf(err.Error())
		return
	}

	res, err := e.repo.CreatePoc(r.Context(), &req)
	if err != nil {
		respondWithError(w, http.StatusForbidden, "Forbidden")
		logging.Logger.Errorf(err.Error())
		return
	}
	// On success
	respondwithJSON(w, 200, res)
}

// swagger:route POST /arb Arbs createArb
// Create new Arb
// responses:
//   200: jsonResponse
//
// swagger:response jsonResponse

// Create Arb
func (e *Poc) CreateArb(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		logging.Logger.Errorf(err.Error())
		return
	}

	if r.Body == nil {
		logging.Logger.Infof("empty body")
	}
	// Create a variable to hold the JSON data
	req := models.Reviews{}

	// Decode the JSON data from the request body
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logging.Logger.Errorf(err.Error())
		return
	}

	res, err := e.repo.CreateArb(r.Context(), &req)
	if err != nil {
		respondWithError(w, http.StatusForbidden, err.Error())
		logging.Logger.Errorf(err.Error())
		return
	}
	// On success
	respondwithJSON(w, 200, res)
}

// SaveTecMember inserts TecMember rows with negative ID values and updates rows with non-negative ID values.
func (e *Poc) SaveTecMember(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		logging.Logger.Errorf(err.Error())
		return
	}

	if r.Body == nil {
		logging.Logger.Infof("empty body")
	}

	var tecMembers []models.SaveTecMember

	err = json.NewDecoder(r.Body).Decode(&tecMembers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logging.Logger.Errorf(err.Error())
		return
	}
	// Iterate over the slice of tecMembers
	for _, tecMember := range tecMembers {
		if tecMember.ID < 0 {
			// Insert TecMember with negative ID
			res, err := e.repo.CreateTecMember(r.Context(), &tecMember)
			if err != nil {
				respondWithError(w, http.StatusForbidden, "Forbidden")
				logging.Logger.Errorf(err.Error())
				return
			}
			// On success
			respondwithJSON(w, 200, res)
		} else {
			// Update TecMember with non-negative ID
			res, err := e.repo.UpdateTecMember(r.Context(), &tecMember)
			if err != nil {
				respondWithError(w, http.StatusForbidden, err.Error())
				logging.Logger.Errorf(err.Error())
				return
			}
			// On success
			respondwithJSON(w, 200, res)
		}
	}
}

// SaveAccSnapshot inserts Accelerator rows with negative ID values and updates rows with non-negative ID values.
func (e *Poc) SaveAccSnapshot(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		logging.Logger.Errorf(err.Error())
		return
	}

	if r.Body == nil {
		logging.Logger.Infof("empty body")
	}

	var accMembers []models.AccSnap

	err = json.NewDecoder(r.Body).Decode(&accMembers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logging.Logger.Errorf(err.Error())
		return
	}
	// Iterate over the slice of accMembers
	for _, accMember := range accMembers {
		if accMember.ID < 0 {
			// Insert accMember with negative ID
			res, err := e.repo.CreateAccSnap(r.Context(), &accMember)
			if err != nil {
				respondWithError(w, http.StatusForbidden, "Forbidden")
				logging.Logger.Errorf(err.Error())
				return
			}
			// On success
			respondwithJSON(w, 200, res)
		} else {
			// Update accMember with non-negative ID
			res, err := e.repo.UpdateAccSnap(r.Context(), &accMember)
			if err != nil {
				respondWithError(w, http.StatusForbidden, err.Error())
				logging.Logger.Errorf(err.Error())
				return
			}
			// On success
			respondwithJSON(w, 200, res)
		}
	}
}

// swagger:route POST /feed Feed createFeed
// Create new feed
// responses:
//   200: jsonResponse
//
// swagger:response jsonResponse

// Create Feed
func (e *Poc) CreateFeed(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form data", http.StatusBadRequest)
		logging.Logger.Errorf(err.Error())
		return
	}

	if r.Body == nil {
		logging.Logger.Infof("empty body")
	}
	// Create a variable to hold the JSON data
	req := models.Feed{}

	// Decode the JSON data from the request body
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	res, err := e.repo.CreateFeed(r.Context(), &req)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		respondWithError(w, http.StatusForbidden, "Forbidden")
		return
	}
	// On success
	respondwithJSON(w, 200, res)
}

// swagger:route PUT /poc/{id} Pocs editPoc
// Update Poc
// responses:
//   200: jsonResponse
//
// swagger:response jsonResponse

// Update poc
func (e *Poc) UpdatePoc(w http.ResponseWriter, r *http.Request) {
	req := models.Poc{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		respondWithError(w, http.StatusBadRequest, err.Error())

		return
	}

	res, err := e.repo.UpdatePoc(r.Context(), &req, id)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		respondWithError(w, http.StatusForbidden, err.Error())
		return
	}
	// On success
	respondwithJSON(w, 200, res)
}

// Update feed
func (e *Poc) UpdateFeed(w http.ResponseWriter, r *http.Request) {
	req := models.Feed{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		respondWithError(w, http.StatusBadRequest, "Bad request")
		return
	}
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		respondWithError(w, http.StatusBadRequest, "Bad request")
		return
	}

	res, err := e.repo.UpdateFeed(r.Context(), &req, id)

	if err != nil {
		respondWithError(w, http.StatusForbidden, "Forbidden")
		logging.Logger.Errorf(err.Error())
		return
	}
	// On success
	respondwithJSON(w, http.StatusAccepted, res)
}

// swagger:route DELETE /poc/{id} Pocs deletePoc
// Delete Poc
// responses:
//   200: jsonResponse
//
// swagger:response jsonResponse

// Delete Poc
func (e *Poc) DeletePoc(w http.ResponseWriter, r *http.Request) {
	// Covert id from str to int64
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Bad request")
		logging.Logger.Errorf(err.Error())
		return
	}

	res, err := e.repo.DeletePoc(r.Context(), id)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		respondWithError(w, http.StatusForbidden, "Forbidden")
		return
	}
	// On succes
	respondwithJSON(w, http.StatusOK, res)
}

// swagger:route DELETE /feed/{id} Pocs deleteFeed
// Delete Feed
// responses:
//   200: jsonResponse
//
// swagger:response jsonResponse

// Delete Feed
func (e *Poc) DeleteFeed(w http.ResponseWriter, r *http.Request) {
	// Covert id from str to int64
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		respondWithError(w, http.StatusBadRequest, "Bad request")
		return
	}

	res, err := e.repo.DeleteFeed(r.Context(), id)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		respondWithError(w, http.StatusForbidden, "Forbidden")
		return
	}
	// On success
	respondwithJSON(w, http.StatusOK, res)
}

// Delete Feed
func (e *Poc) DeleteArb(w http.ResponseWriter, r *http.Request) {
	// Covert id from str to int64
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		respondWithError(w, http.StatusBadRequest, "Bad request")
		return
	}

	res, err := e.repo.DeleteArb(r.Context(), id)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		respondWithError(w, http.StatusForbidden, "Forbidden")
		return
	}
	// On success
	respondwithJSON(w, http.StatusOK, res)
}

// swagger:route GET /techcount TechCount getTechCount
// Get Pocs list
// responses:
//   200: jsonResponse
//
// swagger:response jsonResponse

// Get techcount list
func (e *Poc) GetTechcount(w http.ResponseWriter, r *http.Request) {
	res, err := e.repo.FetchTechCount(r.Context())
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		logging.Logger.Errorf(err.Error())
		return
	}
	// On success
	respondwithJSON(w, 200, res)
}

// swagger:route GET /piechartcount Piechart getPieChart
// Get pocs list
// responses:
//   200: jsonResponse
//
// swagger:response jsonResponse

// Get piechart list
func (e *Poc) GetPieChart(w http.ResponseWriter, r *http.Request) {
	res, err := e.repo.FetchPieChartCount(r.Context())
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		logging.Logger.Errorf(err.Error())
		return
	}
	// On success
	respondwithJSON(w, 200, res)
}

// swagger:route GET /feed Feed getFeed
// Get Feeds list
// responses:
//   200: jsonResponse
//
// swagger:response jsonResponse

// Get Feed list
func (e *Poc) GetFeed(w http.ResponseWriter, r *http.Request) {
	res, err := e.repo.FetchFeeds(r.Context())
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		logging.Logger.Errorf(err.Error())
		return
	}
	// On success
	respondwithJSON(w, 200, res)
}

// Get Feed list
func (e *Poc) GetPillars(w http.ResponseWriter, r *http.Request) {
	res, err := e.repo.FetchPillars(r.Context())
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		logging.Logger.Errorf(err.Error())
		return
	}
	// On success
	respondwithJSON(w, 200, res)
}

// Get Topic list
func (e *Poc) GetTopics(w http.ResponseWriter, r *http.Request) {
	res, err := e.repo.FetchTopics(r.Context())
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		logging.Logger.Errorf(err.Error())
		return
	}
	// On success
	respondwithJSON(w, 200, res)
}

// Get ARB Review list
func (e *Poc) GetARBReviews(w http.ResponseWriter, r *http.Request) {
	res, err := e.repo.FetchArbReviews(r.Context())
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		logging.Logger.Errorf(err.Error())
		return
	}
	// On success
	respondwithJSON(w, 200, res)
}

// Get Tec Member list
func (e *Poc) GetTecMembers(w http.ResponseWriter, r *http.Request) {
	res, err := e.repo.FetchTecMembers(r.Context())
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		logging.Logger.Errorf(err.Error())
		return
	}
	// On success
	respondwithJSON(w, 200, res)
}

// Get Accelerator snap list
func (e *Poc) GetAccSnapshots(w http.ResponseWriter, r *http.Request) {
	res, err := e.repo.FetchAccSnap(r.Context())
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		logging.Logger.Errorf(err.Error())
		return
	}
	// On success
	respondwithJSON(w, 200, res)
}

// swagger:route GET /arb Arb getAR
// Get AR list
// responses:
//   200: jsonResponse
//
// swagger:response jsonResponse

// Get AR list
func (e *Poc) GetArb(w http.ResponseWriter, r *http.Request) {
	res, err := e.repo.FetchArb(r.Context())
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		logging.Logger.Errorf(err.Error())
		return
	}
	// On success
	respondwithJSON(w, 200, res)
}

// swagger:route GET /badge Arb getBadge
// Get BAdge list
// responses:
//   200: jsonResponse
//
// swagger:response jsonResponse

// Get Badge list
func (e *Poc) GetBadges(w http.ResponseWriter, r *http.Request) {
	res, err := e.repo.FetchBadges(r.Context())
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		logging.Logger.Errorf(err.Error())
		return
	}
	// On success
	respondwithJSON(w, 200, res)
}

// swagger:route GET /arbstatus Arb getArbStatus
// Get Arbstatus list
// responses:
//   200: jsonResponse
//
// swagger:response jsonResponse

// Get ArbStatus list
func (e *Poc) GetArbStatus(w http.ResponseWriter, r *http.Request) {
	res, err := e.repo.FetchArbStatus(r.Context())
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		logging.Logger.Errorf(err.Error())
		return
	}
	// On success
	respondwithJSON(w, 200, res)
}

// Update feed
func (e *Poc) UpdateArbStatus(w http.ResponseWriter, r *http.Request) {
	req := models.ArbStatusId{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		respondWithError(w, http.StatusBadRequest, "Bad request")
		return
	}
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		respondWithError(w, http.StatusBadRequest, "Bad request")
		return
	}
	res, err := e.repo.UpdateArbStatus(r.Context(), uint(req.StatusId), id)

	if err != nil {
		respondWithError(w, http.StatusForbidden, "Forbidden")
		logging.Logger.Errorf(err.Error())
		return
	}
	// On success
	respondwithJSON(w, http.StatusAccepted, res)
}

// Update Arb
func (e *Poc) UpdateArb(w http.ResponseWriter, r *http.Request) {
	req := models.Reviews{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		respondWithError(w, http.StatusBadRequest, "Bad request")
		return
	}
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		logging.Logger.Errorf(err.Error())
		respondWithError(w, http.StatusBadRequest, "Bad request")
		return
	}

	res, err := e.repo.UpdateArb(r.Context(), &req, id)

	if err != nil {
		respondWithError(w, http.StatusForbidden, "Forbidden")
		logging.Logger.Errorf(err.Error())
		return
	}
	// On success
	respondwithJSON(w, http.StatusAccepted, res)
}
