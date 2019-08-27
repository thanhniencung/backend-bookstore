package repo_impl

import (
	"bookstore/db"
	"bookstore/model"
	"bookstore/repository"
	"context"
	"database/sql"
	"errors"
	"time"
)

type ProductRepoImpl struct {
	sql *db.Sql
}

func NewProductRepo(sql *db.Sql) repository.ProductRepository {
	return &ProductRepoImpl{
		sql: sql,
	}
}

func (u *ProductRepoImpl) AddProduct(context context.Context, product model.Product) (model.Product, error) {

	// Nên kiểm tra thêm trường hợp cate_id có tồn tại không
	sqlCheckCateId := `select exists(select 1 from cate where cate_id = :$1)`
	var cate = model.Cate{}
	err := u.sql.Db.GetContext(context, &cate, sqlCheckCateId, product.CateId)
	if err != nil && err == sql.ErrNoRows {
		return product, errors.New("Danh mục không tồn tại")
	}

	sqlStatement := `
		  INSERT INTO product(
		  		user_id, product_id, product_name, product_image, quantity, 
		  		sold_items, created_at, updated_at, price, cate_id) 
          VALUES(:user_id, :product_id, :product_name, :product_image, :quantity, 
          		 :sold_items, :created_at, :updated_at, :price, :cate_id)	 
     `

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	_, err = u.sql.Db.NamedExecContext(context, sqlStatement, product)
	return product, err
}

func (u *ProductRepoImpl) UpdateProduct(context context.Context, product model.Product) error {
	sqlStatement := `
		UPDATE product
		SET 
			product_name  = (CASE WHEN LENGTH(:product_name) = 0 THEN product_name ELSE :product_name END),
			product_image = (CASE WHEN LENGTH(:product_image) = 0 THEN product_image ELSE :product_image END),
			quantity 	  = (CASE WHEN :quantity = 0 THEN quantity ELSE :quatity END),
			sold_items 	  = (CASE WHEN :sold_items = 0 THEN sold_items ELSE :sold_items END),
			price 		  = (CASE WHEN :price = 0 THEN price ELSE :price END),
			cate_id 	  = (CASE WHEN LENGTH(:cate_id) = 0 THEN cate_id ELSE :cate_id END),
			updated_at 	  = COALESCE (:updated_at, updated_at)
		WHERE 
			product_id 	   = :product_id 
			AND user_id    = :user_id
	`

	product.UpdatedAt = time.Now()

	result, err := u.sql.Db.NamedExecContext(context, sqlStatement, product)
	if err != nil {
		return err
	}

	count, _ := result.RowsAffected()
	if count == 0 {
		return errors.New("Update thất bại")
	}

	return nil
}

func (u *ProductRepoImpl) DeleteProduct(context context.Context, product model.Product) (error) {
	sqlStatement := ` 
		UPDATE product
		SET deleted_at = $1
		WHERE product_id = $2 AND user_id = $3;
	`
	// Trước khi xoá nên kiểm tra sản phẩm này có thuộc về user này hay không
	result, err := u.sql.Db.ExecContext(context, sqlStatement, time.Now(), product.ProductId, product.UserId)
	count, _ := result.RowsAffected()
	if count == 0 {
		return errors.New("Delete thất bại")
	}
	return err
}

func (u *ProductRepoImpl) SelectProductById(context context.Context, productId string) (model.Product, error) {
	var product model.Product

	row := u.sql.Db.QueryRowxContext(context, "SELECT * FROM product WHERE product_id=$1", productId)
	err := row.Err()
	if err != nil {
		return product, err
	}

	err = row.StructScan(&product)
	if err != nil {
		return product, err
	}

	return product, nil
}

func (c *ProductRepoImpl) SelectAll(context context.Context) ([]model.Product, error) {
	products := []model.Product{}
	err := c.sql.Db.SelectContext(context, &products, "SELECT * FROM product ORDER BY created_at ASC")
	if err != nil {
		return products, err
	}

	return products, nil
}
