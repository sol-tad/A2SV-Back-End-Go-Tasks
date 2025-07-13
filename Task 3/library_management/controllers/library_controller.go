package controllers

import (
	"bufio"
	"fmt"
	"library_management/services"
	"os"
	"strconv"
	"strings"
	"library_management/models"

)

var library=services.NewLibrary()

func StartLibraryConsole(){
	scanner:=bufio.NewScanner(os.Stdin)

	for{
		fmt.Println("\n -----| Library Management System |-----")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. Exit")
		fmt.Print("Choose an option: ")
	

	scanner.Scan()
	choice:=scanner.Text()

    switch choice{
		case "1":
			fmt.Print("Enter Book ID, Title, Author (comma-separated): ")
			scanner.Scan()
			data:=strings.Split(scanner.Text(),",")
			id,_:=strconv.Atoi(strings.TrimSpace(data[0]))
			title:=strings.TrimSpace(data[1])
			author:=strings.TrimSpace(data[2])

			library.AddBook(models.Book{ID: id,Title: title,Author: author,Status: "Available"})
			fmt.Println("You have added the book successfully!")
		
		case "2":
			fmt.Print("Enter Book Id to remove: ")
			scanner.Scan()
			id,_:=strconv.Atoi(scanner.Text())
			library.RemoveBook(id)
			fmt.Println("You have removed the book successfully!")
		
		case "3":
			fmt.Print(" Enter Book ID and Member Id (comma-separated): ")
			scanner.Scan()
			data:=strings.Split(scanner.Text(),",")
			bookID,_:=strconv.Atoi(data[0])
			memberID,_:=strconv.Atoi(data[1])

			if err:=library.BorrowBook(bookID,memberID);err!=nil{
				fmt.Println("Error:->", err)
			}else{
				fmt.Printf("You have it now! Congratulation")
			}
		case "4":
			fmt.Print("Enter Book ID and Member ID (comma-separated): ")
			scanner.Scan()
			data := strings.Split(scanner.Text(), ",")
			bookID, _ := strconv.Atoi(strings.TrimSpace(data[0]))
			memberID, _ := strconv.Atoi(strings.TrimSpace(data[1]))
			if err := library.ReturnBook(bookID, memberID); err != nil {
				fmt.Println("Error:", err)
			}else{
				fmt.Printf("You have returned the Book")
			}
		case "5":
			for _,b:=range library.ListAvailableBooks(){
				if len(library.ListAvailableBooks())!=0{

					fmt.Printf("%d: %s by %s\n",b.ID,b.Title,b.Author)
				}else{
					fmt.Println("No book is Available!")
				}
			}
		case "6":
			fmt.Print("Enter Member Id: ")
			scanner.Scan()
			id,_:=strconv.Atoi(scanner.Text())
			for _,b:=range library.ListBorrowedBooks(id){
				fmt.Printf("%d: %s by %s\n", b.ID, b.Title, b.Author)
			}
		case "7":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid Option.")

	}


}}