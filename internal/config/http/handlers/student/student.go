package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/AkshatKumar12/Rest_API-IN-GO/internal/config/types"
	"github.com/AkshatKumar12/Rest_API-IN-GO/internal/config/utils/response"
	"github.com/go-playground/validator/v10"
)

// CRUD OPERATIONS


func New() http.HandlerFunc {
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

		response.WriteJSON(w,http.StatusCreated,map[string]string{"success":"ok"})
	}
}

