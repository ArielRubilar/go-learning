package repository

type Store interface {
	Save(key string, data []map[string]string) error
}
