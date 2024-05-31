package service

import (
	log "github.com/sirupsen/logrus"
	"lottery_wechat/internal/model"
	"lottery_wechat/internal/repo"
)

func AddPrize(viewPrizeList []*ViewPrize) error {
	prizeList := make([]*model.Prize, 0)
	for _, viewPrize := range viewPrizeList {
		prize := &model.Prize{
			ID:             viewPrize.ID,
			Name:           viewPrize.Name,
			Pic:            viewPrize.Pic,
			Left:           viewPrize.Left,
			Link:           viewPrize.Link,
			Type:           viewPrize.Type,
			Data:           viewPrize.Data,
			Total:          viewPrize.Total,
			IsUse:          viewPrize.IsUse,
			Probability:    viewPrize.Probability,
			ProbabilityMin: viewPrize.ProbabilityMin,
			ProbabilityMax: viewPrize.ProbabilityMax,
		}
		prizeList = append(prizeList, prize)
	}
	if err := repo.AddPrize(prizeList); err != nil {
		log.Errorf("service|AddPrize err:R%v", err)
		return err
	}
	return nil
}
