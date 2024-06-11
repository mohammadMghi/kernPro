package log

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath" 
)

 

func  ReadKernelLogFile(filepath string) (string, error) { 
    file, err := os.Open(filepath)
    if err != nil {
        return "", err
    }
    defer file.Close()

    // Read the file contents
    contents, err := ioutil.ReadAll(file)
    if err != nil {
        return "", err
    }

    return string(contents), nil
}

func WriteToFile(filename, content string) error {
	
	err := os.Mkdir("logFiles", 0755)
	if err != nil && !os.IsExist(err) {
		fmt.Println("Failed to create logFiles directory:", err)

		return err
	}

    path, _ := os.Getwd()
	
	fileLogPath := path + "/logFiles";

    file, err := os.Create(filepath.Join(fileLogPath, filepath.Base(filename)))
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.WriteString(content)
    if err != nil {
        return err
    }

    return nil
}