/*
Representa la declaración en la carpeta actual, no debe que haber
packages distintos en la misma carpeta
*/
package main

import (
	"log"
	"net/http"

	"github.com/Arenoit/Buscador-Goland/database"
	"github.com/Arenoit/Buscador-Goland/models"
	"github.com/Arenoit/Buscador-Goland/routes"
)

/*{{}}*/
type Album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Year   string `json:"year"`
}

/*
// Slicin
	var albums = []Album{
		{ID: "1", Title: "Desarrollador Web", Artist: "programador", Year: "2014"},
		{ID: "2", Title: "Awords Eminem", Artist: "Eminem", Year: "2017"},
		{ID: "3", Title: "Dayanna singing", Artist: "Dayanna", Year: "20008"},
	} //Valores que se llena en el conjunto Users
*/

func main() {
	database.Connection()
	database.DB.AutoMigrate(models.Task{})
	//Depuración de rutas del servidor
	http.HandleFunc("/", routes.Index)
	http.HandleFunc("/crud", routes.Crud)
	http.HandleFunc("/create", routes.Create)
	http.HandleFunc("/project/", routes.Update)
	//API - considerando que HTML solo tiene GET y POST
	http.HandleFunc("/list", routes.GetTasksHandler)      //GET
	http.HandleFunc("/list/", routes.GetTaskHandler)      //GET/{id}
	http.HandleFunc("/insert", routes.PostTaskHandler)    //POST
	http.HandleFunc("/update", routes.PutTaskHandler)     //PUT
	http.HandleFunc("/delete/", routes.DeleteTaskHandler) //DELETE
	//Buscador Autocompletado
	http.HandleFunc("/api/", routes.Searcher)
	// Task Routes
	http.HandleFunc("/style.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "css/style.css")
	})
	http.HandleFunc("/formain.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "js/formain.js")
	})
	http.HandleFunc("/ajax.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "js/ajax.js")
	})
	//Pragma de ejecución
	log.Println("Run server: http://localhost:3000")
	//Levantando servidor
	log.Fatal(http.ListenAndServe(":3000", nil))
}

/* OPERADOR :=
Declara una variable sin tipar manualmente
	nombre1 := "David"
en ves de usar
	var nombre1 string = "David"
	const pi float32 = 3.141592
	const pi float64 = 3.141592
Se deja de user := cuando la variable ya estaba definida
	const pi  = 3.141592
Hay que tener cuidado con las variables dado queno existe undefined en Goland
*/
