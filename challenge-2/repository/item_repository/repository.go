package item_repository

import (
	"challenge-2/entity"
	"challenge-2/pkg/errs"
)

type Repository interface {
	GetItemsByCodes(itemCodes []string) ([]entity.Item, errs.Error)
}
