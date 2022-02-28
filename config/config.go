package config

import (
	"fmt"
)

const (
	Neo4jURI      = "uri"
	Neo4jUsername = "username"
	Neo4jPassword = "password"
	Neo4jRealm    = "realm"
)

// Config represents configuration needed for S3
type Config struct {
	URI      string
	Username string
	Password string
	Realm    string
}

// Parse attempts to parse plugins.Config into a Config struct
func Parse(cfg map[string]string) (Config, error) {
	uri, ok := cfg[Neo4jURI]

	if !ok {
		return Config{}, requiredConfigErr(Neo4jURI)
	}

	username, ok := cfg[Neo4jUsername]

	if !ok {
		return Config{}, requiredConfigErr(Neo4jUsername)
	}

	password, ok := cfg[Neo4jPassword]

	if !ok {
		return Config{}, requiredConfigErr(Neo4jPassword)
	}

	realm, ok := cfg[Neo4jRealm]

	config := Config{
		URI:      uri,
		Username: username,
		Password: password,
		Realm:    realm,
	}

	return config, nil
}

func requiredConfigErr(name string) error {
	return fmt.Errorf("%q config value must be set", name)
}
