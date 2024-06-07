package log

import (
	"io/ioutil"
	"os"
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