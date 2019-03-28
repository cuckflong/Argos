package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/go-ini/ini"
)

const SUCCESS = 0
const PROGRESS = 1
const ERROR = 2
const GET = 3
const POST = 4

var url string
var query string
var charFile string
var identifier string = "ARGOS"
var config string
var method int

var charList = []string{
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
	" ", "{", "}", "#", "$", "!", "@",
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
	flag.StringVar(&config, "config", "", "Config file")
	flag.StringVar(&charFile, "list", "", "The file for character list")
	flag.Parse()
	if config == "" {
		printStatus("Config file not provided", ERROR)
		os.Exit(1)
	}
	cfg, err := ini.Load(config)
	if err != nil {
		printStatus("Config file not found", ERROR)
		os.Exit(1)
	}
	url = cfg.Section("request").Key("url").String()
	method = cfg.Section("request").Key("method").String()

	keys := cfg.Section("cookies").Keys()
	fmt.Println(keys)
}

func printStatus(msg string, status int) {
	switch status {
	case SUCCESS:
		color.Green("[+] %s", msg)
	case PROGRESS:
		color.Blue("[-] %s", msg)
	case ERROR:
		color.Red("[!] %s", msg)
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
	//req, err := http.NewRequest("GET", url, nil)
	//q := req.URL.Query()
	return
}

func postRequest() {
	//req, err := http.NewRequest("POST", url, nil)

	return
}

func main() {

}
