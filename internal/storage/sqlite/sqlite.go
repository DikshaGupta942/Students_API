package sqlite

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/DikshaGupta942/student_API/internal/types"

	"github.com/DikshaGupta942/student_API/internal/config"
	_ "modernc.org/sqlite" // SQLite driver
)

type Sqlite struct {
	Db *sql.DB
}

func (s *Sqlite) GetStudentByID(id int) (*types.Student, error) {
	var student types.Student

	row := s.Db.QueryRow(
		`SELECT id, name, email, age FROM students WHERE id = ?`,
		id,
	)

	err := row.Scan(
		&student.ID,
		&student.Name,
		&student.Email,
		&student.Age,
	)

	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *Sqlite) CreateStudent(student types.Student) error {
	_, err := s.Db.Exec(
		`INSERT INTO students (name, email, age) VALUES (?, ?, ?)`,
		student.Name,
		student.Email,
		*student.Age,
	)
	return err
}

func New(cfg *config.Config) (*Sqlite, error) {
	log.Println("👉 SQLite DB PATH USED BY APP:", cfg.Storagepath)
	db, err := sql.Open("sqlite", cfg.Storagepath)
	if err != nil {
		return nil, err
	}

	// Verify DB connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to sqlite database: %w", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		age INTEGER NOT NULL
	);`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{Db: db}, nil

}

func (s *Sqlite) GetAllStudents() ([]types.Student, error) {
	rows, err := s.Db.Query(
		`SELECT id, name, email, age FROM students`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []types.Student

	for rows.Next() {
		var student types.Student
		if err := rows.Scan(
			&student.ID,
			&student.Name,
			&student.Email,
			&student.Age,
		); err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}
