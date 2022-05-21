package main

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func main() {

	movies = append(movies, Movie{ID: "1", Isbn: "1234", Title: "Naruto", Director: &Director{Firstname: "Sushi", Lastname: "Makamoti"}})

	// load routes
	LoadRoutes()

}
