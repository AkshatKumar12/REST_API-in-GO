package storage

type storage interface{
	CreateStudent(name string,email string, age int)(int, error)
}