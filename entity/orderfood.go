package entity

type Orderfood struct {
	ID         int `xorm:"pk autoincr" json:"id"`
	MenufoodID int `xorm:"notnull" json:"menufood_id"`
	Order_id   int `xorm:"notnull" json:"order_id"`
	Order_num  int `xorm:"notnull default(0)" json:"order_num"`
}


type OrderfoodIns struct {
	ID 				int `json:"id"`
	MenufoodID  	int `json:"menufood_id"`
	Order_id 		int `json:"order_id"`
	Order_num 		int `json:"order_num"`
	MenufoodName 	string `json:"name"`
	Price           float64 `json:"price"`
}