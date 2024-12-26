package day9

import "strconv"

func Level1(input []string) (sum int) {
	blocks := parse(input[0])
	rearranged := rearrange(blocks)
	for i, v := range rearranged {
		sum += i * v
	}
	return
}

func Level2(input []string) int {
	return 0
}

func rearrange(blocks []string) (rearranged []int) {
	j := len(blocks) - 1
	for i, v := range blocks {
		if i > j {
			return
		}
		for v == "." {
			v = blocks[j]
			j--
		}
		n, _ := strconv.Atoi(v)
		rearranged = append(rearranged, n)
	}
	return
}

func parse(input string) (blocks []string) {
	for i, v := range input {
		size, _ := strconv.Atoi(string(v))
		val := "."
		if i%2 == 0 {
			val = strconv.Itoa(i / 2)
		}
		for j := 0; j < size; j++ {
			blocks = append(blocks, val)
		}
	}
	return
}
