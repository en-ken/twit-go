package main

import (
	"fmt"
	"os"
)

func writeError(format string, a ...interface{}) {
	if len(a) <= 0 {
		os.Stderr.Write([]byte(format))
	} else {
		os.Stderr.Write([]byte(fmt.Sprintf(format, a...)))
	}
}

func main() {
	if len(os.Args) <= 1 {
		writeError("subcommand is necessary.")
		return
	}

	subcmd := os.Args[1]
	switch subcmd {
	case "auth":
		if len(os.Args) != 4 {
			writeError("usage:twit-go auth [consumerKey] [consumerSecret]")
			return
		}
		consumerKey := os.Args[2]
		consumerSecret := os.Args[3]
		procAuth(consumerKey, consumerSecret)
	case "list":
		procList()
	case "post":
		if len(os.Args) != 3 {
			writeError("usage:twit-go post [content]")
			return
		}
		msg := os.Args[2]
		procPost(msg)
	default:
		writeError("Unkown command.")
	}
}

func procAuth(consumerKey, consumerSecret string) {
	auth, err := newAuth(consumerKey, consumerSecret)
	if err != nil {
		writeError("%v\n", err)
		return
	}

	url, err := auth.getAuthzURL()
	if err != nil {
		writeError("%v\n", err)
		return
	}
	fmt.Printf("Access this url and Authorize: %v\n", url)
	fmt.Printf("Enter PIN on the website:\n")
	var pin string
	fmt.Scan(&pin)
	token, secret, err := auth.getTokenAndSecret(pin)
	if err != nil {
		writeError("%v\n", err)
		return
	}

	if err := save(consumerKey, consumerSecret, token, secret); err != nil {
		writeError("%v\n", err)
		return
	}
	fmt.Println("Authorized.")
}

func procList() {
	cred, err := load()
	if err != nil {
		writeError("%v\n", err)
	}
	req := newRequest(cred)
	tweets, err := req.list()
	if err != nil {
		writeError("%v\n", err)
	}

	//display
	for _, t := range tweets {
		fmt.Printf("%v (@%v)\n", t.User.ScreenName, t.User.Name)
		fmt.Printf(" %v\n", t.Text)
		fmt.Printf("------------------------------\n")
	}
}

func procPost(content string) {
	cred, err := load()
	if err != nil {
		writeError("%v\n", err)
	}
	req := newRequest(cred)
	req.post(content)

	fmt.Println("Posted.")
}
