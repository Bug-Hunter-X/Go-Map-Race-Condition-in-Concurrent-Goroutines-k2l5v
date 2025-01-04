```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        m := make(map[int]int)
        var mu sync.Mutex // add mutex

        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        mu.Lock() // lock before accessing the map
                        m[i] = i
                        mu.Unlock() // unlock after accessing the map
                }(i)
        }

        wg.Wait()
        fmt.Println(len(m))
}
```