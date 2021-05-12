package converter

import (
	"bytes"
	"fmt"
	"strings"
)

const SEPARATOR = "___AUTO-KZ___"

type Converter struct {
}

func (c *Converter) FromFrontendToJWT(authData string) map[string]interface{} {
	entries  := strings.Split(authData, SEPARATOR)

	m := make(map[string]interface{})

	for _, e := range entries {
		parts := strings.Split(e, "=")
		m[parts[0]] = parts[1]
	}

	return m
}

func (c *Converter) FromJWTToFrontend(jwt map[string]interface{}) string{
	b := new(bytes.Buffer)

	for key, value := range jwt {
		fmt.Fprintf(b, "%s=%s%v", key, value, SEPARATOR)
	}

	return b.String()
}