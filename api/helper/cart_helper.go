package helper

import (
	"api/entity"
)

// func CheckerCartTshirtId(tshirt_id uint, data entity.Cart) uint {
// 	if tshirt_id != 0 {
// 		data.TShirtID = tshirt_id
// 	} else {
// 		tshirt_id = data.TShirtID
// 		data.TShirtID = tshirt_id
// 	}

// 	return tshirt_id
// }

// func CheckerCartUserId(user_id uint, data entity.Cart) uint {
// 	if user_id != 0 {
// 		data.UserID = user_id
// 	} else {
// 		user_id = data.UserID
// 		data.UserID = user_id
// 	}

// 	return user_id
// }

func CheckCartStatus(status string, data entity.Cart) string {
	if status != "" {
		data.Status = status
	} else {
		status = data.Status
		data.Status = status
	}

	return status
}

func CheckCartQuantity(quantity int, data entity.Cart) int {
	if quantity != 0 {
		data.Quantity = quantity
	} else {
		quantity = data.Quantity
		data.Quantity = quantity
	}

	return quantity
}
