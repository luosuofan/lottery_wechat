package repo

import (
	log "github.com/sirupsen/logrus"
	"lottery_wechat/internal/model"
	"lottery_wechat/internal/pkg/constant"
	"lottery_wechat/internal/pkg/gormcli"
)

func AddPrize(prizeList []*model.Prize) error {
	db := gormcli.GetDb()
	if err := db.Model(&model.Prize{}).Create(prizeList).Error; err != nil {
		log.Errorf("repo|add prize err:%v", err)
		return err
	}
	log.Infof("repo|add prize success")
	return nil
}

func GetPrizeList() ([]*model.Prize, error) {
	db := gormcli.GetDb()
	var prizeList []*model.Prize
	err := db.Model(&model.Prize{}).Where("is_use = ?", constant.PrizeInuse).Find(&prizeList).Error //查询正在使用中的奖品类型保存到prizeList
	if err != nil {
		log.Errorf("repo|GetPrzieList err:%v", err)
		return nil, err
	}
	return prizeList, nil
}

func SavePriize(prize *model.Prize) error {
	db := gormcli.GetDb()
	if err := db.Model(&model.Prize{}).Where("id=?", prize.ID).Save(prize).Error; err != nil {
		log.Errorf("repo|SavePrize err:%v", err)
		return err
	}
	return nil
}
