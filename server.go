package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type Command struct {
	Cmd   string `json:"Cmd"`
	Os    string `json:"Os"`
	Stdin string `json:"Stdin"`
}

type Out struct {
	Stdout string `json:"stdout"`
	Stderr string `json:"stderr"`
}

type Error struct {
	Err string `json:"err"`
}

func main() {

	http.HandleFunc("/api/v1/remote-execution", Handler)

	println("started")
	log.Fatal(http.ListenAndServeTLS(":8080", "cert/localhost.crt", "cert/localhost.key", nil))
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var get Command

	err := json.NewDecoder(r.Body).Decode(&get)
	if err != nil {
		fmt.Println("Problem with decode", err)
	}

	var ex *exec.Cmd

	args := strings.Split(get.Cmd, " ")
	if get.Cmd == "" {
		WriteError(w, "поле cmd пустое")
		return
	}

	if get.Os == "" {
		WriteError(w, "поле os пустое")
		return
	}

	switch get.Os {
	default:
		fmt.Printf("os %s not supported", get.Os)
		return
	case "windows":
		// 'wine cmd /c комманды'
		args = append([]string{"cmd", "/c"}, args...)

		ex = exec.Command("wine", args...)
	case "linux":
		ex = exec.Command(args[0], args[1:]...)
	}

	// делаем из строки io.Reader
	ex.Stdin = strings.NewReader(get.Stdin)

	// будем в буфферы записывать, а потом в строчки назад
	var wout, werr bytes.Buffer

	ex.Stdout = &wout
	ex.Stderr = &werr

	err = ex.Run()
	if err != nil {
		// добавляем в конец потока ошибок
		werr.WriteString(err.Error())
	}

	err = json.NewEncoder(w).Encode(Out{
		// получаем данные из интерфейсов
		Stdout: wout.String(),
		Stderr: werr.String(),
	})
	if err != nil {
		fmt.Println("Problem with encode")
	}
}

func WriteError(w http.ResponseWriter, err string) error {
	return json.NewEncoder(w).Encode(Error{err})
}
