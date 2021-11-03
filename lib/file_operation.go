package lib

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func FileReadLine(filepath string, ch chan string, quit chan bool) {
	conf := ConfigReader("../config/proc_stream.yaml")
	tick := time.Tick(time.Duration(conf.ProcConf.Process.Cycle_msec) * time.Millisecond)
	for {
		select {
		case <-tick:
			stat, err := os.Stat(filepath)
			if err != nil {
				log.Fatal(err)
			}
			if stat.Size() == 0 {
				continue
			}

			fd, err := os.Open(filepath)
			if err != nil {
				log.Fatal(err)
				continue
			}
			scanner := bufio.NewScanner(fd)

			for scanner.Scan() {
				ch <- scanner.Text()
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintln(os.Stderr, "reading standard input:", err)
			}

			data := []byte("")
			err = ioutil.WriteFile(filepath, data, 0644)
			if err != nil {
				log.Fatal(err)
			}
		case <-quit:
			return
		}
	}
}

func FileReader(s_filepath string) []byte {
	bytes, err := ioutil.ReadFile(s_filepath)
	if err != nil {
		fmt.Println("Error Occured: ", err)
		panic(err)
	}

	return bytes
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func CreateFile(filepath string) {
	// check if file exists
	var _, err = os.Stat(filepath)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(filepath)
		if isError(err) {
			return
		}
		defer file.Close()
	}
}
func DeleteFile(filepath string) {
	// delete file
	var err = os.Remove(filepath)
	if isError(err) {
		return
	}
}

func EraseFile(filepath string) {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(filepath, os.O_RDWR, 0644)
	if isError(err) {
		return
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString("")
	if isError(err) {
		return
	}
	err = file.Sync()
	if isError(err) {
		return
	}
}
