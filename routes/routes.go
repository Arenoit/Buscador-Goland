package routes

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

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
func Dates(fecha string) string {
	onlyDate, err := time.Parse(time.RFC3339, fecha)
	if err != nil {
		onlyDate, err = time.Parse("2006-01-02", fecha)
		if err != nil {
			log.Println("Formato de fecha no válido")
			return ""
		}
	}
	return onlyDate.Format("2006-01-02")
}

var templates = template.Must(template.ParseGlob("templates/*")) //Enlaza la ruta de otros archivos HTML
// Función http que responde a la petición del cliente
func Index(w http.ResponseWriter, r *http.Request) {
	erro := r.ParseForm()
	if erro != nil {
		http.Error(w, "Error al analizar el formulario", http.StatusInternalServerError)
		return
	}
	var tasks []models.Task
	seeker := r.Form.Get("seeker")
	if seeker != "" {
		database.DB.Where("Titulo = ?", seeker).Find(&tasks)
	} else {
		database.DB.Find(&tasks)
		// Formatea la fecha en cada tarea
		for i := range tasks {
			tasks[i].Fecha = Dates(tasks[i].Fecha)
		}
	}
	templates.ExecuteTemplate(w, "index", tasks)
}
func Crud(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	database.DB.Find(&tasks)
	templates.ExecuteTemplate(w, "crud", tasks)
}
func Create(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "create", nil)
}
func Update(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	id := params("/project/", r)
	database.DB.First(&task, id)
	fecha := Dates(task.Fecha)
	task.Fecha = fecha
	templates.ExecuteTemplate(w, "update", task)
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

	createdTask := database.DB.Create(&task)
	if createdTask.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(createdTask.Error.Error()))
	}
	//json.NewDecoder(r.Body).Decode(&task)
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

func Searcher(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path[len("/api/"):]
	parts := strings.Split(path, "/")
	seeker := parts[0]
	erro := r.ParseForm()
	if erro != nil {
		http.Error(w, "Error al analizar el formulario", http.StatusInternalServerError)
		return
	}
	var items []string
	var query string = "SELECT DISTINCT titulo FROM tasks WHERE "

	for i, keyword := range seeker {
		if string(keyword) != "" {
			query += "(titulo LIKE '%" + string(keyword) + "%') OR " +
				"(autor LIKE '%" + string(keyword) + "%')"

			if i < len(seeker)-1 {
				query += " OR "
			}
		}
	}
	query += ";"
	rows, err := database.DB.Raw(query).Rows()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var titulo string
		if err := rows.Scan(&titulo); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		items = append(items, titulo)
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Serializar los resultados en formato JSON
	jsonData, err := json.Marshal(items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Establecer el encabezado Content-Type en application/json
	w.Header().Set("Content-Type", "application/json")
	// Escribir los datos JSON en el cuerpo de la respuesta HTTP
	w.Write(jsonData)
}
