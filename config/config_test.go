package config

var exampleConfig = map[string]string{
	"uri":      "https://example.com",
	"username": "some-user",
	"password": "some-password",
	"realm":    "some-realm",
}

func configWith(pairs ...string) map[string]string {
	cfg := make(map[string]string)

	for key, value := range exampleConfig {
		cfg[key] = value
	}

	for i := 0; i < len(pairs); i += 2 {
		key := pairs[i]
		value := pairs[i+1]
		cfg[key] = value
	}

	return cfg
}

func configWithout(keys ...string) map[string]string {
	cfg := make(map[string]string)

	for key, value := range exampleConfig {
		cfg[key] = value
	}

	for _, key := range keys {
		delete(cfg, key)
	}

	return cfg
}
