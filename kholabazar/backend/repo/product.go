package repo

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type Product struct {
	ID          int     `json:"id" db:"id"`
	Name        string  `json:"name" db:"name"`
	Image       string  `json:"image" db:"image"`
	Price       float64 `json:"price" db:"price"`
	Description string  `json:"description" db:"description"`
	Category    string  `json:"category" db:"category"`
}

type ProductRepo interface {
	List() ([]*Product, error)
	Create(product Product) (*Product, error)
	Get(ID int) (*Product, error)
	Update(product Product) (*Product, error)
	Delete(ID int) error
}

type productRepo struct {
	db *sqlx.DB
}

/* Constructor */
func NewProductRepo(db *sqlx.DB) *productRepo {
	return &productRepo{
		db: db,
	}

}

func (r *productRepo) Create(p Product) (*Product, error) {
	query := `
	INSERT INTO products(name,image,price,description,category)
	VALUES(
	$1, $2, $3, $4, $5
	)
	RETURNING id
	`
	row := r.db.QueryRow(query, p.Name, p.Image, p.Price, p.Description, p.Category)
	err := row.Scan(&p.ID)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
func (r *productRepo) List() ([]*Product, error) {
	var productList []*Product
	query := `
	SELECT id,name,description,price,image,category FROM products
	`
	err := r.db.Select(&productList, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return productList, nil

}
func (r *productRepo) Get(ID int) (*Product, error) {
	var product Product
	query := `
	SELECT id,name,price,description,image,category FROM products WHERE id=$1
	`
	err := r.db.Get(&product, query, ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}
func (r *productRepo) Update(p Product) (*Product, error) {
	query := `UPDATE products 
	SET 
	name=$1, 
	image=$2, 
	price=$3, 
	description=$4,
	category=$5
	WHERE 
	id=$6
	`
	row := r.db.QueryRow(query, p.Name, p.Image, p.Price, p.Description, p.Category, p.ID)
	err := row.Err()
	if err != nil {
		return nil, err
	}
	return &p, nil
}
func (r *productRepo) Delete(ID int) error {
	query := `
	DELETE FROM products WHERE id=$1
	`
	_, err := r.db.Exec(query, ID)
	if err != nil {
		return err
	}
	return nil

}
