package main

import (
	"net/http"

	"asyons.com/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.NotFound = http.FileServer(http.Dir("public"))
	router.GET("/", controllers.Index)
	router.GET("/item/:id", basicAuth(controllers.Detail))
	router.GET("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)
	http.ListenAndServe(":8000", router)
}

func basicAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if ok := check(r); ok {
			h(w, r, ps)
		} else {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

func check(r *http.Request) bool {
	cookie, _ := r.Cookie("username")
	if cookie == nil {
		return false
	}
	return true
}
