package database

import (
	"database/sql"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

const listOrders = "SELECT id, price, tax, final_price FROM orders"

func (r *OrderRepository) List() ([]entity.Order, error) {
	rows, err := r.Db.Query(listOrders)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orders []entity.Order
	for rows.Next() {
		var i entity.Order
		if err := rows.Scan(&i.ID, &i.Price, &i.Tax, &i.FinalPrice); err != nil {
			return nil, err
		}
		orders = append(orders, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return orders, nil
}
