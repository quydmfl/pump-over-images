package main

import (
	"context"

	"freelancer.com/pump-over-images/fetcher/internal/tank"
	"freelancer.com/pump-over-images/fetcher/pkg/fetch"
)

func main() {
	ctx := context.Background()
	fetcher := fetch.NewFetch()
	_ = tank.NewTank(ctx, *fetcher)

	// go-routine in here //
}
