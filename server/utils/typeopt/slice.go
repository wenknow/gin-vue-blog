package typeopt

func MaxSliceInt(l []int) (max int, index int) {
	for k, v := range l {
		if v > max {
			max = v
			index = k
		}
	}
	return
}

func MinSliceInt(l []int) (min int, index int) {
	for k, v := range l {
		if v > min {
			min = v
			index = k
		}
	}
	return
}

func InUintArray(slice []uint, val uint) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func InArray(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func IntersectUint(s1, s2 []uint) (result []uint) {
	if s1 == nil || s2 == nil {
		return result
	}
	if len(s1) == 0 || len(s2) == 0 {
		return result
	}
	tempMap := make(map[uint]uint)
	for _, v := range s1 {
		tempMap[v]++
	}
	for _, v := range s2 {
		if c, ok := tempMap[v]; ok && c > 0 {
			result = append(result, v)
			tempMap[v]--
		}
	}
	return result
}
