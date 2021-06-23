package routes

import (
	"github.com/gorilla/mux"

	"github.com/netyut1703/shopingProject/server"
)

func InitRoutes(s *server.Server) *mux.Router {

	r := mux.NewRouter()

	// arrange our route
	r.Handle("/api/v1/info", NoAuthHandler(s, s.ApiInfo)).Methods("GET")

	r.Handle("/api/v1/health-check", NoAuthHandler(s, s.GetHealth)).Methods("GET")
	// user routes
	r.Handle("/api/v1/user/buyer/signup", NoAuthHandler(s, s.Signup)).Methods("POST")
	r.Handle("/api/v1/user/buyer/verifyOtp", NoAuthHandler(s, s.BuyerMailVerifyOtp)).Methods("POST") //

}
