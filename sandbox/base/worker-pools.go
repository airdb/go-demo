package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "processing job", j)
        time.Sleep(time.Second )
        results <- j * 2
    }
}

func main() {

    job_count := 90
    jobs := make(chan int, 10)
    results := make(chan int, 10)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }


    for j := 1; j <= job_count; j++ {
        jobs <- j
    }
    close(jobs)

    for a := 1; a <= job_count; a++ {
        <-results
        fmt.Println("result===")
    }

}
