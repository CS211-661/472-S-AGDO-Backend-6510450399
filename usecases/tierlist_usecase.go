package usecases

import (
	"github.com/ppwlsw/sa-project-backend/domain/entities"
	"github.com/ppwlsw/sa-project-backend/usecases/repositories"
)

type TierListUsecase interface {
	GetDiscountPercentByUserID(id int) (*entities.TierList, error)
	InitialTierList() error
}

type TierListService struct {
	repo repositories.TierListRepository
}

func InitiateTierListService(repo repositories.TierListRepository) *TierListService {
	return &TierListService{repo: repo}
}

func (tls *TierListService) GetDiscountPercentByUserID(id int) (*entities.TierList, error) {

	discount, err := tls.repo.GetDiscountPercentByUserID(id)
	if err != nil {
		return &entities.TierList{}, err
	}

	return discount, nil

}
func (tls *TierListService) InitialTierList() error {

	for i := 1; i <= 3; i++ {
		tier := i
		discount := i * 10

		var tierList = entities.TierList{
			Tier:            tier,
			DiscountPercent: float64(discount),
		}
		_, err := tls.repo.InitialTierList(tierList.Tier, tierList.DiscountPercent)

		if err != nil {
			return err
		}

	}

	return nil

}
