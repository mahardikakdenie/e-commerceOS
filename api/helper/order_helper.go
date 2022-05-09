package helper

import (
	"api/entity"
)

// func CheckerTshirtId(tshirt_id *uint, data entity.Order) *uint {
// 	if tshirt_id != nil {
// 		data.TShirtID = tshirt_id
// 	} else {
// 		tshirt_id = data.TShirtID
// 		data.TShirtID = tshirt_id
// 	}

// 	return tshirt_id
// }

func CheckStatus(status string, data entity.Order) string {
	if status != "" {
		data.Status = status
	} else {
		status = data.Status
		data.Status = status
	}

	return status
}

func CheckQuantity(quantity int, data entity.Order) int {
	if quantity != 0 {
		data.Quantity = quantity
	} else {
		quantity = data.Quantity
		data.Quantity = quantity
	}

	return quantity
}

func CheckAddress(address string, data entity.Order) string {
	if address != "" {
		data.Address = address
	} else {
		address = data.Address
		data.Address = address
	}

	return address
}

func CheckCustomerForeignKey(customer_id *uint, data entity.Order) *uint {
	if customer_id != nil {
		data.CustomerID = customer_id
		customer_id = data.CustomerID
	} else {
		customer_id = nil
	}

	return customer_id
}
