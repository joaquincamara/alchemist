package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

var templ = template.Must(template.New("png").Parse(templateStr))

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	goHT(path)

	//initServer()
}

// Listen for changes at the current dir
func goHT(dirPath string) {
	doneChan := make(chan bool)

	go func(doneChan chan bool) {
		defer func() {
			doneChan <- true
		}()

		err := watchDir(dirPath)
		if err != nil {
			log.Println(err)
		}

		fmt.Println("Working dir Changed")
	}(doneChan)

	<-doneChan
}

func watchDir(dirPath string) error {
	initialStat, err := os.Stat(dirPath)
	if err != nil {
		return err
	}

	for {
		stat, err := os.Stat(dirPath)
		if err != nil {
			return err
		}

		if stat.Size() != initialStat.Size() || stat.ModTime() != initialStat.ModTime() {
			break
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}

func initServer() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(png))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func png(w http.ResponseWriter, req *http.Request) {
	templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="https://user-images.githubusercontent.com/26718123/94979651-96461b00-04e9-11eb-94d6-660af9663975.{{.}}" />
<br>
<br>
<br>
{{end}}
<form action="/" name=f method="GET">
    <input maxLength=1024 size=70 name=s value="" title="Text to QR Encode">
    <input type=submit value="Show QR" name=qr>
</form>
</body>
</html>
`
