package hamletgrpc

import "github.com/davidbetz/hamletgo"

// ToGRPCDTO converts a Hamlet object into a Hamlet GRPC DTO
func ToGRPCDTO(book *hamlet.Book) *Book {
	return &Book{
		Id:      book.ID,
		Title:   book.Title,
		Authors: book.Authors,
		Metadata: &Metadata{
			Pages:   int32(book.Metadata.Pages),
			Genre:   book.Metadata.Genre,
			Summary: book.Metadata.Summary,
		},
		Published: book.Published,
		Editor:    book.Editor,
	}
}

// FromGRPCDTO converts a Hamlet GRPC DTO into a Hamlet object
func FromGRPCDTO(book *Book) hamlet.Book {
	return hamlet.Book{
		ID:      book.Id,
		Title:   book.Title,
		Authors: book.Authors,
		Metadata: hamlet.Metadata{
			Pages:   int(book.Metadata.Pages),
			Genre:   book.Metadata.Genre,
			Summary: book.Metadata.Summary,
		},
		Published: book.Published,
		Editor:    book.Editor,
	}
}
