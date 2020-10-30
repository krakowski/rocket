package rocket

type AuthService service

type AuthRequest struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Status string   `json:"status"`
	Data   AuthData `json:"data"`
}

type AuthData struct {
	AuthToken string `json:"authToken"`
	UserId    string `json:"userId"`
}
