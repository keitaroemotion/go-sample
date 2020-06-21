package main

import "fmt"

const SECRET_FILE = "/usr/etc/hmt/credentials"
const SECRET_KEY  = "/usr/etc/hmt/.secret"

func main() {
    var x = filterNegativeElements([]int{3, -1, 1, 14}, []int{}, 0)
    fmt.Println(x)
}

func filterNegativeElements(numbers [] int, positiveNumbers []int, cursor int) []int {
    if(cursor == len(numbers)) { return positiveNumbers }
    var number = numbers[cursor]
    if(number > 0) {
        positiveNumbers = append(positiveNumbers, number)
    }
    return filterNegativeElements(numbers, positiveNumbers, cursor + 1)
}

func getSumOfSquares(numbers []int, sum int) int {
    var head = numbers[0]
    sum      = sum + head * head
    if(len(numbers) == 1) {
        return sum
    }
    var tail = numbers[1:]
    return head + getSumOfSquares(tail, sum)
}
