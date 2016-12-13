package containers_test

type db interface {
	GetProduct(productCode string) product
	GetProductImages(productCode string) []Image
	GetProductDescription(productCode string) string
}

type product struct {
	Name string
}

type Image struct {
	Url string
}

func getDb() db {
	return db(nil)
}
