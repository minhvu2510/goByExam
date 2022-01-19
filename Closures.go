package main

import "fmt"

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {

    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq()
    fmt.Println(newInts())
}
}
// => Closure có thẻ truye cạp vào 1 biến theo dõi số lần nó thay đổi nhưng ko có code nào ngoài initSeq dc truy cập đến nó