package database

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}
func List() []Product {
	return productList
}
func Get(productID int) *Product {
	for _, product := range productList {
		if product.ID == productID {
			return &product
		}
	}
	return nil
}
func Update(product Product) {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
		}
	}
}
func Delete(productID int) {
	var newProductList []Product
	for _, p := range productList {
		if p.ID != productID {
			newProductList = append(newProductList, p)
		}
	}
	productList = newProductList
}

func init() {
	p1 := Product{
		ID:          1,
		Name:        "Wireless Headphones",
		Image:       "https://images.unsplash.com/photo-1505740420928-5e560c06d30e?w=300&h=200&fit=crop",
		Price:       99.99,
		Description: "High-quality wireless headphones with noise cancellation",
		Category:    "Electronics",
	}
	p2 := Product{
		ID:          2,
		Name:        "Cotton T-Shir",
		Image:       "https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=300&h=200&fit=crop",
		Price:       19.99,
		Description: "Comfortable 100% cotton t-shirt in various colors",
		Category:    "Clothing",
	}
	p3 := Product{
		ID:          3,
		Name:        "Coffee Mug",
		Image:       "https://images.unsplash.com/photo-1520031473529-2c06dea61853?q=80&w=687&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
		Price:       19.99,
		Description: "Ceramic coffee mug perfect for your morning brew",
		Category:    "Home",
	}
	p4 := Product{
		ID:          4,
		Name:        "Running Shoes",
		Image:       "https://images.unsplash.com/photo-1542291026-7eec264c27ff?w=300&h=200&fit=crop",
		Price:       79.99,
		Description: "Lightweight running shoes for all terrains",
		Category:    "Sports",
	}
	productList = append(productList, p1)
	productList = append(productList, p2)
	productList = append(productList, p3)
	productList = append(productList, p4)
}
