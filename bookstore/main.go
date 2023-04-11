package main

import (
	"log"
	// "os"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"github.com/gorilla/mux"

	"encoding/json"
    "net/http"
    "strconv"
)

type Book struct{
	gorm.Model
	ID uint `gorm:"primary_key" json:"id"`
	Title string `json:"title"`
	Description string `json:"desc"`
	Price float64 `json:"price"`
	// Author *Author `json:"author"`
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
    var books []Book
    DB.Find(&books)
    json.NewEncoder(w).Encode(books)
}

func getBookByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var book Book
    DB.First(&book, id)
    json.NewEncoder(w).Encode(book)
}

func createBook(w http.ResponseWriter, r *http.Request) {
    var book Book
    json.NewDecoder(r.Body).Decode(&book)
    DB.Create(&book)
    json.NewEncoder(w).Encode(book)
}

func updateBookByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var book Book
    DB.First(&book, id)
    json.NewDecoder(r.Body).Decode(&book)
    DB.Save(&book)
    json.NewEncoder(w).Encode(book)
}

func deleteBookByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var book Book
    DB.Delete(&book, id)
    json.NewEncoder(w).Encode(book)
}

func searchBookByTitle(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var books []Book
    DB.Where("title LIKE ?", "%"+params["title"]+"%").Find(&books)
    json.NewEncoder(w).Encode(books)
}

var DB *gorm.DB


func init(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	

	var errr error
	dsn := "host=trumpet.db.elephantsql.com user=vgcudlrb password=4piKqE43gj_hwq3m30TIXv0kPVejsdZl dbname=vgcudlrb port=5432 sslmode=disable"
	DB, errr = gorm.Open(postgres.Open(dsn), &gorm.Config{}) 
	if errr != nil {
		log.Fatal("Failed to connect to Database")
	}

    DB.AutoMigrate(&Book{})
}

func main() {
	

	r := mux.NewRouter()

    r.HandleFunc("/books", getAllBooks).Methods("GET")
    r.HandleFunc("/books/{id}", getBookByID).Methods("GET")
    r.HandleFunc("/books", createBook).Methods("POST")
    r.HandleFunc("/books/{id}", updateBookByID).Methods("PUT")
    r.HandleFunc("/books/{id}", deleteBookByID).Methods("DELETE")
    r.HandleFunc("/books/search/{title}", searchBookByTitle).Methods("GET")
    // r.HandleFunc("/books/sort/{order}", sortBooksByCost).Methods("GET")


    log.Println("Listening on :8080...")
    log.Fatal(http.ListenAndServe(":8080", r))
}

