package jsonUtil

import "reflect"

// ValueTypesInMap is a MapValueType slice to store all the types for values in a map.
type ValueTypesInMap []MapValueType

// Find returns the MapValueType that matches the search criteria. It returns false, if no match found.
func (v ValueTypesInMap) Find(parentKey string, key string, noOfSliceLvs int) (out MapValueType, found bool) {
	for _, x := range v {
		if parentKey == x.ParentKey && key == x.ClosestKey && noOfSliceLvs == x.NoOfSliceLevels {
			return x, true
		}
	}
	return MapValueType{}, false
}

// Delete removes a MapValueType from the slice.
func (v ValueTypesInMap) Delete(parentKey string, key string, noOfSliceLvs int) ValueTypesInMap {
	vv := ValueTypesInMap{}
	for _, x := range v {
		if !(parentKey == x.ParentKey && key == x.ClosestKey && noOfSliceLvs == x.NoOfSliceLevels) {
			vv = append(vv, x)
		}
	}
	return vv
}

// Append appends a MapValueType to ValueTypesInMap. If the is a MapValueType with the same path already, 1) the existing one's Type == reflect.Invalid, delete the existing one, append the new one and return appended = true; 2) otherwise, the new one will not be appended. It will return the existing one with appended = false.
// In short, it treat the schema of the JSON as fixed and let other mechanism to handle dynamic JSON schema, if needed.
func (v ValueTypesInMap) Append(m MapValueType) (fromTypes MapValueType, appended bool) {
	o, found := v.Find(m.ParentKey, m.ClosestKey, m.NoOfSliceLevels)
	if found {
		if o.Type == reflect.Invalid {
			v.Delete(m.ParentKey, m.ClosestKey, m.NoOfSliceLevels)
		} else {
			return MapValueType{}, false
		}
	}
	v = append(v, m)
	return m, true
}
