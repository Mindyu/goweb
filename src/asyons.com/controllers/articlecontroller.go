package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"asyons.com/services"
	"github.com/julienschmidt/httprouter"
)

// Index is yes
func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	temp, _ := template.ParseFiles("public/views/index.html", "public/views/_header.html", "public/views/_toper.html", "public/views/_list.html", "public/views/_footer.html")
	result, err := services.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	err = temp.Execute(w, result)
	if err != nil {
		fmt.Fprintf(w, "%q", err)
	}
}

// Detail is yes
func Detail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	temp, _ := template.ParseFiles("public/views/detail.html", "public/views/_header.html", "public/views/_toper.html", "public/views/_footer.html")
	result, err := services.Detail(ps.ByName("id"))
	if err != nil {
		log.Fatal(err)
	}

	err = temp.Execute(w, result)
	if err != nil {
		fmt.Fprintf(w, "%q", err)
	}
}

// Login is yes
func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, 1)
	cookie := http.Cookie{Name: "username", Value: "jimmy", Expires: expiration}
	http.SetCookie(w, &cookie)
}

// Logout is yes
func Logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, -1)
	cookie := http.Cookie{Name: "username", Value: "jimmy", Expires: expiration}
	http.SetCookie(w, &cookie)
}
