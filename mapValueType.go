package jsonUtil

import "reflect"

// We use the word "memo" here to indicate a data structure that stores all the value type information we known.
// For a value in a map, its location in a map can be either under a key directly or under a key and some levels of slices. It's like a path for the value.
// In a value type memo, which only stores the type for value. If the value is in a slice, there are two possible cases. 1) the slice is homogeneous, there is no need to know it's index since every value in the slice has the same type. 2) the slice is heterogeneous, we have to explore it every time we meet it. It needs extra work to find out it's value pattern over time. Because of this, we leave the pattern exploration function out for now. The only thing to do at this moment is to go through everything in the slice each time we meet it and do not store its type to memo since we do not know its pattern and are not able to be sure the pattern will remain the same in the future.
// Based on above cases, we only care about the homogeneous slice, which the index of the value is not a concern.
// So the path from the closest key to the value can be key/noOfSliceLevels. It's better to add the parent key to identify value's location in case there are keys with same names.

// MapValueType stores the path to a value and the type of the value. If it's under root key, the ParentKey = "".
type MapValueType struct {
	ParentKey             string
	ClosestKey            string
	ParentNoOfSliceLevels int
	NoOfSliceLevels       int
	Type                  reflect.Kind
}

// IsForSameNode finds out if the given MapValueType is the same node as the receiver.
func (m MapValueType) IsForSameNode(mm MapValueType) bool {
	if m.ParentKey == mm.ParentKey && m.ClosestKey == mm.ClosestKey && m.ParentNoOfSliceLevels == mm.ParentNoOfSliceLevels && m.NoOfSliceLevels == mm.NoOfSliceLevels {
		return true
	}
	return false
}

// IsInstantUnderSameKey finds out if the given MapValueType is directly under the same key as the receiver.
func (m MapValueType) IsInstantUnderSameKey(mm MapValueType) bool {
	if m.ParentKey == mm.ParentKey && m.ClosestKey == mm.ClosestKey {
		return true
	}
	return false
}

// ValueTypeForPath gets MapValueType for a value.
func ValueTypeForPath(parentKey string, key string, parentNoOfSliceLevels int, noOfSliceLvs int, v interface{}) MapValueType {
	return MapValueType{parentKey, key, parentNoOfSliceLevels, noOfSliceLvs, TypeForValue(v)}
}

// ValueTypesForMap gets all type info in the form of MapValueType for a map.
func ValueTypesForMap(parentKey string, parentNoOfSliceLevels int, noOfSliceLvs int, m map[string]interface{}) []MapValueType {
	r := []MapValueType{}
	for k, v := range m {
		r = append(r, ValueTypeForPath(parentKey, k, parentNoOfSliceLevels, noOfSliceLvs, v))
	}
	return r
}

// ValueTypeForSlice gets type info in the form of MapValueType for a homogeneous slice.
func ValueTypeForSlice(parentKey string, key string, parentNoOfSliceLevels int, noOfSliceLvs int, s []interface{}) MapValueType {
	if len(s) > 0 {
		return ValueTypeForPath(parentKey, key, parentNoOfSliceLevels, noOfSliceLvs, s[1])
	}
	return MapValueType{parentKey, key, parentNoOfSliceLevels, noOfSliceLvs, reflect.Invalid}
}
