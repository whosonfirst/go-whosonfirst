package parquet

type Record struct {
	Id         int64 `parquet:"id,type=INT64"`
	Geometry   []byte `parquet:"geometry,type=BYTE_ARRAY,encoding=PLAIN"`
	Properties string `parquet:"properties,type=BYTE_ARRAY,encoding=JSON"`
}