package oauthProxy

import (
	"encoding/json"
	"fmt"
)

func (p *Proxy) Login(username, password string) (string, error) {
	loginURL := p.BaseURL + p.OAuthURL

	dto := make(map[string]interface{})

	dto["grant_type"] = p.OAuthGrantType
	dto["username"] = username
	dto["password"] = password
	dto["domain"] = p.DomainURL
	dto["refresh_token"] = ""

	result, statusCode := p.Client.Post(loginURL, dto, nil)
	if statusCode != 200 {
		return "", fmt.Errorf(result)
	}

	// From JWT To Frontend
	var jwt map[string]interface{}
	json.Unmarshal([]byte(result), &jwt)

	authData := p.Converter.FromJWTToFrontend(jwt)

	return authData, nil

}