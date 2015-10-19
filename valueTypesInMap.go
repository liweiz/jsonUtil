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

// Append appends a MapValueType to ValueTypesInMap. It treats the schema of the JSON as fixed and let other mechanism to handle dynamic JSON schema, if needed. A new slice is always returned. the only difference is the appended value that indicates if the value is really appended.
func (v ValueTypesInMap) Append(m MapValueType) (out ValueTypesInMap, appended bool) {
	o, found := v.Find(m.ParentKey, m.ClosestKey, m.NoOfSliceLevels)
	if found {
		if o.Type == reflect.Invalid {
			if m.Type == reflect.Invalid {
				return v, false
			}
			v = v.Delete(m.ParentKey, m.ClosestKey, m.NoOfSliceLevels)
		} else {
			return v, false
		}
	}
	v = append(v, m)
	return v, true
}
