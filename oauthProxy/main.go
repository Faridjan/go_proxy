package oauthProxy

import (
	"github.com/Faridjan/go_proxy/client"
	"github.com/Faridjan/go_proxy/converter"
)

type Proxy struct {
	BaseURL string
	CheckURL string
	OAuthURL string
	LogoutURL string
	LogoutAllURL string
	DomainURL string

	OAuthGrantType string
	OAuthRefreshGrantType string

	Client client.NetClientInterface
	Converter converter.JwtConverterInterface
}

func (p *Proxy) Init()  {
	p.Client.Init()
}