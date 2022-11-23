package ports

import "github.com/TheCozyHangout/website/backend/internal/core/domain"

type MySQLPort interface {
	GetCardByID(id string) (domain.Card, error)
	GetAllCards() ([]domain.Card, error)
}
