package oauthProxy

import (
	"encoding/json"
	"fmt"
)

func (p *Proxy) Refresh(authData string) (string, error) {
	jwt := p.Converter.FromFrontendToJWT(authData)
	rToken := jwt["RefreshToken"].(string)

	loginURL :=  p.BaseURL + p.OAuthURL

	dto := make(map[string]interface{})

	dto["grant_type"] = p.OAuthGrantType
	dto["username"] = ""
	dto["password"] = ""
	dto["domain"] = p.DomainURL
	dto["refresh_token"] = rToken

	result, statusCode := p.Client.Post(loginURL, dto, nil)
	if statusCode != 200 {
		return "", fmt.Errorf(result)
	}

	// From JWT To Frontend
	var newJWT map[string]interface{}
	json.Unmarshal([]byte(result), &newJWT)

	newAuthData := p.Converter.FromJWTToFrontend(newJWT)

	return newAuthData, nil
}