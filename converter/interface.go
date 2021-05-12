package converter

type JwtConverterInterface interface {
	FromFrontendToJWT(authData string) map[string]interface{}
	FromJWTToFrontend(map[string]interface{}) string
}