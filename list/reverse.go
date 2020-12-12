package list

func Reverse(l []int) {
	for i, j := 0, len(l)-1; j > i; i, j = i+1, j-1 {
		l[i], l[j] = l[j], l[i]
	}
}
