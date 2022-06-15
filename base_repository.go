package main

type BaseRepositoryAggregation interface {
	Count(groupBy string) (int64, error)
	GroupBy(result interface{}, selectBy string, whereBy map[string]interface{}, groupBy string) error
}

type BaseRepositoryPaging interface {
	//	Limit specify the max number of records to retrieve
	//	Offset specify the number of records to skip before starting to return the records
	Paging(itemPerPage int, page int) (interface{}, error)
}

type BaseRepositoryRaw interface {
	Query(result interface{}, sql string, vals ...interface{}) error
}
