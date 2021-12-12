package week11

//News is produced by the service by adding an ID to the article
type News struct {
	// News identification Number
	ID int `json:"id"`
	//Main text of the news
	Body string `json:"body"`
	//Category of the news
	Category string `json:"catagory"`
}
