package main

import (
	"net"
	"bufio"
	"fmt"
	"encoding/json"
	"strings"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Name      string
	Phone     string
	Timestamp time.Time
}

func main() {

	
	maxWait := time.Duration(5 * time.Second)
	session, err := mgo.DialWithTimeout("localhost:27017", maxWait)
	if err != nil {
			panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	
	c := session.DB("test").C("person")
        
	
	
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8001")

	// accept connection on port
	conn, _ := ln.Accept()

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))
		// sample process for string received
		newmessage := strings.TrimSpace(message)
		title := strings.Title(newmessage)
		// send new string back to client
		//conn.Write([]byte(title + "\n"))

		fmt.Print(title, " . ")
		
		
		if title == "Ale"{
			// Query All
			var results []Person
			err = c.Find(bson.M{"name": "Ale"}).Sort("-timestamp").All(&results)
			
			if err != nil {
				panic(err)
			}
			
			out, err := json.Marshal(results)
			if err != nil {
				panic (err)
			}

			//fmt.Println(string(out))
			//fmt.Println("Results All: ", results)
			
			conn.Write([]byte(string(out) + "\n"))
		} else if title == "Cla"{
			// Query All
			var results []Person
			err = c.Find(bson.M{"name": "Cla"}).Sort("-timestamp").All(&results)
			
			if err != nil {
				panic(err)
			}
			
			out, err := json.Marshal(results)
			if err != nil {
				panic (err)
			}

			//fmt.Println(string(out))
			//fmt.Println("Results All: ", results)
			
			conn.Write([]byte(string(out) + "\n"))
		} else if len(title) > 0 {
			err = c.Insert(&Person{Name: title, Phone: "+55 53 6273 2173", Timestamp: time.Now()})

			if err != nil {
				panic(err)
			}
			
			// Query All
			var results []Person
			err = c.Find(bson.M{"name": title}).Sort("-timestamp").All(&results)
			
			if err != nil {
				panic(err)
			}
			
			out, err := json.Marshal(results)
			if err != nil {
				panic (err)
			}

			//fmt.Println(string(out))
			//fmt.Println("Results All: ", results)
			
			conn.Write([]byte(string(out) + "\n"))
		} else {
			title = "Data not found"
			conn.Write([]byte(title + "\n"))
			break
		}
	}
}