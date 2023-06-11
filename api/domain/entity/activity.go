package entity

type Activity struct {
	Title       string        `json:"title"`
	Description string        `json:"description"`
	StartDate   string        `json:"start_date"`
	FinishDate  string        `json:"finish_date"`
	Thumbnail   string        `json:"thumbnail"`
	IsPublic    bool          `json:"is_public"`
	URLs        []ActivityURL `json:"urls"`
}

type ListActivities []ListActivity

type ListActivity struct {
	ActivityID  string        `json:"activity_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	StartDate   string        `json:"start_date"`
	FinishDate  string        `json:"finish_date"`
	Thumbnail   string        `json:"thumbnail"`
	IsPublic    bool          `json:"is_public"`
	URLs        []ActivityURL `json:"urls"`
}

type ActivityURL struct {
	URLDescription string `json:"url_description"`
	URL            string `json:"url"`
}
