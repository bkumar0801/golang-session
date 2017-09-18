package function

func Add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func MakeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
