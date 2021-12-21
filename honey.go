package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"time"
)

func connection(conn net.Conn, f *os.File, ip string, port string) {

	//send ftp banner
	conn.Write([]byte("220 (vsFTPd 2.3.4)\r\n"))

	//send user
	conn.Write([]byte("331 User name OK, Enter PASS command\r\n"))

	//read user
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	//delete all null characters
	user := bytes.Trim(buf, "\x00")
	//delete \r\n
	user = user[:len(user)-2]
	//print from the sixth character
	fmt.Println("[+]User:", string(user[5:]))

	//send password
	conn.Write([]byte("230 User logged in, proceed\r\n"))

	//read password
	buff := make([]byte, 1024)
	_, err = conn.Read(buff)
	if err != nil {
		fmt.Println(buff)
		return
	}

	//delete all null characters
	pass := bytes.Trim(buff, "\x00")
	//delete \r\n
	pass = pass[:len(pass)-2]
	//print from the sixth character
	fmt.Println("[+]Password:", string(pass[5:]))

	//send syst
	conn.Write([]byte("215 UNIX Type: L8\r\n"))

	//send pwd
	conn.Write([]byte("257 \"/\" is current directory\r\n"))

	//create a json struct to log the connection
	f.WriteString(fmt.Sprintf("{\"username\":\"%s\",\"password\":\"%s\",\"ip\":\"%s\",\"port\":\"%s\"hour\":\"%s\"date\":\"%s}\n", string(user[5:]), string(pass[5:]), ip, port, time.Now().Format("15:04:05"), time.Now().Format("02/01/2006")))
}

func main() {

	//create a file honeypot.log to log all ftp connections
	f, err := os.OpenFile("honeypot.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	fmt.Println("[+]Honeypot started in port 2121")

	//listener in port 2121
	ln, err := net.Listen("tcp", ":2121")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		//wait a connection
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("[-]Error accepting connection")
			return
		}

		fmt.Println("[+]Connection accepted", conn.RemoteAddr())

		ip := conn.RemoteAddr().String()[:len(conn.RemoteAddr().String())-6]

		port := conn.RemoteAddr().String()[len(conn.RemoteAddr().String())-5:]

		//create a goroutine to handle the connection
		go connection(conn, f, ip, port)

	}
}
