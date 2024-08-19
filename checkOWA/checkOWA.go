package main

import (
	"fmt"
	"goInpy/net/http"
	"os"
	"strings"
)

func CheckOWA(url, username, password string) {
	url1 := "https://" + url + "/owa/auth.owa"

	var headers = map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36",
	}
	payload := fmt.Sprintf("destination=https://%s/owa&flags=4&forcedownlevel=0&username=%s&password=%s&passwordText=&isUtf8=1", url, username, password)
	BProxy := "127.0.0.1:8080"

	r, _ := http.Post(url1, BProxy, headers, strings.NewReader(payload), false)
	defer r.Body.Close()

	// if _, in := r.Header.Get("Cookie")["X-OWA-CANARY"]; in {
	if strings.Contains(r.Header.Get("Cookie"), "X-OWA-CANARY") {
		fmt.Printf("[+] Valid:%s  %s\n", username, password)
	} else {
		fmt.Println("[!] Login error")
	}
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println(`[!] Wrong parameter
	[checkOWA]
	Use to check the valid account of Exchange by connecting to OWA.
	Author: Gr%1m  (rewrite by golang__from 3gstudent)
	Usage:`)
		fmt.Printf("\t%s <url> <user> <password>\n", os.Args[0])
		fmt.Printf("Eg.")
		fmt.Printf("%s 192.168.1.1 user1 password1\n", os.Args[0])
		return
	} else {
		fmt.Println(os.Args[1], os.Args[2], os.Args[3])
		CheckOWA(os.Args[1], os.Args[2], os.Args[3])
	}

}
