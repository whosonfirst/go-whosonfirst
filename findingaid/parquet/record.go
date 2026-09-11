package parquet

type ParquetRecord struct {
	Id   int64  `parquet:"id"`
	Repo string `parquet:"repo"`
}
