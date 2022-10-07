package atoi

//https://leetcode.com/problems/string-to-integer-atoi/
func myAtoi(s string) int {
	right, left := 0, -1
loop:
	for ; right < len(s); right++ {
		switch {
		case s[right] == ' ' && left == -1:
			continue
		case isSign(s[right]) &&
			right < len(s)-1 &&
			!isDigit(s[right+1]) &&
			left == -1:
			return 0
		case !isDigit(s[right]) && left > -1:
			break loop
		case !isDigit(s[right]) && !isSign(s[right]):
			return 0
		case isDigit(s[right]) && left == -1:
			left = right
		}
	}

	if left == -1 {
		return 0
	}
	if left != 0 && s[left-1] == '-' {
		left--
	}

	return convert(s[left:right])
}

// convert stolen from strconv.Atoi and modified
func convert(s string) int {
	const (
		max = 2147483647
		min = -2147483648
	)
	s0 := s
	if s[0] == '-' || s[0] == '+' {
		s = s[1:]
	}

	n := 0
	for _, ch := range []byte(s) {
		ch -= '0'
		if ch > 9 {
			return 0
		}
		n = n*10 + int(ch)

		if s0[0] == '-' {
			if -n <= min {
				return min
			}
		}
		if n > max {
			return max
		}
	}

	if s0[0] == '-' {
		n = -n
	}
	return n
}

func isSign(a uint8) bool {
	if a == '-' || a == '+' {
		return true
	}
	return false
}

func isDigit(a uint8) bool {
	if a >= '0' && a <= '9' {
		return true
	}
	return false
}
