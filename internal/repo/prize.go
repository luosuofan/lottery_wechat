package repo

import (
	log "github.com/sirupsen/logrus"
	"lottery_wechat/internal/model"
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
