// Function in action at http://play.golang.org/p/KU4Jn1Qkcx

func ArrayMin(a []int, min int) int {	
	if (len(a) == 0) {
		return min
	} else {
		if (a[0] < min) {
			min = a[0]
		}
	}
	return ArrayMin(a[1:], min)
}