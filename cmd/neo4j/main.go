package main

import (
	neo4j "github.com/ahamidi/conduit-neo4j-plugin"
	"github.com/ahamidi/conduit-neo4j-plugin/source"
	sdk "github.com/conduitio/connector-plugin-sdk"
)

func main() {
	sdk.Serve(neo4j.Specification, source.NewSource, nil)
}
