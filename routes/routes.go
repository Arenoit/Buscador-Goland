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

type ID struct {
	ID int
}

func params(rute string, r *http.Request) ID {
	path := r.URL.Path[len(rute):]
	parts := strings.Split(path, "/")
	ids, _ := strconv.Atoi(parts[0])
	id := ID{
		ID: ids,
	}
	return id
}

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
func Update(w http.ResponseWriter, r *http.Request) {
	id := params("/project/", r)
	templates.ExecuteTemplate(w, "update", id)
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
	id := params("/list/", r)
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
	// Es necesario parsear el formulario HTML porque sino no se reconoce el campo fecha por Formulario
	erro := r.ParseForm()
	if erro != nil {
		http.Error(w, "Error al analizar el formulario", http.StatusInternalServerError)
		return
	}
	// Obtención de datos por HTML
	var task models.Task //Call struct
	task.Fecha = r.Form.Get("fecha")
	task.Titulo = r.Form.Get("titulo")
	task.Autor = r.Form.Get("autor")
	json.NewDecoder(r.Body).Decode(&task)
	/*//Vale por POSTMAN pero no por HTML
	fecha, err := time.Parse(time.RFC3339, task.Fecha)
	if err != nil {
		fecha, err = time.Parse("2006-01-02", task.Fecha)
		if err != nil {
			http.Error(w, "Formato de fecha no válido", http.StatusBadRequest)
			return
		}
	}
	task.Fecha = fecha.Format("2006-01-02")
	*/
	createdTask := database.DB.Create(&task)
	if createdTask.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(createdTask.Error.Error()))
	}
	json.NewEncoder(w).Encode(&task)
}

// router.HandleFunc("/update", handlers.CreateTask).Methods("PUT")
func PutTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	erro := r.ParseForm()
	if erro != nil {
		http.Error(w, "Error al analizar el formulario", http.StatusInternalServerError)
		return
	}
	// Obtención de datos por HTML
	var task models.Task
	id, _ := strconv.Atoi(r.Form.Get("ID"))
	fecha := ""
	if r.Form.Get("fecha") != "" {
		fecha = r.Form.Get("fecha")
	} else {
		fecha = "0001-01-01"
	}

	titulo := r.Form.Get("titulo")
	autor := r.Form.Get("autor")
	database.DB.First(&task, id)
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	// Actualizar la tarea en la base de datos
	updates := map[string]interface{}{
		"Fecha":  fecha,
		"Titulo": titulo,
		"Autor":  autor,
	}

	// Actualizar la tarea en la base de datos
	if err := database.DB.Model(&task).Where("id = ?", id).Updates(updates).Error; err != nil {
		http.Error(w, "Error al actualizar la tarea en la base de datos", http.StatusInternalServerError)
		return
	}
	// Respondemos con un código de estado 200 (OK)
	w.WriteHeader(http.StatusOK)
	// Ahora decodificamos el cuerpo de la solicitud
	json.NewEncoder(w).Encode(&task)
}

// router.HandleFunc("/delete/{id}", handlers.DeleteTask).Methods("DELETE")
func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	var task models.Task //Call struct
	id := params("/delete/", r)
	database.DB.First(&task, id)
	if task.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	database.DB.Unscoped().Delete(&task)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted"))
}
