package basics

func SwapTwoNumbers(a, b int) (int, int) {
	a, b = b, a
	return a, b
}
