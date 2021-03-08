package pkg

// Contains ...
func Contains(target interface{}, list interface{}) bool {
	switch list.(type) {
	case []int:
		revert := list.([]int)
		for _, r := range revert {
			if target == r {
				return true
			}
		}
		return false
	case []string:
		revert := list.([]string)
		for _, r := range revert {
			if target == r {
				return true
			}
		}
		return false
	default:
		return false
	}
}
