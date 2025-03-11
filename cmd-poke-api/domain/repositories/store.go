package repository

type Store interface {
	Save(key string, data []byte) error
}
