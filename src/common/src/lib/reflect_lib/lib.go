package reflect_lib

import (
	"errors"
	"reflect"
	"strings"
)

func RMGetVal(val any, key string) (reflect.Value, error) {
	mapValue := reflect.ValueOf(val)
	//mapType := mapValue.Type()
	keySlice := strings.Split(key, ".")
	kSliceLen := len(keySlice)

	var value reflect.Value
	for k := 0; kSliceLen > k; k++ {
		for i := 0; i < mapValue.Len(); i++ {
			mkey := mapValue.MapKeys()[i]

			if keySlice[k] == mkey.String() {
				value = mapValue.MapIndex(mkey)
				if kSliceLen-1 > k && reflect.TypeOf(value.Interface()).Kind() == reflect.Map {
					return RMGetVal(value.Interface(), strings.Join(keySlice[k+1:], "."))
				}
			}

		}
	}

	if value.Kind() == reflect.Invalid {
		return value, errors.New(value.Kind().String())
	}
	return reflect.ValueOf(value.Interface()), nil
}

func Prop(val any, name string) (reflect.Value, error) {
	return RMGetVal(val, name)
}

func PropString(val any, name string) (*string, error) {
	if val, err := Prop(val, name); err != nil {
		return nil, err
	} else if val.Kind() != reflect.String {
		return nil, errors.New(val.String())
	} else {
		v := val.String()
		return &v, err
	}
}

func PropBoll(val any, name string) (*bool, error) {
	if val, err := Prop(val, name); err != nil {
		return nil, err
	} else if val.Kind() != reflect.Bool {
		return nil, errors.New(val.String())
	} else {
		v := val.Bool()
		return &v, err
	}
}

func PropInt(val any, name string) (*int, error) {
	if val, err := Prop(val, name); err != nil {
		return nil, err
	} else if val.Kind() != reflect.Int {
		return nil, errors.New(val.String())
	} else {
		v := int(val.Int())
		return &v, err
	}
}

func PropUint(val any, name string) (*uint, error) {
	if val, err := Prop(val, name); err != nil {
		return nil, err
	} else if val.Kind() != reflect.Uint {
		return nil, errors.New(val.String())
	} else {
		v := uint(val.Int())
		return &v, err
	}
}

func PropInterface(val any, name string) (*any, error) {
	if val, err := Prop(val, name); err != nil {
		return nil, err
	} else if val.Kind() != reflect.Interface {
		return nil, errors.New(val.String())
	} else {
		v := val.Interface()
		return &v, err
	}
}

func PropMap(val any, name string) (*map[string]any, error) {
	if val, err := Prop(val, name); err != nil {
		return nil, err
	} else if val.Kind() != reflect.Map {
		return nil, errors.New(val.String())
	} else {
		v := val.Interface().(map[string]any)
		return &v, err
	}
}
