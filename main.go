package main

import (
	"github.com/junhaideng/51job/job"
	"flag"
	"os"

	"go.uber.org/zap"
)


var filename string 
var keyword string 

func init(){
	flag.StringVar(&filename, "f", "job.csv", "爬取的数据存放文件名")
	flag.StringVar(&keyword, "k", "go", "keyword")
}


func main() {
	flag.Parse()
	f, err := os.Create(filename)
	if err != nil {
		println(err)
		return
	}
	j := job.NewJob(zap.NewExample(), f, keyword)
	j.Run()
}
