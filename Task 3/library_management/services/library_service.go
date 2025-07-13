package services

import(
	"errors"
	"library_management/models"
)

type Library struct{
	Books map[int]models.Book
	Members map[int]models.Member
}

type LibraryManager interface{
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int,memberID int) error
	ReturnBook(bookID int,memberId int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int)[]models.Book
}

func NewLibrary() *Library  {
	return &Library{
		Books: make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

func (lib *Library) AddBook(book models.Book){
	lib.Books[book.ID]=book
}

func (lib *Library) RemoveBook(bookID int)  {
	delete(lib.Books,bookID)
	
}

func (lib *Library) BorrowBook( bookID int,memberID int) error {
	book,ok:=lib.Books[bookID]

	if !ok{
		return errors.New("book not found")
	}
	if book.Status!="Available"{
		return errors.New("book is already borrowed")
	}

	member,ok:=lib.Members[memberID]
	if!ok{
		member=models.Member{ID: memberID,Name: "Member"}
	}
	book.Status="Borrowed"
	lib.Books[bookID]=book

	member.BorrowedBooks=append(member.BorrowedBooks, book)
	lib.Members[memberID]=member

	return nil
}

func (lib *Library) ReturnBook(bookID int,memberId int) error{
	book,ok:= lib.Books[bookID]

	if !ok{
		return errors.New("book not found")
	}
	member,ok:=lib.Members[memberId]
	if !ok{
		return errors.New("member is not found")
	}
	book.Status="Available"
	lib.Books[bookID]=book

	updatedBooks:=[]models.Book{}

	for _,b:=range member.BorrowedBooks{
		if b.ID!=bookID{
			updatedBooks=append(updatedBooks, b)
		}
	}
	member.BorrowedBooks=updatedBooks
	lib.Members[memberId]=member
	return nil
}

func (lib *Library) ListAvailableBooks() []models.Book{
	available:=[]models.Book{}
	for _,book:=range lib.Books{
		if book.Status=="Available"{
			available=append(available, book)
		}
	}

	return available
}


func (lib *Library) ListBorrowedBooks(memberID int)[]models.Book{
	member,ok:=lib.Members[memberID]
	if!ok{
		return []models.Book{}
	}
	return member.BorrowedBooks
}