package main

import (
	"fmt"
	"flag"

)

var ip1 string
var subnetIp string
var hits int


func init(){
	flag.StringVar(&ip1, "IpAddress1", "10.2.1.1","Ip Address")
	flag.StringVar(&subnetIp, "Subnet", "255.255.0.0","Subnet Mask")
	flag.IntVar(&hits, "hits", 20,"Number of hits")

	flag.Parse()
}

func main(){
	fmt.Println("calling main()")
	fmt.Println("value passed: ", ip1)
	fmt.Println("value passed: ", subnetIp)
	fmt.Println("value passed: ", hits)

}
