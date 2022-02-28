package source

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/ahamidi/conduit-neo4j-plugin/config"

	sdk "github.com/conduitio/connector-plugin-sdk"
)

func TestSource_Lifecycle(t *testing.T) {
	query := `
		MATCH (m:Movie)
		WHERE m.released > 2000
		RETURN m`

	cfg := map[string]string{
		config.Neo4jURI:        os.Getenv("NEO4J_URI"),
		config.Neo4jUsername:   os.Getenv("NEO4J_USERNAME"),
		config.Neo4jPassword:   os.Getenv("NEO4J_PASSWORD"),
		Query:                  query,
		ConfigKeyPollingPeriod: "100ms",
	}

	ctx := context.Background()
	source := &Source{}
	err := source.Configure(context.Background(), cfg)
	if err != nil {
		t.Fatal(err)
	}
	err = source.Open(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	r, err := source.Read(ctx)
	if err != nil && err.Error() != sdk.ErrBackoffRetry.Error() {
		t.Fatalf("expected a BackoffRetry error, got: %v", err)
	}

	var movie map[string]interface{}
	err = json.Unmarshal(r.Payload.Bytes(), &movie)
	if err != nil {
		t.Fatalf("expected a no error, got: %v", err)
	}

	if label := movie["labels"].([]interface{})[0].(string); label != "Movie" {
		t.Fatalf("expected blah, got: %s", label)
	}

	err = source.Teardown(ctx)
	if err != nil {
		t.Fatalf("expected a no error, got: %v", err)
	}
}
