package main

import (
	"flag"
	"fmt"
	"github.com/Unknwon/goconfig"
	"log"
	"os"
)

type Mysql struct {
	Host string
	User string
	Pwd  string
	Db   string
}

type System struct {
	Model string
	Logfile string
	Port string
}

var (
	Mc Mysql = Mysql{}
	Sc System = System{}
	Config *string = flag.String("config", "config.ini", "使用 -config=uconfig.ini文件路径")
)

func ConfReadFile(filename string) {
	cfg, err := goconfig.LoadConfigFile(filename)
	if err != nil {
		panic("错误")
	}

	sys, _ := cfg.GetSection("system")
	Sc.Model = sys["model"]
	Sc.Logfile = sys["logfile"]
	Sc.Port = sys["port"]

	mysql, _ := cfg.GetSection("mysql")
	Mc.Db = mysql["db"]
	Mc.Host = mysql["host"]
	Mc.Pwd = mysql["pwd"]
	Mc.User = mysql["user"]
}

//读取配置文件，判断是否为调试模式
func init() {
	ConfReadFile(*Config)
	if Sc.Model == "release" {
		outfile, err := os.OpenFile(Sc.Logfile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println(*outfile, "open failed")
			os.Exit(1)
		}
		log.SetOutput(outfile) //设置log的输出文件，不设置log输出默认为stdout
		log.SetPrefix("TRACE: ")
		log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	}
}