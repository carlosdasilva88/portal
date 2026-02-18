package validators

import (
	"errors"
	"portal/internal/models"
)

func ProductValidate(p models.Product) error {
	if p.Price < 0 {
		return errors.New("O valor do produto não pode ser negativo")
	}
	return nil
}
