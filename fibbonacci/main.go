package main

func fib(n int) int {
	x, y := 0, 1

	fib := func() int {
		a := x + y
		x, y = y, a
		return a
	}

	if n < 1 {
		return n
	}

	for i := 2; i <= n; i++ {
		fib()
	}

	return y
}
