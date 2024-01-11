package main

import (
	"GRPC/GRPC"
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8000"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client instance
	client := GRPC.NewBookServiceClient(conn)

	// Example calls to the gRPC service methods
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Get Books
	booksResponse, err := client.GetBooks(ctx, &GRPC.NoRequest{})
	if err != nil {
		log.Fatalf("Error calling getBooks: %v", err)
	}
	fmt.Println("Books:", booksResponse.Books)

	// Create Book
	newBook := &GRPC.Book{
		Isbn:  "123456789",
		Title: "New Book",
		Author: &GRPC.Author{
			Firstname: "kave",
			Lastname:  "call",
		},
	}
	createResponse, err := client.CreateBook(ctx, newBook)
	if err != nil {
		log.Fatalf("Error calling createBook times 1: %v", err)
	}
	fmt.Println("Created Book:", createResponse.Book)

	newBook2 := &GRPC.Book{
		Isbn:  "12346789",
		Title: "New Book2",
		Author: &GRPC.Author{
			Firstname: "joseph",
			Lastname:  "christ",
		},
	}
	createResponse2, err := client.CreateBook(ctx, newBook2)
	if err != nil {
		log.Fatalf("Error calling createBook times 1: %v", err)
	}
	fmt.Println("Created Book:", createResponse2.Book)

	// Get Book by ID
	bookID := "1"
	bookResponse, err := client.GetBook(ctx, &GRPC.BookRequest{Id: bookID})
	if err != nil {
		log.Fatalf("Error calling getBook: %v", err)
	}
	fmt.Println("Retrieved Book:", bookResponse.Book)

	//Update book
	booksResponse, err = client.UpdateBook(ctx, &GRPC.BookRequest{Id: bookID})
	if err != nil {
		log.Fatalf("Error calling getBook: %v", err)
	}
	fmt.Println("Update Book", booksResponse.Books)
	// Delete Book
	deleteResponse, err := client.DeleteBook(ctx, &GRPC.BookRequest{Id: bookID})
	if err != nil {
		log.Fatalf("Error calling deleteBook: %v", err)
	}
	fmt.Println("Deleted Books:", deleteResponse.Books)
}
