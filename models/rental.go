package models

type Rental struct {
	ID            uint    `json:"id" gorm:"primary_key"`
	BookID        int     `json:"book_id"`
	UserID        int     `json:"user_id"`
	IsReturned    bool    `json:"is_returned"`
	PayableAmount float64 `json:"payable_amount"`
	LateReturnFee float64 `json:"late_return_fee"`
	ReturnDate    string  `json:"return_date"`
	Book          Book    `gorm:"foreinKey:BookID"`
	User          User    `gorm:"foreignKey:UserID"`
	CreatedAt     string  `gorm:"autoCreateTime"`
	UpdatedAt     string  `gorm:"autoUpdateTime"`
}
