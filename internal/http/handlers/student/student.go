package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	//"syscall/js"

	"github.com/DikshaGupta942/student_API/internal/storage/sqlite"
	"github.com/DikshaGupta942/student_API/internal/types"
	"github.com/DikshaGupta942/student_API/internal/utils/response"
	"github.com/go-playground/validator/v10"
)

// func New() http.HandlerFunc {
func GetByID(storage *sqlite.Sqlite) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			response.WriteJSON(
				w,
				http.StatusBadRequest,
				response.GeneralError(fmt.Errorf("invalid student id")),
			)
			return
		}

		student, err := storage.GetStudentByID(id)
		if err != nil {
			response.WriteJSON(
				w,
				http.StatusNotFound,
				response.GeneralError(fmt.Errorf("student not found")),
			)
			return
		}

		response.WriteJSON(w, http.StatusOK, student)
	}
}

func New(storage *sqlite.Sqlite) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("empty request body")))
			return
		}

		if err != nil {
			response.WriteJSON(w, http.StatusBadRequest, response.GeneralError(fmt.Errorf("invalid request payload")))
			// 	http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		//Validation of Request Body can be added here

		if err := validator.New().Struct(student); err != nil {

			validateErrs := err.(validator.ValidationErrors)
			response.WriteJSON(w, http.StatusBadRequest, response.ValidationError(validateErrs))
			return
		}
		//retrun

		//slog.Info("Creating Student record")

		//w.Write([]byte("Welcome to Student API"))

		//response.WriteJSON(w, http.StatusCreated, map[string]string{"Success": "Student record created successfully"})
		// if r.Method != http.MethodGet {
		// 	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		// 	return
		slog.Info("Creating Student record")

		if err := storage.CreateStudent(student); err != nil {
			slog.Error("Failed to create student", slog.String("error", err.Error()))
			response.WriteJSON(
				w,
				http.StatusInternalServerError,
				response.GeneralError(fmt.Errorf("failed to create student")),
			)
			return
		}

		response.WriteJSON(
			w,
			http.StatusCreated,
			map[string]string{"Success": "Student record is added and created successfully"},
		)

	}

}
func GetAll(storage *sqlite.Sqlite) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		students, err := storage.GetAllStudents()
		if err != nil {
			response.WriteJSON(
				w,
				http.StatusInternalServerError,
				response.GeneralError(fmt.Errorf("failed to fetch students")),
			)
			return
		}

		response.WriteJSON(w, http.StatusOK, students)
	}
}
