package main

import (

	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//buatDB()
	buatTabel()
}

func koneksiMysql() (db *sql.DB){
	dbDriver := "mysql"
    dbUser := "root"
    dbPass := ""
    db, err := sql.Open(dbDriver, dbUser + ":" + dbPass + "@/")
    if err != nil {
        panic(err.Error())
    }
	return db
}

func buatDB(){

	db := koneksiMysql()

	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS bwastarups")
	
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("berhasil terhubung dengan database")
	}
}

func koneksiDB() (db *sql.DB) {

    dbDriver := "mysql"
    dbUser := "root"
    dbPass := ""
    dbName := "bwastartups"
    db, err := sql.Open(dbDriver, dbUser+":"+  dbPass + "@/" + dbName)
    if err != nil {
        panic(err.Error())
    }
	return db
}

func buatTabel(){

	db := koneksiDB()

	_, err := db.Exec("USE bwastartups")
	
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("berhasil memilih database")
	}

	// stmt, err := db.Prepare(`CREATE TABLE users (
	// 	id int NOT NULL AUTO_INCREMENT, 
	// 	name varchar(50), 
	// 	ocuption varchar(50), 
	// 	email varchar(50),
	// 	passwordhash varchar(255),
	// 	avatar varchar(255),
	// 	created_at datetime,
	// 	updated_at datetime,
	// 	PRIMARY KEY (id))`)

	// stmt, err := db.Prepare(`CREATE TABLE campaign (
	// 	id int NOT NULL AUTO_INCREMENT, 
	// 	userid int,
	// 	name varchar(50), 
	// 	short_description varchar(100), 
	// 	description varchar(255),
	// 	goal_ammount int,
	// 	current_ammount int,
	// 	perks text,
	// 	slug varchar(100),
	// 	created_at datetime,
	// 	updated_at datetime,
	// 	FOREIGN KEY (userid) REFERENCES users(id),
	// 	PRIMARY KEY (id))`)

	// stmt, err := db.Prepare(`CREATE TABLE gambarcampaign (
	// 	id int NOT NULL AUTO_INCREMENT, 
	// 	campaignid int,
	// 	filename varchar(100),
	// 	isprimary tinyint,
	// 	created_at datetime,
	// 	updated_at datetime,
	// 	FOREIGN KEY (campaignid) REFERENCES campaign(id),
	// 	PRIMARY KEY (id))`)

	stmt, err := db.Prepare(`CREATE TABLE transaction (
		id int NOT NULL AUTO_INCREMENT, 
		campaignid int,
		userid int,
		amount int,
		status varchar(15),
		code varchar(20),
		created_at datetime,
		updated_at datetime,
		FOREIGN KEY (campaignid) REFERENCES campaign(id),
		FOREIGN KEY (userid) REFERENCES users(id),
		PRIMARY KEY (id))`)
	

	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("berhasil membuat tabel lur..")
	}
	defer db.Close()
}