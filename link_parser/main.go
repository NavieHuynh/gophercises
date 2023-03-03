package main

func main() {
	links := GetLinksFromFile("./ex4.html")

	for _, link := range links {
		link.Print()
	}
}
