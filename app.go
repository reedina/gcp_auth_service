package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/reedina/gcp_auth_service/ctrl"
	"github.com/reedina/gcp_auth_service/model"
	"github.com/rs/cors"

	//Initialize mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//App  (TYPE)
type App struct {
	Router *mux.Router
}

//InitializeApplication - Init router, db connection and restful routes
func (a *App) InitializeApplication(user, password, url, dbname string) {

	model.ConnectDB(user, password, url, dbname)
	a.Router = mux.NewRouter()
	a.InitializeRoutes()
}

//InitializeRoutes - Declare all application routes
func (a *App) InitializeRoutes() {

	//model.Account struct
	a.Router.HandleFunc("/api/account", ctrl.CreateAccount).Methods("POST")
	a.Router.HandleFunc("/api/accounts", ctrl.GetAccounts).Methods("GET")
	a.Router.HandleFunc("/api/account/{id:[0-9]+}", ctrl.GetAccount).Methods("GET")
	a.Router.HandleFunc("/api/account/{email}", ctrl.GetAccountByEmail).Methods("GET")
	a.Router.HandleFunc("/api/account/{id:[0-9]+}", ctrl.UpdateAccount).Methods("PUT")
	a.Router.HandleFunc("/api/account/{id:[0-9]+}", ctrl.DeleteAccount).Methods("DELETE")

	//model.Role struct
	a.Router.HandleFunc("/api/role", ctrl.CreateRole).Methods("POST")
	a.Router.HandleFunc("/api/roles", ctrl.GetRoles).Methods("GET")
	a.Router.HandleFunc("/api/role/{id:[0-9]+}", ctrl.GetRole).Methods("GET")
	a.Router.HandleFunc("/api/role/{id:[0-9]+}", ctrl.UpdateRole).Methods("PUT")
	a.Router.HandleFunc("/api/role/{id:[0-9]+}", ctrl.DeleteRole).Methods("DELETE")

	/*
		//model.Project struct
		a.Router.HandleFunc("/api/project", ctrl.CreateProject).Methods("POST")
		a.Router.HandleFunc("/api/projects", ctrl.GetProjects).Methods("GET")
		a.Router.HandleFunc("/api/projects/team/name/{name}", ctrl.GetProjectsByTeamName).Methods("GET")
		a.Router.HandleFunc("/api/projects/team/id/{id:[0-9]+}", ctrl.GetProjectsByTeamID).Methods("GET")
		a.Router.HandleFunc("/api/project/{id:[0-9]+}", ctrl.GetProject).Methods("GET")
		a.Router.HandleFunc("/api/project/{id:[0-9]+}", ctrl.UpdateProject).Methods("PUT")
		a.Router.HandleFunc("/api/project/{id:[0-9]+}", ctrl.DeleteProject).Methods("DELETE")
	*/
}

//RunApplication - Start the HTTP server
func (a *App) RunApplication(addr string) {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})

	log.Fatal(http.ListenAndServe(addr, c.Handler(a.Router)))
}
