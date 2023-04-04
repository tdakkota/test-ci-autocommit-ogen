package main

import (
	"os"

	"github.com/ogen-go/ogen"
	"github.com/ogen-go/ogen/gen"
	"github.com/ogen-go/ogen/gen/genfs"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	file := "openapi.yaml"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	spec, err := ogen.Parse(data)
	if err != nil {
		panic(err)
	}

	g, err := gen.NewGenerator(spec, gen.Options{Logger: logger.Named("gen")})
	if err != nil {
		panic(err)
	}

	dir := "./api"
	if err := os.MkdirAll(dir, 0o750); err != nil {
		panic(err)
	}

	if err := g.WriteSource(genfs.FormattedSource{Root: dir}, "api"); err != nil {
		panic(err)
	}
}
