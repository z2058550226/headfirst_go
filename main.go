package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// guess()
	readFile()
}

func readFile() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func guess() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println(target)
	for i := 10; i > 0; i-- {
		fmt.Println("Make a guess, left times is: ", i)
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)

		inputNum, err := strconv.Atoi(input)
		if inputNum == target {
			fmt.Println("right")
			break
		} else if inputNum > target {
			fmt.Println("bigger")
		} else if inputNum < target {
			fmt.Println("less")
		}
	}

}

func outputGrade() {
	fmt.Println("Enter a Grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	fmt.Println("grade is: ", grade)

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("status: ", status)
}

func main2() {
	fileInfo, err := os.Stat("go.mod")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}
