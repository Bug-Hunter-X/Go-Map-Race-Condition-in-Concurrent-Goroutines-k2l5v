```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        m := make(map[int]int)

        for i := 0; i < 1000; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        m[i] = i
                }(i)
        }

        wg.Wait()
        fmt.Println(len(m))
}
```