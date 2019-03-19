package main

import (
	"flag"
	"fmt"
)

var url string
var query string
var charList []string

func banner() {
	fmt.Println()
	fmt.Println("\t █████╗ ██████╗  ██████╗  ██████╗ ███████╗\t /^ ^\\")
	fmt.Println("\t██╔══██╗██╔══██╗██╔════╝ ██╔═══██╗██╔════╝\t/ 0 0 \\")
	fmt.Println("\t███████║██████╔╝██║  ███╗██║   ██║███████╗\tV\\ Y /V")
	fmt.Println("\t██╔══██║██╔══██╗██║   ██║██║   ██║╚════██║\t / - \\")
	fmt.Println("\t██║  ██║██║  ██║╚██████╔╝╚██████╔╝███████║\t |    \\")
	fmt.Println("\t╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚══════╝\t || (__V")
	fmt.Println()
}

func init() {
	banner()
	flag.StringVar(&url, "url", "", "The target URL or Domain")
	flag.StringVar("")
}

func main() {

}
