package selenium

type Cache struct {
	Token *Token
}

func NewCache() *Cache {
	return &Cache{Token: &Token{}}
}
