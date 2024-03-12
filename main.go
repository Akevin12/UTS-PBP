package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/modul2/controllers"
)

func main() {
	router := mux.NewRouter()

	//USER
	// router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	// // router.HandleFunc("/users", controllers.InsertUser).Methods("POST")
	// router.HandleFunc("/users", controllers.InsertUserGorm).Methods("POST")
	// router.HandleFunc("/users/{user_id}", controllers.DeleteUser).Methods("DELETE")
	// router.HandleFunc("/users", controllers.UpdateUser).Methods("PUT")

	//Room
	router.HandleFunc("/getAllRoom", controllers.GetAllRoom).Methods("GET")
	router.HandleFunc("/getRoomDetail", controllers.GetRoomDetail).Methods("GET")
	router.HandleFunc("/insertRoom", controllers.InsertRoom).Methods("POST")
	router.HandleFunc("/deleteRoom/{room_id}", controllers.DeleteRoom).Methods("DELETE")

	http.Handle("/", router)
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}
