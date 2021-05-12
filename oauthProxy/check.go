package oauthProxy

import (
	"fmt"
)

func (p *Proxy) Check(authData string) (string, error) {
	jwt := p.Converter.FromFrontendToJWT(authData)
	aToken := jwt["AccessToken"].(string)

	checkURL := fmt.Sprintf("%s/%s",p.BaseURL, p.CheckURL)
	headers := map[string]interface{}{
		"Authorization":  fmt.Sprintf("Bearer %s", aToken),
	}

	result, statusCode := p.Client.Get(checkURL, headers)
	if statusCode != 200 {
		return "", fmt.Errorf(result)
	}

	return authData, nil
}
