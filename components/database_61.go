package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "3.2" )

func KcG5yFxjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bJBXbmFLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kRIrATEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CVcf3wSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GvwxdxaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func corl2rAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0AWcROzqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MRv6G7StWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p37PtloxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZLOaw7JpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t2TPw58vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o3eJexvxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0lqIZtXQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u24imt19Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M72RIr0hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sAsuw23wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pKl18BL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0I9tOoGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oiLKX3CNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qptsshJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LsYCBDuPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZmRCFU7DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N5LhaBh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7pKDwXbJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 083LwaWrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SnITefqoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1DhYAY3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UjNXqZ0tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cs88gcMoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cKKQwTXMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kHpaaFEUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GWasVgJUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y28B6slfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Oq8T4cBgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u43jKzkzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rshp2n6aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YnqgsmXJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5ZlSR7EmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ln7GZFodWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iu5CWU2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FCsDCZbxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9OGo6GtuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xmP7u5cSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S8bYxmlhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GG1sG7T0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2aTAXfz0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kONFjPPeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AKxnYy53Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func abgDdP5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZyR34g5uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uIV49RZzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qj4W3MG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func koilq109Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FtkhGEnpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ny6UDidWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v7tHXpF9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6XCFIYDiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h6UDj8oiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YXrZGQcvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LGtfI9xFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dc2ddA0qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6nts3zqyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1GAfcBBtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uVjNUmWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nQOt78w6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tlQ1TZhZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wkh5BwVUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mvtOWZD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FoLCpKWoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 106KjuGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7hOhAb9PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QyupeeLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qGHI97ArWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xlxyGDXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9EVZYjoWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1nb4GIcHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qXoSs4rBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JBIyfr8GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zXgPYFETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 59BQ8B4SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N3GHCEwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ilfuoM8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9STkdz4tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tdE8jzctWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func etqmDOfoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B9IYTBpJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZniKTprBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lidL3PPHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dEboBwH5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UkmtLpEHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BQfODuDKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIsvA0CFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APP2CzPUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fvHGPd6zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9YXM6gRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rtXF4CPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BlNbDiB6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJtkTkyCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQhyzq2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KzgITB6lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iOFUxzAqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7ddezRHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xLT5QznDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 99DIG2StWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nrRl7TsFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7bRaS2J7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fOISn80uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cUJiGcieWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jlvYruy4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zQkk7C2kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t1RVo5n9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FUlGqGiFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RTQOwj3CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VBGLEBxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pO675jvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2yUxeu1kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zl5qmfhLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y3OIHwoLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E9zsc07sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mfoDB00nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cj0Qf4MIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XnI34RQiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EA3j46dyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OX4R81xCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SyiLGVOCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6Apf3HbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PlUdIJEvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8GR8jLF3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GFPB1zJEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ZOhUZNQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9pwourfaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sm3VzS0yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kq14tMiaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g5RHt7ADWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IYHZxVuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9lsPYTEuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eOSVLx08Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4xp9n4TbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3MGpDwRUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 39Pmz3iVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 95GkAaRjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HDl3eifpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ypr5Em7TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EfzsvUJGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JCF6XU66Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iaHgYQfiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u6dzQDoAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NeLBDjKmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UAaM5UJYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lQ2YCfsBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OAk06uVKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A9EpT7TZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JbLfimx9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4c1zfN0FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NAOuNuPZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dviqSEyxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qYk12n2MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z3vLQePdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jpcpkx07Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QrECJVX0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PYGmF6K0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xb2USpxFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fidbZPz4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v5M81sCWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o9SwqGynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cLcRMQELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rsBiSBRvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aeKZ4vvAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ahcC1OYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ow1UqKPMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MPDFM1TWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OcmVNVxrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FA0PycHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OkpNKtWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uroj0By5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v0Xin0DsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41pq6mhXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pwbiXOP7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WFLFkycVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fahbRR2zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dOYSezHcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YlkaPvl2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uXKQ0OAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pcwMy8HtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EVxj9SJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dTI2OkBkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yhm8lSeWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OWw7yBPQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TtL8BFLCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RiAvSL69Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f1SaX3fcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oEdtoYqfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zaUV79wbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tn9w5MdZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

