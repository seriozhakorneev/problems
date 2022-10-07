package zigzagconvertion

// https://leetcode.com/problems/zigzag-conversion/
func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	result, index := make([]byte, len(s)), 0

	y := numRows - 2
	x := numRows + y
	z := x
	for i := 0; i < numRows; i, z = i+1, z-2 {
		// first col
		if i < len(s) {
			result[index], index = s[i], index+1
		}

		for j := i; j < len(s); j += x {
			// diagonal
			if i > 0 && i < numRows-1 && j+z < len(s) {
				result[index], index = s[j+z], index+1
			}
			// row
			if j+x < len(s) {
				result[index], index = s[j+x], index+1
			}
		}
	}

	return string(result)
}
