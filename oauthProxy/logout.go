package oauthProxy

import (
	"fmt"
)

func (p *Proxy) Logout(authData string) (bool, error) {
	jwt := p.Converter.FromFrontendToJWT(authData)
	aToken := jwt["AccessToken"].(string)

	logoutURL := p.BaseURL + p.LogoutURL

	headers := map[string]interface{}{
		"Authorization":  fmt.Sprintf("Bearer %s", aToken),
	}

	result, statusCode := p.Client.Post(logoutURL, nil, headers)
	if statusCode != 200 {
		return false, fmt.Errorf(result)
	}

	return true, nil
}