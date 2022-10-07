package presentation

import (
	"fmt"
	"log"
	"net/http"

	"github.com/flightlogteam/userservice/src/device"
	"github.com/flightlogteam/userservice/src/user"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/klyngen/jsend"
)

type Api struct {
	userService   user.Service
	deviceService device.FlyingDeviceService
	router        *mux.Router
}

func NewApi(userService user.Service, deviceService device.FlyingDeviceService) Api {
	router := mux.NewRouter()

	api := Api{
		userService:   userService,
		deviceService: deviceService,
		router:        router,
	}

	api.mountRoutes(router)
	return api
}

func (a *Api) Start(port string) {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	log.Printf("Starting server on port: %v", port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(headersOk, originsOk, methodsOk)(a.router)))
}

func (a *Api) mountRoutes(router *mux.Router) {
	router.HandleFunc("/device/{userId}", a.handleGetDevicesByUser).Methods("GET")
	router.HandleFunc("/device", a.handlePostDevice).Methods("POST")
	router.HandleFunc("/device", a.handlePutDevice).Methods("PUT")
}

func (a *Api) handleGetDevicesByUser(w http.ResponseWriter, r *http.Request) {
	// TODO Get the userID from the token
	userId := mux.Vars(r)["userId"]

	fmt.Println(r.Header)

	if len(userId) == 0 {
		jsend.FormatResponse(w, "No user-id present. should be device/{userId}", jsend.BadRequest)
		return
	}

	user, err := a.userService.UserById(userId)

	if err != nil {
		jsend.FormatResponse(w, "Unable to fetch the user by that ID", jsend.InternalServerError)
		return
	}

	if user == nil {
		jsend.FormatResponse(w, "No such user exists", jsend.BadRequest)
		return
	}

	devices, err := a.deviceService.GetDeviceByUser(userId)

	if err != nil {
		jsend.FormatResponse(w, "Unable to fetch the user by that ID", jsend.InternalServerError)
	}

	jsend.FormatResponse(w, devices, jsend.Success)

}

func (a *Api) handlePostDevice(w http.ResponseWriter, r *http.Request) {
	// TODO Get the userID from the token
}

func (a *Api) handlePutDevice(w http.ResponseWriter, r *http.Request) {
	// TODO Get the userID from the token
}
