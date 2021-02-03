package main

import "fmt"

func transform(subject int, loop_size int) int {
    v := 1

    for step := 0; step < loop_size; step++ {
        v = (v * subject) % 20201227
    }

    return v
}

func search(subject int, wanted int) int {
    v := 1

    for step := 0; ; step++ {
        v = (v * subject) % 20201227

        if v == wanted {
            return step + 1
        }
    }
}

func main() {
    fmt.Println(transform(8252394, search(7, 6269621)))
}
