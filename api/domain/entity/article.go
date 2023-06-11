package entity

type PostArticle struct {
	Title     string   `json:"title"`
	Subtitle  string   `json:"subtitle"`
	Body      string   `json:"body"`
	Thumbnail string   `json:"thumbnail"`
	Tags      []string `json:"tags"`
	IsPublic  bool     `json:"is_public"`
}

type GetArticle struct {
	Title     string   `json:"title"`
	Subtitle  string   `json:"subtitle"`
	Body      string   `json:"body"`
	Thumbnail string   `json:"thumbnail"`
	Tags      []string `json:"tags"`
	IsPublic  bool     `json:"is_public"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

type UpdateArticle struct {
	Title     string   `json:"title"`
	Subtitle  string   `json:"subtitle"`
	Body      string   `json:"body"`
	Thumbnail string   `json:"thumbnail"`
	Tags      []string `json:"tags"`
	IsPublic  bool     `json:"is_public"`
}

type ListArticle struct {
	ArticleID string   `json:"article_id"`
	Title     string   `json:"title"`
	Subtitle  string   `json:"subtitle"`
	Thumbnail string   `json:"thumbnail"`
	Tags      []string `json:"tags"`
	IsPublic  bool     `json:"is_public"`
}

type ArticleID struct {
	ArticleID string `json:"article_id"`
}
