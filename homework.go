package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    testCaseCount   := getUserNumericalInput()
    arrayStackCount := getUserNumericalInput()
    getSumOfSquares(testCaseCount, arrayStackCount)
}

func getUserNumberArrayInput() []int {
    var numbers = strings.Fields(strings.TrimSpace(getUserInput()))
    fmt.Printf("> %s\n" ,numbers)
    return toIntList(numbers, []int{}, 0)
}

func toIntList(ss []string, xs []int, cursor int) []int {
    if(cursor == len(ss)) { return xs }
    number := ss[cursor]
    xs = append(xs, toInt(number))
    return toIntList(ss, xs, cursor + 1) 
}

func toInt(x string) int {
    var count = strings.TrimSpace(x)
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
    var number = numbers[cursor]
    if(number > 0) {
        positiveNumbers = append(positiveNumbers, number)
    }
    return getPositiveNumbers(numbers, positiveNumbers, cursor + 1)
}

func getSumOfSquares(testCaseCount int, arrayStackCount int) {
    if(testCaseCount == 0) {
        return
    }
    var positiveNumbers = getPositiveNumbers(getUserNumberArrayInput(), []int{}, 0)
    var sumOfSquares    = _getSumOfSquares(positiveNumbers, 0, arrayStackCount)
    fmt.Println(sumOfSquares)
    getSumOfSquares(testCaseCount - 1, arrayStackCount)
}

func _getSumOfSquares(numbers []int, sum int, arrayStackCount int) int {
    if(len(numbers) == 0 || arrayStackCount == 0) {
        return sum
    }
    var head = numbers[0]
    var tail = numbers[1:]
    return _getSumOfSquares(tail, sum + (head * head), arrayStackCount - 1)
}
