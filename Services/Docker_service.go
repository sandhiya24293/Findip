package Services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
	//"strings"
)

type Dockercp struct {
	ADDSOURCE string
	ADDDESTIN string
}

type Exposeport struct {
	Port     string
	Protocal string
}
type WriteDocker struct {
	FROM       string
	ENTRYPOINT string
	EXPOSE     []Exposeport
	RUN        []string
	ENV        []Dockercp
	VOLUME     []string
	ADD        []Dockercp
	COPY       []Dockercp
	CMD        string
	USER       string
	WORKDIR    string
}

func CreateDocker(w http.ResponseWriter, r *http.Request) {
	createFile()
	w.Write([]byte("Successfully created"))
}

func createFile() {

}

func WriteDockerFile(w http.ResponseWriter, r *http.Request) {

	//Create Folder and Dockerfile
	t := time.Now()
	folder := t.Format("20060102150405")
	newpath := filepath.Join("Docker/", folder, "/")
	os.MkdirAll(newpath, os.ModePerm)
	var path = "Docker/" + folder + "/Dockerfile"

	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> done creating file", path)

	var Dockercmd WriteDocker
	err1 := json.NewDecoder(r.Body).Decode(&Dockercmd)
	if err1 != nil {
		fmt.Println("Error on Get particular details", err1)
	}

	// open file using READ & WRITE permission

	var file, err2 = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(err2) {
		return
	}
	defer file.Close()

	if Dockercmd.FROM != "" {
		Fromstring := "FROM " + Dockercmd.FROM + "\n"
		_, err = file.WriteString(Fromstring)
	}

	if Dockercmd.ENTRYPOINT != "" {
		Fromstring := "ENTRYPOINT " + "[" + Dockercmd.ENTRYPOINT + "]" + "\n"
		_, err = file.WriteString(Fromstring)

	}

	if len(Dockercmd.EXPOSE) != 0 {

		for _, v := range Dockercmd.EXPOSE {
			fmt.Println(v.Port, v.Protocal)
			exposestring := "EXPOSE " + v.Port + v.Protocal + "\n"
			_, err = file.WriteString(exposestring)

		}

	}

	if len(Dockercmd.RUN) != 0 {

		for _, v := range Dockercmd.RUN {

			exposestring := "RUN " + v + "\n"
			_, err = file.WriteString(exposestring)

		}

	}

	if len(Dockercmd.ENV) != 0 {

		for _, v := range Dockercmd.ENV {
			fmt.Println(v.ADDSOURCE, v.ADDDESTIN)
			exposestring := "ENV " + v.ADDSOURCE + " " + v.ADDDESTIN + "\n"
			_, err = file.WriteString(exposestring)

		}

	}
	fmt.Println("volume of len", len(Dockercmd.VOLUME))
	if len(Dockercmd.VOLUME) != 0 {

		for _, v := range Dockercmd.VOLUME {

			exposestring := "VOLUME " + v + "\n"
			_, err = file.WriteString(exposestring)

		}

	}

	if len(Dockercmd.ADD) != 0 {

		for _, v := range Dockercmd.ADD {
			fmt.Println(v.ADDSOURCE, v.ADDDESTIN)
			exposestring := "ADD " + v.ADDSOURCE + " " + v.ADDDESTIN + "\n"
			_, err = file.WriteString(exposestring)

		}

	}

	if len(Dockercmd.COPY) != 0 {

		for _, v := range Dockercmd.COPY {
			fmt.Println(v.ADDSOURCE, v.ADDDESTIN)
			exposestring := "COPY " + v.ADDSOURCE + " " + v.ADDDESTIN + "\n"
			_, err = file.WriteString(exposestring)

		}

	}

	if Dockercmd.CMD != "" {
		cmdstring := "CMD " + Dockercmd.CMD + "\n"
		_, err = file.WriteString(cmdstring)
	}

	if Dockercmd.USER != "" {
		userstring := "USER " + Dockercmd.USER + "\n"
		_, err = file.WriteString(userstring)
	}

	if Dockercmd.WORKDIR != "" {
		workdirstring := "WORKDIR " + Dockercmd.WORKDIR + "\n"
		_, err = file.WriteString(workdirstring)
	}
	// write some text line-by-line to file
	if isError(err) {
		return
	}

	// save changes
	err = file.Sync()
	if isError(err) {
		return
	}

	fmt.Println("==> done writing to file")

	var file1, error1 = os.OpenFile(path, os.O_RDWR, 0644)
	if isError(error1) {
		return
	}
	defer file1.Close()

	// read file, line by line
	var text = make([]byte, 1024)
	for {
		_, error1 = file1.Read(text)

		// break if finally arrived at end of file
		if error1 == io.EOF {
			break
		}

		// break if error occured
		if error1 != nil && error1 != io.EOF {
			isError(error1)
			break
		}
	}

	fmt.Println("==> done reading from file")
	w.Write(text)

}

//func ReadDockerFile(w http.ResponseWriter, r *http.Request) {
//	// re-open file
//	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
//	if isError(err) {
//		return
//	}
//	defer file.Close()

//	// read file, line by line
//	var text = make([]byte, 1024)
//	for {
//		_, err = file.Read(text)

//		// break if finally arrived at end of file
//		if err == io.EOF {
//			break
//		}

//		// break if error occured
//		if err != nil && err != io.EOF {
//			isError(err)
//			break
//		}
//	}

//	fmt.Println("==> done reading from file")

//	w.Write([]byte(string(text)))
//}

//func Deletedockerfile(w http.ResponseWriter, r *http.Request) {
//	// delete file
//	var err = os.Remove(path)
//	if isError(err) {
//		return
//	}

//	fmt.Println("==> done deleting file")
//}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
