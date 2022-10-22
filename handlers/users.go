package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	dto "waysfood/dto/result"
	usersdto "waysfood/dto/users"
	"waysfood/models"
	"waysfood/pkg/mysql/bcrypt"
	"waysfood/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handlerUser struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handlerUser {
	return &handlerUser{UserRepository}
}

// Find Users

func (h *handlerUser) FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepository.FindUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// loop for image
	for i, p := range users {
		users[i].Image = os.Getenv("PATH_FILE") + p.Image
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: users}
	json.NewEncoder(w).Encode(response)
}

// Get User

func (h *handlerUser) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	// for image
	user.Image = os.Getenv("PATH_FILE") + user.Image

	w.WriteHeader(http.StatusOK)
	// response := dto.SuccessResult{Code: "success", Data: user}
	response := dto.SuccessResult{Code: http.StatusOK, Data: user}
	json.NewEncoder(w).Encode(response)
}

// Create User

func (h *handlerUser) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// for image
	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	// request := new(usersdto.CreateUserRequest)
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
	// 	json.NewEncoder(w).Encode(response)
	// 	return
	// }

	request := usersdto.CreateUserRequest{
		FullName: r.FormValue("fullName"),
		Email:    r.FormValue("email"),
		Phone:    r.FormValue("phone"),
		Location: r.FormValue("location"),
		Image:    r.FormValue("Image"),
		Password: r.FormValue("password"),
		Gender:   r.FormValue("gender"),
		Role:     r.FormValue("role"),
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// for bcrypt password
	password, err := bcrypt.HashingPassword(request.Password)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	// Dilihat di model yg dibuat
	//  dan dto models
	user := models.User{
		FullName: request.FullName,
		Email:    request.Email,
		Password: password,
		Phone:    request.Phone,
		Gender:   request.Gender,
		Location: request.Location,
		Image:    filename,
		Role:     request.Role,
	}

	data, err := h.UserRepository.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	user, _ = h.UserRepository.GetUser(user.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

// Update user
func (h *handlerUser) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile")
	filename := dataContex.(string)

	//take pattern data submission
	// request := new(usersdto.UpdateUserRequest)

	request := usersdto.UpdateUserRequest{
		FullName: r.FormValue("fullName"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
		Phone:    r.FormValue("phone"),
		Gender:   r.FormValue("gender"),
		Location: r.FormValue("location"),
		Image:    filename,
		Role:     r.FormValue("role"),
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	// user := models.User{}
	user, err := h.UserRepository.GetUser(int(id))

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	if request.FullName != "" {
		user.FullName = request.FullName
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Password != "" {
		user.Password = password
	}

	if request.Phone != "" {
		user.Phone = request.Phone
	}

	if request.Gender != "" {
		user.Gender = request.Gender
	}

	if request.Location != "" {
		user.Location = request.Location
	}

	if request.Image != "" {
		user.Image = request.Image
	}

	if request.Role != "" {
		user.Role = request.Role
	}

	data, err := h.UserRepository.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}
	user, _ = h.UserRepository.GetUser(user.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

// Delete User

func (h *handlerUser) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.UserRepository.DeleteUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Code: http.StatusOK, Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponse(u models.User) usersdto.UserResponse {
	return usersdto.UserResponse{
		ID:       u.ID,
		Fullname: u.FullName,
		Email:    u.Email,
		Password: u.Password,
		Phone:    u.Phone,
		Location: u.Location,
		Image:    u.Image,
		Role:     u.Role,
		Gender:   u.Gender,
	}
}
