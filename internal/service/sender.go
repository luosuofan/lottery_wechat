package service //定义一个接口

import (
	log "github.com/sirupsen/logrus"
	"lottery_wechat/internal/model"
	"lottery_wechat/internal/pkg/constant"
	"lottery_wechat/internal/repo"
)

type PrizeSender interface {
	SendPrize(prize *model.Prize) (bool, string)
}
type CoinSender struct{}

func (c *CoinSender) SendPrize(prize *model.Prize) (bool, string) {
	//奖品无限量
	if prize.Total == 0 {
		return true, prize.Data
	}
	//奖品数量不足
	if prize.Left < 0 {
		return false, "奖品已发完"
	}
	prize.Left--
	if err := repo.SavePriize(prize); err != nil {
		log.Errorf("service|CoinSender|SavePrize err:%v", err)
		return false, err.Error()
	}
	return true, prize.Data
}

type CouponSender struct{}

func (c *CouponSender) SendPrize(prize *model.Prize) (bool, string) {
	//奖品无限量
	if prize.Total == 0 {
		return true, prize.Data
	}
	//奖品数量不足
	if prize.Left < 0 {
		return false, "奖品已发完"
	}
	prize.Left--
	if err := repo.SavePriize(prize); err != nil {
		log.Errorf("service|CouponSender|SavePrize err:%v", err)
		return false, err.Error()
	}
	return true, prize.Data
}

type SmallEntitySender struct{}

func (c *SmallEntitySender) SendPrize(prize *model.Prize) (bool, string) {
	//奖品无限量
	if prize.Total == 0 {
		return true, prize.Data
	}
	//奖品数量不足
	if prize.Left < 0 {
		return false, "奖品已发完"
	}
	prize.Left--
	if err := repo.SavePriize(prize); err != nil {
		log.Errorf("service|SmallEntitySender|SavePrize err:%v", err)
		return false, err.Error()
	}
	return true, prize.Data
}

type LargeEntitySender struct{}

func (c *LargeEntitySender) SendPrize(prize *model.Prize) (bool, string) {
	//奖品无限量
	if prize.Total == 0 {
		return true, prize.Data
	}
	//奖品数量不足
	if prize.Left < 0 {
		return false, "奖品已发完"
	}
	prize.Left--
	if err := repo.SavePriize(prize); err != nil {
		log.Errorf("service|LargeEntitySender|SavePrize err:%v", err)
		return false, err.Error()
	}
	return true, prize.Data
}

var PrizeSenderMap = map[int]PrizeSender{
	constant.PrizeTypeCoin:        &CoinSender{},
	constant.PrizeTypeCoupon:      &CouponSender{},
	constant.PrizeTypeSmallEntity: &SmallEntitySender{},
	constant.PrizeTypeLargeEntity: &LargeEntitySender{},
}
