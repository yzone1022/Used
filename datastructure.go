package main

type Item struct{
	Itemid			int				`json:"itemid"`
	Itemname 		string 			`json:"itemname"`
	Itemimage		string			`json:"itemimage"`
	Price 			float32			`json:"price"`
	Sellername  	string 			`json:"sellername"`
	Description 	string 			`json:"description"`
}
type Order struct {
	OrderId			   int 				`json:"orderId"`
	Buyername 		   string 			`json:"buyername"`
	Sellername 		   string 			`json:"sellername"`
	Itemname 		   string 			`json:"itemname"`
	Itemimage          string			`json:"itemimage"`
	Totalprice 		   float32 			`json:"totalprice"`
	ShipTo             string       	`json:"shipTo,omitempty"`
	Address            string      		`json:"address,omitempty"`
	City               string       	`json:"city,omitempty"`
	Country            string       	`json:"country,omitempty`
	State              string       	`json:"state,omitempty"`
	PostalCode         string       	`json:"postalCode,omitempty"`
	ShipmentNumber     string       	`json:"shipmentNumber,omitempty"`
	CreateTime         string       	`json:"createTime, omitempty"`
	ConfirmTime        string       	`json:"confirmTime, omitempty"`//卖家Accept  or  Reject
	CancelTime         string       	`json:"cancelTime, omitempty"`//买家Cancel
	DeliveryTime       string       	`json:"deliveryTime, omitempty"`
	CompleteTime       string       	`json:"completeTime, omitempty"`//订单完成、成功退货或者成功收货
	Orderstatus        int 				`json:"orderstatus"`//
}

type User struct{
	Userid             int          `json:"userid"`
	Username 		   string 		`json:"username"`
	Password     	   string		`json:"password"`
	Nickame 		   string 		`json:"nickname"`
	Balance 		   float32 			`json:"balance"`
	Introduction       string 		`json:"introduction"`
	Userimage		   string 		`json:"userimage"`
}