// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package internal // import "github.com/liatrio/liatrio-otel-collector/receiver/gitproviderreceiver/internal"

import (
	"context"

	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

// fake fix 2
type ScraperFactory interface {
	// Create the default configuration for the sub sccraper.
	CreateDefaultConfig() Config
	// Create a scraper based on the configuration passed or return an error if not valid.
	CreateMetricsScraper(ctx context.Context, params receiver.Settings, cfg Config) (scraperhelper.Scraper, error)
}

type Config interface {
}

type ScraperConfig struct {
}
