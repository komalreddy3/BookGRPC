package main

import (
	"GRPC/GRPC"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

type server struct {
	GRPC.UnimplementedBookServiceServer
}

var books []GRPC.Book

//Get Books
func (s *server) GetBooks(ctx context.Context, req *GRPC.NoRequest) (*GRPC.BookListResponse, error) {
	for _, item := range books {
		fmt.Println(item.Id)
		fmt.Println(item.Isbn)
		fmt.Println(item.Title)
	}
	// Convert the books slice to a slice of pointers to GRPC.Book
	var booksPtr []*GRPC.Book
	for i := range books {
		booksPtr = append(booksPtr, &books[i])
	}
	return &GRPC.BookListResponse{Books: booksPtr}, nil
}

//Get Book
func (s *server) GetBook(ctx context.Context, req *GRPC.BookRequest) (*GRPC.BookResponse, error) {
	for _, item := range books {
		if item.Id == req.Id {
			return &GRPC.BookResponse{Book: &item}, nil
		}
	}

	return &GRPC.BookResponse{}, nil
}

// Create Book

func (s *server) CreateBook(ctx context.Context, req *GRPC.Book) (*GRPC.BookResponse, error) {
	req.Id = strconv.Itoa(rand.Intn(10000000))
	books = append(books, *req)
	fmt.Println(len(books))
	return &GRPC.BookResponse{Book: req}, nil
}

//Update Book
func (s *server) UpdateBook(ctx context.Context, req *GRPC.BookRequest) (*GRPC.BookListResponse, error) {
	for index, item := range books {
		if item.Id == req.Id {
			books = append(books[:index], books[index+1:]...)
			books = append(books, GRPC.Book{Id: req.Id, Isbn: "67363763726732", Title: "Updated book"})

			var booksPtr []*GRPC.Book
			for i := range books {
				booksPtr = append(booksPtr, &books[i])
			}
			return &GRPC.BookListResponse{Books: booksPtr}, nil
		}
	}
	var booksPtr []*GRPC.Book
	for i := range books {
		booksPtr = append(booksPtr, &books[i])
	}
	return &GRPC.BookListResponse{Books: booksPtr}, nil
}

//delete book
func (s *server) DeleteBook(ctx context.Context, req *GRPC.BookRequest) (*GRPC.BookListResponse, error) {
	for index, item := range books {
		if item.Id == req.Id {
			books = append(books[:index], books[index+1:]...)
			fmt.Println(len(books))

			var booksPtr []*GRPC.Book
			for i := range books {
				booksPtr = append(booksPtr, &books[i])
			}
			return &GRPC.BookListResponse{Books: booksPtr}, nil
		}
	}
	fmt.Println(len(books))

	var booksPtr []*GRPC.Book
	for i := range books {
		booksPtr = append(booksPtr, &books[i])
	}
	return &GRPC.BookListResponse{Books: booksPtr}, nil
}

func main() {
	// mock data
	books = append(books, GRPC.Book{Id: "1", Isbn: "6727672", Title: "Book one", Author: &GRPC.Author{
		Firstname: "john",
		Lastname:  "doe",
	}})
	books = append(books, GRPC.Book{Id: "2", Isbn: "67787672", Title: "Book two", Author: &GRPC.Author{
		Firstname: "steve",
		Lastname:  "mith",
	}})
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	GRPC.RegisterBookServiceServer(s, &server{})
	log.Println("Server of grpc running on localhost:8000")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
