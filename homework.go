package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    testCaseCount := getUserNumericalInput()
    getTestArgumentSet(testCaseCount)
}

func getTestArgumentSet(testCaseCount int) {
    if(testCaseCount == 0) { return }
    arrayStackCount := getUserNumericalInput()
    positiveNumbers := getPositiveNumbers(getUserNumberArrayInput(), []int{}, 0)
    sumOfSquares    := getSumOfSquares(positiveNumbers[0:arrayStackCount-1], 0)
    fmt.Printf("%d\n", sumOfSquares)
    getTestArgumentSet(testCaseCount - 1)
}

func getUserNumberArrayInput() []int {
    numbers := strings.Fields(strings.TrimSpace(getUserInput()))
    return toIntList(numbers, []int{}, 0)
}

func toIntList(ss []string, xs []int, cursor int) []int {
    if(cursor == len(ss)) { return xs }
    number := ss[cursor]
    xs = append(xs, toInt(number))
    return toIntList(ss, xs, cursor + 1) 
}

func toInt(x string) int {
    count         := strings.TrimSpace(x)
    countInt, err := strconv.Atoi(count)
    if err != nil {
    }
    return countInt
}

func getUserNumericalInput() int {
    return toInt(getUserInput())
}

func getUserInput() string {
    reader  := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    return text
}

func getPositiveNumbers(numbers [] int, positiveNumbers []int, cursor int) []int {
    if(cursor == len(numbers)) { return positiveNumbers }
    number := numbers[cursor]
    if(number > 0) {
        positiveNumbers = append(positiveNumbers, number)
    }
    return getPositiveNumbers(numbers, positiveNumbers, cursor + 1)
}

func getSumOfSquares(numbers []int, sum int) int {
    if(len(numbers) == 0) {
        return sum
    }
    head := numbers[0]
    tail := numbers[1:]
    return getSumOfSquares(tail, sum + (head * head))
}
