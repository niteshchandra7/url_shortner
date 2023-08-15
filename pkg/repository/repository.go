package repository

type DatabaseRepo interface {
	Exists(link string) bool
}
