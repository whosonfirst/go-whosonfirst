package export

import (
	"context"
	"fmt"
	"net/url"
	"sort"
	"strings"

	"github.com/aaronland/go-roster"
)

// Exporter is the interface for customizing the export of Who's On First records.
type Exporter interface {
	// Export will perform all the steps necessary to "export" (as in create or update) a Who's On First feature record taking care to ensure correct formatting, default values and validation. It returns a boolean value indicating whether the feature was changed during the export process.
	Export(context.Context, []byte) (bool, []byte, error)
}

var exporter_roster roster.Roster

// ExporterInitializationFunc constructs an [Exporter] from a configuration URI.
type ExporterInitializationFunc func(ctx context.Context, uri string) (Exporter, error)

// RegisterExporter associates a URI scheme with a constructor used by [NewExporter].
func RegisterExporter(ctx context.Context, scheme string, init_func ExporterInitializationFunc) error {

	err := ensureExporterRoster()

	if err != nil {
		return err
	}

	return exporter_roster.Register(ctx, scheme, init_func)
}

func ensureExporterRoster() error {

	if exporter_roster == nil {

		r, err := roster.NewDefaultRoster()

		if err != nil {
			return err
		}

		exporter_roster = r
	}

	return nil
}

// NewExporter returns a new `Exporter` instance derived from 'uri'.
func NewExporter(ctx context.Context, uri string) (Exporter, error) {

	u, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	scheme := u.Scheme

	i, err := exporter_roster.Driver(ctx, scheme)

	if err != nil {
		return nil, err
	}

	init_func := i.(ExporterInitializationFunc)
	return init_func(ctx, uri)
}

// ExporterSchemes returns list of registered `Exporter` schemes which have been registered.
func ExporterSchemes() []string {

	ctx := context.Background()
	schemes := []string{}

	err := ensureExporterRoster()

	if err != nil {
		return schemes
	}

	for _, dr := range exporter_roster.Drivers(ctx) {
		scheme := fmt.Sprintf("%s://", strings.ToLower(dr))
		schemes = append(schemes, scheme)
	}

	sort.Strings(schemes)
	return schemes
}
