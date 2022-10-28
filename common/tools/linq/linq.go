package linq

import (
	linq "github.com/ahmetb/go-linq/v3"
)

type customQuery struct {
	linq.Query
}

func From(source interface{}) customQuery {
	return customQuery{
		linq.From(source),
	}
}

func (q customQuery) FirstTOrDefault(predicateFn interface{}, defaultValue interface{}) interface{} {
	first := q.FirstWithT(predicateFn)
	if first == nil {
		return defaultValue
	}
	return first
}

func (q customQuery) IndexOfTOrDefault(predicateFn interface{}, defaultIndex int) int {
	index := q.IndexOfT(predicateFn)
	if index < 0 {
		return defaultIndex
	}
	return index
}
