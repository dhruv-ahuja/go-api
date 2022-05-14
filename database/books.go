// contains the code relating to the CRUD operations to be perfomed
// on the database

package database

type Book struct {
	// all field tags must be exported through capitalization
	// if they are to be used in the JSON encodings
	ID int `json:"id"`
	// `omitempty` excludes the field from the JSON encoding if its empty
	ISBN   int    `json:"isbn,omitempty"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Year   int    `json:"year,omitempty"`
}
