package dom


// Book model of book details
type Book struct {
	ID         		uint64 			`json:"book_id" gorm:"primaryKey"`
	BookName       	string 			`json:"book_name" gorm:"not null;unique" validate:"required"`
	Author       	string 			`json:"author_id" gorm:"not null" validate:"required"`
	NoOfCopies     	uint64 			`json:"no_of_copies" gorm:"not null" validate:"required"`
	Year          	string 			`json:"year" gorm:"not null" validate:"required,len=4"`
	Publications 	string 			`json:"publications_id" gorm:"not null" validate:"required"`
	Category     	string 			`json:"category_id" gorm:"not null" validate:"required"`
	Description    	string 			`json:"description" gorm:"not null" validate:"required"`
}