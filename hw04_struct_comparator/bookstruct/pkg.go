package bookstruct

type Book struct {
	iD     int
	title  string
	author string
	year   int
	size   int
	rate   float32
}

func (b Book) ID() int {
	return b.iD
}

func (b Book) Title() string {
	return b.title
}

func (b Book) Author() string {
	return b.author
}

func (b Book) Year() int {
	return b.year
}

func (b Book) Size() int {
	return b.size
}

func (b Book) Rate() float32 {
	return b.rate
}

func (b *Book) SetID(id int) {
	b.iD = id
}

func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}

func (b *Book) SetYear(year int) {
	b.year = year
}

func (b *Book) SetSize(size int) {
	b.size = size
}

func (b *Book) SetRate(rate float32) {
	b.rate = rate
}

func NewBook(id int, title string, author string, year int, size int, rate float32) *Book {
	return &Book{
		iD:     id,
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rate:   rate,
	}
}
