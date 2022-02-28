package neo4j

import (
	"github.com/ahamidi/conduit-neo4j-plugin/config"
	"github.com/ahamidi/conduit-neo4j-plugin/source"
	sdk "github.com/conduitio/connector-plugin-sdk"
)

type Spec struct{}

// Specification returns the Plugin's Specification.
func Specification() sdk.Specification {
	return sdk.Specification{
		Name:    "neo4j",
		Summary: "A Neo4j source and destination plugin for Conduit, written in Go.",
		Version: "v0.0.1",
		Author:  "Ali Hamidi",
		SourceParams: map[string]sdk.Parameter{
			config.Neo4jURI: {
				Default:     "",
				Required:    true,
				Description: "The URI of the Neo4j database.",
			},
			config.Neo4jUsername: {
				Default:     "",
				Required:    true,
				Description: "The username for the Neo4j database.",
			},
			config.Neo4jPassword: {
				Default:     "",
				Required:    true,
				Description: "The password for the Neo4j database.",
			},
			config.Neo4jRealm: {
				Default:     "",
				Required:    false,
				Description: "The realm for the Neo4j database.",
			},
			source.Query: {
				Default:     "",
				Required:    true,
				Description: "The Cypher query to execute.",
			},
			source.ConfigKeyPollingPeriod: {
				Default:     source.DefaultPollingPeriod,
				Required:    false,
				Description: "polling period for the CDC mode, formatted as a time.Duration string.",
			},
		},
	}
}
