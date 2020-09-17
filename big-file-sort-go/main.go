package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
)

func main() {
	merge(splitandsort())
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func generatebigfile() {
	file, err := os.Create("./bigfile.txt")
	check(err)

	buf := bufio.NewWriter(file)
	defer buf.Flush()
	for i := 0; i < 1000000; i++ {
		randIntStr := strconv.FormatInt(rand.Int63(), 10)
		buf.WriteString(randIntStr)
		buf.WriteString("\n")
	}
}

func splitandsort() []string {
	file, err := os.Open("./bigfile.txt")
	check(err)

	var (
		count     = 0
		filecount = 0
		br        = bufio.NewReader(file)
		files     []string
		nums      []int64
	)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		i, err := strconv.ParseInt(string(a), 10, 64)
		check(err)
		nums = append(nums, i)
		count++

		if count >= 10000 {
			sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

			filecount++
			filename := fmt.Sprintf("./file_%d.txt", filecount)
			files = append(files, filename)

			file, err := os.Create(filename)
			check(err)

			bw := bufio.NewWriter(file)
			for _, i := range nums {
				_, err := bw.WriteString(strconv.FormatInt(i, 10) + "\n")
				check(err)
			}
			check(bw.Flush())

			count = 0
			nums = nums[:0]
		}
	}

	return files
}

func merge(files []string) {
	fmt.Println(files)
	var readers = make([]*bufio.Reader, len(files))
	for i, filename := range files {
		fr, err := os.Open(filename)
		check(err)
		readers[i] = bufio.NewReader(fr)
	}

	fw, err := os.Create("./bigfile_sorted.txt")
	check(err)
	var bw = bufio.NewWriter(fw)

	var nums []int64
	for i, reader := range readers {
		num, ok := read(reader)
		if ok {
			nums = append(nums, num)
		} else {
			readers = append(readers[:i], readers[i+1:]...)
		}
	}

	for len(readers) > 0 {
		index, num := min(nums)
		bw.WriteString(strconv.FormatInt(num, 10))
		bw.WriteString("\n")
		bw.Flush()

		num, ok := read(readers[index])
		if ok {
			nums[index] = num
		} else {
			nums = append(nums[:index], nums[index+1:]...)
			readers = append(readers[:index], readers[index+1:]...)
		}
	}
}

func min(nums []int64) (int, int64) {
	var m int64 = math.MaxInt64
	var index = 0

	for i, num := range nums {
		if num < m {
			m = num
			index = i
		}
	}

	return index, m
}

func read(reader *bufio.Reader) (int64, bool) {
	numBytes, _, err := reader.ReadLine()
	if err == io.EOF {
		return 0, false
	}

	num, err := strconv.ParseInt(string(numBytes), 10, 64)
	check(err)

	return num, true
}
