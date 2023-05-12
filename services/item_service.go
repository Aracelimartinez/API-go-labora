package services

import (
	"labora-api/models"
	"errors"
)

var Items []models.Item

// Obtiene todos os items
func GetItems() ([]models.Item, error) {
	Items = nil

	rows, err := Db.Query("SELECT * FROM items")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item

		err := rows.Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price)

		if err != nil {
			return nil, err
		}

		Items = append(Items, item)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return Items, nil
}

// Busca un item por Id
func GetItem (id string) (*models.Item, error) {
	var item models.Item

	stmt, err := Db.Prepare("SELECT * FROM items WHERE id = $1")
	if err != nil {
	    return &models.Item{}, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price)
    if err != nil {
      return &models.Item{}, err
    }

	return &item, nil
}

// Crea y valida un nuevo item
func CreateItem(newItem models.Item) (int, error) {

	var err error
	if newItem.Product == "" || newItem.CustomerName == "" || newItem.OrderDate == "" || newItem.Quantity == 0 || newItem.Price == 0 {
		err = errors.New("Todos los campos son obligatorios")
		return 0, err
	}

	stmt, err := Db.Prepare("INSERT INTO public.items(customer_name, order_date, product, quantity, price) VALUES ($1, $2, $3, $4, $5) RETURNING id")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	var newItemID int
	err = stmt.QueryRow(newItem.CustomerName, newItem.OrderDate, newItem.Product, newItem.Quantity, newItem.Price).Scan(&newItemID)
	if err != nil {
		return 0, err
	}

	return newItemID, nil
}

// Valida y actualiza un item existente
func UpdateItem(updatedItem models.Item) (int64, error) {
	var err error
	if updatedItem.Product == "" || updatedItem.CustomerName == "" || updatedItem.OrderDate == "" || updatedItem.Quantity == 0 || updatedItem.Price == 0 {
		err = errors.New("Todos los campos son obligatorios")

		return 0, err
	}

	stmt, err := Db.Prepare("UPDATE items	SET customer_name = $1,	order_date = $2, product = $3, quantity = $4, price = $5 WHERE id = $6")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(updatedItem.CustomerName, updatedItem.OrderDate, updatedItem.Product, updatedItem.Quantity, updatedItem.Price, updatedItem.ID )
	if err != nil {
		return  0 , err
	}

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		return  0 , err
	}

	return rowsUpdated, err
}

//Elimina un item
func DeleteItem(id string) (int64, error) {
	var err error

	stmt, err := Db.Prepare("DELETE FROM items WHERE id = $1")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
			return 0, err
	}

	rowsAffected, err := result.RowsAffected()
    if err != nil {
        return 0, err
    }

	return rowsAffected, nil
}

func SearchItem(product string) ([]models.Item, error) {
	Items = nil

	stmt, err := Db.Prepare("SELECT * FROM items WHERE product = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(product)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Item

		err := rows.Scan(&item.ID, &item.CustomerName, &item.OrderDate, &item.Product, &item.Quantity, &item.Price)

		if err != nil {
			return nil, err
		}

		Items = append(Items, item)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return Items, nil
}
