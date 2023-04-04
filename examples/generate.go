package examples

import (
	_ "github.com/go-faster/errors"

	_ "github.com/tdakkota/test-ci-autocommit-ogen"
)

// Fully supported:
//
//go:generate go run github.com/tdakkota/test-ci-autocommit-ogen/cmd/tool ../openapi.yaml
