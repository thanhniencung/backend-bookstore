package repo_impl

import (
	"bookstore/db"
	"bookstore/encrypt"
	"bookstore/model"
	"bookstore/repository"
	"context"
	"database/sql"
	"errors"
	"time"
	"fmt"
)

type OrderRepoImpl struct {
	sql *db.Sql
}

func NewOrderRepo(sql *db.Sql) repository.OrderRepository {
	return &OrderRepoImpl{
		sql: sql,
	}
}

// 1. Insert 1 record vào order table
// chú ý trước khi tạo 1 record trong order table thì kiểm tra user hiện tại đã có
// order hay chưa, nếu có rồi thì ko làm gì cả, chưa có thì mới tạo mới
// 2. insert 1 record vào card table
func (o *OrderRepoImpl) AddToCard(context context.Context, userId string, card model.Card) (int, error) {
	sqlCheckOrder := `select * from orders where user_id = $1 and status = $2`
	var orderRow = model.Order{}
	err := o.sql.Db.GetContext(context, &orderRow, sqlCheckOrder, userId, model.ORDERING.String())
	if err != nil && err == sql.ErrNoRows {
		// Tạo 1 order mới với staus = ORDERING
		sqlInsertOrderStatement := `
		  INSERT INTO orders(user_id, order_id, status, updated_at) 
          VALUES(:user_id, :order_id, :status, :updated_at)
     	`
		orderRow.UserId = userId
		orderRow.OrderId = encrypt.UUID()
		orderRow.UpdatedAt = time.Now()
		orderRow.Status = model.ORDERING.String()

		_, err := o.sql.Db.NamedExecContext(context, sqlInsertOrderStatement, orderRow)
		if err != nil {
			return 0, err
		}
	}

	fmt.Println("ORDER_ID = ", orderRow.OrderId)

	// Kiểm tra xem product id có tồn tại trong shopping chưa, nếu có rồi thì update quantity
	sqlCheckCard := `select * from card where product_id = $1 and order_id = $2`
	var cardRow = model.Card{}
	err = o.sql.Db.GetContext(context, &cardRow, sqlCheckCard, card.ProductId, orderRow.OrderId)
	fmt.Println("OOO = ", err)
	if err != nil && err == sql.ErrNoRows {

		fmt.Println("OOO INSERT = ", orderRow.OrderId)

		fmt.Println("LOI 1: ", err.Error())
		sqlInsertCardStatement := `
		  INSERT INTO card(order_id, product_id, product_name, product_image, quantity, price) 
          VALUES(:order_id, :product_id, :product_name, :product_image, :quantity, :price)
     	`
		card.OrderId = orderRow.OrderId
		card.Quantity = 1

		_, err = o.sql.Db.NamedExecContext(context, sqlInsertCardStatement, card)
		if err != nil {
			fmt.Println("LOI 2: ", err.Error())
			return 0, err
		}
	}

	// Nếu đã tạo order cho user này rồi thì chỉ update table card thôi
	sqlUpdateCardStatement := `
		UPDATE card
		SET quantity = :quantity
		WHERE product_id = :product_id
	`
	card.Quantity = cardRow.Quantity + 1;

	_, err = o.sql.Db.NamedExecContext(context, sqlUpdateCardStatement, card)
	if err != nil {
		fmt.Println("LOI 3: ", err.Error())
		return 0, err
	}

	var total int;
	err = o.sql.Db.QueryRowxContext(context,
		"SELECT COALESCE(SUM(quantity), 0) AS total FROM card WHERE order_id=$1", orderRow.OrderId).Scan(&total)
	if err != nil {
		fmt.Println("LOI 4: ", err.Error())
		fmt.Println(err.Error())
		return 0, err
	}

	fmt.Println("TOTAL 1 = ", total)

	return total, nil
}

func (o *OrderRepoImpl) UpdateStateOrder(context context.Context, order model.Order) error {
	// status, order_id, user_id
	sqlStatement := `
		UPDATE orders
		SET 
			status = :status,
			updated_at = :updated_at
		WHERE 
			user_id = :user_id 
			AND order_id = :order_id
	`

	order.UpdatedAt = time.Now()
	result, err := o.sql.Db.NamedExecContext(context, sqlStatement, order)
	if err != nil {
		return err
	}

	count, _ := result.RowsAffected()
	if count == 0 {
		return errors.New("Update thất bại")
	}

	return nil
}

func (o *OrderRepoImpl) UpdateQuantityOrder(context context.Context, userId string, orderId string, quantity int, productId string) error {
		// status, order_id, user_id
	sqlStatement := `
		UPDATE card
		SET 
			quantity = :quantity
		WHERE 
			order_id IN (
				SELECT order_id 
				FROM orders 
				WHERE 
					user_id = :user_id and 
					order_id = :order_id and
					status = 'ORDERING'
			) 
			AND order_id = :order_id
			AND product_id = :product_id
	`

	fmt.Println("USER_ID = ",userId )

	result, err := o.sql.Db.NamedExecContext(context, sqlStatement, 
									map[string]interface{}{
				            "quantity": quantity,
				            "user_id": userId,
				            "product_id": productId,
				            "order_id": orderId,
				          })

	if err != nil {
		return err
	}

	count, _ := result.RowsAffected()
	if count == 0 {
		return errors.New("Update thất bại")
	}

	return nil
}

func (o *OrderRepoImpl) CountShoppingCard(context context.Context, userId string) (model.OrderCount, error) {
	// Order của ai thì người đó mới được xem thông tin : orders.user_id = $2
	sqlCountStatement := `
		SELECT
		   orders.order_id,	
		   SUM(card.quantity) AS total
		FROM
		   orders
		INNER JOIN card 
		ON 
		  orders.user_id = $1 AND 
		  orders.order_id = card.order_id AND 
		  orders.status = 'ORDERING'
		GROUP BY
		   orders.order_id
	`

	row := o.sql.Db.QueryRowxContext(context, sqlCountStatement, userId)

	orderCount := model.OrderCount{}

	err := row.Err()
	if err != nil {
		fmt.Println("SQL ", err)
		return orderCount, err
	}

	err = row.StructScan(&orderCount)
	if err != nil {
		orderCount.Total = -1
		return orderCount, err
	}

	return orderCount, nil
}

func (o *OrderRepoImpl) ShoppingCard(context context.Context, userId string, orderId string) (model.Order, error) {
	sqlShoppingCard := `
		SELECT
		   orders.order_id,
		   card.product_id,
		   card.product_name,
		   card.product_image,
		   card.quantity,
		   card.price
		FROM
		   orders
		INNER JOIN card 
		ON 
      orders.user_id = $1 AND 
		  orders.order_id = $2 AND 
		  orders.order_id = card.order_id AND 
		  orders.status = 'ORDERING'
	`
	orders := model.Order{}
	cards := []model.Card{}

	err := o.sql.Db.SelectContext(context, &cards, sqlShoppingCard, userId, orderId)
	if err != nil {
		return orders, err
	}

	var sum float64 = 0
	for _, card := range cards {
		sum += card.Price * float64(card.Quantity)
	}

	orders.Total = sum
	orders.Items = cards

	return orders, nil
}

func (o *OrderRepoImpl) ListOrder(context context.Context) ([]model.Order, error) {
	sqlOrders := `
		SELECT
		   orders.user_id,
			 orders.order_id,
			 orders.updated_at,
			 orders.status,
			 SUM(card.total) as total
		FROM
		   orders
		INNER JOIN (
			SELECT 
				card.order_id,
				(card.price * SUM(card.quantity)) as total
			FROM card
			GROUP BY 
				card.order_id, card.price
		) card
		ON 
		  orders.order_id = card.order_id 
		GROUP BY 
		  orders.user_id,
			orders.order_id,
			orders.updated_at, 
			orders.status
	`
	orders := []model.Order{}

	err := o.sql.Db.SelectContext(context, &orders, sqlOrders)
	if err != nil {
		return orders, err
	}

	return orders, nil
}



