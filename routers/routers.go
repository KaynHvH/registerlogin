package routers

import (
	"github.com/gorilla/mux"
	"registerlogin/auth"
	"registerlogin/middleware"
)

func InitRouters(router *mux.Router) {
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.AuthenticationMiddleware)

	router.HandleFunc("/register", auth.RegisterHandler).Methods("POST")
	router.HandleFunc("/login", auth.LoginHandler).Methods("POST")
	router.HandleFunc("/deleteacc", auth.DeleteAccountHandler).Methods("DELETE")
	router.HandleFunc("/changepass", auth.ChangePasswordHandler).Methods("PUT")
}
