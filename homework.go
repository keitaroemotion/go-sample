package main

import "fmt"

const SECRET_FILE = "/usr/etc/hmt/credentials"
const SECRET_KEY  = "/usr/etc/hmt/.secret"

func main() {
    readCredentials(SECRET_FILE)
    // action 1: read credentials
    //  - decrypt credentials - asks for user input (pass) to decrypt
    //  - parse and get required credentials
    //  - set on the clipboard

    // action 2: write credentials
    // - enrypt credentials
    fmt.Println("hello world")
}

func readCredentials(file string) {
    // file content should be ciphered, so need to decrypt it first.
}

func getSumOfSquares(numbers []int, sum int) {
}
