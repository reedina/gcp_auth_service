package ctrl

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/mail"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/reedina/gcp_auth_service/model"
)

//CreateAccount (POST)
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account model.Account

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&account); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	//Does the Account Resource Exist ?
	if model.DoesAccountResourceExist(&account) == true {
		respondWithError(w, http.StatusConflict, "Resource already exists")
		return
	}

	// Does Role Resource Exist ?
	if model.DoesRoleIDExist(account.Role.ID) != true {
		respondWithError(w, http.StatusBadRequest, "Role ID does not exist")
		return
	}
	//Resource does not exist, go ahead and create resource
	if err := model.CreateAccount(&account); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Get Role name for Role id
	role := model.Role{}
	role.ID = account.Role.ID
	model.GetRole(&role)
	account.Role.Name = role.Name

	respondWithJSON(w, http.StatusCreated, account)
}

//GetAccounts  (GET)
func GetAccounts(w http.ResponseWriter, r *http.Request) {

	accounts, err := model.GetAccounts()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, accounts)
}

//GetAccount (GET)
func GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Account ID")
		return
	}

	account := model.Account{ID: id}
	if err := model.GetAccount(&account); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Account not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, account)
}

//GetAccountByEmail (GET)
func GetAccountByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email, err := mail.ParseAddress(vars["email"])
	//accountEmail := vars["email"]

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Email Address")
		return
	}

	account := model.Account{}
	account.Email = email.Address

	if err := model.GetAccountByEmail(&account); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Account not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, account)
}

//GetAccountsByRoleName (GET)
func GetAccountsByRoleName(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountRoleName := vars["name"]

	account := model.Account{}
	account.Role.Name = accountRoleName

	users, err := model.GetAccountsByRoleName(&account)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Role Name not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, users)
}

//GetAccountsByRoleID (GET)
func GetAccountsByRoleID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User ID")
		return
	}

	account := model.Account{}
	account.Role.ID = id

	accounts, err := model.GetAccountsByRoleID(&account)

	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Role ID not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, accounts)
}

//UpdateAccount (PUT)
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Account ID")
		return
	}

	var account model.Account

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&account); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
		return
	}
	defer r.Body.Close()
	account.ID = id

	// Does Role Resource Exist ?
	if model.DoesRoleIDExist(account.Role.ID) != true {
		respondWithError(w, http.StatusBadRequest, "Role ID does not exist")
		return
	}
	// Does Account Resource Exist ?
	if model.DoesAccountIDExist(account.ID) != true {
		respondWithError(w, http.StatusBadRequest, "Account ID does not exist")
		return
	}
	// Does Account Email exists for another ID
	if model.DoesAccountEmailExistForAnotherID(account.Email, account.ID) == true {
		respondWithError(w, http.StatusBadRequest, "Account Email Exists for another Account ID")
		return
	}
	if err := model.UpdateAccount(&account); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Get Role name for Role id
	role := model.Role{}
	role.ID = account.Role.ID
	model.GetRole(&role)
	account.Role.Name = role.Name
	respondWithJSON(w, http.StatusOK, account)
}

//DeleteAccount (DELETE)
func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Account ID")
		return
	}
	account := model.Account{ID: id}

	if err := model.DeleteAccount(&account); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
