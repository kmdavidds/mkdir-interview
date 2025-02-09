package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("argument missing: {filePath}")
		return
	}
	filePath := args[1]

	names, err := readLinesTXT("input.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("missing input.txt")
		return
	}

	dirName := "output"
	err = os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, name := range names {
		formattedName := fmt.Sprintf("%d. %s", i+1, name)
		err = os.Mkdir(fmt.Sprintf("%s/%s", dirName, formattedName), 0755)
		if err != nil {
			fmt.Println(err)
		}

		err = copyFile(filePath, fmt.Sprintf("%s/%s/%s", dirName, formattedName, filePath))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func readLinesTXT(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return lines, nil
}

func copyFile(srcPath string, dstPath string) error {
	srcFile, err := os.Open(srcPath)
    if err != nil {
        return err
    }
    defer srcFile.Close()

    destFile, err := os.Create(dstPath)
    if err != nil {
        return err
    }
    defer destFile.Close()

    _, err = io.Copy(destFile, srcFile)
    if err != nil {
        return err
    }

	return nil
}
