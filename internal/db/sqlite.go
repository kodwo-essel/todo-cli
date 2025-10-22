package db


import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"


	_ "modernc.org/sqlite"
)


func DefaultDBPath() string {
	// simple default in user's home directory
	home, err := os.UserHomeDir()
	if err != nil {
		home = "."
	}
	dir := filepath.Join(home, ".local", "share", "todo")
	_ = os.MkdirAll(dir, 0o755)
	return filepath.Join(dir, "todo.db")
}


func Open(path string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s?_foreign_keys=1", path)
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}
	if err := migrate(db); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}


func migrate(db *sql.DB) error {
	schema := `
	PRAGMA foreign_keys = ON;


	CREATE TABLE IF NOT EXISTS tasks (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	description TEXT,
	priority TEXT DEFAULT 'medium',
	status TEXT DEFAULT 'pending',
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME,
	due_at DATETIME,
	reminder_at DATETIME,
	notes TEXT
	);


	CREATE TABLE IF NOT EXISTS tags (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT UNIQUE NOT NULL
	);


	CREATE TABLE IF NOT EXISTS task_tags (
	task_id INTEGER NOT NULL,
	tag_id INTEGER NOT NULL,
	PRIMARY KEY (task_id, tag_id),
	FOREIGN KEY (task_id) REFERENCES tasks(id) ON DELETE CASCADE,
	FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
	);


	CREATE INDEX IF NOT EXISTS idx_tasks_status ON tasks(status);
	CREATE INDEX IF NOT EXISTS idx_tasks_due ON tasks(due_at);
	`
	_, err := db.Exec(schema)
	return err
}