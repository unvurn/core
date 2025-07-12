package values

import "reflect"

func IsNilOrZero[T any](v T) bool {
	switch (reflect.TypeOf(v)).Kind() {
	case reflect.Slice, reflect.Map:
		return reflect.ValueOf(v).Len() == 0
	case reflect.Pointer:
		return reflect.ValueOf(v).IsNil() || reflect.ValueOf(v).Elem().IsZero()
	default:
		return reflect.ValueOf(v).IsZero()
	}
}
