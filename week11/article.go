package week11

//Article is news in its raw form
//Article is processed by the service to create news
type Article struct {
	//Body is the main content of the article
	Body string `json:"body"`
	//Category is the theme for the article
	Category string `json:"category"`
}

//Articles is the colletion type for Article
type Articles []Article
