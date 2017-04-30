package main

import (
	"log"
	"runtime"
	"sync"
	"time"
    "github.com/henson/ProxyPool/api"
	"github.com/henson/ProxyPool/getter"
	"github.com/henson/ProxyPool/models"
	"github.com/leozvc/ProxyPool/storage"
	"github.com/leozvc/ProxyPool/file"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ipChan := make(chan *models.IP, 1000)
	conn := storage.NewStorage()

	// Start HTTP
	go func() {
		api.Run()
	}()

	//write to file
	go func() {
		file.Run()
	}()

	// Check the IPs in DB
	go func() {
		storage.CheckProxyDB()
	}()

	// Check the IPs in channel
	for i := 0; i < 100; i++ {
		go func() {
			for {
				storage.CheckProxy(<-ipChan)
			}
		}()
	}

	// Start getters to scraper IP and put it in channel
	for {
		x := conn.Count()
		log.Printf("Chan: %v, IP: %v\n", len(ipChan), x)
		if len(ipChan) < 200 {
			go run(ipChan)
		}
		time.Sleep(10 * time.Minute)
	}
}


func run(ipChan chan<- *models.IP) {
	var wg sync.WaitGroup
	funs := []func() []*models.IP{
		getter.Data5u,
		//getter.IP66,
		getter.KDL,
		getter.GBJ,
		getter.Xici,
		getter.XDL,
		getter.IP181,
		getter.YDL,
		//getter.PLP,
	}
	for _, f := range funs {
		wg.Add(1)
		go func(f func() []*models.IP) {
			temp := f()
			for _, v := range temp {
				ipChan <- v
			}
			wg.Done()
		}(f)
	}
	wg.Wait()
	log.Println("All getters finished.")
}
