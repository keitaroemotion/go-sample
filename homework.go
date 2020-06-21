package main

import "fmt"

const SECRET_FILE = "/usr/etc/hmt/credentials"
const SECRET_KEY  = "/usr/etc/hmt/.secret"

func main() {
    var positiveNumbers = getPositiveNumbers([]int{3, -1, 1, 14}, []int{}, 0)
    var sumOfSquares    = getSumOfSquares(positiveNumbers, 0)
    fmt.Println(sumOfSquares)
}

func getPositiveNumbers(numbers [] int, positiveNumbers []int, cursor int) []int {
    if(cursor == len(numbers)) { return positiveNumbers }
    var number = numbers[cursor]
    if(number > 0) {
        positiveNumbers = append(positiveNumbers, number)
    }
    return getPositiveNumbers(numbers, positiveNumbers, cursor + 1)
}

func getSumOfSquares(numbers []int, sum int) int {
    if(len(numbers) == 0) {
        return sum
    }
    var head = numbers[0]
    var tail = numbers[1:]
    return getSumOfSquares(tail, sum + (head * head))
}
