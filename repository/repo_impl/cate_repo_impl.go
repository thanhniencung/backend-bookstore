package repo_impl

import (
	"bookstore/db"
	"bookstore/model"
	"bookstore/repository"
	"context"
	"errors"
	"time"
)

type CateRepoImpl struct {
	sql *db.Sql
}

func NewCateRepo(sql *db.Sql) repository.CateRepository {
	return &CateRepoImpl{
		sql: sql,
	}
}

func (c *CateRepoImpl) AddCate(context context.Context, cate model.Cate) (model.Cate, error) {
	sqlStatement := `
		  INSERT INTO cate(cate_id, cate_name, created_at, updated_at) 
          VALUES(:cate_id, :cate_name, :created_at, :updated_at)
     `

	cate.CreatedAt = time.Now()
	cate.UpdatedAt = time.Now()

	_, err := c.sql.Db.NamedExecContext(context, sqlStatement, cate)
	return cate, err
}

func (c *CateRepoImpl) UpdateCate(context context.Context, cate model.Cate) error {
	sqlStatement := `
		UPDATE cate
		SET 
			cate_name = :cate_name
			deleted_at = :deleted_at
		WHERE cate_id = :cate_id AND LENGTH(:cate_name) > 0
	`
	cate.UpdatedAt = time.Now()

	result, err := c.sql.Db.NamedExecContext(context, sqlStatement, cate)
	if err != nil {
		return err
	}

	count, _ := result.RowsAffected()
	if count == 0 {
		return errors.New("Update thất bại")
	}

	return nil
}

func (c *CateRepoImpl) DeleteCate(context context.Context, cateId string) (error) {
	sqlStatement := ` 
		UPDATE cate
		SET deleted_at = $1
		WHERE cate_id = $2;
	`
	// Trước khi xoá nên kiểm tra sản phẩm này có thuộc về user này hay không
	result, err := c.sql.Db.ExecContext(context, sqlStatement, time.Now(), cateId)
	count, _ := result.RowsAffected()
	if count == 0 {
		return errors.New("Delete thất bại")
	}
	return err
}

func (c *CateRepoImpl) SelectCateById(context context.Context, cateId string) (model.Cate, error) {
	var cate model.Cate

	row := c.sql.Db.QueryRowxContext(context, "SELECT * FROM cate WHERE cate_id=$1", cateId)
	err := row.Err()
	if err != nil {
		return cate, err
	}

	err = row.StructScan(&cate)
	if err != nil {
		return cate, err
	}

	return cate, nil
}

func (c *CateRepoImpl) SelectAll(context context.Context) ([]model.Cate, error) {
	cates := []model.Cate{}
	err := c.sql.Db.SelectContext(context, &cates, "SELECT * FROM cate ORDER BY created_at ASC")
	if err != nil {
		return cates, err
	}

	return cates, nil
}
