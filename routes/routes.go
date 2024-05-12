package routes

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/Arenoit/Buscador-Goland/database"
	"github.com/Arenoit/Buscador-Goland/models"
)

var templates = template.Must(template.ParseGlob("templates/*")) //Enlaza la ruta de otros archivos HTML
// Función http que responde a la petición del cliente
func Index(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index", nil)
}
func Crud(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "crud", nil)
}
func Create(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create", nil)
}

type ID struct {
	ID int
}

func params(r *http.Request) ID {
	path := r.URL.Path[len("/list/"):]
	parts := strings.Split(path, "/")
	ids, _ := strconv.Atoi(parts[0])
	id := ID{
		ID: ids,
	}
	return id
}

// router.HandleFunc("/list", handlers.GetTasks).Methods("GET")
func GetTasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	var tasks []models.Task
	database.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

// router.HandleFunc("/list/{id}", handlers.GetOneTask).Methods("GET")
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	var task models.Task //Call struct
	id := params(r)
	database.DB.First(&task, id)
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	json.NewEncoder(w).Encode(&task)
}

// router.HandleFunc("/insert", handlers.CreateTask).Methods("POST")
func PostTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	var task models.Task //Call struct
	json.NewDecoder(r.Body).Decode(&task)

	createdTask := database.DB.Create(&task)
	err := createdTask.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	json.NewEncoder(w).Encode(&task)
}

// router.HandleFunc("/delete/{id}", handlers.DeleteTask).Methods("DELETE")
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	var task models.Task //Call struct
	id := params(r)
	database.DB.First(&task, id)
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	database.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusOK)
}
