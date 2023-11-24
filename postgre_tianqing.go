package main

import (
	"database/sql"
	"flag"
	"fmt"

	// "strings"
	// "io/ioutil"
	_ "github.com/lib/pq"
	// "log"
	// "os"
)

var (
	//version      bool
	host          string
	port          int
	user          string
	password      string
	dbname        string
	sqlquser      bool
	sqlqcomputer  bool
	sqladduser    bool
	sqlupdateuser bool
	adduser       string
	addpass       string
	userid        string
)

func connectDB() *sql.DB {
	fmt.Println("[+] Get the config")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Println("[*] psqlInfo:" + psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("[+] Successfully connected!")
	return db

}

func quser(db *sql.DB) {
	var id, name, passwd, create_time, login_time, login_ip, last_login_time, last_login_ip, pass_changetime string
	fmt.Println("[*] Querying skylar UserLists")
	rows, err := db.Query(`SELECT id,name,passwd,create_time,login_time,login_ip,last_login_time,last_login_ip,pass_changetime FROM public.user;`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &passwd, &create_time, &login_time, &login_ip, &last_login_time, &last_login_ip, &pass_changetime)
		if err != nil {
			//  fmt.Println(err)
		}
		fmt.Println(id, "\t", name, "\t", passwd, "\t", create_time, "\t", login_time, "\t", login_ip, "\t", last_login_time, "\t", last_login_ip, "\t", pass_changetime)

	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("\r\n" + "Query User Lists Successfully!")
}

func qcomputerlist(db *sql.DB) {
	var id, mid, ip, report_ip, name, domain, mac, nickname, work_group, sd_engine, ie_ver, os, os_bit, login_user, display_name string
	fmt.Println("[*] Querying skylar Computers Lists")
	rows, err := db.Query(`SELECT id,mid,ip,report_ip,name,domain,mac,nickname,work_group,sd_engine,ie_ver,os,os_bit,login_user,display_name FROM public.client;`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &mid, &ip, &report_ip, &name, &domain, &mac, &nickname, &work_group, &sd_engine, &ie_ver, &os, &os_bit, &login_user, &display_name)
		if err != nil {
			//  fmt.Println(err)
		}
		fmt.Println(id, "\t", mid, "\t", ip, "\t", report_ip, "\t", name, "\t", domain, "\t", mac, "\t", nickname, "\t", work_group, "\t", sd_engine, "\t", ie_ver, "\t", os, "\t", os_bit, "\t", login_user, "\t", display_name)

	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("\r\n" + "Query Computer Lists Successfully!")
}

func addquser(db *sql.DB) {
	fmt.Println("[*] Add Manage Account")
	rows, err := db.Query(`INSERT INTO "public"."user" ("name", "passwd", "super_user") VALUES ('` + adduser + `', '` + addpass + `', '1');`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan()
		if err != nil {
			//fmt.Println(err)
		}
	}
	fmt.Println("===================================================")
	fmt.Println("\r\n" + "默认账号: " + adduser)
	fmt.Println("默认密码: " + addpass)
	fmt.Println("默认密码: Admin12345")
	fmt.Println("Add User Successfully!")
	fmt.Println("===================================================")
}

func updateuser(db *sql.DB) {
	fmt.Println("[*] Change the password of the account")
	rows, err := db.Query(`UPDATE "public"."user" SET "passwd" = '` + addpass + `' WHERE "id" = '` + userid + `';`)

	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan()
		if err != nil {
			//fmt.Println(err)
		}
	}
	fmt.Println("===================================================")
	fmt.Println("\r\n" + "更改用户密码")
	fmt.Println("pass hash: " + addpass)
	fmt.Println("默认密码: Admin12345")
	fmt.Println("Change the password of the account Successfully!")
	fmt.Println("===================================================")
}

func main() {
	flag.StringVar(&host, "h", "127.0.0.1", "database address")
	flag.IntVar(&port, "P", 5360, "database port")
	flag.StringVar(&user, "u", "postgres", "database username")
	flag.StringVar(&password, "p", "postgres", "database password")
	flag.StringVar(&dbname, "d", "skylar", "database dbname")
	flag.BoolVar(&sqlquser, "qusers", false, "query qax skylar UserLists")
	flag.BoolVar(&sqlqcomputer, "qcomputers", false, "query qax skylar Computer Lists")
	flag.BoolVar(&sqladduser, "add", false, "Add manage account")
	flag.BoolVar(&sqlupdateuser, "update", false, "Change the password of the account")
	flag.StringVar(&adduser, "user", "test", "Add user username")
	flag.StringVar(&addpass, "pass", "a71a36a92c71ba476faa632b812bf636", "Add user password (默认密码Admin12345)")
	flag.StringVar(&userid, "id", "5", " the user id")

	flag.Parse()
	flag.PrintDefaults()
	db := connectDB()

	if sqlquser {
		quser(db)
	}
	if sqlqcomputer {
		qcomputerlist(db)
	}
	if sqladduser {
		addquser(db)
	}
	if sqlupdateuser {
		updateuser(db)
	}

}
