package week11

type story struct {
	Body     string `json:"body"`
	Catagory string `json:"catagory"`
}

type stories []story
