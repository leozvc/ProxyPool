package file

import (
	"log"
	"strings"
    "time"
	"io/ioutil"
	"github.com/leozvc/ProxyPool/storage"
	"github.com/leozvc/ProxyPool/util"
)

// VERSION for this program
//const VERSION = "/v1"

// Run for request
func Run() {
    
    
	Config := util.NewConfig()
	for {
	    //自动刷新代理列表
	    go GetProxys(Config.Output_file.Filepath)
		
        time.Sleep(time.Duration(Config.Output_file.Interval) * time.Second)
	}

}

// ProxyHandler .
func GetProxys(fp string) {
	filename := fp
 
	proxys := storage.ProxyAll()
	var pl []string
	for _, p := range proxys {
            types := strings.Split(p.Type, ",")
            //log.Println(types[0], p.Data, p.ID)
            s := strings.Join([]string{types[0],"://", p.Data}, "")
            pl = append(pl, s) 
	}
 	var d1 = []byte(strings.Join(pl, "\r\n"));
        err := ioutil.WriteFile(filename, d1, 0777)
	if err == nil {
	    log.Println("更新代理文件成功,总条数",len(proxys))    
	}

}
