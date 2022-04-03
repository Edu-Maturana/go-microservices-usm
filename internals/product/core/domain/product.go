package domain

type Product struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Stock       uint8  `json:"stock" validate:"required,gte=0"`
	Price       uint16 `json:"price" validate:"required,gte=0"`
}
