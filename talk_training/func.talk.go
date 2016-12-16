package talk_training

import (
	"database/sql"
	"fmt"
)

var db *sql.DB = nil

func Init() {
	dbcon, err := sql.Open("postgres",
		"user=techacademy password=123qwe!@#QWE dbname=tokopedia-talk host=192.168.100.126 port=5432 sslmode=disable")
	if err!=nil {
		return
	}
	db = dbcon
}

func Clean() {
	if db!= nil {
		db.Close()
	}
}

func GetTalks(productId int) []Talks {
	//fmt.Println("productId=", productId)
	//query := "select talk_id, product_id, message, create_time from ws_talk where product_id = ? "
	rows, err := db.Query("select talk_id, product_id, message, create_time from ws_talk where product_id = $1 order by talk_id desc limit 100", productId)

	if err!=nil {
		fmt.Println("GetTalks err:", err)
		return nil
	}

	var listTalks = make([]Talks, 10)
	var i = 0
	for rows.Next() {
		talk := Talks{}
		err := rows.Scan(&talk.ID, &talk.ProductID, &talk.Message, &talk.CreateTime)
		if err!= nil {
			fmt.Println("Err scanning row:", err)
			return nil
		}

		listTalks = append(listTalks, talk)
		i++
	}
	return listTalks
}

func AddTalks(userId int, shopId int, productId int, message string, createBy int) bool {
	err := db.QueryRow("insert into ws_talk (user_id, shop_id, product_id, message, create_by) " +
		"values ($1, $2, $3, $4, $5)",
		userId, shopId, productId, message, createBy)

	if err!=nil {
		fmt.Println("AddTalks err:", err)
		return false
	}
	return true;
}