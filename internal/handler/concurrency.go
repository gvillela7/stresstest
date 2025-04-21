package handler

import (
	"fmt"
	tableshow "github.com/gvillela7/stresstest/internal/model"
	"github.com/useinsider/go-pkg/insrequester"
	"sync"
	"time"
)

var (
	start time.Time
)
var Report tableshow.Report

func Concurrency(url string, requests, concurrency int) {
	numJobs := requests
	numWorkers := concurrency
	start = time.Now()
	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(url, start, jobs, &wg)
	}

	close(jobs)
	wg.Wait()
	timeExec := time.Since(start)
	tableshow.TableShow(timeExec, numJobs, Report.Status200, Report.Status404, Report.Status429, Report.Status500)
}
func worker(url string, start time.Time, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Worker processing job %d\n", j)
		makeRequest(url, start)
	}
}
func makeRequest(url string, start time.Time) {
	requester := insrequester.NewRequester().Load()
	response, _ := requester.Get(insrequester.RequestEntity{Endpoint: url})

	Report.Start = start
	switch response.StatusCode {
	case 200:
		Report.Status200 += 1
	case 404:
		Report.Status404 += 1
	case 429:
		Report.Status429 += 1
	case 500:
		Report.Status500 += 1
	}
}
