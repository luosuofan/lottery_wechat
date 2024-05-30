package service

type ViewPrize struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	Pic            string `json:"pic"`
	Link           string `gorm:"column:link" json:"link" form:"link"`
	Type           int64  `gorm:"column:type" json:"type" form:"type"`
	Data           string `gorm:"column:data" json:"data" form:"data"`
	Total          int64  `gorm:"column:total" json:"total" form:"total"`
	Left           int64  `gorm:"column:left" json:"left" form:"left"`
	IsUse          int64  `gorm:"column:is_use" json:"is_use" form:"is_use"`
	Probability    int64  `gorm:"column:probability" json:"probability" form:"probability"`
	ProbabilityMax int64  `gorm:"column:probability_max" json:"probability_max" form:"probability_max"`
	ProbabilityMin int64  `gorm:"column:probability_min" json:"probability_min" form:"probability_min"`
}
type InitPrizeReq struct {
	ViewPrizeList []*ViewPrize `json:"view_prize_list"`
}
