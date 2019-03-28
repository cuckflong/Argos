package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/go-ini/ini"
)

const SUCCESS = 0
const PROGRESS = 1
const ERROR = 2

type header struct {
	name  string
	value string
}

type cookie struct {
	name  string
	value string
}

type data struct {
	name  string
	value string
}

var headers []header
var cookies []cookie
var dataList []data
var client *http.Client
var result bytes.Buffer
var urlPath string
var charFile string
var identifier string = "ARGOS"
var valid string
var invalid string
var config string
var method string

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
	urlPath = cfg.Section("basic").Key("url").String()
	method = cfg.Section("basic").Key("method").String()
	valid = cfg.Section("response").Key("valid").String()
	invalid = cfg.Section("response").Key("invalid").String()
	if urlPath == "" || method == "" {
		printStatus("Please provide URL and request method", ERROR)
		os.Exit(1)
	}
	if valid == "" && invalid == "" {
		printStatus("Please provide valid or invalid response", ERROR)
		os.Exit(1)
	}
	for _, key := range cfg.Section("headers").KeyStrings() {
		hd := header{name: key, value: cfg.Section("headers").Key(key).String()}
		headers = append(headers, hd)
	}
	for _, key := range cfg.Section("cookies").KeyStrings() {
		ck := cookie{name: key, value: cfg.Section("cookies").Key(key).String()}
		cookies = append(cookies, ck)
	}
	for _, key := range cfg.Section("data").KeyStrings() {
		dt := data{name: key, value: cfg.Section("data").Key(key).String()}
		dataList = append(dataList, dt)
	}

	getCharList()
}

func printStatus(msg string, status int) {
	switch status {
	case SUCCESS:
		color.Green("[+] %s", msg)
	case PROGRESS:
		color.Yellow("[-] %s", msg)
	case ERROR:
		color.Red("[!] %s", msg)
	}
}

func getCharList() {
	if charFile == "" {
		printStatus("Character list not provided, using default list", PROGRESS)
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

func getRequest(payload string) bool {
	req, err := http.NewRequest("GET", urlPath, nil)
	if err != nil {
		printStatus("Error creating GET request", ERROR)
		os.Exit(1)
	}
	for _, hd := range headers {
		new_name := strings.Replace(hd.name, identifier, payload, 1)
		new_value := strings.Replace(hd.value, identifier, payload, 1)
		req.Header.Add(new_name, new_value)
	}
	for _, ck := range cookies {
		new_name := strings.Replace(ck.name, identifier, payload, 1)
		new_value := strings.Replace(ck.value, identifier, payload, 1)
		c := http.Cookie{Name: new_name, Value: new_value}
		req.AddCookie(&c)
	}
	q := req.URL.Query()
	for _, dt := range dataList {
		new_name := strings.Replace(dt.name, identifier, payload, 1)
		new_value := strings.Replace(dt.value, identifier, payload, 1)
		q.Add(new_name, new_value)
	}
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		printStatus("Error performing GET request", ERROR)
		os.Exit(1)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(respBody), valid) && !(invalid != "" && strings.Contains(string(respBody), invalid)) {
		return true
	}
	return false
}

func postRequest(payload string) bool {
	var body bytes.Buffer
	for index, dt := range dataList {
		new_name := strings.Replace(dt.name, identifier, payload, 1)
		new_value := strings.Replace(dt.value, identifier, payload, 1)
		if index != 0 {
			body.WriteString("&")
		}
		body.WriteString(new_name + "=" + new_value)
	}
	req, err := http.NewRequest("POST", urlPath, bytes.NewBuffer(body.Bytes()))
	if err != nil {
		printStatus("Error creating POST request", ERROR)
		os.Exit(1)
	}
	for _, hd := range headers {
		new_name := strings.Replace(hd.name, identifier, payload, 1)
		new_value := strings.Replace(hd.value, identifier, payload, 1)
		req.Header.Add(new_name, new_value)
	}
	for _, ck := range cookies {
		new_name := strings.Replace(ck.name, identifier, payload, 1)
		new_value := strings.Replace(ck.value, identifier, payload, 1)
		c := http.Cookie{Name: new_name, Value: new_value}
		req.AddCookie(&c)
	}
	resp, err := client.Do(req)
	if err != nil {
		printStatus("Error performing POST request", ERROR)
		os.Exit(1)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if strings.Contains(string(respBody), valid) && !(invalid != "" && strings.Contains(string(respBody), invalid)) {
		return true
	}
	return false
}

func main() {
	var proxyStr string = "http://localhost:8080"
	proxyURL, _ := url.Parse(proxyStr)
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	client = &http.Client{
		Transport: transport,
	}
	var found bool
	for {
		found = false
		for _, chr := range charList {
			switch method {
			case "get":
				if getRequest(result.String() + chr) {
					result.WriteString(chr)
					found = true
				}
			case "post":
				if postRequest(result.String() + chr) {
					result.WriteString(chr)
					found = true
				}
			}
			if found {
				printStatus(result.String(), SUCCESS)
				break
			}
		}
		if !found {
			if result.String() == "" {
				printStatus("No result found", ERROR)
				os.Exit(1)
			} else {
				printStatus("Result found", SUCCESS)
				printStatus(result.String(), SUCCESS)
				os.Exit(0)
			}
		}
	}
}
