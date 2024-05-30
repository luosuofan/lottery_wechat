package api

import (
	"lottery_wechat/internal/pkg/constant"
	"lottery_wechat/internal/service"
	"testing"
)

func TestAddPrize(t *testing.T) { //做test固定写法t *testing.T
	prizeList := make([]*service.ViewPrize, 4)
	prizeTypeCoin := service.ViewPrize{
		ID:             1,
		Name:           "Q币",
		Pic:            "http://",
		Link:           "http://q.qq.com",
		Type:           constant.PrizeTypeCoin,
		Data:           "100Q币",
		Total:          20000,
		Left:           20000,
		IsUse:          1,
		Probability:    5000,
		ProbabilityMin: 0,
		ProbabilityMax: 0,
	}
	prizeList[0] = &prizeTypeCoin

	prizeTypeSmallEntity := service.ViewPrize{
		ID:             2,
		Name:           "充电宝",
		Pic:            "",
		Link:           "",
		Type:           constant.PrizeTypeSmallEntity,
		Data:           "",
		Total:          100,
		Left:           100,
		IsUse:          1,
		Probability:    100, //百分之1中奖
		ProbabilityMin: 0,
		ProbabilityMax: 0,
	}
	prizeList[1] = &prizeTypeSmallEntity

	prizeTypeLargeEntity := service.ViewPrize{
		ID:             3,
		Name:           "iphone14",
		Pic:            "",
		Link:           "",
		Type:           constant.PrizeTypeLargeEntity,
		Data:           "",
		Total:          10,
		Left:           10,
		IsUse:          1,
		Probability:    10, //百分之1中奖
		ProbabilityMin: 0,
		ProbabilityMax: 0,
	}
	prizeList[2] = &prizeTypeLargeEntity

	prizeTypeCoupon := service.ViewPrize{
		ID:             4,
		Name:           "优惠券满100减10元",
		Pic:            "",
		Link:           "",
		Type:           constant.PrizeTypeCoupon,
		Data:           "黄焖鸡外卖优惠券",
		Total:          5000,
		Left:           5000,
		IsUse:          1,
		Probability:    3000, //百分之1中奖
		ProbabilityMin: 0,
		ProbabilityMax: 0,
	}
	prizeList[3] = &prizeTypeCoupon

	var start int64 = 0
	for _, prize := range prizeList {
		if prize.IsUse != constant.PrizeInuse {
			continue
		}
		prize.ProbabilityMin = start
		prize.ProbabilityMax = start + prize.Probability

	}
}
