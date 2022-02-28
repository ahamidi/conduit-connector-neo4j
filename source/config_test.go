package source

import (
	"testing"
	"time"
)

var exampleConfig = map[string]string{
	"uri":      "https://example.com",
	"username": "some-user",
	"password": "some-password",
	"realm":    "some-realm",
	"query":    "some-query",
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

func TestPollingPeriod(t *testing.T) {
	c, err := Parse(configWith("polling-period", "5s"))

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if c.PollingPeriod != 5*time.Second {
		t.Fatalf("expected Polling Period to be %q, got %q", "5s", c.PollingPeriod)
	}
}
