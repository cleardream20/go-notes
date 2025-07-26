package main

import (
    "fmt"
    "os"

    "golang.org/x/net/html"
)

func main() {
	fmt.Println(WinAndWin("win"))
	fmt.Println(add(1, 2))

	fmt.Printf("%T\n",add)
	getLinks()

	win := "win"
    winAndWin := doubleWin(win)
    fmt.Println(winAndWin(win))
}

func declare(win string) {
	fmt.Println(win)
}

func WinAndWin(win string) (string, string) {
	declare(win)
	return win, win
}

func add(a, b int) (res int) {
	res = a + b
	return // 等价于return a + b
}

func getLinks() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func doubleWin(win string) func(string) string {
    return func(win string) string {
        return win + win
    }
}



