package entity

import "time"

type KeyValueStorage interface {
	Get(key string) (string, error)
	Set(key, value string, expire *time.Duration) error
	Remove(key string) error
}
