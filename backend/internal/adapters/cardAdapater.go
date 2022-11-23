package adapters

// Implements methods from the CardPort interface to the MySQlAdapter struct

import (
	"fmt"

	"github.com/TheCozyHangout/website/backend/internal/core/domain"
)

func (adapter *MySQLAdapter) GetCardByID(id string) (domain.Card, error) {
	str := fmt.Sprintf("SELECT * FROM `cards` WHERE `identifier` = '%s';", id)

	result, err := adapter.client.Queryx(str)

	if err != nil {
		return domain.Card{}, err
	}

	card := domain.Card{}
	for result.Next() {
		err := result.StructScan(&card)

		if err != nil {
			return domain.Card{}, err
		}
	}

	return card, nil
}

func (adapter *MySQLAdapter) GetAllCards() ([]domain.Card, error) {
	result, err := adapter.client.Queryx("SELECT * FROM `cards`")
	var cards []domain.Card

	if err != nil {
		return cards, err
	}

	card := domain.Card{}

	for result.Next() {
		err := result.StructScan(&card)

		if err != nil {
			return cards, err
		}

		cards = append(cards, card)
	}

	return cards, nil
}
