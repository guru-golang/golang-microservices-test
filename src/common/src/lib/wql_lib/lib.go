package wql_lib

import (
	"car-rent-platform/backend/common/src/lib/assertion_lib"
	"car-rent-platform/backend/common/src/lib/net_lib"
	"car-rent-platform/backend/common/src/lib/reflect_lib"
	"errors"
	"github.com/gin-gonic/gin"
)

type FilterInput struct {
	FilterMeta    map[string]string `json:"filterMeta" form:"filterMeta" validate:"required"`
	QueryMeta     map[string]string `json:"queryMeta" form:"queryMeta"`
	OrderMeta     map[string]string `json:"orderMeta" form:"orderMeta"`
	AttributeMeta map[string]string `json:"attributeMeta" form:"attributeMeta"`
	IncludeMeta   map[string]string `json:"includeMeta" form:"includeMeta"`
}

func (f *FilterInput) GinScan(ctx *gin.Context) (*FilterInput, error) {
	if filterMeta, ok := ctx.GetQueryMap("FilterMeta"); ok {
		f.FilterMeta = filterMeta
	} else {
		return f, errors.New("\"filterMeta\" is required")
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
	return f, nil
}
func (f *FilterInput) NetScan(ctx *net_lib.Context) (*FilterInput, error) {
	data := ctx.Msg.Data()
	swap := data.(map[string]any)

	if filterMeta, err := reflect_lib.PropMap(swap, "filterMeta"); err == nil {
		f.FilterMeta = assertion_lib.MapAnyToString(filterMeta)
	} else {
		return f, errors.New("\"filterMeta\" is required")
	}
	if queryMeta, err := reflect_lib.PropMap(swap, "queryMeta"); err == nil {
		f.QueryMeta = assertion_lib.MapAnyToString(queryMeta)
	} else {
		f.QueryMeta = map[string]string{"page": "1", "count": "50"}
	}
	if orderMeta, err := reflect_lib.PropMap(swap, "orderMeta"); err == nil {
		f.OrderMeta = assertion_lib.MapAnyToString(orderMeta)
	} else {
		//return errors.New("\"queryMeta\" is required"), f
	}
	if attributeMeta, err := reflect_lib.PropMap(swap, "attributeMeta"); err == nil {
		f.AttributeMeta = assertion_lib.MapAnyToString(attributeMeta)
	} else {
		//return errors.New("\"attributeMeta\" is required"), f
	}
	if includeMeta, err := reflect_lib.PropMap(swap, "includeMeta"); err == nil {
		f.IncludeMeta = assertion_lib.MapAnyToString(includeMeta)
	} else {
		//return errors.New("\"includeMeta\" is required"), f
	}
	return f, nil
}
