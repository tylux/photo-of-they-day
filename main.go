package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/rwcarlsen/goexif/exif"
	gomail "gopkg.in/gomail.v2"
)

var (
	smtpServer   = flag.String("m", "smtp.gmail.com", "outbound smtp server, defaults to smtp.gmail.com")
	smtpPort     = flag.Int("i", 587, "SMTP Port, defaults to 587")
	smtpUsername = flag.String("u", "", "SMTP username")
	smtpPassword = flag.String("p", "", "SMTP password")
	photosPath   = flag.String("d", "", "Path to photos directory")
	sendTo       = flag.String("w", "", "Addresses to send to")
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func sendMail(p string) {

	fname := p

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	tm, _ := x.DateTime()
	timestamp := tm.String()

	//message body
	message := fmt.Sprintf("<h3>Photo of the day!</h3></br><p>Taken on %s</p>", timestamp)

	m := gomail.NewMessage()
	m.SetHeader("From", *smtpUsername)
	m.SetHeader("To", *sendTo)
	m.SetHeader("Subject", "Random Photo of the day!")
	m.SetBody("text/html", message)
	m.Attach(p)

	d := gomail.NewDialer(*smtpServer, *smtpPort, *smtpUsername, *smtpPassword)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()
	searchDir := *photosPath

	//recursivly walk through all subdirectories and return filepaths
	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		//ignore directories
		if f.IsDir() {
			return nil
		}
		fileList = append(fileList, path)
		return err
	})

	if e != nil {
		panic(e)
	}

	//get random photo based on number of photos
	n := len(fileList)
	randomfile := random(0, n)
	randomPhoto := fileList[randomfile]

	//email photo and extract out timestamp
	fmt.Println("\nSending photo: ", randomPhoto)
	sendMail(randomPhoto)
}
