package sqlite

import (
	"database/sql"

	"github.com/AkshatKumar12/Rest_API-IN-GO/internal/config"
	_"github.com/mattn/go-sqlite3"	// _ means used indirectlt

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