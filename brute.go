package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var url string
var query string
var charFile string
var charList = [...]string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
	" ", "{", "}",
}

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
	flag.StringVar(&query, "query", "", "The query for blind SQL injection")
	flag.StringVar(&charFile, "list", "", "The file for character list")
	flag.Parse()
}

func getCharList() {
	if charFile == "" {
		return
	}
	file, err := os.Open(charFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

}

func getRequest() {
	return
}

func postRequest() {
	return
}

func main() {

}
