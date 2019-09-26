// Go-StatusMonitor
// Copyright (c) 2019, Joshua Sing
// All rights reserved.
// License: BSD 2-Clause
package main

import(
	"database/sql"
	"fmt"
	"net"
	"time"

	"gopkg.in/gomail.v2"

	_ "github.com/go-sql-driver/mysql"
)

const (
   debug = false

   mysqladdress = "localhost" // MySQL address
   mysqlusername = "username" // MySQL username
   mysqlpassword = "password" // MySQL password
   mysqldatabase = "database" // MySQL database
   mysqltable = "statusmonitor" // MySQL table

   statusUp = 1 // The status code that will be put in the MySQL table if the server is up.
   statusDown = 2 // The status code that will be put in the MySQL table if the server is down.

   emailEnabled = true
   emailToAddress = "fred@example.com" // Who should we send the email to?
   emailFromAddress = "status@example.com" // Email from address
   emailSubject = "[Alert] Some servers may be down!" // Email subject
   emailMessage = "The following servers seem to be down: " // Email Message (First Message) HTML / TEXT
   emailMessage2 = "<br><br>From Go-StatusMonitor!" // Email Message (Last Message) HTML / TEXT
   emailServerAddress = "localhost" // Address of your mail server (You should be able to use Gmail)
   emailServerPort = 465 // Port of your mail server
   emailServerUsername = "status@example.com" // Mail server username
   emailServerPassword = "qwertyIsABADPassw0rd!" // Mail server password

   checkTime = 30 // How many seconds should we wait before checking the status of servers again?
)

var servers = map[string]string{ // Add more servers to check by adding a new line!
	"ServerName": "www.google.com:443",
	"ServerName2": "localhost:80",
 }
var serversdown = "none" // Don't change!
var emailed = false // Don't change!
const ver = "1.0.0-RELEASE" // Version number

func addServerDown(name string){ // Add a server to the serversdown list.
	if(serversdown == "none"){
		serversdown = name + ", "
		return
	}
	serversdown = serversdown + name + ", "
}

func checkServer(name, addr string) { // Checks a server
	fmt.Printf("[Info] Checking %s (%s)\n", name, addr)
	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil{
		fmt.Printf("[Error] Fail to connect to server: %v\n", err)
		updateStatus(name, statusDown)
		addServerDown(name)
		return
	}
	defer conn.Close()
	tcpConn, ok := conn.(*net.TCPConn)
	if !ok {
			fmt.Printf("[Error] conn is not a *net.TCPConn: %T\n", conn)
			return
	}
	fmt.Printf("[Info] The server (%v) is up! =)\n", tcpConn.RemoteAddr())
  updateStatus(name, statusUp)
}

func updateStatus(name string, status int) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s", mysqlusername, mysqlpassword, mysqladdress, mysqldatabase))
	if err != nil {
			fmt.Printf("[Error] Failed to connect to MySQL: %v\n", err)
			return
	}
	defer db.Close()
	if debug {
		fmt.Printf("[Info] Successfully connected to MySQL\n")
	}
	result, err := db.Exec("UPDATE "+mysqltable" SET Status = ? WHERE Server_Name = ?", status, name)
	if err != nil {
		fmt.Println(err)
		return
	}
	rows, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	if rows != 1 {
		fmt.Printf("[Info] Did not effect 1 row! (The value in the database may already be set to %v)\n", status)
		return
	}

}

func send(body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", emailFromAddress)
	m.SetHeader("To", emailToAddress)
	m.SetHeader("Subject", emailSubject)
	m.SetBody("text/html", emailMessage+" "+body+" "+emailMessage2)
	d := gomail.NewDialer(emailServerAddress, emailServerPort, emailServerUsername, emailServerPassword)
	if err := d.DialAndSend(m); err != nil {
	    panic(err)
	}else{
		fmt.Println("[Info] Sent Email!")
	}
}

func main(){
	fmt.Println("Go-StatusMonitor Started!")
	ticker := time.NewTicker(checkTime * time.Second)
	ticker2 := time.NewTicker(checkTime * 2 * time.Second)
	for{
		for name, addr := range servers {
			checkServer(name, addr)
		}
		if(serversdown == "none"){
			fmt.Println("[Info] All servers are up! Not sending email!")
		}else{
			if(!emailed){
				send(serversdown)
				emailed = true
			}
		}
		serversdown = "none"
		<-ticker.C
	}
	for{
		if(emailed == true){
			emailed = false
		}
		<-ticker2.C
	}
}
