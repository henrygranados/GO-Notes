package main

import ("fmt"
"net/http"
"github.com/julienschmidt/httprouter")

func main(){
	 
    router := httprouter.New()
    router.GET("/", IndexPage)
    router.GET("/about", AboutPage)
    http.ListenAndServe(":8081", router)
}


func IndexPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}


func AboutPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
      fmt.Println(r.URL.Path) // gets the path or route being used
      fmt.Fprintf(w, "All contacts")
}

