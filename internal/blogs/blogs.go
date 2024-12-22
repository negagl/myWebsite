package blogs

type Blog struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var blogs = []Blog{
	{1, "First Blog", "This is the first blog post"},
}
