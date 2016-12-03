package main

func isPossibleTriangle(a, b, c int) bool {
	if a > b {
		if a > c {
			if b+c > a {
				return true
			}
		} else {
			if a+b > c {
				return true
			}
		}
	} else {
		if b > c {
			if a+c > b {
				return true
			}
		} else {
			if a+b > c {
				return true
			}
		}
	}

	return false
}
