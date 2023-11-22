package routehandler

import (
	"backend/controller"
	"backend/dbconfig"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlePocRoutes(r *mux.Router, db *dbconfig.DB) {
	pocHandler := controller.NewPocHandler(db)

	//Poc APIs
	r.HandleFunc("/poc", pocHandler.CreatePoc).Methods(http.MethodPost)
	r.HandleFunc("/poc", pocHandler.GetPocsList).Methods(http.MethodGet)
	r.HandleFunc("/poc/{id}", pocHandler.GetPocById).Methods(http.MethodGet)
	r.HandleFunc("/poc/{id}", pocHandler.UpdatePoc).Methods(http.MethodPut)
	r.HandleFunc("/poc/{id}", pocHandler.DeletePoc).Methods(http.MethodDelete)

	r.HandleFunc("/techcount", pocHandler.GetTechcount).Methods(http.MethodGet)
	r.HandleFunc("/piechartcount", pocHandler.GetPieChart).Methods(http.MethodGet)

	//feed APIs
	r.HandleFunc("/feed", pocHandler.GetFeed).Methods(http.MethodGet)
	r.HandleFunc("/feed", pocHandler.CreateFeed).Methods(http.MethodPost)
	r.HandleFunc("/feed/{id}", pocHandler.UpdateFeed).Methods(http.MethodPut)
	r.HandleFunc("/feed/{id}", pocHandler.DeleteFeed).Methods(http.MethodDelete)

	//Architecture Review APIs
	r.HandleFunc("/arb", pocHandler.GetArb).Methods(http.MethodGet)
	r.HandleFunc("/arb", pocHandler.CreateArb).Methods(http.MethodPost)
	r.HandleFunc("/arb/{id}", pocHandler.UpdateArb).Methods(http.MethodPut)

	// r.HandleFunc("/arb/{id}", pocHandler.UpdateArb).Methods(http.MethodPut)
	r.HandleFunc("/badge", pocHandler.GetBadges).Methods(http.MethodGet)
	r.HandleFunc("/arbstatus", pocHandler.GetArbStatus).Methods(http.MethodGet)
	r.HandleFunc("/arbstatus/{id}", pocHandler.UpdateArbStatus).Methods(http.MethodPut)
	r.HandleFunc("/arb/{id}", pocHandler.DeleteArb).Methods(http.MethodDelete)

	//Review Report
	r.HandleFunc("/pillars", pocHandler.GetPillars).Methods(http.MethodGet)
	r.HandleFunc("/topics", pocHandler.GetTopics).Methods(http.MethodGet)
	r.HandleFunc("/arbreview", pocHandler.GetARBReviews).Methods(http.MethodGet)

	r.HandleFunc("/tecmember", pocHandler.GetTecMembers).Methods(http.MethodGet)
	r.HandleFunc("/tecmember", pocHandler.SaveTecMember).Methods(http.MethodPost)
	r.HandleFunc("/tectimeline/{id}", pocHandler.GetTecActivity).Methods(http.MethodGet)
	r.HandleFunc("/tectimeline", pocHandler.SaveTecActivity).Methods(http.MethodPost)
	//Accelerator
	r.HandleFunc("/accsnap", pocHandler.GetAccSnapshots).Methods(http.MethodGet)
	r.HandleFunc("/accsnap", pocHandler.SaveAccSnapshot).Methods(http.MethodPost)
	r.HandleFunc("/acctimeline/{id}", pocHandler.GetAccActivity).Methods(http.MethodGet)
	r.HandleFunc("/acctimeline", pocHandler.SaveAccActivity).Methods(http.MethodPost)
}
