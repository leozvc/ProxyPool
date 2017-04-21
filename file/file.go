package file

import (
	"log"

	"github.com/leozvc/ProxyPool/storage"
)

// VERSION for this program
//const VERSION = "/v1"

// Run for request
func Run() {
	//for {
	//	time.Sleep(time.Duration(2) * time.Second)
	//自动刷新代理列表
	go GetProxys()
	//}

}

// ProxyHandler .
func GetProxys() {

	a := storage.ProxyAll()
	for _, r := range a {

		log.Println(r.Type, r.Data, r.ID)
	}

}
