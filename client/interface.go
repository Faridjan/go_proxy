package client

type NetClientInterface interface {
	Get(URL string, header map[string]interface{}) (result string, statusCode int)
	Post(URL string, body, header map[string]interface{}) (result string, statusCode int)

	Init()
}