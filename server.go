package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func upload(c echo.Context) error {
	mpf, _ := c.MultipartForm()
	for k, _ := range mpf.Value {
		fmt.Println(k)
	}
	for k, _ := range mpf.File {
		fmt.Println(k)
	}
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("1")
	src, err := file.Open()
	if err != nil {
		return err
	}
	fmt.Println("2")
	defer src.Close()

	filename := file.Filename

	t := time.Now()

	saveFilename := strings.Split(filename, ".")[0] + fmt.Sprint(t.UnixNano()) + ".apk"
	fixedFilename := strings.Split(filename, ".")[0] + "_fixed.apk"
	fmt.Println("3")

	dst, err := os.Create(saveFilename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	//Run script on file and output progress

	var stdoutBuf, stderrBuf bytes.Buffer

	cmd := exec.Command("./addSecurityExceptions.sh", saveFilename)

	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Waiting for command to finish...")

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()

	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}

	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)

	return c.Attachment(strings.Split(saveFilename, ".")[0]+"_new.apk", fixedFilename)
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public")
	e.POST("/aase/upload", upload)

	e.Logger.Fatal(e.Start(":1323"))
}
