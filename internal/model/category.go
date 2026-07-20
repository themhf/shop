package model

type Category struct {
	ID          int64
	ParentID    *int64
	Title       string
	Slug        string
	Priority    int
	Level       int
	IsActive    bool
	Description string
}
