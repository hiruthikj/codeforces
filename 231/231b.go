package main

import "fmt"
import "math"

func computeInitialCards(n, d, l int) ([]int, int) {
	initial_cards := make([]int, n)

	min_odd_sum := (n+1)/2
	max_odd_sum := ((n+1)/2) * l
	
	if d+(n/2)*l < min_odd_sum {
		//println(d+(n/2)*l, min_odd_sum)
		return []int{}, -1
	}

	if d+(n/2) > max_odd_sum {
		//println(d+(n/2), max_odd_sum)
		//println("max_odd_sum")
		return []int{}, -1
	}

	curr_even_num := int(math.Max(math.Ceil(float64(min_odd_sum-d) / float64(n/2)), 1))
	even_pos_sum := curr_even_num * (n/2)
	
	//println(curr_even_num, even_pos_sum)

	curr_odd_num := (even_pos_sum + d) / ((n+1)/2)
	curr_odd_sum := curr_odd_num * ((n+1)/2)

//	println(curr_odd_num, curr_odd_sum)
	for i := range initial_cards {
		if i%2 == 1 {
			initial_cards[i] = curr_even_num
		} else {
			initial_cards[i] = curr_odd_num
		}
	}

//	fmt.Println(initial_cards, "f")
	curr_pos := 0
	for curr_odd_sum < even_pos_sum + d {
		initial_cards[curr_pos] += 1
		curr_odd_sum++
		curr_pos += 2
		if curr_pos >= n {
			curr_pos = 0
		}
	}

	return initial_cards, 0

}

func main() {
	var n, d, l int

	fmt.Scanf("%d %d %d\n", &n, &d, &l)
	
	result, errNo := computeInitialCards(n,d,l)
	if errNo == -1 {
		fmt.Println(-1)
	} else {
		res := fmt.Sprintf("%v", result)
		fmt.Println(res[1: len(res)-1])
	}
}
