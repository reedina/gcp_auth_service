package ctrl

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reedina/gcp_auth_service/model"
)

//CreateRole (POST)
func CreateRole(w http.ResponseWriter, r *http.Request) {
	var role model.Role

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&role); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	//Does the Role Resource Exist ?
	if model.DoesRoleResourceExist(&role) == true {
		respondWithError(w, http.StatusConflict, "Resource already exists")
		return
	}

	//Resource does not exist, go ahead and create resource
	if err := model.CreateRole(&role); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, role)
}

//GetRoles  (GET)
func GetRoles(w http.ResponseWriter, r *http.Request) {

	roles, err := model.GetRoles()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, roles)
}

//GetRole (GET)
func GetRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Role ID")
		return
	}

	role := model.Role{ID: id}
	if err := model.GetRole(&role); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Role not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, role)
}

//UpdateRole (PUT)
func UpdateRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Role ID")
		return
	}

	var role model.Role

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&role); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	role.ID = id

	// Does Role Resource Exist ?
	if model.DoesRoleIDExist(role.ID) != true {
		respondWithError(w, http.StatusBadRequest, "Role ID does not exist")
		return
	}
	// Does Role Email exists for another ID
	if model.DoesRoleNameExistForAnotherID(role.Name, role.ID) == true {
		respondWithError(w, http.StatusBadRequest, "Role Name Exists for another Role ID")
		return
	}
	if err := model.UpdateRole(&role); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, role)
}

//DeleteRole (DELETE)
func DeleteRole(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Role ID")
		return
	}
	role := model.Role{ID: id}

	if err := model.DeleteRole(&role); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
