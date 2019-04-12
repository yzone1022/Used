package main

import (
	"fmt"
	"net/http"
    "io"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"encoding/json"
	"path"

	"os"
)
const (
	Image_DIR = "./images"
)
type DbWorker struct {
	//mysql data source name
	Dsn string
}
var (
	dbhostsip  = "127.0.0.1"
	dbusername = "root"
	dbpassowrd = "yz199510220"
	dbname     = "user"
)

type LoginMessage struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResult struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
type RegisterResult struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
type PostItemResult struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
type EditResult struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
type OrderResult struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
type FindBackMessage struct{
	Username     string  	`json:"username"`
}
type ItemSellername struct{
	Sellername   string 	`json:"sellername"`
}
type OrderSellername struct{
	Sellername   string   	`json:"sellername"`
}
type OrderBuyername struct{
	Buyername   string   	`json:"buyername"`
}
type OrderStatusChange struct{
	OrderId     	int    		`json:"orderId"`
	Time        	string 		`json:"time"`
	ShipmentNumber 	string 		`json:"shipmentNumber,omitempty"`
}
type ItemId struct{
	Itemid    int   `json:"itemid"`
}
type OrderId struct{
	Orderid   int   `json:"orderid"`
}
func main() {
	//http.HandleFunc("/echo", echo)
	//http.HandleFunc("/", home)
	fmt.Println("welcome to USED")
	http.HandleFunc("/used/login/", LoginHandler)
	http.HandleFunc("/used/register/", RegisterHandler)
	http.HandleFunc("/used/findbackpassword/", FindBackPassword)
	http.HandleFunc("/used/images/", ImageHandler)
	http.HandleFunc("/used/getallitems/", GetAllitems)
	http.HandleFunc("/used/getmyitems/",GetMyitems)
	http.HandleFunc("/used/getitemdetail/",GetItemDetail)
	http.HandleFunc("/used/postitem/",PostItem)
	//首先需要单独上传图片才能发布商品
	http.HandleFunc("/used/uploadimage/",UploadHandle)
	http.HandleFunc("/used/getsoldorder/",GetSoldOrder)
	http.HandleFunc("/used/getpurchasedorder/",GetPurchasedOrder)
	http.HandleFunc("/used/orderdetail/",GetOrderDetail)
	http.HandleFunc("/used/createorder/",CreateOrder)
	http.HandleFunc("/used/cancelorder/",CancelOrder)
	http.HandleFunc("/used/confirmorder/",ConfirmOrder)
	http.HandleFunc("/used/deliveryorder/",DeliveryOrder)
	http.HandleFunc("/used/completeorder/",CompleteOrder)
	http.HandleFunc("/used/getuserinfo/",GetUserInformation)
	http.HandleFunc("/used/edituserinfo/",EditPersonalInfo)
	http.ListenAndServe(":9001", nil)//监听端口
}
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("0")
	db, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		panic(err)
		return
	}
	var result LoginResult
	//Login
	if r.Method == "POST" {
		res, _:= ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", res)
		var s LoginMessage
		json.Unmarshal([]byte(res), &s)
		fmt.Println(s.Username)
		fmt.Println(s.Password)
		_,err:=QueryLoginData(db, s.Username,s.Password)
		if err!=nil{
			fmt.Print(err)
			result.Code = 100
			result.Message = "fail to login"
		}else{
			result.Code = 101
			result.Message = "succeed to login"
		}
		//向客户端返回JSON数据
		bytes, _ := json.Marshal(result)
		fmt.Fprint(w, string(bytes))
	}
	defer db.Close()
}
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("1")
	db, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		panic(err)
		return
	}
	var result RegisterResult
	//register
	if r.Method == "POST" {
		res, _:= ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", res)
		//结构已知，解析到结构体
		fmt.Printf("%s\n", res)
		var s LoginMessage;
		json.Unmarshal([]byte(res), &s)
		fmt.Println(s.Username);
		fmt.Println(s.Password);
		err:= InserRegisterData(db, s.Username,s.Password)
		if err!=nil{
			fmt.Print(err)
			result.Code = 200
			result.Message = "fail to register"
		}else{
			result.Code = 201
			result.Message = "succeed to register"
		}
		//向客户端返回JSON数据
		bytes, _ := json.Marshal(result)
		fmt.Fprint(w, string(bytes))
	}
	defer db.Close()
}
func FindBackPassword(w http.ResponseWriter, r *http.Request){
	//召回密码
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("0")
	db, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		panic(err)
		return
	}
	var result LoginResult
	//findbackpassword
	if r.Method == "POST" {
		res, _:= ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", res)
		var s FindBackMessage
		json.Unmarshal([]byte(res), &s)
		fmt.Println(s.Username)
		var password string
		password,err:=QueryPassword(db,s.Username)
		if err!=nil{
			fmt.Print(err)
			result.Code = 102
			result.Message = "fail to find back password"
		}else{
			result.Code = 103
			result.Message = password
		}
		//向客户端返回JSON数据
		bytes, _ := json.Marshal(result)
		fmt.Fprint(w, string(bytes))
	}
	defer db.Close()
}
func ImageHandler(w http.ResponseWriter, r *http.Request){
	_, imagename := path.Split(r.URL.Path)
	fmt.Println(r.URL.Path)
	fmt.Println(imagename)
	imagePath := Image_DIR + "/" +imagename
	fmt.Println(imagePath)
	if exists := isExists(imagePath); !exists {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)// ?
}
func isExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return os.IsExist(err)
}
func LogOutHandler(w http.ResponseWriter, r *http.Request){
	//登出账户
}
func GetAllitems(w http.ResponseWriter, r *http.Request){
	//获取所有商品
	r.ParseForm() //解析参数，默认是不会解析的
	db, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		panic(err)
		return
	}
	if r.Method == "GET" {
		fmt.Println("l")
		result,err:=QueryItems(db)
		if err!=nil{
			fmt.Println(err)
			return
		}
		bytes, _ := json.Marshal(result)
		fmt.Println(result)
		fmt.Fprint(w, string(bytes))
	}
	defer db.Close()
}
func GetMyitems(w http.ResponseWriter, r *http.Request){
	//获取我的商品
	r.ParseForm() //解析参数，默认是不会解析的
	db, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		panic(err)
		return
	}
	if r.Method == "POST" {
		fmt.Println("l")
		res, _:= ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", res)
		var s ItemSellername
		json.Unmarshal([]byte(res), &s)
		fmt.Println(s.Sellername)
		result,err:=QueryMyItems(db,s.Sellername)
		fmt.Println(result)
		if err!=nil{
			fmt.Println(err)
			return
		}
		bytes, _ := json.Marshal(result)
		fmt.Fprint(w, string(bytes))
	}
	defer db.Close()
}
func GetItemDetail(w http.ResponseWriter, r *http.Request){
	//获取商品详情
	r.ParseForm() //解析参数，默认是不会解析的
	db, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		panic(err)
		return
	}
	if r.Method == "POST" {
		fmt.Println("l")
		res, _:= ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", res)
		var s ItemId
		json.Unmarshal([]byte(res), &s)
		fmt.Println(s.Itemid)
		result,err:=QueryItemDetail(db,s.Itemid)
		fmt.Println(result)
		if err!=nil{
			fmt.Println(err)
			return
		}
		bytes, _ := json.Marshal(result)
		fmt.Fprint(w, string(bytes))
	}
	defer db.Close()
}
func GetPurchasedOrder(w http.ResponseWriter, r *http.Request){
	//获取已购买订单列表
	r.ParseForm() //解析参数，默认是不会解析的
	db, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		panic(err)
		return
	}
	if r.Method == "POST" {
		fmt.Println("l")
		res, _:= ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", res)
		var s OrderBuyername
		json.Unmarshal([]byte(res), &s)
		fmt.Println(s.Buyername)
		result,err:=QueryPurchasedOrders(db,s.Buyername)
		fmt.Println(result)
		if err!=nil{
			fmt.Println(err)
			return
		}
		bytes, _ := json.Marshal(result)
		fmt.Fprint(w, string(bytes))
	}
	defer db.Close()
}
func GetSoldOrder(w http.ResponseWriter, r *http.Request){
	//获取已售出订单列表
	r.ParseForm() //解析参数，默认是不会解析的
	db, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		panic(err)
		return
	}
	if r.Method == "POST" {
		fmt.Println("l")
		res, _:= ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", res)
		var s OrderSellername
		json.Unmarshal([]byte(res), &s)
		fmt.Println(s.Sellername)
		result,err:=QuerySoldOrders(db,s.Sellername)
		fmt.Println(result)
		if err!=nil{
			fmt.Println(err)
			return
		}
		bytes, _ := json.Marshal(result)
		fmt.Fprint(w, string(bytes))
	}
	defer db.Close()
}
func GetOrderDetail(w http.ResponseWriter, r *http.Request){
	//获取订单详情
	r.ParseForm() //解析参数，默认是不会解析的
	db, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		panic(err)
		return
	}
	if r.Method == "POST" {
		fmt.Println("l")
		res, _:= ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", res)
		var s OrderId
		json.Unmarshal([]byte(res), &s)
		fmt.Println(s.Orderid)
		result,err:=QueryOrderDetail(db,s.Orderid)
		fmt.Println(result)
		if err!=nil{
			fmt.Println(err)
			return
		}
		bytes, _ := json.Marshal(result)
		fmt.Fprint(w, string(bytes))
	}
	defer db.Close()
}
func PostItem(w http.ResponseWriter, r *http.Request){
	//发布商品
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("1")
	db, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		panic(err)
		return
	}
	var result PostItemResult
	//register
	if r.Method == "POST" {
		res, _:= ioutil.ReadAll(r.Body)
		r.Body.Close()
		//fmt.Printf("%s\n", res)
		////结构已知，解析到结构体
		fmt.Printf("%s\n", res)
		var s Item;
		s.Itemid = 0
		json.Unmarshal([]byte(res), &s)
		fmt.Println(s);
		err:= InsertItemData(db,s)
		if err!=nil{
			fmt.Print(err)
			result.Code = 300
			result.Message = "fail to post"
		}else{
			result.Code = 301
			result.Message = "succeed to post"
		}
		//向客户端返回JSON数据
		bytes, _ := json.Marshal(result)
		fmt.Fprint(w, string(bytes))
	}
	defer db.Close()
}
func UploadHandle(w http.ResponseWriter, r *http.Request) {
	//上传图片
	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		filename := h.Filename
		fmt.Println("h.Filename:", h.Filename)
		fmt.Println("f", f)
		defer f.Close()
		t, err := os.Create(Image_DIR + "/" + filename)
		if err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
		defer t.Close()

		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(),
				http.StatusInternalServerError)
			return
		}
	}
}
func CreateOrder(w http.ResponseWriter, r *http.Request){
	//创建订单
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("1")
	db, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		panic(err)
		return
	}
	var result OrderResult
	if r.Method == "POST" {
		res, _:= ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", res)
		//结构已知，解析到结构体
		fmt.Printf("%s\n", res)
		var s Order;
		//orderId 设置为0
		s.OrderId = 0
		json.Unmarshal([]byte(res), &s)
		fmt.Println(s);
		//移动端添加判断余额够不够
		err:= InsertOrderData(db,s)
		if err!=nil{
			fmt.Print(err)
			result.Code = 400
			result.Message = "fail to createorder"
		}else{
			result.Code = 401
			result.Message = "succeed to createorder"
		}
		//向客户端返回JSON数据
		bytes, _ := json.Marshal(result)
		fmt.Fprint(w, string(bytes))
	}
	defer db.Close()
}
func CancelOrder(w http.ResponseWriter, r *http.Request){
	//取消订单
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("1")
	db, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		panic(err)
		return
	}
	var result OrderResult
	if r.Method == "POST" {
		res, _:= ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", res)
		//结构已知，解析到结构体
		fmt.Printf("%s\n", res)
		var s OrderStatusChange;
		//orderId 设置为0
		json.Unmarshal([]byte(res), &s)
		fmt.Println(s);
		//移动端添加判断余额够不够
		err:= Update_CancelOrder(db,s.OrderId,s.Time)
		if err!=nil{
			fmt.Print(err)
			result.Code = 500
			result.Message = "fail to cancelorder"
		}else{
			result.Code = 501
			result.Message = "succeed to cancelorder"
		}
		//向客户端返回JSON数据
		bytes, _ := json.Marshal(result)
		fmt.Fprint(w, string(bytes))
	}
	defer db.Close()

}
func ConfirmOrder(w http.ResponseWriter, r *http.Request){
	//确认订单
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("1")
	db, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		panic(err)
		return
	}
	var result OrderResult
	if r.Method == "POST" {
		res, _:= ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", res)
		//结构已知，解析到结构体
		fmt.Printf("%s\n", res)
		var s OrderStatusChange;
		//orderId 设置为0
		json.Unmarshal([]byte(res), &s)
		fmt.Println(s);
		//移动端添加判断余额够不够
		err:= Update_ConfirmOrder(db,s.OrderId,s.Time)
		if err!=nil{
			fmt.Print(err)
			result.Code = 600
			result.Message = "fail to confirmorder"
		}else{
			result.Code = 601
			result.Message = "succeed to confirmorder"
		}
		//向客户端返回JSON数据
		bytes, _ := json.Marshal(result)
		fmt.Fprint(w, string(bytes))
	}
	defer db.Close()
}
func DeliveryOrder(w http.ResponseWriter, r *http.Request){
	//发货
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("1")
	db, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		panic(err)
		return
	}
	var result OrderResult
	if r.Method == "POST" {
		res, _:= ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", res)
		//结构已知，解析到结构体
		fmt.Printf("%s\n", res)
		var s OrderStatusChange;
		json.Unmarshal([]byte(res), &s)
		fmt.Println(s);
		//移动端添加判断余额够不够
		err:= Update_DeliveryOrder(db,s.OrderId,s.Time,s.ShipmentNumber)
		if err!=nil{
			fmt.Print(err)
			result.Code = 700
			result.Message = "fail to deliveryorder"
		}else{
			result.Code = 701
			result.Message = "succeed to deliveryorder"
		}
		//向客户端返回JSON数据
		bytes, _ := json.Marshal(result)
		fmt.Fprint(w, string(bytes))
	}
	defer db.Close()
}
func CompleteOrder(w http.ResponseWriter, r *http.Request){
	//完成订单
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("1")
	db, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		panic(err)
		return
	}
	var result OrderResult
	if r.Method == "POST" {
		res, _:= ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", res)
		//结构已知，解析到结构体
		fmt.Printf("%s\n", res)
		var s OrderStatusChange;
		json.Unmarshal([]byte(res), &s)
		fmt.Println(s);
		//移动端添加判断余额够不够
		err:= Update_CompleteOrder(db,s.OrderId,s.Time)
		if err!=nil{
			fmt.Print(err)
			result.Code = 800
			result.Message = "fail to completeyorder"
		}else{
			result.Code = 801
			result.Message = "succeed to completeorder"
		}
		//向客户端返回JSON数据
		bytes, _ := json.Marshal(result)
		fmt.Fprint(w, string(bytes))
	}
	defer db.Close()
}
func GetUserInformation(w http.ResponseWriter, r *http.Request){
	//获取用户信息
	//包含查看余额
	r.ParseForm() //解析参数，默认是不会解析的
	db, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		panic(err)
		return
	}
	if r.Method == "POST" {
		fmt.Println("l")
		res, _:= ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", res)
		var s FindBackMessage
		json.Unmarshal([]byte(res), &s)
		fmt.Println(s.Username)
		result,err:=QueryUserDetail(db,s.Username)
		fmt.Println(result)
		if err!=nil{
			fmt.Println(err)
			return
		}
		bytes, _ := json.Marshal(result)
		fmt.Fprint(w, string(bytes))
	}
	defer db.Close()
}
func EditPersonalInfo(w http.ResponseWriter, r *http.Request){
	//编辑个人信息
	r.ParseForm() //解析参数，默认是不会解析的
	fmt.Println("1")
	db, err := sql.Open("mysql", dbusername+":"+dbpassowrd+"@tcp("+dbhostsip+")/"+dbname)
	if err != nil {
		panic(err)
		return
	}
	var result EditResult
	//register
	if r.Method == "POST" {
		res, _:= ioutil.ReadAll(r.Body)
		r.Body.Close()
		fmt.Printf("%s\n", res)
		//结构已知，解析到结构体
		fmt.Printf("%s\n", res)
		var s User;
		json.Unmarshal([]byte(res), &s)
		fmt.Println(s);
		err:= EditUserInfo(db,s)
		if err!=nil{
			fmt.Print(err)
			result.Code = 900
			result.Message = "fail to edit"
		}else{
			result.Code = 901
			result.Message = "succeed to edit"
		}
		//向客户端返回JSON数据
		bytes, _ := json.Marshal(result)
		fmt.Fprint(w, string(bytes))
	}
	defer db.Close()
}



