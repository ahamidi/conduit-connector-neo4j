package source

import (
	"context"
	"encoding/json"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j/dbtype"
	"log"
	"strconv"
	"time"

	sdk "github.com/conduitio/connector-plugin-sdk"
)

// Source connector
type Source struct {
	sdk.UnimplementedSource
	config  Config
	client  neo4j.Driver
	session neo4j.Session
}

func NewSource() sdk.Source {
	return &Source{}
}

// Configure parses and stores the configurations
// returns an error in case of invalid config
func (s *Source) Configure(ctx context.Context, cfg map[string]string) error {
	config2, err := Parse(cfg)
	if err != nil {
		return err
	}

	s.config = config2

	return nil
}

// Open prepare the plugin to start sending records from the given position
func (s *Source) Open(ctx context.Context, rp sdk.Position) error {
	c, err := neo4j.NewDriver(s.config.URI, neo4j.BasicAuth(s.config.Username, s.config.Password, s.config.Realm))
	if err != nil {
		return err
	}

	s.client = c
	s.session = s.client.NewSession(neo4j.SessionConfig{})
	return nil
}

// Read gets the next object from Neo4j
func (s *Source) Read(ctx context.Context) (sdk.Record, error) {
	r, err := s.session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		result, err := tx.Run(s.config.Query, nil)
		if err != nil {
			return nil, err
		}

		for result.Next() {
			for _, v := range result.Record().Values {
				n := v.(dbtype.Node)
				return nodeToRecord(n)
			}
			log.Println("outer loop")
		}
		return nil, result.Err()
	})

	rr := r.(*sdk.Record)
	return *rr, err
}

// Teardown is called when the connector is stopped
func (s *Source) Teardown(ctx context.Context) error {
	if s.client != nil {
		s.client.Close()
	}
	if s.session != nil {
		s.session.Close()
	}
	return nil
}

// Ack ...
func (s *Source) Ack(ctx context.Context, position sdk.Position) error {
	return nil // no ack needed
}

func nodeToRecord(n dbtype.Node) (*sdk.Record, error) {
	pos := sdk.Position(strconv.Itoa(int(n.Id)))
	key := sdk.RawData(strconv.Itoa(int(n.Id)))
	payload := map[string]interface{}{
		"labels": n.Labels,
		"props":  n.Props,
	}
	jPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	r := &sdk.Record{
		Position:  pos,
		Metadata:  nil,
		CreatedAt: time.Time{},
		Key:       key,
		Payload:   sdk.RawData(jPayload),
	}
	return r, nil
}
