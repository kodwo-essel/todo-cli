package store

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"todo-cli/internal/model"
)

// SQLiteStore handles SQLite database operations.
type SQLiteStore struct {
	DB *sql.DB
}

// NewSQLiteStore opens (or creates) a SQLite database file.
func NewSQLiteStore(path string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	store := &SQLiteStore{DB: db}
	if err := store.init(); err != nil {
		return nil, err
	}
	return store, nil
}

// init creates tables if they don’t exist.
func (s *SQLiteStore) init() error {
	schema := `
	CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		priority TEXT,
		status TEXT,
		created_at DATETIME,
		updated_at DATETIME,
		due_at DATETIME,
		reminder_at DATETIME,
		notes TEXT,
		tags TEXT
	);
	`
	_, err := s.DB.Exec(schema)
	return err
}

// AddTask inserts a new task into the database.
func (s *SQLiteStore) AddTask(t *model.Task) error {
	if err := t.Validate(); err != nil {
		return err
	}
	t.CreatedAt = time.Now()
	query := `
	INSERT INTO tasks (title, description, priority, status, created_at, due_at, reminder_at, notes, tags)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);
	`
	_, err := s.DB.Exec(query, t.Title, t.Description, t.Priority, t.Status,
		t.CreatedAt, t.DueAt, t.ReminderAt, t.Notes, t.TagsString())
	return err
}

// UpdateTask updates an existing task.
func (s *SQLiteStore) UpdateTask(t *model.Task) error {
	now := time.Now()
	t.UpdatedAt = &now

	query := `
		UPDATE tasks
		SET title = ?, description = ?, priority = ?, status = ?, updated_at = ?, due_at = ?, reminder_at = ?, notes = ?, tags = ?
		WHERE id = ?;
	`

	result, err := s.DB.Exec(query,
		t.Title,
		t.Description,
		t.Priority,
		t.Status,
		t.UpdatedAt,
		t.DueAt,
		t.ReminderAt,
		t.Notes,
		t.TagsString(),
		t.ID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("no task found with id %d", t.ID)
	}

	return nil
}


// DeleteTask removes a task by ID.
func (s *SQLiteStore) DeleteTask(id int) error {
	result, err := s.DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("no task found with id %d", id)
	}

	return nil
}


// GetTask retrieves a task by ID.
func (s *SQLiteStore) GetTask(id int) (*model.Task, error) {
	row := s.DB.QueryRow(`SELECT id, title, description, priority, status, created_at, updated_at, due_at, reminder_at, notes, tags FROM tasks WHERE id = ?;`, id)
	return scanTask(row)
}

// ListTasks retrieves all tasks, optionally filtered by status or priority.
func (s *SQLiteStore) ListTasks(filters map[string]string) ([]*model.Task, error) {
	where := []string{}
	args := []any{}

	for k, v := range filters {
		where = append(where, fmt.Sprintf("%s = ?", k))
		args = append(args, v)
	}

	query := `SELECT id, title, description, priority, status, created_at, updated_at, due_at, reminder_at, notes, tags FROM tasks`
	if len(where) > 0 {
		query += " WHERE " + strings.Join(where, " AND ")
	}
	query += " ORDER BY due_at ASC, created_at DESC;"

	rows, err := s.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*model.Task
	for rows.Next() {
		task, err := scanTask(rows)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// MarkComplete updates task status to completed.
func (s *SQLiteStore) MarkComplete(id int) error {
	query := `UPDATE tasks SET status = 'completed' WHERE id = ?`
	res, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("no task found with id %d", id)
	}

	return nil
}

// ArchiveCompleted moves completed tasks to archived status.
func (s *SQLiteStore) ArchiveCompleted() (int64, error) {
	res, err := s.DB.Exec(`UPDATE tasks SET status = ?, updated_at = ? WHERE status = ?;`, model.StatusArchived, time.Now(), model.StatusCompleted)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

// scanTask reads task data from SQL row.
func scanTask(scanner interface {
	Scan(dest ...any) error
}) (*model.Task, error) {
	var t model.Task
	var tags string
	var updatedAt, dueAt, reminderAt sql.NullTime

	err := scanner.Scan(&t.ID, &t.Title, &t.Description, &t.Priority, &t.Status,
		&t.CreatedAt, &updatedAt, &dueAt, &reminderAt, &t.Notes, &tags)
	if err != nil {
		return nil, err
	}

	if updatedAt.Valid {
		t.UpdatedAt = &updatedAt.Time
	}
	if dueAt.Valid {
		t.DueAt = &dueAt.Time
	}
	if reminderAt.Valid {
		t.ReminderAt = &reminderAt.Time
	}
	if tags != "" {
		t.Tags = strings.Split(tags, ",")
	}
	return &t, nil
}
