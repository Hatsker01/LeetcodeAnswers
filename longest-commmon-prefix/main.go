func longestCommonPrefix(strs []string) string {
	var st string = ""

	var min int = len(strs[0])
	if len(strs) == 1 {
		st = string(strs[0])
		return st
	}
	for i := 0; i < len(strs); i++ {
		if len(strs[i]) < min {
			min = len(strs[i])
		}
	}
	for i := 0; i < min; i++ {
		var t bool = false
		for j := 1; j < len(strs); j++ {
			if (strs[0])[i] == (strs[j])[i] {
				// st = st + string((strs[0])[i])
				t = true
			} else {
				t = false
				break
			}
		}
		if t {
			st = st + string((strs[0])[i])
		} else {
			break
		}

	}

	return st
}