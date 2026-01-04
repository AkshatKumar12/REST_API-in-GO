package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/AkshatKumar12/Rest_API-IN-GO/internal/config"
	"github.com/AkshatKumar12/Rest_API-IN-GO/internal/config/types"
	_ "github.com/mattn/go-sqlite3" // _ means used indirectlt
)

type SqLite struct {
	Db *sql.DB
}

func New(cfg *config.Config)(*SqLite,error){
	db,err:= sql.Open("sqlite3",cfg.Storage_path)

	if err != nil{
		return nil,err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER
	)`)

	if err != nil{
		return nil,err
	}

	return &SqLite{
		Db: db,
	},nil
}

func (s *SqLite)CreateStudent(name string, email string, age int)(int64, error){
	
	stat, err := s.Db.Prepare("INSERT INTO students (name,email,age) VALUES(?,?,?)")

	if err != nil{
		return 0,err
	}
	defer stat.Close()

	result,err := stat.Exec(name,email,age)
	lastid,err := result.LastInsertId()		// returns the last affected cell
	if err != nil{
		return 0,err					// return empty value
	}

	return lastid,nil
}

func (s*SqLite) GetStudentById(id int64) (types.Student,error){
	stmt, err := s.Db.Prepare("SELECT id,name,email,age FROM students where ID = ? limit 1")//stmt = statement
	if err != nil{
		return types.Student{},err
	}
	defer stmt.Close()

	var student types.Student

	err = stmt.QueryRow(id).Scan(&student.Id,&student.Name,&student.Email,&student.Age)

	if err != nil{

		if err == sql.ErrNoRows{
			return types.Student{},fmt.Errorf("no student found with id = %d",id)
		}
		return types.Student{},fmt.Errorf("query error: %d",err)
	}
	return student,nil
}