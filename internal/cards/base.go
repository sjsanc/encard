package cards

type Card interface {
	Update(key string) bool
}
