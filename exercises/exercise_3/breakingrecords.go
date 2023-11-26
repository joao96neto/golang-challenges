package breakingrecords

var Input = []int32{12, 24, 10, 24}

func Execute(scores []int32) []int32 {
	value := make([]int32, 2)
	min, max, minRes, maxRes := int32(0), int32(0), int32(0), int32(0)

	for i, score := range scores {
		if i == 0 {
			min = score
			max = score
			continue
		}
		if score > max {
			max = score
			maxRes++
		}
		if score < min {
			min = score
			minRes++
		}
	}
	value[1] = minRes
	value[0] = maxRes
	return value
}
