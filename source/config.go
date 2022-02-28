package source

import (
	"fmt"
	"time"

	"github.com/ahamidi/conduit-neo4j-plugin/config"
)

const (
	// ConfigKeyPollingPeriod is the config name for the S3 CDC polling period
	ConfigKeyPollingPeriod = "polling-period"

	// DefaultPollingPeriod is the value assumed for the pooling period when the
	// config omits the polling period parameter
	DefaultPollingPeriod = "1s"

	// The Neo4j cypher query to run
	Query = "query"
)

// Config represents source configuration with S3 configurations
type Config struct {
	config.Config
	PollingPeriod time.Duration
	Query         string
}

// Parse attempts to parse the configurations into a Config struct that Source could utilize
func Parse(cfg map[string]string) (Config, error) {
	common, err := config.Parse(cfg)
	if err != nil {
		return Config{}, err
	}

	pollingPeriodString, exists := cfg[ConfigKeyPollingPeriod]
	if !exists || pollingPeriodString == "" {
		pollingPeriodString = DefaultPollingPeriod
	}
	pollingPeriod, err := time.ParseDuration(pollingPeriodString)
	if err != nil {
		return Config{}, fmt.Errorf(
			"%q config value should be a valid duration",
			ConfigKeyPollingPeriod,
		)
	}
	if pollingPeriod <= 0 {
		return Config{}, fmt.Errorf(
			"%q config value should be positive, got %s",
			ConfigKeyPollingPeriod,
			pollingPeriod,
		)
	}

	if query, exists := cfg[Query]; !exists || query == "" {
		return Config{}, fmt.Errorf("%q config value is required", Query)
	}

	query := cfg[Query]

	sourceConfig := Config{
		Config:        common,
		PollingPeriod: pollingPeriod,
		Query:         query,
	}

	return sourceConfig, nil
}
