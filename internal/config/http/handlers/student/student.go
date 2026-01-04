package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/AkshatKumar12/Rest_API-IN-GO/internal/config/storage"
	"github.com/AkshatKumar12/Rest_API-IN-GO/internal/config/types"
	"github.com/AkshatKumar12/Rest_API-IN-GO/internal/config/utils/response"
	"github.com/go-playground/validator/v10"
)

// CRUD OPERATIONS


func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Creating a new student")

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err,io.EOF){
			response.WriteJSON(w,http.StatusBadRequest,response.GeneralError(fmt.Errorf("empty body hai bhaiya")))
			return
		}

		if err != nil {
			response.WriteJSON(w,http.StatusBadRequest,response.GeneralError(err))
		}

		// 	REQUEST VALIDATOIN
		if err := validator.New().Struct(student);err  != nil{
			validateErrs := err.(validator.ValidationErrors)
			response.WriteJSON(w,http.StatusBadRequest,response.ValidationError(validateErrs))
			return 
		}


		lastid,err := storage.CreateStudent(
			student.Name,
			student.Email,
			student.Age,
		)

		slog.Info("user created successfully",slog.String("userId",fmt.Sprint(lastid)))

		if err != nil{
			response.WriteJSON(w,http.StatusInternalServerError,err)
			return
		}


		response.WriteJSON(w,http.StatusCreated,map[string]int64{"id":lastid})
	}
}

