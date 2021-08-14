package main

// ItemService exposes methods to CRU (Create, Read, Update)
type ItemService interface {
	Create(key string, value string) error
	Read(key string) (string, error)
	Update(key string, value string) error
}
