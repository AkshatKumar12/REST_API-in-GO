package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/AkshatKumar12/Rest_API-IN-GO/internal/config/types"
	"github.com/AkshatKumar12/Rest_API-IN-GO/internal/config/utils/response"
)

// CRUD OPERATIONS

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err,io.EOF){
			response.WriteJSON(w,http.StatusBadRequest,err.Error())
			return
		}

		slog.Info("Creating a new student")
		response.WriteJSON(w,http.StatusCreated,map[string]string{"success":"ok"})
	}
}
