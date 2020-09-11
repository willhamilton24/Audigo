package datatypes

type Track struct {
	title string
	artist string
	runtime string
}

type Album struct {
	title string
	artist string
	runtime string
	tracks []Track
	release string
}

