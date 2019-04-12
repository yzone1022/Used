package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)
var OrderStatus = map[string]int32{
	"cancel order":  		0,
	"place order":     	1,
	"confrim order":   	2,
	"delivery order":		3,
	"complete order":		4,
}

//显示数据库
/*DB中执行SQL通过Exec和Query方法，查询操作是通过Query完成，它会返回一个sql.Rows的结果集，包含一个游标用来遍历查询结果；
Exec方法返回的是sql.Result对象，用于检测操作结果，及被影响记录数*/
//register
func InserRegisterData(db *sql.DB, registername string, registerpassword string)(error){
	stmt,_:=db.Prepare("INSERT INTO `userlogin`(`username`,`password`) VALUES(?, ?)")
	defer stmt.Close()
	result,err:=stmt.Exec(registername,registerpassword)
	//result, err := db.Exec("INSERT INTO `userlogin`(`username`,`password`) VALUES(registername, registerpassword)")
	if err != nil {
		fmt.Println("insert data failed:", err.Error())
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("fetch last insert id failed:", err.Error())
		return err
	}
	fmt.Println("insert new record", id)
	return nil
}
//login
func QueryLoginData(db *sql.DB, loginusername string, loginpassword string )(int,error){
	var password string
	err := db.QueryRow("select password from userlogin where username = ?", loginusername).Scan(&password)
	if err != nil {
		fmt.Println("fetech data failed:", err.Error())
		return 0,err
	}
	if(password == loginpassword){
		return 1,nil
	}
	fmt.Println(password)
	return 0,err
}
//findback Password
func QueryPassword(db *sql.DB, loginusername string )(string,error){
	var password string
	err := db.QueryRow("select password from userlogin where username = ?", loginusername).Scan(&password)
	if err != nil {
		fmt.Println("fetech data failed:", err.Error())
		return "fetech data failed",err
	}
	return password,nil
}
//get all items
func QueryItems(db *sql.DB)([]Item,error){
	items :=new([]Item)
	rows, err := db.Query("SELECT * FROM item")
	if err != nil {
		fmt.Println("fetech data failed:", err.Error())
		return nil,err
	}
	for rows.Next() {
		var itemid int
		var price float32
		var itemname, itemimage,sellername,description string
		rows.Scan(&itemid, &itemname, &itemimage,&price,&sellername,&description)
		item:=new(Item)
		{
			item.Itemid = itemid
			item.Itemname = itemname
			item.Itemimage= itemimage
			item.Price = price
			item.Sellername = sellername
			item.Description = description
		}
		*items = append(*items, *item)
	}
	return *items,nil
}
//get my items
func QueryMyItems(db *sql.DB,sellername string)([]Item,error){
	items :=new([]Item)
	rows, err := db.Query("SELECT itemid,itemname,itemimage,price,description FROM item where sellername = ?",sellername)
	if err != nil {
		fmt.Println("fetech data failed:", err.Error())
		return nil,err
	}
	for rows.Next() {
		var itemid int
		var price float32
		var itemname, itemimage,description string
		rows.Scan(&itemid, &itemname, &itemimage,&price,&description)
		item:=new(Item)
		{
			item.Itemid = itemid
			item.Itemname = itemname
			item.Itemimage= itemimage
			item.Price = price
			item.Sellername = sellername
			item.Description = description
		}
		fmt.Println(item)
		*items = append(*items, *item)
	}
	return *items,nil
}
//get itemdetail
func QueryItemDetail(db *sql.DB,itemid int)(Item,error){
	var item Item
	err:=db.QueryRow("SELECT itemname,itemimage,price,sellername,description FROM item where itemid = ?",itemid).Scan(&item.Itemname,item.Itemimage,item.Price,item.Sellername,item.Description)
	if err!=nil{
		fmt.Println("fetech data failed:", err.Error())
		return item,err
	}
	return item,nil
}
//post Item 移动端打包订单json时 先设置itemId为0
func InsertItemData(db *sql.DB,item Item)(error){
	stmt,_:=db.Prepare("INSERT INTO `item`(`itemname`,`price`,`sellername`,`description`) VALUES(?,?,?,?)")
	defer stmt.Close()
	result,err:=stmt.Exec(item.Itemname,item.Price,item.Sellername,item.Description)
	//result, err := db.Exec("INSERT INTO `userlogin`(`username`,`password`) VALUES(registername, registerpassword)")
	if err != nil {
		fmt.Println("insert data failed:", err.Error())
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("fetch last insert id failed:", err.Error())
		return err
	}
	fmt.Println("insert new record", id)
	return nil
}
func QuerySoldOrders(db *sql.DB,sellername string)([]Order,error){
	orders :=new([]Order)
	rows, err := db.Query("SELECT orderid,buyername,itemname,itemimage,totalprice,shipTo,address,city,country,state,postalCode,shipmentNumber,createTime,confirmTime,cancelTime,deliveryTime,completeTime,orderstatus FROM `order` where sellername = ?",sellername)
	if err != nil {
		fmt.Println("fetech data failed:", err.Error())
		return nil,err
	}
	for rows.Next() {
		var orderid,orderstatus int
		var totalprice float32
		var buyername,itemname,itemimage,shipTo,address,city,country,state,postalCode,shipmentNumber,createTime,confirmTime,cancelTime,deliveryTime,completeTime string
		rows.Scan(&orderid,&buyername,&itemname,&itemimage,&totalprice,&shipTo,&address,&city,&country,&state,&postalCode,&shipmentNumber,&createTime,&confirmTime,&cancelTime,&deliveryTime,&completeTime,&orderstatus)
		order:=new(Order)
		{
			order.OrderId = orderid
			order.Buyername = buyername
			order.Sellername = sellername
			order.Itemname = itemname
			order.Itemimage = itemimage
			order.Totalprice = totalprice
			order.ShipTo = shipTo
			order.Address = address
			order.City = city
			order.Country = country
			order.State = state
			order.PostalCode = postalCode
			order.ShipmentNumber = shipmentNumber
			order.CreateTime = createTime
			order.ConfirmTime = confirmTime
			order.CancelTime = cancelTime
			order.DeliveryTime = deliveryTime
			order.CompleteTime = completeTime
			order.Orderstatus = orderstatus
		}
		*orders = append(*orders, *order)
	}
	return *orders,nil
}
func QueryPurchasedOrders(db *sql.DB,buyername string)([]Order,error){
	orders :=new([]Order)
	rows, err := db.Query("SELECT orderid,sellername,itemname,itemimage,totalprice,shipTo,address,city,country,state,postalCode,shipmentNumber,createTime,confirmTime,cancelTime,deliveryTime,completeTime,orderstatus FROM `order` where buyername = ?",buyername)
	if err != nil {
		fmt.Println("fetech data failed:", err.Error())
		return nil,err
	}
	for rows.Next() {
		var orderid,orderstatus int
		var totalprice float32
		var sellername,itemname,itemimage,shipTo,address,city,country,state,postalCode,shipmentNumber,createTime,confirmTime,cancelTime,deliveryTime,completeTime string
		rows.Scan(&orderid,&sellername,&itemname,&itemimage,&totalprice,&shipTo,&address,&city,&country,&state,&postalCode,&shipmentNumber,&createTime,&confirmTime,&cancelTime,&deliveryTime,&completeTime,&orderstatus)
		order:=new(Order)
		{
			order.OrderId = orderid
			order.Buyername = buyername
			order.Sellername = sellername
			order.Itemname = itemname
			order.Itemimage = itemimage
			order.Totalprice = totalprice
			order.ShipTo = shipTo
			order.Address = address
			order.City = city
			order.Country = country
			order.State = state
			order.PostalCode = postalCode
			order.ShipmentNumber = shipmentNumber
			order.CreateTime = createTime
			order.ConfirmTime = confirmTime
			order.CancelTime = cancelTime
			order.DeliveryTime = deliveryTime
			order.CompleteTime = completeTime
			order.Orderstatus = orderstatus
		}
		*orders = append(*orders, *order)
	}
	return *orders,nil
}
//createorder  移动端打包订单json时 先设置orderId为0
func InsertOrderData(db *sql.DB,order Order)error{
	stmt,_:=db.Prepare("INSERT INTO `order`(`buyername`,`sellername`,`itemname`,`itemimage`,`totalprice`,`shipTo`,`address`,`city`,`country`,`state`,`postalCode`,`createTime`,`orderstatus`) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)")
	defer stmt.Close()
	result,err:=stmt.Exec(order.Buyername,order.Sellername,order.Itemname,order.Itemimage,order.Totalprice,order.ShipTo,order.Address,order.City,order.Country,order.State,order.PostalCode,order.CreateTime,order.Orderstatus)
	//result, err := db.Exec("INSERT INTO `userlogin`(`username`,`password`) VALUES(registername, registerpassword)")
	if err != nil {
		fmt.Println("insert data failed:", err.Error())
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("fetch last insert id failed:", err.Error())
		return err
	}
	fmt.Println("insert new record", id)
	//要对用户余额做出改变
	var buyer User
	buyer,err= QueryUserDetail(db,order.Buyername)
	if err!=nil{
		fmt.Println("获取买家失败：",err)
		return err
	}
	buyer.Balance=buyer.Balance - order.Totalprice
	err = EditUserInfo(db,buyer)
	fmt.Print(buyer.Balance)
	if err!=nil{
		fmt.Println("扣款失败:",err)
		return err
	}
	return nil
}
//get orderdetail
func QueryOrderDetail(db *sql.DB,orderid int)(Order,error){
	var order Order
	err:=db.QueryRow("SELECT orderid,buyername,sellername,itemname,itemimage,totalprice,shipTo,address,city,country,state,postalCode,shipmentNumber,createTime,confirmTime,cancelTime,deliveryTime,completeTime,orderstatus FROM `order` where orderid = ?",orderid).Scan(&order.OrderId,&order.Buyername,&order.Sellername,&order.Itemname,&order.Itemimage,&order.Totalprice,&order.ShipTo,&order.Address,&order.City,&order.Country,&order.State,&order.PostalCode,&order.ShipmentNumber,&order.CreateTime,&order.ConfirmTime,&order.CancelTime,&order.DeliveryTime,&order.CompleteTime,&order.Orderstatus)
	if err!=nil{
		fmt.Println("fetech data failed:", err.Error())
		return order,err
	}
	return order,nil
}

//cancel order
func Update_CancelOrder(db *sql.DB, orderId int, canceltime string)error{
	result, err := db.Exec("UPDATE `order` SET `cancelTime`=? ,`orderstatus`=? WHERE `orderid`=? ", canceltime,0, orderId)
	if err != nil {
		fmt.Println("update data failed:", err.Error())
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		fmt.Println("fetch row affected failed:", err.Error())
		return err
	}
	fmt.Println("update recors number", num)
	//取消订单 钱要退回给买家
	var order Order
	order, err= QueryOrderDetail(db,orderId)
	if err!=nil{
		fmt.Println("获取订单详情失败",err)
		return err
	}
	var buyer User
	buyer,err= QueryUserDetail(db,order.Buyername)
	if err!=nil{
		fmt.Println("获取买家失败：",err)
		return err
	}
	buyer.Balance=buyer.Balance + order.Totalprice
	fmt.Print(buyer.Balance)
	err = EditUserInfo(db,buyer)
	if err!=nil{
		fmt.Println("退款失败:",err)
		return err
	}
	return nil
}
//confirm order
func Update_ConfirmOrder(db *sql.DB, orderId int, confirmtime string)error{
	result, err := db.Exec("UPDATE `order` SET `confirmTime`=? ,`orderstatus`=? WHERE `orderid`=?", confirmtime,2,orderId)
	if err != nil {
		fmt.Println("update data failed:", err.Error())
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		fmt.Println("fetch row affected failed:", err.Error())
		return err
	}
	fmt.Println("update recors number", num)
	return nil
}
//delivery order
func Update_DeliveryOrder(db *sql.DB, orderId int, deliverytime string, shipmentnumber string)error{
	result, err := db.Exec("UPDATE `order` SET `deliveryTime`=? ,`shipmentNumber`=? ,`orderstatus`=? WHERE `orderid`=? ", deliverytime,shipmentnumber,3,orderId)
	if err != nil {
		fmt.Println("update data failed:", err.Error())
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		fmt.Println("fetch row affected failed:", err.Error())
		return err
	}
	fmt.Println("update recors number", num)
	return nil
}
//complete order
func Update_CompleteOrder(db *sql.DB, orderId int, completetime string)error{
	result, err := db.Exec("UPDATE `order` SET `completeTime`=? ,`orderstatus`=?  WHERE `orderid`=?", completetime,4, orderId)
	if err != nil {
		fmt.Println("update data failed:", err.Error())
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		fmt.Println("fetch row affected failed:", err.Error())
		return err
	}
	fmt.Println("update recors number", num)
	//完成订单后 钱要打到卖家帐上
	var order Order
	order, err= QueryOrderDetail(db,orderId)
	if err!=nil{
		fmt.Println("获取订单详情失败",err)
		return err
	}
	var seller User
	seller,err= QueryUserDetail(db,order.Sellername)
	if err!=nil{
		fmt.Println("获取卖家失败：",err)
		return err
	}
	seller.Balance=seller.Balance + order.Totalprice
	err = EditUserInfo(db,seller)
	if err!=nil{
		fmt.Println("打钱失败:",err)
		return err
	}
	return nil
}
//GET USER INFOMATION
func QueryUserDetail(db *sql.DB,username string)(User,error){
	var user User
	err:=db.QueryRow("SELECT userid,username,password,nickname,balance,introduction,userimage FROM userlogin where username = ?",username).Scan(&user.Userid,&user.Username,&user.Password,&user.Nickame,&user.Balance,&user.Introduction,&user.Userimage)
	if err!=nil{
		fmt.Println("fetech data failed:", err.Error())
		return user,err
	}
	return user,nil
}
func EditUserInfo(db *sql.DB,user User)error{
	result, err := db.Exec("UPDATE `userlogin` SET `password`=? ,`nickname`=?, `balance`=? ,`introduction`=? ,`userimage`=? WHERE `username`=?",user.Password,user.Nickame,user.Balance,user.Introduction,user.Userimage,user.Username )
	if err != nil {
		fmt.Println("update data failed:", err.Error())
		return err
	}
	num, err := result.RowsAffected()
	if err != nil {
		fmt.Println("fetch row affected failed:", err.Error())
		return err
	}
	fmt.Println("update recors number", num)
	return nil
}

func ShowData(db *sql.DB){
	rows, err := db.Query("SELECT * FROM userlogin")
	//fmt.Println("rows:",*rows)
	if err != nil {
		fmt.Println("fetech data failed:", err.Error())
		return
	}
	for rows.Next() {
		var uid int
		var username, password string
		rows.Scan(&uid, &username, &password)
		fmt.Println("uid:", uid, "username:", username, "password:", password)
	}
}

func QueryData(db *sql.DB){
	var password string
	err := db.QueryRow("select password from userlogin where username = ?", "roc").Scan(&password)
	if err != nil {
		fmt.Println("fetech data failed:", err.Error())
	}
	fmt.Println(password)
}
//更新数据库字段内容
func UpdataData(db *sql.DB){
	result, err := db.Exec("UPDATE `userlogin` SET `password`=? WHERE `username`=?", "1231", "rocky")
	if err != nil {
		fmt.Println("update data failed:", err.Error())
		return
	}
	num, err := result.RowsAffected()
	if err != nil {
		fmt.Println("fetch row affected failed:", err.Error())
		return
	}
	fmt.Println("update recors number", num)
}

//删除一条数据
func DeleteData(db *sql.DB){
	result, err := db.Exec("DELETE FROM `userlogin` WHERE `username`=? ", "roc")
	if err != nil {
		fmt.Println("delete data failed:", err.Error())
		return
	}
	num, err := result.RowsAffected()
	if err != nil {
		fmt.Println("fetch row affected failed:", err.Error())
		return
	}
	fmt.Println("delete record number", num)
}



