package main

import (
	"bufio"
	"fmt"

	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/tr2vil/its_manager/lib"
)

func fileReader(s_filepath string) []byte {
	bytes, err := ioutil.ReadFile(s_filepath)
	if err != nil {
		fmt.Println("Error Occured: ", err)
		panic(err)
	}

	return bytes
}

func fileReadLine(filepath string, ch chan string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Printf("Read Line: %s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	ch <- scanner.Text()
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func createFile(filepath string) {
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
func deleteFile(filepath string) {
	// delete file
	var err = os.Remove(filepath)
	if isError(err) {
		return
	}
}

func eraseFile(filepath string) {
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

func fileReadLine_2(filepath string, wg *sync.WaitGroup, ch chan string, quit chan bool) {
	tick := time.Tick(time.Second)
	//terminate := time.After(5 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Println("Tick")
			ch <- "Tick"
			/*
				case <-terminate:
					fmt.Println("Terminate")
					wg.Done()
					quit <- true
					return
			*/
		}
	}
}

func main() {
	conf := lib.ConfigReader("../config/proc_stream.yaml")
	var wg sync.WaitGroup
	fmt.Println(conf.CommonConf.Log.Level)
	fmt.Println(conf.ProcConf.Process.Data_path)

	ch := make(chan string, 10)
	quit := make(chan bool)
	wg.Add(1)
	go fileReadLine_2(conf.ProcConf.Process.Data_path, &wg, ch, quit)
	fmt.Println("Start for Loop")
	for {
		select {
		case res := <-ch:
			fmt.Println("1-", res)
			/*
				case res := <-quit:
					fmt.Println("2-", res)
					if res == true {
						fmt.Println("Terminate for-loop!!!")
					}
			*/
		}
	}
	wg.Wait()
	fmt.Println("End for Loop")

	/*
		ch := make(chan string, 1000)
		go fileReadLine(conf.ProcConf.Process.Data_path, ch)
		for {
			select {
			case res := <-ch:
				fmt.Println("=================================")
				fmt.Printf(res)
				fmt.Println("=================================")
			case <-time.After(time.Duration(conf.ProcConf.Process.Cycle_msec) * time.Millisecond):
				fmt.Println("Time out!")
			}
			//fmt.Printf(<-ch)
			//time.Sleep(time.Duration(conf.ProcConf.Process.Cycle_msec) * time.Millisecond)
			//fmt.Println("222")
		}
		defer close(ch)
	*/

	/*
		ch := make(chan string, 10)
		for {
			go func() {
				file, err := os.Open(conf.ProcConf.Process.Data_path)
				if err != nil {
					log.Fatal(err)
				}
				defer file.Close()
				scanner := bufio.NewScanner(file)
				for scanner.Scan() {
					fmt.Printf("Read Line: %s\n", scanner.Text())
					ch <- scanner.Text()
				}
				if err := scanner.Err(); err != nil {
					fmt.Fprintln(os.Stderr, "reading standard input:", err)
				}

				//ch <- scanner.Text()
				eraseFile(conf.ProcConf.Process.Data_path)
			}()
			select {
			case res := <-ch:
				fmt.Println(res)
			case <-time.After(time.Duration(conf.ProcConf.Process.Cycle_msec) * time.Millisecond):
				fmt.Println("Time out Occured!")
			}
			//time.Sleep(time.Duration(conf.ProcConf.Process.Cycle_msec) * time.Millisecond)
		}
	*/
}
