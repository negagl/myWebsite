package main

import "fmt"

func HomePageContent() string {
	return "<h1>welcome to my site</h1>"
}

func BlogsContent() string {
	listItems := ""
	for _, blog := range blogs {
		listItems += fmt.Sprintf("<li>%s</li>", blog)
	}

	return fmt.Sprintf(`
		<h1>My Blog</h1>
		<ul>
			%s
		</ul>
		`, listItems)
}
