package parquet

type Record struct {
	Id         int64  `parquet:"id,type=INT64"`
	ParentId   int64  `parquet:"parent_id,type="INT64"`
	Placetype  string `parquet:"placetype,type=dict,zstd"`
	Country    string `parquet:"country,type=dict,zstd"`
	Geometry   []byte `parquet:"geometry,geometry"`
	Properties []byte `parquet:"properties,json"`
}

/*

memory D SELECT id, parent_id, placetype, properties."wof:name" AS name, ST_Centroid(geometry) FROM read_parquet('test.parquet') WHERE properties->>'$.sfomuseum:placetype'='airport' AND country = 'CA' LIMIT 5;
┌───────────┬────────────┬───────────┬────────────────────────────────────┬─────────────────────────────────────┐
│    id     │ parent_id  │ placetype │                name                │        st_centroid(geometry)        │
│   int64   │   int64    │  varchar  │                json                │        geometry('ogc:crs84')        │
├───────────┼────────────┼───────────┼────────────────────────────────────┼─────────────────────────────────────┤
│ 102554351 │  101736475 │ campus    │ "Montreal-Pierre Elliott Trudeau I │ POINT (-73.7436343824761 45.4673683 │
│           │            │           │ nternational Airport"              │ 15411896)                           │
├───────────┼────────────┼───────────┼────────────────────────────────────┼─────────────────────────────────────┤
│ 102554473 │         -1 │ campus    │ "Edmonton International Airport"   │ POINT (-113.58489936246232 53.30860 │
│           │            │           │                                    │ 794988534)                          │
├───────────┼────────────┼───────────┼────────────────────────────────────┼─────────────────────────────────────┤
│ 102554519 │   85682123 │ campus    │ "Goose Bay Airport"                │ POINT (-60.425833 53.319167)        │
├───────────┼────────────┼───────────┼────────────────────────────────────┼─────────────────────────────────────┤
│ 102554947 │ 1108964229 │ campus    │ "Calgary International Airport"    │ POINT (-114.00691541254936 51.12049 │
│           │            │           │                                    │ 9437034056)                         │
├───────────┼────────────┼───────────┼────────────────────────────────────┼─────────────────────────────────────┤
│ 102555091 │ 1108905413 │ campus    │ "Toronto Lester B Pearson Internat │ POINT (-79.63120282896976 43.676962 │
│           │            │           │ ional Airport"                     │ 210845375)                          │
└───────────┴────────────┴───────────┴────────────────────────────────────┴─────────────────────────────────────┘

*/
