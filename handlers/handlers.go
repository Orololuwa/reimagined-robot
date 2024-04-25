package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/orololuwa/reimagined-robot/models"
	"github.com/orololuwa/reimagined-robot/repository"
)

type Repository struct {
	User repository.UserRepo
}

var Repo *Repository

func NewHandler(db *sql.DB) {
	r := &Repository{
		User: repository.NewUserRepo(db),
	}
	Repo = r
}

func (m *Repository) CreateAUser(w http.ResponseWriter, r *http.Request) {
	type userBody struct {
		FirstName string `json:"firstName"`
		LastName string `json:"lastName"`
		Email string `json:"email"`
		Password string `json:"password"`
	}

	var body userBody
	
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		responseMap := map[string]interface{}{"message": "error decoding requset body",}

		jsonData, err := json.Marshal(responseMap)
		if err != nil {
			return
		}

		w.Write(jsonData)
		return
	}

	user := models.User{
		FirstName: body.FirstName,
		LastName: body.LastName,
		Email: body.Email,
		Password: body.Password,
	}
	
	id, err := m.User.CreateAUser(user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)		
		responseMap := map[string]interface{}{"message": "error creating user",}

		jsonData, err := json.Marshal(responseMap)
		if err != nil {
			return
		}

		w.Write(jsonData)
		return
	}

	response := map[string]interface{}{"message": "user created successfully", "data": id}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (m *Repository) GetAUser(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		responseMap := map[string]interface{}{"message": "error decoding id",}

		jsonData, err := json.Marshal(responseMap)
		if err != nil {
			return
		}

		w.Write(jsonData)
		return
	}
	
	user, err := m.User.GetAUser(id)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		responseMap := map[string]interface{}{"message": "error getting user",}

		jsonData, err := json.Marshal(responseMap)
		if err != nil {
			return
		}

		w.Write(jsonData)
		return
	}

	response := map[string]interface{}{"message": "user created successfully", "data": user}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}