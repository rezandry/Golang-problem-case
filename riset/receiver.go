package main

import (
	"log"
	"sync"
)

const Limit = 2

func multiply(a, b int) int {
	return a * b
}

type Struct struct {
	kota  []string
	kotac string
	url   string
}

func main() {
	log.SetFlags(log.Ltime) // format log output hh:mm:ss

	wg := sync.WaitGroup{}
	queue := make(chan Struct)

	for worker := 0; worker < Limit; worker++ {
		wg.Add(1)

		go func(worker int) {
			defer wg.Done()

			for work := range queue {
				doWork(worker, work) // blocking wait for work
			}
		}(worker)
	}
	var s Struct
	s.kota = []string{"Solo", "Karanganyar", "Sragen"}
	s.url = "sayareza.com"
	for _, k := range s.kota {
		log.Printf("Work %s enqueued\n", k)
		s.kotac = k
		queue <- s
	}

	close(queue)

	wg.Wait()
}

func doWork(i int, s Struct) {
	log.Printf("Worker %d working on %s\n", i, s.kotac)
	log.Printf("Results %d working is %s\n", i, s.url)
}
