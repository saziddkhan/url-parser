package main

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"strings"
)

func main() {

	var links = []string{"https://www.w3schools.com/go/index.php"}
	for _, link := range links {

		fmt.Println("URL:", link)

		u, err := url.Parse(link)
		if err != nil {
			log.Println(err)
			continue
		}

		parserURL(u)
		fmt.Println(strings.Repeat("#", 50))
		fmt.Println()
	}
}

func parserURL(u *url.URL) {
	fmt.Println("Scheme:", u.Scheme)
	if u.Opaque != "" {
		fmt.Println("Opaque:", u.Opaque)
	}
	if u.User != nil {
		fmt.Println("Username:", u.User.Username())
		if pwd, ok := u.User.Password(); ok {
			fmt.Println("Password:", pwd)
		}
	}
	if u.Host != "" {
		if host, port, err := net.SplitHostPort(u.Host); err == nil {
			fmt.Println("Host:", host)
			fmt.Println("Port:", port)
		} else {
			fmt.Println("Host:", u.Host)
		}
	}
	if u.Path != "" {
		fmt.Println("Path:", u.Path)
	}
	if u.RawQuery != "" {
		fmt.Println("RawQuery:", u.RawQuery)
		m, err := url.ParseQuery(u.RawQuery)
		if err == nil {
			for k, v := range m {
				fmt.Printf("Key: %q Values: %q\n", k, v)
			}
		}
	}
	if u.Fragment != "" {
		fmt.Println("Fragment:", u.Fragment)
	}
}
