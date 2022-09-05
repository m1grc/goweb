package main 
import ( 
	"fmt" 
	"log" 
	"net/http"
	"os" 
)

func main() {
	 http.HandleFunc("/", handler)
	 fmt.Printf("Serveur Web Golang - Version 0.0.16\n\n") 
	 log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil)) 
}

func handler(w http.ResponseWriter, r *http.Request) { 
	log.Printf("Requête de %s", r.RemoteAddr) 
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(w, "<b>Salutations de Kubernetes sur le pod",name, "</b><br><br><br>") 
	fmt.Fprintln(w, "version 0.0.16")
}
