package mylib

func Add(x, y int) int {
	return x + y
}

func Swap(x, y string) (string, string) {
	return y, x
}

func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func Foundation() {
	a, b := 3, 4
	println("add:", Add(a, b))

	s, t := "hello", "world"
	u, v := Swap(s, t)
	println("swap:", u, v)

	x, y := Split(17)
	println("split:", x, y)
}
