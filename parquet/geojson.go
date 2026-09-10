package parquet

import (
	"encoding/json"
	"fmt"
	"io"

	_ "github.com/whosonfirst/go-whosonfirst/v4/geojson"

	"github.com/paulmach/orb/encoding/wkb"
	"github.com/paulmach/orb/geojson"
)

func RecordsAsFeatureCollection(records []*Record) (*geojson.FeatureCollection, error) {

	features := make([]*geojson.Feature, len(records))

	for i, r := range records {

		f, err := r.AsGeoJSON()

		if err != nil {
			return nil, err
		}

		features[i] = f
	}

	fc := geojson.NewFeatureCollection()
	fc.Features = features

	return fc, nil
}

func RecordFromGeoJSONReader(r io.Reader) (*Record, error) {

	body, err := io.ReadAll(r)

	if err != nil {
		return nil, err
	}

	return RecordFromGeoJSONBytes(body)
}

func RecordFromGeoJSONBytes(body []byte) (*Record, error) {

	f, err := geojson.UnmarshalFeature(body)

	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal feature, %w", err)
	}

	return RecordFromGeoJSONFeature(f)
}

func RecordFromGeoJSONFeature(f *geojson.Feature) (*Record, error) {

	geom, err := wkb.Marshal(f.Geometry, wkb.DefaultByteOrder)

	if err != nil {
		return nil, fmt.Errorf("Failed to marshal geometry, %w", err)
	}

	id_fl64 := f.Properties.MustFloat64("wof:id", -1)

	if id_fl64 == -1 {
		return nil, fmt.Errorf("Failed to determine wof:id")
	}

	pid_fl64 := f.Properties.MustFloat64("wof:parent_id", -1)

	id := int64(id_fl64)
	pid := int64(pid_fl64)

	pt := f.Properties.MustString("wof:placetype", "custom")
	co := f.Properties.MustString("wof:country", "XY")

	props, err := json.Marshal(f.Properties)

	record := &Record{
		Id:         id,
		AltLabel:   "",
		ParentId:   pid,
		Placetype:  pt,
		Country:    co,
		Geometry:   geom,
		Properties: props,
	}

	// See notes in feature/alt.IsAlt

	allowed_alt_labels := []string{
		"src:alt_label",
		"wof:alt_label",
	}

	for _, prop := range allowed_alt_labels {

		label := f.Properties.MustString(prop, "")

		if label != "" {
			record.AltLabel = label
			break
		}
	}

	return record, nil
}
