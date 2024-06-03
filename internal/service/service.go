package service

import (
	log "github.com/sirupsen/logrus"
	"lottery_wechat/internal/model"
	"lottery_wechat/internal/pkg/constant"
	"lottery_wechat/internal/repo"
	"math/rand"
	"time"
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

func GetPrizeList() ([]*model.Prize, error) {
	prizeList, err := repo.GetPrizeList()
	if err != nil {
		log.Errorf("service|GetPrizeList err: %v", err)
		return nil, err
	}
	return prizeList, nil
}
func GetWinner() map[string]interface{} {
	code := luckyCode()
	log.Infof("service|GetWinner code::%d\n", code)
	var ok bool
	res := make(map[string]interface{})
	res["中奖信息"] = "未中奖"
	prizeList, err := repo.GetPrizeList()
	if err != nil {
		log.Errorf("service.GetPrizeList err:%v", err)
		return nil
	}
	for _, prize := range prizeList {
		if prize.IsUse != constant.PrizeInuse || (prize.Total > 0 && prize.Left < 0) { //不在使用中，或者产品总数大于0但是剩余量小于0则continue，不够了
			continue
		}
		if prize.ProbabilityMin <= int64(code) && prize.ProbabilityMax > int64(code) {
			var profile string
			ok, profile = PrizeSenderMap[int(prize.Type)].SendPrize(prize)
			if ok {
				delete(res, "中奖信息")
				res["success"] = ok
				res["id"] = prize.ID
				res["name"] = prize.Link
				res["profile"] = profile
				break
			}
		}
	}
	return res
}

func luckyCode() int32 {
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Int31n(constant.ProbabilityLimit) //返回一个0-9999的随机数
	return code
}
