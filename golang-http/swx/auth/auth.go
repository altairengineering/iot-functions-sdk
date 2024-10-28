package auth

import (
	"encoding/json"
	"function/swx"
	"github.com/go-resty/resty/v2"
	"strings"
	"time"
)

const (
	tokenEndpoint  = "/oauth2/token"
	revokeEndpoint = "/oauth2/revoke"

	defaultTimeout = 3 * time.Second
)

// Token represents a SmartWorks token.
type Token struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	revokeFunc   func() error
}

// Revoke revokes this Token.
func (t *Token) Revoke() error {
	if t.revokeFunc == nil {
		return &swx.TokenRevokeError{Message: "token cannot be revoked"}
	}
	return t.revokeFunc()
}

// GetToken requests an OAuth 2.0 Bearer token from SmartWorks using the
// client_credentials grant.
func GetToken(clientID, clientSecret string, scopes []string) (*Token, error) {
	client := resty.New()
	client.SetTimeout(defaultTimeout)

	payload := map[string]string{
		"grant_type":    "client_credentials",
		"client_id":     clientID,
		"client_secret": clientSecret,
		"scope":         strings.Join(scopes, " "),
	}

	url := swx.GetApiUrl() + tokenEndpoint

	var token *Token
	resp, err := client.R().
		SetFormData(payload).
		SetResult(&token).
		Post(url)

	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		if resp.StatusCode() >= 500 {
			return nil, unmarshalError(resp.Body(), &swx.ErrorResponse{})
		}
		if resp.StatusCode() >= 400 {
			err = unmarshalError(resp.Body(), &swx.OAuth2Error{
				Err: swx.ResponseStatus{Status: resp.StatusCode()},
			})
			return nil, err
		}
	}

	token.revokeFunc = func() error {
		return RevokeToken(token.AccessToken, clientID, clientSecret)
	}
	return token, nil
}

// RevokeToken revokes the given SmartWork's OAuth2 access token.
func RevokeToken(accessToken, clientID, clientSecret string) error {
	client := resty.New()
	client.SetTimeout(defaultTimeout)

	payload := map[string]string{
		"token":     accessToken,
		"client_id": clientID,
	}

	if clientSecret != "" {
		payload["client_secret"] = clientSecret
	}

	url := swx.GetApiUrl() + revokeEndpoint

	resp, err := client.R().
		SetFormData(payload).
		SetResult(&Token{}).
		SetError(&swx.OAuth2Error{}).
		Post(url)

	if err != nil {
		return err
	}
	if resp.IsError() {
		return resp.Error().(*swx.OAuth2Error)
	}

	return nil
}

func unmarshalError(data []byte, target error) error {
	if err := json.Unmarshal(data, target); err != nil {
		return err
	}
	return target
}
