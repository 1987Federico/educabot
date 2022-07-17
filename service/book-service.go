package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"github.com/ydhnwb/golang_api/dto"
	"github.com/ydhnwb/golang_api/entity"
	"github.com/ydhnwb/golang_api/repository"
)

//BookService is a ....
type BookService interface {
	Insert(b dto.BookCreateDTO) entity.Driver
	Update(b dto.BookUpdateDTO) entity.Driver
	Delete(b entity.Driver)
	All() []entity.Driver
	FindByID(bookID uint64) entity.Driver
	IsAllowedToEdit(userID string, bookID uint64) bool
}

type bookService struct {
	bookRepository repository.DriverRepository
}

//NewBookService .....
func NewBookService(bookRepo repository.DriverRepository) BookService {
	return &bookService{
		bookRepository: bookRepo,
	}
}

func (service *bookService) Insert(b dto.BookCreateDTO) entity.Driver {
	book := entity.Driver{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.bookRepository.InsertDriver(book)
	return res
}

func (service *bookService) Update(b dto.BookUpdateDTO) entity.Driver {
	book := entity.Driver{}
	err := smapping.FillStruct(&book, smapping.MapFields(&b))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.bookRepository.UpdateDriver(book)
	return res
}

func (service *bookService) Delete(b entity.Driver) {
	service.bookRepository.DeleteDriver(b)
}

func (service *bookService) All() []entity.Driver {
	return service.bookRepository.AllDriver()
}

func (service *bookService) FindByID(bookID uint64) entity.Driver {
	return service.bookRepository.FindDriverByID(bookID)
}

func (service *bookService) IsAllowedToEdit(userID string, bookID uint64) bool {
	b := service.bookRepository.FindDriverByID(bookID)
	id := fmt.Sprintf("%v", b.UserID)
	return userID == id
}
