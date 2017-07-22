package main

import ("fmt"
"net/http"
"log")

func main(){

	http.HandleFunc("/", handler)
	http.HandleFunc("/contacts", contacts)
	//http.ListenAndServe(":9090", nil)

	err := http.ListenAndServe(":9090", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func handler(w http.ResponseWriter, r *http.Request){
     fmt.Fprint(w, "Welcome to Home Page!\n")	
}

func contacts(w http.ResponseWriter, r *http.Request){ 
    
     fmt.Println(r.URL.Path) // gets the path or current route being used
     fmt.Fprintf(w, "CUSTOMERS\n")

      m := make(map[string]int)

      m["Mike"] = 20
      m["James"] = 30
      m["Mark"] = 25

     for k, v := range m {
         fmt.Fprint(w, "Name: ", k, "\nAge: ", v, "\n------------\n")
      }
}

