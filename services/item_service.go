package services

import (
	"labora-api/models"
	"errors"
)

var Items []models.Item

func GetItems() ([]models.Item, error) {
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
func CreateItem(newItem models.Item) (int64, error) {

	var err error
	if newItem.Product == "" || newItem.CustomerName == "" || newItem.OrderDate == 0 || newItem.Quantity == "" || newItem.Price == "" || newItem.Details == "" {
		err = errors.New("Todos los campos son obligatorios")

		return 0, err
	}

	stmt, err := Db.Prepare("INSERT INTO public.items(customer_name, order_date, product, quantity, price) VALUES ($1, $2, $3, $4, $5, $6)",)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(newItem.CustomerName, newItem.OrderDate, newItem.Product, newItem.Quantity, newItem.Price)
	if err != nil {
		return 0, err
	}

	newItemID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return newItemID, nil
}


// func CreateItem(newItem Item) (Item, error) {
// 	// Se realiza la consulta SQL necesaria para crear el nuevo item en la base de datos.
// 	query := "INSERT INTO items(name, description, price) VALUES($1, $2, $3, $4, $5, $6) RETURNING customer_name, order_date, product, quantity, price, details"
// 	row := db.QueryRow(query, newItem.Name, newItem.Description, newItem.Price)

// 	// Se crean las variables necesarias para almacenar los datos de la fila creada.
// 	var id int
// 	var name, description string
// 	var price float64

// 	// Se escanean los valores de la fila creada y se almacenan en las variables correspondientes.
// 	err := row.Scan(&id, &name, &description, &price)
// 	if err != nil {
// 			return Item{}, err
// 	}

// 	// Se devuelve el objeto creado como respuesta a la solicitud.
// 	return Item{id, name, description, price}, nil
// }
