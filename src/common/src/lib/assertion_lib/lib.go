package assertion_lib

import (
	"fmt"
	"reflect"
)

func MapAnyToString(obj any) map[string]string {
	swap := make(map[string]string)
	mapValue := reflect.ValueOf(obj)

	if reflect.ValueOf(obj).Kind() != reflect.Map {
		return swap
	}

	for i := 0; i < mapValue.Len(); i++ {
		key := mapValue.MapKeys()[i]
		value := mapValue.MapIndex(key)
		swap[key.String()] = fmt.Sprintf("%v", value.Interface())
	}

	return swap
}
