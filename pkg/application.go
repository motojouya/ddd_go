package pkg

import (
	"github.com/alecthomas/kong"
)

var app struct {
	Serve     ServeCmd     `cmd:"" help:"start server"`
	Work      WorkCmd      `cmd:"" help:"start job worker"`
	Aggregate AggregateCmd `embed:""`
}

func Execute() {
	ctx := kong.Parse(&app)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
