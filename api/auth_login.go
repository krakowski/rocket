package rocket

import (
	"fmt"
	"net/http"
)

const (
	loginPath string = "api/v1/login"
)

func (auth *AuthService) Login(username string, password string) error {

	var response AuthResponse
	resp, err := auth.resty.R().
		SetBody(AuthRequest{Username: username, Password: password}).
		SetResult(&response).
		SetError(&LastError).
		Post(loginPath)

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("Authentication failed : %s", resp)
	}

	auth.resty.SetHeader("X-Auth-Token", response.Data.AuthToken)
	auth.resty.SetHeader("X-User-Id", response.Data.UserId)

	return nil
}
