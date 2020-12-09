package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	nums := []uint{}
	{
		bin := bufio.NewReader(os.Stdin)
		for {
			line, err := bin.ReadString('\n')
			if len(line) > 0 {
				line = strings.TrimRight(line, "\r\n")
				num, err := strconv.ParseUint(line, 10, 64)
				if err != nil {
					panic(err)
				}
				nums = append(nums, uint(num))
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
	}

	const WINDOW = 25
	var first uint
find_first:
	for o := WINDOW; o < len(nums); o++ {
		num := nums[o]
		for i, a := range nums[o-WINDOW : o] {
			for e, b := range nums[o-WINDOW : o] {
				if i == e {
					continue
				}
				if a+b == num {
					continue find_first
				}
			}
		}
		first = num
		break
	}
	fmt.Println("first:", first)

find_weakness:
	for w := 2; ; w++ {
		for o := w; o < len(nums); o++ {
			var sum uint
			for _, num := range nums[o-w : o] {
				sum += num
			}
			if sum == first {
				weakness := make([]uint, w)
				copy(weakness, nums[o-w:o])
				sort.Sort(UintSlice(weakness))
				fmt.Println("second:", weakness[0]+weakness[len(weakness)-1])
				break find_weakness
			}
		}
	}
}

type UintSlice []uint

func (p UintSlice) Len() int           { return len(p) }
func (p UintSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p UintSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
