package models

type ServiceFee struct {
	ID             uint    `json:"id" gorm:"primary_key"`
	LateServiceFee float64 `json:"late_service_fee"`
	InitiationFee  float64 `json:"initiation_fee"`
	LostFee        float64 `json:"lost_fee"`
	CreatedAt      string  `gorm:"autoCreateTime"`
	UpdatedAt      string  `gorm:"autoUpdateTime"`
}
