package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

const SECRET_FILE = "/usr/etc/hmt/credentials"
const SECRET_KEY  = "/usr/etc/hmt/.secret"

func main() {
    testCaseCount   := getUserNumericalInput()
    arrayStackCount := getUserNumericalInput()
    getSumOfSquares(testCaseCount, arrayStackCount)
}

func getUserNumericalInput() int {
    var count = getUserInput()
    countInt, err := strconv.Atoi(count)
    if err != nil {
    }
    return countInt
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

func getSumOfSquares(testCaseCount int, arrayStackCount int) int {
    var positiveNumbers = getPositiveNumbers([]int{9, 6, -53, 32, 16}, []int{}, 0)
    var sumOfSquares    = _getSumOfSquares(positiveNumbers, 0)
    fmt.Println(sumOfSquares)
    return sumOfSquares
}

func _getSumOfSquares(numbers []int, sum int) int {
    if(len(numbers) == 0) {
        return sum
    }
    var head = numbers[0]
    var tail = numbers[1:]
    return _getSumOfSquares(tail, sum + (head * head))
}
