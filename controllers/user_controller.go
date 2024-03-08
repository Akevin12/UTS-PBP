package controllers

import (
	"database/sql"
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
			participants.id,
            participants.id_account,
            accounts.username
        FROM 
            rooms
        JOIN 
            participants ON rooms.id = participants.id_room
        JOIN 
            accounts ON participants.id_account = accounts.id
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

	var detaileds []m.DetailRoom

	for rows.Next() {
		var detailed m.DetailRoom
		err := rows.Scan(&detailed.Room.ID, &detailed.Room.RoomName, &detailed.Participants.ID, &detailed.Participants.AccountID, &detailed.Account.Username)
		if err != nil {
			log.Println(err)
			sendErrorResponse(w, "Something went wrong, please try again")
			return
		}
		detaileds = append(detaileds, detailed)
	}

	w.Header().Set("Content-Type", "application/json")
	response := m.DetailRoomResponse{Status: 200, Message: "Success", Data: detaileds}
	json.NewEncoder(w).Encode(response)
}

func InsertRoom(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	roomName := r.URL.Query()["room_name"]
	gameID := r.URL.Query()["game_id"]

	var roomID int
	var maxPlayers int

	// ngecek roomnya
	err := db.QueryRow("SELECT rooms.id, games.max_player FROM Rooms INNER JOIN Games ON games.id = rooms.id_game WHERE rooms.room_name = ? AND games.id = ?", roomName, gameID).Scan(&roomID, &maxPlayers)
	if err != nil {
		if err == sql.ErrNoRows {
			sendErrorResponse(w, "Room or game not found")
			return
		}
		log.Println(err)
		sendErrorResponse(w, "Something went wrong, please try again")
		return
	}

	// check kalo udah penuh roomnya
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM Participants WHERE id_room = ?", roomID).Scan(&count)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Something went wrong, please try again")
		return
	}

	if count >= maxPlayers {
		sendErrorResponse(w, "Maximum number of players reached")
		return
	}

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
