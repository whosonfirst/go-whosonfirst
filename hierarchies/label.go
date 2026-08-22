package hierarchies

import (
	"fmt"
	"log/slog"
)

type AncestorIdsForLabelOptions struct {
	Hierarchies []map[string]int64
	Placetype   string
	ParentId    int64
}

// Id may be an int64 or a string – both are acceptable for the map values.
type Id interface{ ~int64 | ~string }

// ──────────────────────────────────────────────────────────────────────────────
//
//	Generic options struct
//
// ──────────────────────────────────────────────────────────────────────────────
type AncestorIdsForLabelOptionsGeneric[T Id] struct {
	Hierarchies []map[string]T // one or more hierarchy maps
	Placetype   string
	ParentId    T // must be the same type that is used in the hierarchy maps
}

// ──────────────────────────────────────────────────────────────────────────────
//
//	Generic merge‑common‑hierarchies helper
//
// ──────────────────────────────────────────────────────────────────────────────
func MergeCommonHierarchiesGeneric[T Id](hierarchies []map[string]T) map[string]T {
	if len(hierarchies) == 0 {
		return nil
	}

	// Keep only the keys that appear in *every* map and that have the same value.
	merged := make(map[string]T)
	for k, v := range hierarchies[0] {
		common := true
		for _, h := range hierarchies[1:] {
			if val, ok := h[k]; !ok || val != v {
				common = false
				break
			}
		}
		if common {
			merged[k] = v
		}
	}
	return merged
}

// ──────────────────────────────────────────────────────────────────────────────
//
//	Generic AncestorIdsForLabel implementation
//
// ──────────────────────────────────────────────────────────────────────────────
func AncestorIdsForLabelGeneric[T Id](opts *AncestorIdsForLabelOptionsGeneric[T]) []T {
	// The result slice – the same type that is stored in the hierarchy maps.
	nameIDs := make([]T, 0)

	// “continent”, “empire” and “country” are leaf nodes – they have no
	// ancestors to return, so we simply return an empty slice.
	switch opts.Placetype {
	case "continent", "empire", "country":
		// nothing to do
	default:

		// --------------------------------------------------------------------
		//  Build the ancestor lineage for the requested placetype.
		// --------------------------------------------------------------------
		var lineage []string
		switch opts.Placetype {
		case "macroregion", "region":
			lineage = []string{"country"}
		case "macrocounty", "county":
			lineage = []string{"region", "country"}
		case "metroarea", "localadmin", "locality":
			lineage = []string{"region", "country"}
		case "borough", "campus", "postalcode":
			lineage = []string{"localadmin", "locality", "region", "country"}
		case "macrohood", "neighbourhood":
			lineage = []string{"localadmin", "locality", "region", "country"}
		case "microhood":
			lineage = []string{"localadmin", "locality", "region", "country"}
		case "venue":
			lineage = []string{"neighbourhood", "localadmin", "locality", "region", "country"}
		case "building":
			lineage = []string{"campus", "locality", "region", "country"}
		case "wing":
			lineage = []string{"building", "campus", "locality", "country"}
		case "concourse":
			lineage = []string{"wing", "building", "campus", "locality", "country"}
		case "arcade":
			lineage = []string{"concourse", "wing", "building", "campus", "country"}
		case "enclosure":
			lineage = []string{"arcade", "concourse", "wing", "campus", "country"}
		default:
			slog.Warn("Unsupported placetype", "placetype", opts.Placetype)
		}

		// --------------------------------------------------------------------
		//  Pick the “principal” hierarchy map:
		//  • If only one hierarchy is supplied → that one.
		//  • If a ParentId is supplied → the hierarchy that contains that id.
		//  • Otherwise → merge the common keys of all hierarchies.
		// --------------------------------------------------------------------
		var useHier map[string]T

		switch len(opts.Hierarchies) {
		case 0: // nothing – we’ll warn later
		case 1:
			useHier = opts.Hierarchies[0]
		default:
			// If ParentId has a meaningful value, try to find the hierarchy that
			// contains that id.  The comparison works for both int64 and string.
			for _, h := range opts.Hierarchies {
				for _, hid := range h {
					if hid == opts.ParentId {
						useHier = h
						break
					}
				}
				if useHier != nil {
					break
				}
			}

			if useHier == nil {
				// No hierarchy contains ParentId – fall back to the merged view.
				useHier = MergeCommonHierarchiesGeneric(opts.Hierarchies)
			}
		}

		// If we still have no hierarchy, default to the first one – the caller
		// should be aware that this may be incorrect.
		if useHier == nil {
			slog.Warn("Failed to determine principal (or merged) hierarchy, falling back to the first one")
			if len(opts.Hierarchies) > 0 {
				useHier = opts.Hierarchies[0]
			}
		}

		// --------------------------------------------------------------------
		//  Walk the lineage, pulling the id for each ancestor from the chosen
		//  hierarchy map.
		// --------------------------------------------------------------------
		for _, pt := range lineage {
			hierKey := fmt.Sprintf("%s_id", pt)
			if nameID, exists := useHier[hierKey]; exists {
				nameIDs = append(nameIDs, nameID)
			}
		}
	}

	return nameIDs
}

func AncestorIdsForLabel(opts *AncestorIdsForLabelOptions) []int64 {

	genericOpts := &AncestorIdsForLabelOptionsGeneric[int64]{
		Hierarchies: opts.Hierarchies,
		Placetype:   opts.Placetype,
		ParentId:    opts.ParentId,
	}

	return AncestorIdsForLabelGeneric[int64](genericOpts)
}

func x_AncestorIdsForLabel(opts *AncestorIdsForLabelOptions) []int64 {

	name_ids := make([]int64, 0)

	switch opts.Placetype {
	case "continent", "empire", "country":
		// pass
	default:

		var lineage []string

		switch opts.Placetype {
		case "macroregion", "region":
			lineage = []string{
				"country",
			}
		case "macrocounty", "county":
			lineage = []string{
				"region",
				"country",
			}
		case "metroarea", "localadmin", "locality":

			lineage = []string{
				"region",
				"country",
			}

		case "borough", "campus", "postalcode":

			lineage = []string{
				"localadmin",
				"locality",
				"region",
				"country",
			}

		case "macrohood", "neighbourhood":

			lineage = []string{
				"localadmin",
				"locality",
				"region",
				"country",
			}

		case "microhood":

			lineage = []string{
				// "neighbourhood",
				"localadmin",
				"locality",
				"region",
				"country",
			}

		case "venue":

			lineage = []string{
				"neighbourhood",
				"localadmin",
				"locality",
				"region",
				"country",
			}

		case "building":

			lineage = []string{
				"campus",
				"locality",
				"region",
				"country",
			}

		case "wing":

			lineage = []string{
				"building",
				"campus",
				"locality",
				// "region",
				"country",
			}

		case "concourse":

			lineage = []string{
				"wing",
				"building",
				"campus",
				"locality",
				// "region",
				"country",
			}

		case "arcade":

			lineage = []string{
				"concourse",
				"wing",
				"building",
				"campus",
				// "locality",
				// "region",
				"country",
			}

		case "enclosure":

			lineage = []string{
				"arcade",
				"concourse",
				"wing",
				// "building",
				"campus",
				// "locality",
				// "region",
				"country",
			}

		default:
			slog.Warn("Unsupported placetype", "placetype", opts.Placetype)
		}

		var use_hier map[string]int64

		switch len(opts.Hierarchies) {
		case 0:
			// pass - len(hiers) checked above
		case 1:
			use_hier = opts.Hierarchies[0]
		default:

			switch opts.ParentId >= 0 {
			case true:

				for _, h := range opts.Hierarchies {

					for _, hid := range h {

						if hid == opts.ParentId {
							use_hier = h
							break
						}
					}

					if use_hier != nil {
						break
					}
				}

			default:
				use_hier = MergeCommonHierarchies(opts.Hierarchies)
			}
		}

		if use_hier == nil {
			slog.Warn("Failed to determine principal (or merged) hierarchy, simply pulling the first one (which may be incorrect)")
			use_hier = opts.Hierarchies[0]
		}

		for _, pt := range lineage {

			hier_k := fmt.Sprintf("%s_id", pt)
			name_id, exists := use_hier[hier_k]

			if exists {
				name_ids = append(name_ids, name_id)
			}
		}
	}

	return name_ids
}
