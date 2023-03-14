package main

import	(
	"fmt"
	// "os"
	"log"
	"net/http"
	"html/template"
	"strconv"
)

type User struct{
	login string
	password string
	
}

type Item struct{
	
	Name string
	ID int
	Price float64
	Aver float64
	Ratings []int
}

type PeoplePageData struct {
    Wares []Item
}
var ratings1 = []int{3, 4, 4}
var ratings2 = []int{5, 2, 4}
var ratings3 = []int{5, 5, 5}
var sysUsers []User
var sysWare = []Item{ {"iPhone", 103, 499.99, aver(ratings1), ratings1}, {"Samsung", 104, 329.99, aver(ratings2), ratings2 }, {"Pixel", 104, 319.99, aver(ratings3), ratings3 } } // {"Xiaomi", 105, 289.99 }, {"ASUS", 106, 299.99 }   }
//  {"iPhone", 103, 499.99} , {"Samsung", 104, 199.99 }, {"Pixel", 104, 319.99 } }

func aver ( it []int )  float64 {
	sum :=0

	for i:=0;i<len(it);i++ {
		sum = sum + it[i]
	}
	return float64( float64(sum) / float64(len(it)) )
}



func registrationHandler(w http.ResponseWriter, r *http.Request) { //u *users
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	tmpl, _:= template.ParseFiles( "front/registration.html" )
	tmpl.Execute(w, nil)

	log := r.FormValue("login")
	pass := r.FormValue("password")

	User1 := User{}
	User1.login = log
	User1.password = pass



	sysUsers = append(sysUsers, User1)
}



func loginHandler(w http.ResponseWriter, r *http.Request ) { //u users, wa warehouse 
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	tmpl, _:= template.ParseFiles( "front/login.html" )
	tmpl.Execute(w, nil)

	

}

func chekH(w http.ResponseWriter, r *http.Request){
	log := r.FormValue("login")
	pass := r.FormValue("password")

	for i:=0;i<len(sysUsers);i++ {
		if (log == sysUsers[i].login) && (pass == sysUsers[i].password) {
			// fmt.Fprintf(w, "Name = %s\n", log)
			// fmt.Fprintf(w, "Address = %s\n", pass)
			// managementHandler(w,r)
			http.Redirect(w, r, "http://localhost:8080/management", http.StatusSeeOther)	
			break
		}	
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)	

}

func managementHandler(w http.ResponseWriter, r *http.Request){ //w warehouse
	if r.URL.Path !="/management" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	tmpl, _:= template.ParseFiles( "front/management.html" )
	tmpl.Execute(w, nil)


}


func listHandler(w http.ResponseWriter, r *http.Request ){ //w warehouse
	if r.URL.Path !="/list" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	// sysWare = []Item{ {"Samsung", 104, 199.99 }, {"iPhone", 103, 499.99 } , {"Pixel", 104, 319.99 } }
	for i:=0;i<len(sysWare);i++ {
		sysWare[i].Aver = aver(sysWare[i].Ratings)
	}


	tmpl, _:= template.ParseFiles( "front/list.html" )
	tmpl.Execute(w, sysWare)

	

}

func searchHandler(w http.ResponseWriter, r *http.Request){ //w warehouse
	if r.URL.Path !="/search" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	var sysWare2 = []Item{}
	searc := r.FormValue("findme")
	for i:=0;i<len(sysWare);i++ {

		if sysWare[i].Name == searc {
			sysWare2 = append(sysWare2, sysWare[i])
		}
		
	}

	tmpl, _:= template.ParseFiles( "front/search.html" )
	tmpl.Execute(w, sysWare2)


}

func filterHandler(w http.ResponseWriter, r *http.Request){ //w warehouse
	if r.URL.Path !="/filter" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}


	var sysWare2 = []Item{}
	lowe := r.FormValue("lower")
	highe := r.FormValue("higher")

	lower, err := strconv.ParseFloat(lowe, 64)
	higher, err := strconv.ParseFloat(highe, 64)

	if err == nil {
		// fmt.Println(5) // 3.1415927410125732
	}

	for i:=0;i<len(sysWare);i++ {

		if sysWare[i].Price >= lower && sysWare[i].Price <= higher {
			sysWare2 = append(sysWare2, sysWare[i])
		}
		
	}


	tmpl, _:= template.ParseFiles( "front/filter.html" )
	tmpl.Execute(w, sysWare2)
}

func ratingHandler( w http.ResponseWriter, r *http.Request ){ //w warehouse
	if r.URL.Path !="/rating" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	name := r.FormValue("name")
	ratin := r.FormValue("rating")

	rating , err := strconv.Atoi(ratin)
	if err != nil{
		// fmt.Println(5)
	}

	for i:=0;i<len(sysWare);i++ {
		if sysWare[i].Name == name {
			sysWare[i].Ratings = append(sysWare[i].Ratings, rating)
		} 

	}

	tmpl, _:= template.ParseFiles( "front/rating.html" )
	tmpl.Execute(w, nil)


}



func helloHandler(w http.ResponseWriter, r *http.Request ){  //us users, w warehouse
	if r.URL.Path !="/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	// if r.Method !="/GET" {
	// 	http.Error(w, "method is not suppoerted", http.StatusNotFound)
	// 	return
	// }
	tmpl, _:= template.ParseFiles( "front/hello.html" )
	tmpl.Execute(w, nil)

		
}



func main(){
	sysUsers  = []User{{"tileu", "tileu"}}

	
	
	fileServer := http.FileServer(http.Dir("./front"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/registration", registrationHandler)
	http.HandleFunc("/management", managementHandler)
	http.HandleFunc("/chek", chekH)

	http.HandleFunc("/list", listHandler )
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/filter", filterHandler)
	http.HandleFunc("/rating", ratingHandler)



	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe(":8080", nil) 
	if err != nil {
		log.Fatal(err)
	}

}