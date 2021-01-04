import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hola mundo desde el servidor web con GO")
}

func Planets(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Pagina de los planetas")
}
