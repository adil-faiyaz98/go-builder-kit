package builders


// BuilderUtil provides utility functions for builders
type BuilderUtil struct{}

// NewBuilderUtil creates a new BuilderUtil
func NewBuilderUtil() *BuilderUtil {
	return &BuilderUtil{}
}

// DefaultBuilderUtil is the default BuilderUtil instance
var DefaultBuilderUtil = NewBuilderUtil()

// IsEmpty checks if a slice or map is empty
func (u *BuilderUtil) IsEmpty(v interface{}) bool {
	if v == nil {
		return true
	}

	// Type switch for common container types
	switch val := v.(type) {
	case string:
		return val == ""
	case []string:
		return len(val) == 0
	case []int:
		return len(val) == 0
	case []interface{}:
		return len(val) == 0
	case map[string]string:
		return len(val) == 0
	case map[string]interface{}:
		return len(val) == 0
	}

	// For other types, we can't easily check without reflection
	return false
}

// MergeStringMaps merges two string maps
func (u *BuilderUtil) MergeStringMaps(m1, m2 map[string]string) map[string]string {
	result := make(map[string]string)

	// Copy all entries from m1
	for k, v := range m1 {
		result[k] = v
	}

	// Copy all entries from m2, overwriting any duplicates from m1
	for k, v := range m2 {
		result[k] = v
	}

	return result
}

// MergeStringSlices merges two string slices
func (u *BuilderUtil) MergeStringSlices(s1, s2 []string) []string {
	result := make([]string, len(s1)+len(s2))

	// Copy all entries from s1
	copy(result, s1)

	// Copy all entries from s2
	copy(result[len(s1):], s2)

	return result
}