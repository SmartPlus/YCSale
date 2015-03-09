package main

import (
	"YCSale/web"
	"flag"
	"fmt"
)

var PORT int

func main() {
	flag.IntVar(&PORT, "port", 8080, "port number")
	flag.Parse()
	web.Init().RunOnAddr(fmt.Sprintf(":%v", PORT))
}
