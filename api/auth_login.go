package rocket

const (
	loginPath string = "api/v1/login"
)

func (auth *AuthService) Login(username string, password string) error {

	var response AuthResponse
	_, err := auth.resty.R().
		SetBody(AuthRequest{Username: username, Password: password}).
		SetResult(&response).
		Post(loginPath)

	if err != nil {
		return err
	}

	auth.resty.SetHeader("X-Auth-Token", response.Data.AuthToken)
	auth.resty.SetHeader("X-User-Id", response.Data.UserId)

	return nil
}
