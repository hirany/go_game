package main

import "fmt"

func output(l int, dist ...int) {
	for i := 0; i < l*l; i++ {
		fmt.Print(dist[i], " ")
		if i%l == l-1 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func find(l int, dist ...int) bool {
	l *= l
	for i := 0; i < l; i++ {
		if dist[i] == 0 {
			return false
		}
	}
	return true
}

func sca(x, y, l int, dist ...int) int {
	fmt.Scan(&x, &y)
	fmt.Println()
	if x < 0 || l-1 < x || y < 0 || l-1 < y {
		return 0
	}
	rev(x, y, l, dist...)
	return 1
}

func rev(x, y, l int, dist ...int) {
	if 0 < x {
		dist[x+y*l-1] = (dist[x+y*l-1] + 1) % 2
	}
	if x < l {
		dist[x+y*l+1] = (dist[x+y*l+1] + 1) % 2
	}
	if 0 < y {
		dist[x+y*l-l] = (dist[x+y*l-l] + 1) % 2
	}
	if y < l {
		dist[x+y*l+l] = (dist[x+y*l+l] + 1) % 2
	}
}

func main() {
	dist := make([]int, 100)
	var x, y int
	var l int
	fmt.Print("level = ")
	fmt.Scan(&l)
	output(l, dist...)
	for i := 1; true; i++ {
		if sca(x, y, l, dist...) == 0 {
			fmt.Println("miss")
			break
		}
		output(l, dist...)
		if find(l, dist...) {
			fmt.Println("count =", i)
			fmt.Println("complete!!")
			break
		}
	}
}
