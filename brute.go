package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const SUCCESS = 0
const PROGRESS = 1
const ERROR = 2

var url string
var query string
var charFile string
var identifier string = "ARGOS"
var post bool

var charList = []string{
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
	flag.BoolVar(&post, "post", false, "Use POST request")
	flag.Parse()
}

func printStatus(msg string, status int) {
	switch status {
	case SUCCESS:

	}
}

func getCharList() {
	if charFile == "" {
		return
	}
	file, err := os.Open(charFile)
	if err != nil {
		log.Fatal(err)
	}
	charList = nil
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		charList = append(charList, scanner.Text())
	}
	defer file.Close()
}

func getRequest() {
	resp, err := http.Get(url)
	if err != nil {

	}
	return
}

func postRequest() {
	return
}

func main() {

}
