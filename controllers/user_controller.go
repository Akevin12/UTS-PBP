package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	m "github.com/modul2/model"
)

// GET, INSERT, UPDATE, DELETE USERS
func GetAllRoom(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	query := "SELECT rooms.id, rooms.room_name FROM rooms"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Something went wrong, please try again")
		return
	}

	var room m.Rooms
	var rooms []m.Rooms
	for rows.Next() {
		if err := rows.Scan(&room.ID, &room.RoomName); err != nil {
			log.Println(err)
			sendErrorResponse(w, "Something went wrong, please try again")
			return
		} else {
			rooms = append(rooms, room)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	var response m.RoomsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = rooms
	json.NewEncoder(w).Encode(response)

}

func GetRoomDetail(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	query := `
        SELECT 
            rooms.id,
            rooms.room_name,
            participants.id AS participant_id,
            participants.id_account,
            accounts.username
        FROM 
            rooms
        JOIN 
            participants ON rooms.id = participants.id_room
        JOIN 
            accounts ON participants.id_account = id_account
    `

	id := r.URL.Query().Get("id")
	if id != "" {
		query += " WHERE rooms.id='" + id + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Something went wrong, please try again")
		return
	}

	var detailed m.DetailRoom
	var detaileds []m.DetailRoom

	for rows.Next() {
		if err := rows.Scan(&detailed.ID, &detailed.Room.RoomName, &detailed.Participants.ID, &detailed.Participants.AccountID, &detailed.Account.Username); err != nil {
			log.Println(err)
			sendErrorResponse(w, "Something went wrong, please try again")
			return
		} else {
			detaileds = append(detaileds, detailed)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	var response m.DetailRoomResponse
	response.Status = 200
	response.Message = "Success"
	response.Data = detaileds
	json.NewEncoder(w).Encode(response)
}

func InsertRoom(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	roomName := r.URL.Query()["room_name"]
	gameID := r.URL.Query()["game_id"]

	_, errQuery := db.Exec("INSERT INTO rooms(room_name, game_id)values (?,?)",
		roomName,
		gameID,
	)
	var response m.RoomsResponse

	if errQuery == nil {
		response.Status = 200
		response.Message = "Success"
	} else {
		response.Status = 400
		response.Message = "Insert Failed"
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendSuccessRespon(w http.ResponseWriter) {
	var response m.RoomsResponse
	response.Status = 200
	response.Message = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendErrorResponse(w http.ResponseWriter, message string) {
	var response m.ErrorResponse
	response.Status = 400
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
