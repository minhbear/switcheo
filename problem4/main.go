package main

// Complexity O(1)
func sum_to_n_a(n int) int {
	if n == 0 {
		return 0
	} else if n > 0 {
		return ((n + 1) * n) / 2
	} else {
		return -(((-n + 2) * -n) / 2)
	}
}

// Complexity O(n)
func sum_to_n_b(n int) int {
	if n == 0 {
		return 0
	} else if n > 0 {
		sum := 0

		for i := 1; i <= n; i++ {
			sum += i
		}

		return sum
	} else {
		sum := 0

		for k := n; k <= -2; k++ {
			sum += k
		}

		return sum
	}
}

// Complexity O(n)
func sum_to_n_c(n int) int {
	if n == 0 {
		return 0
	} else if n > 0 {
		return n + sum_recursive(1, n-1)
	} else {
		return n + sum_recursive(-2, n+1)
	}
}

func sum_recursive(last_value int, current_value int) int {
	if last_value == current_value {
		return last_value
	} else {
		var step int

		if last_value == 1 {
			step = -1
		} else {
			step = 1
		}

		return current_value + sum_recursive(last_value, current_value + step)
	}
}

func main() {
	println(sum_to_n_c(-5))
}
