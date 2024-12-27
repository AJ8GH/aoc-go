package day9

import (
	"strconv"
)

type blockType string
type block struct {
	id        int
	size      int
	blockType blockType
}

const (
	free blockType = "free"
	file blockType = "file"
)

func Level1(input []string) (sum int) {
	blocks, _ := parse(input[0])
	rearranged := rearrangeBlocks(blocks)
	for i, v := range rearranged {
		sum += i * v
	}
	return
}

func Level2(input []string) (sum int) {
	_, blocks := parse(input[0])
	rearranged := rearrangeFiles(blocks)
	i := 0
	for _, v := range rearranged {
		switch v.blockType {
		case free:
			i += v.size
		case file:
			for v.size > 0 {
				sum += i * v.id
				v.size--
				i++
			}
		}
	}
	return
}

func rearrangeFiles(files []block) (rearranged []block) {
	moved := map[int]bool{}
	for _, v := range files {
		switch {
		case moved[v.id]:
			continue
		case v.blockType == file:
			rearranged = append(rearranged, v)
			moved[v.id] = true
		case v.blockType == free:
		inner:
			for i := len(files) - 1; i >= 0; i-- {
				fileToMove := files[i]
				switch {
				case moved[fileToMove.id] ||
					fileToMove.blockType == free ||
					fileToMove.size > v.size:
					continue
				default:
					rearranged = append(rearranged, fileToMove)
					v.size -= fileToMove.size
					moved[fileToMove.id] = true
					files[i] = block{id: -1, size: fileToMove.size, blockType: free}
					if v.size == 0 {
						break inner
					}
				}
			}
			if v.size > 0 {
				rearranged = append(rearranged, v)
			}
		}
	}
	return
}

func rearrangeBlocks(blocks []string) (rearranged []int) {
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

func parse(input string) (blocks []string, fileBlocks []block) {
	for i, v := range input {
		size, _ := strconv.Atoi(string(v))
		val := "."
		blockType := free
		id := -1
		if i%2 == 0 {
			id = i / 2
			val = strconv.Itoa(id)
			blockType = file
		}
		fileBlocks = append(fileBlocks, block{id: id, size: size, blockType: blockType})
		for j := 0; j < size; j++ {
			blocks = append(blocks, val)
		}
	}
	return
}
