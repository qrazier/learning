package main

import (
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {

	//var s string
	// connect to this socket
	conn, _ := net.Dial("tcp", "127.0.0.1:8001")
	for { 
		// read in input from Stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Person Name: ")
		text, _ := reader.ReadString('\n')
		// send to socket
		fmt.Fprintf(conn, text + "\n")
		// listen for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		
		if strings.TrimSpace(message) == "Data not found"{
			break
		} else {
			//fmt.Print("Result: " + message)
			
			result := strings.Split(message, ",")
			for i := range result{
				fmt.Println(result[i])
			}
			
			ex := bufio.NewReader(os.Stdin)
			fmt.Print("Do you want to insert a new Person? [Y/N]")
			
			s, _ := ex.ReadString('\n')
			if strings.TrimRight(s, "\n") == "Y" || strings.TrimRight(s, "\n") == "y" {
				
				// read in input from Stdin
				reader = bufio.NewReader(os.Stdin)
				fmt.Print("Enter Person Name: ")
				new, _ := reader.ReadString('\n')
				
				// send to socket
				fmt.Fprintf(conn, new + "\n")
				
				// listen for reply
				mesej, _ := bufio.NewReader(conn).ReadString('\n')
				//fmt.Print("New person: " + mesej)
				result = strings.Split(mesej, ",")
				for i := range result{
					fmt.Println(result[i])
				}
			}
		}
	}
}