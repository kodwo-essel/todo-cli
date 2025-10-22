package model

import (
	"errors"
	"strings"
	"time"

	"github.com/araddon/dateparse"
)

const (
	StatusPending   = "pending"
	StatusCompleted = "completed"
	StatusArchived  = "archived"
)

var validPriorities = []string{"low", "medium", "high"}

// Task represents a todo task.
type Task struct {
	ID          int
	Title       string
	Description string
	Priority    string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DueAt       *time.Time
	ReminderAt  *time.Time
	Notes       string
	Tags        []string
}

// IsCompleted checks if a task is completed.
func (t *Task) IsCompleted() bool {
	return t.Status == StatusCompleted
}

// TagsString returns tags as a comma-separated string.
func (t *Task) TagsString() string {
	if len(t.Tags) == 0 {
		return ""
	}
	return strings.Join(t.Tags, ",")
}

// Validate checks that the task data is valid before saving.
func (t *Task) Validate() error {
	if strings.TrimSpace(t.Title) == "" {
		return errors.New("title cannot be empty")
	}
	if !isValidPriority(t.Priority) {
		return errors.New("invalid priority, must be one of: low, medium, high")
	}
	if t.Status == "" {
		t.Status = StatusPending
	}
	return nil
}

func isValidPriority(p string) bool {
	for _, v := range validPriorities {
		if v == strings.ToLower(p) {
			return true
		}
	}
	return false
}

// ParseHumanDate attempts to parse common human-friendly date phrases.
func ParseHumanDate(s string) (time.Time, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	now := time.Now()

	switch s {
	case "now":
		return now, nil
	case "today":
		return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()), nil
	case "tomorrow":
		t := now.AddDate(0, 0, 1)
		return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, now.Location()), nil
	case "next week":
		t := now.AddDate(0, 0, 7)
		return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, now.Location()), nil
	default:
		parsed, err := dateparse.ParseAny(s)
		if err != nil {
			return time.Time{}, errors.New("invalid date format")
		}
		return parsed, nil
	}
}
