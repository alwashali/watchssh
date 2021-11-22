package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	cmd := flag.String("cmd", "", "Command to execute")
	username := flag.String("username", "", "Username")
	password := flag.String("password", "", "Password")
	IPAddress := flag.String("ip", "", "IP Address remote server")
	interval := flag.Int("interval", 3, "Interval in seconds")
	outfile := flag.String("outfile", "output.txt", "Result file name")
	//key := flag.String("key", "", "key")
	flag.Parse()

	if *cmd == "" {
		flag.Usage()
		log.Fatal("Need a command to monitor")

	}
	if *username == "" {
		flag.Usage()
		log.Fatal("username is needed for login")
	}
	if *password == "" {
		flag.Usage()
		log.Fatal("Password is needed for login")
	}
	if *IPAddress == "" {
		flag.Usage()
		log.Fatal("Need IPAddress")
	}

	// if *key != "" {
	// 	client, err := DialWithKey(*IPAddress+":22", *username, *key)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	defer client.Close()
	// }

	client, err := DialWithPasswd(*IPAddress+":22", *username, *password)
	if err != nil {
		fmt.Println(err)
	}
	defer client.Close()

	for {
		out, err := client.Cmd(*cmd).Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))

		f, err := os.OpenFile(*outfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}

		defer f.Close()
		TimeNow := time.Now()

		timestampedOutput := TimeNow.Format("2006-01-02 15:04:05:") + "\n" + string(out)
		if _, err = f.WriteString(timestampedOutput + "\n"); err != nil {
			panic(err)
		}
		//sleep
		time.Sleep(time.Second * time.Duration(*interval))
	}

}
