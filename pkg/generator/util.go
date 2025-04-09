package generator

// This file contains utility templates for generated builders

// UtilTemplate contains utility functions to include in generated builders
const UtilTemplate = `
// BuilderUtil provides utility functions for builders
type BuilderUtil struct{}

// NewBuilderUtil creates a new BuilderUtil
func NewBuilderUtil() *BuilderUtil {
	return &BuilderUtil{}
}

// DefaultBuilderUtil is the default BuilderUtil instance
var DefaultBuilderUtil = NewBuilderUtil()

// IsZero checks if a value is the zero value for its type
func (u *BuilderUtil) IsZero(v interface{}) bool {
	if v == nil {
		return true
	}

	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Slice, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func:
		return val.IsNil()
	case reflect.Array:
		zero := true
		for i := 0; i < val.Len(); i++ {
			zero = zero && u.IsZero(val.Index(i).Interface())
		}
		return zero
	case reflect.Struct:
		zero := true
		for i := 0; i < val.NumField(); i++ {
			zero = zero && u.IsZero(val.Field(i).Interface())
		}
		return zero
	}

	return val.Interface() == reflect.Zero(val.Type()).Interface()
}

// DeepCopy creates a deep copy of a value
func (u *BuilderUtil) DeepCopy(v interface{}) interface{} {
	if v == nil {
		return nil
	}

	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Ptr:
		if val.IsNil() {
			return nil
		}
		newVal := reflect.New(val.Elem().Type())
		newVal.Elem().Set(reflect.ValueOf(u.DeepCopy(val.Elem().Interface())))
		return newVal.Interface()
	case reflect.Slice:
		if val.IsNil() {
			return nil
		}
		newSlice := reflect.MakeSlice(val.Type(), val.Len(), val.Cap())
		for i := 0; i < val.Len(); i++ {
			newSlice.Index(i).Set(reflect.ValueOf(u.DeepCopy(val.Index(i).Interface())))
		}
		return newSlice.Interface()
	case reflect.Map:
		if val.IsNil() {
			return nil
		}
		newMap := reflect.MakeMap(val.Type())
		for _, key := range val.MapKeys() {
			newMap.SetMapIndex(key, reflect.ValueOf(u.DeepCopy(val.MapIndex(key).Interface())))
		}
		return newMap.Interface()
	case reflect.Struct:
		newStruct := reflect.New(val.Type()).Elem()
		for i := 0; i < val.NumField(); i++ {
			newStruct.Field(i).Set(reflect.ValueOf(u.DeepCopy(val.Field(i).Interface())))
		}
		return newStruct.Interface()
	default:
		return v
	}
}

// MergeSlices merges two slices of the same type
func (u *BuilderUtil) MergeSlices(s1, s2 interface{}) interface{} {
	v1 := reflect.ValueOf(s1)
	v2 := reflect.ValueOf(s2)

	if v1.Kind() != reflect.Slice || v2.Kind() != reflect.Slice {
		panic(fmt.Sprintf("both arguments must be slices, got %T and %T", s1, s2))
	}

	if v1.Type() != v2.Type() {
		panic(fmt.Sprintf("slices must be of the same type, got %v and %v", v1.Type(), v2.Type()))
	}

	result := reflect.MakeSlice(v1.Type(), v1.Len()+v2.Len(), v1.Len()+v2.Len())

	for i := 0; i < v1.Len(); i++ {
		result.Index(i).Set(v1.Index(i))
	}

	for i := 0; i < v2.Len(); i++ {
		result.Index(v1.Len()+i).Set(v2.Index(i))
	}

	return result.Interface()
}

// MergeMaps merges two maps of the same type
func (u *BuilderUtil) MergeMaps(m1, m2 interface{}) interface{} {
	v1 := reflect.ValueOf(m1)
	v2 := reflect.ValueOf(m2)

	if v1.Kind() != reflect.Map || v2.Kind() != reflect.Map {
		panic(fmt.Sprintf("both arguments must be maps, got %T and %T", m1, m2))
	}

	if v1.Type() != v2.Type() {
		panic(fmt.Sprintf("maps must be of the same type, got %v and %v", v1.Type(), v2.Type()))
	}

	result := reflect.MakeMap(v1.Type())

	for _, key := range v1.MapKeys() {
		result.SetMapIndex(key, v1.MapIndex(key))
	}

	for _, key := range v2.MapKeys() {
		result.SetMapIndex(key, v2.MapIndex(key))
	}

	return result.Interface()
}
`
