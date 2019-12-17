package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func routine(x uint64, y uint64, wg *sync.WaitGroup, in <-chan bool) uint64 {
    defer wg.Done()
    var sum uint64 = 0
    for i := uint64(0); i < x; i++ {
        sum += x * y
        time.Sleep(1 * time.Second)
    }
    <-in
    return sum
}

func main() {
    t1 := time.Now()
    wait := sync.WaitGroup{}
    ch := make(chan bool, 4)
    for i:=0;i<10;i++{
        ch<-true
        wait.Add(1)
        go routine(uint64(5+i),uint64(500+i),&wait,ch)
        fmt.Println("current exist routine nums: ", runtime.NumGoroutine())
    }
    wait.Wait()
    fmt.Println("current exist routine nums: ", runtime.NumGoroutine())
    //fmt.Println(sum)
    elapsed := time.Since(t1)
    fmt.Println("spend time: ", elapsed)
}
