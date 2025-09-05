package dto

type Album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   int    `json:"year"`
}

var Albums = []Album{
	{ID: "1", Title: "The Slim Shady LP", Artist: "Eminem", Year: 1999},
	{ID: "2", Title: "Recovery", Artist: "Eminem", Year: 2010},
	{ID: "3", Title: "Encore", Artist: "Eminem", Year: 2004},
}
