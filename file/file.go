package file

import (
	"log"
	"strings"
	"io/ioutil"
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
	filename := "./testfile.txt"
 
	proxys := storage.ProxyAll()
	var pl []string
	for _, p := range proxys {
            types := strings.Split(p.Type, ",")
            log.Println(types[0], p.Data, p.ID)
            s := strings.Join([]string{types[0], p.Data}, "")
            pl = append(pl, s) 
	}
 	var d1 = []byte(strings.Join(pl, "\r\n"));
        err := ioutil.WriteFile(filename, d1, 0777)
	if err != nil {
	    log.Println("写入文件成功")    
	}

}
