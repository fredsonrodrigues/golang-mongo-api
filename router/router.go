package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	. "golang-api/conf/DAO"
)

// Utils
var dao = ArticlesDAO{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}
func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
//

func GetAll(w http.ResponseWriter, r *http.Request) {
	article, err := dao.GetAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, article)
}

func GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	article, err := dao.GetByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, article)
}