package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// ReadFileLines takes a file name and returns a slice of strings,
// each containing one line from the file.
func ReadFileLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// RotateLeft returns a new slice where the elements are shifted left by n.
// The elements that "fall off" the left are appended at the end.
func RotateLeft(s []int, n int) []int {
	if len(s) == 0 {
		return s
	}
	n = n % len(s)
	if n < 0 {
		n += len(s)
	}

	// s = [0..n-1 | n..end]
	// result = [n..end | 0..n-1]
	return append(s[n:], s[:n]...)
}

func RotateLeftInPlace(s []int, n int) {
	t := RotateLeft(s, n)
	copy(s, t)
}

// RotateRight returns a new slice where the elements are shifted right by n.
// The elements that "fall off" the right are prepended at the start.
func RotateRight(s []int, n int) []int {
	if len(s) == 0 {
		return s
	}
	n = n % len(s)
	if n < 0 {
		n += len(s)
	}

	// rotating right by n is the same as rotating left by len(s)-n
	k := len(s) - n

	return append(s[k:], s[:k]...)
}

func RotateRightInPlace(s []int, n int) {
	t := RotateRight(s, n)
	copy(s, t)
}

func parseMove(s string) (dir rune, steps int, ok bool) {
	if len(s) < 2 {
		return 0, 0, false // too short
	}

	// First character must be 'L' or 'R'
	dir = rune(s[0])
	if dir != 'L' && dir != 'R' {
		return 0, 0, false
	}

	// The rest should be 1–2 digits (or more, if you like)
	n, err := strconv.Atoi(s[1:])
	if err != nil {
		return 0, 0, false
	}

	return dir, n, true
}

func main() {
	wd, _ := os.Getwd()
	fmt.Println("Working directory:", wd)

	filename := "input.txt"
	fmt.Println("Trying to open:", filename)

	lines, err := ReadFileLines("day-1-secret-entrance/input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	nums := make([]int, 100)
	for i := 0; i < 100; i++ {
		nums[i] = i
	}
	RotateLeftInPlace(nums, 50)

	counter := 0

	for _, line := range lines {
		move, steps, ok := parseMove(line)
		if !ok {
			fmt.Printf("Error in parseMove: %s\n", line)

		}
		switch move {
		case 'L':
			RotateLeftInPlace(nums, steps)
		case 'R':
			RotateRightInPlace(nums, steps)
		}

		if nums[0] == 0 {
			counter += 1
		}
	}

	fmt.Printf("The result is: %d", counter)
}
