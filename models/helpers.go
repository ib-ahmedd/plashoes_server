package models

func getSortOrder(sort string) string {
	switch sort {
	case "popularity":
		return "sold DESC"

	case "average rating":
		return "rating DESC"

	case "latest":
		return "date_arrived DESC"

	case "price: low to high":
		return "price ASC"

	case "price: high to low":
		return "price DESC"

	case "Default sorting":
		return "id ASC"
	default:
		return ""
	}
}
