package projects

type Project struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Status      string `json:"status"`
}

var Projects = []Project{
	{ID: 1, Title: "Project 1", Description: "This is the first project", URL: "https://github.com/project1", Status: "started"},
}
