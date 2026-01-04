package storage

import(
	"github.com/AkshatKumar12/Rest_API-IN-GO/internal/config/types"
)
type Storage interface{
	CreateStudent(name string,email string, age int)(int64, error)
	GetStudentById(id int64) (types.Student,error)
}