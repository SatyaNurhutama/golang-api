package service

import (
	"fmt"
	"golang_api/dto"
	"golang_api/entity"
	"golang_api/repository"
	"log"

	"github.com/mashingan/smapping"
)

type BookService interface {
	Insert(book dto.BookCreateDTO) entity.Book
	Update(book dto.BookUpdateDTO) entity.Book
	Delete(book entity.Book)
	All() []entity.Book
	FindById(bookID uint64) entity.Book
	IsAllowedToEdit(userID string, bookdID uint64) bool
}

type bookService struct {
	bookRepository repository.BookRepository
}

func NewBookService(bookRepository repository.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRepository,
	}
}

func (service *bookService) Insert(book dto.BookCreateDTO) entity.Book {
	bookCreate := entity.Book{}
	err := smapping.FillStruct(&bookCreate, smapping.MapFields(&book))
	if err != nil {
		log.Fatalf("Failed to map %v", err)
	}
	response := service.bookRepository.InsertBook(bookCreate)
	return response
}

func (service *bookService) Update(book dto.BookUpdateDTO) entity.Book {
	bookToUpdate := entity.Book{}
	err := smapping.FillStruct(&bookToUpdate, smapping.MapFields(&book))
	if err != nil {
		log.Fatalf("Failed to map %v", err)
	}
	response := service.bookRepository.UpdateBook(bookToUpdate)
	return response
}

func (service *bookService) Delete(book entity.Book) {
	service.bookRepository.DeleteBook(book)
}

func (service *bookService) All() []entity.Book {
	return service.bookRepository.AllBook()
}

func (service *bookService) FindById(bookID uint64) entity.Book {
	return service.bookRepository.FindBookByID(bookID)
}

func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
	book := service.bookRepository.FindBookByID(bookID)
	id := fmt.Sprintf("%v", book.UserID)
	return userID == id
}
