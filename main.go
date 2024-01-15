package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/proxy"
	"io"
	"log"
	"net/http"
)

func main() {
	target := flag.String("target", "https://eth0.me/", "URL to get")
	proxyAddr := flag.String("proxy", "127.0.0.1:1080", "SOCKS5 proxy address to use")
	username := flag.String("user", "", "username for SOCKS5 proxy")
	password := flag.String("pass", "", "password for SOCKS5 proxy")
	flag.Parse()

	auth := proxy.Auth{
		User:     *username,
		Password: *password,
	}

	dialer, err := proxy.SOCKS5("tcp", *proxyAddr, &auth, nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			Dial: dialer.Dial,
		},
	}

	r, err := client.Get(*target)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
