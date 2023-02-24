package wql_lib

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type FilterInput struct {
	FilterMeta    map[string]string `json:"filterMeta" form:"filterMeta"`
	QueryMeta     map[string]string `json:"queryMeta" form:"queryMeta"`
	OrderMeta     map[string]string `json:"orderMeta" form:"orderMeta"`
	AttributeMeta map[string]string `json:"attributeMeta" form:"attributeMeta"`
	IncludeMeta   map[string]string `json:"includeMeta" form:"includeMeta"`
}

func (f *FilterInput) Scan(ctx *gin.Context) (error, *FilterInput) {
	if filterMeta, ok := ctx.GetQueryMap("filterMeta"); ok {
		f.FilterMeta = filterMeta
	} else {
		return errors.New("\"filterMeta\" is required"), f
	}
	if queryMeta, ok := ctx.GetQueryMap("queryMeta"); ok {
		f.QueryMeta = queryMeta
	} else {
		f.QueryMeta = map[string]string{"page": "1", "count": "50"}
	}
	if orderMeta, ok := ctx.GetQueryMap("orderMeta"); ok {
		f.OrderMeta = orderMeta
	} else {
		//return errors.New("\"queryMeta\" is required"), f
	}
	if attributeMeta, ok := ctx.GetQueryMap("attributeMeta"); ok {
		f.AttributeMeta = attributeMeta
	} else {
		//return errors.New("\"attributeMeta\" is required"), f
	}
	if includeMeta, ok := ctx.GetQueryMap("includeMeta"); ok {
		f.IncludeMeta = includeMeta
	} else {
		//return errors.New("\"includeMeta\" is required"), f
	}
	return nil, f
}
