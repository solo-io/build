package main

import (
	"context"

	"github.com/solo-io/build/pkg/cli"

	"github.com/solo-io/go-utils/contextutils"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	if err := cli.App(ctx, "vtest").Execute(); err != nil {
		contextutils.LoggerFrom(ctx).Fatalw("error while running build cli", zap.Error(err))
	}

}
