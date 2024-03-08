package model

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type Games struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	MaxPlayers int    `json:"max_players"`
}

type Rooms struct {
	ID       int    `json:"id"`
	RoomName string `json:"room_name"`
	GameID   int    `json:"game_id"`
}

type Participants struct {
	ID        int `json:"id"`
	RoomID    int `json:"room_id"`
	AccountID int `json:"account_id"`
}

type DetailRoom struct {
	ID           int          `json:"id"`
	Room         Rooms        `json:"room"`
	Participants Participants `json:"participants"`
	Account      Account      `json:"account"`
}

type DetailRoomResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    []DetailRoom `json:"data"`
}

type RoomResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Rooms  `json:"data"`
}

type RoomsResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    []Rooms `json:"data"`
}

// type UsersResponse struct {
// 	Status  int     `json:"status"`
// 	Message string  `json:"message"`
// 	Data    []Users `json:"data"`
// }

type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
