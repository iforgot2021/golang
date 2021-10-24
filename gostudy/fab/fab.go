package fab

func Fab(num int) int {
	switch num {
	case 1, 2:
		return 1
	default:
		return Fab(num-1) + Fab(num-2)
	}
}

func FabList(num int) []int {
	res := make([]int, 0, num)
	for i := 1; i <= num; i++ {
		res = append(res, Fab(i))
	}
	return res
}
