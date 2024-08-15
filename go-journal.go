package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	// "time"
)

// this is temp while deving
// This should be stored in the program files equivalent
var STOREPATH string = "./dev"

var HEADERSTRING string = fmt.Sprintf(`
=====================================================
Welcome to your Journal. Today's date is <Date here>.
Lay that shit out i guess:
=====================================================
`)

const (
	RIGHT_ARROW = 67
	LEFT_ARROW  = 68
	UP_ARROW    = 65
	DOWN_ARROW  = 66
)

type MenuLocation struct {
	InitialStep int
}

func InitMenu() MenuLocation {
	menu := MenuLocation{}
	menu.InitialStep = 0
	return menu
}

type SelectionItem struct {
	Path     string
	Name     string
	Date     string
	OnSelect func()
}

func NewSelectionItem(filename string, name string, date string, callback func()) SelectionItem {

	sitem := SelectionItem{}
	sitem.Path = fmt.Sprintf("%s/%s.txt", STOREPATH, filename)
	sitem.Name = name
	sitem.Date = date
	sitem.OnSelect = callback

	return sitem
}

func RenderHeader() {
	fmt.Print(HEADERSTRING)
}

func ReceiveEntry() {

	quit := make(chan bool)
	skey := make(chan string)
	defer close(quit)
	defer close(skey)
	// go devPrintTTY(in, out, quit)
	go readStdIn(quit)
	go devPrintTTY(quit, skey)

	for {
		select {
		case <-quit:
			break
		default:
			continue
		}
	}
}

func readStdIn(quit chan<- bool) {
	entname := "test-entry"
	input := bufio.NewReader(os.Stdin)
	entbytes := []byte("==================\n")

	for {

		fmt.Print(">>> ")

		line, err := input.ReadString('\n')
		check(err)

		linebite := []byte(line)
		entbytes = append(entbytes, linebite...)

		if len(linebite) == 1 && linebite[0] == 10 {
			// EOF
			quit <- true
			break
		}
	}

	fmt.Println(entbytes)
	err := os.WriteFile(fmt.Sprintf("%s/%s.txt", STOREPATH, entname), entbytes, 0644)
	check(err)

}

func devPrintTTY(quit <-chan bool, skey chan<- string) {
	b := make([]byte, 3)
	for {
		select {
		case <-quit:
			fmt.Println("quitting")
			return
		default:
			os.Stdin.Read(b)
			bslice := b[:]
			keycode := bslice[len(bslice)-1]
			fmt.Println(keycode)
			switch {
			case reflect.DeepEqual(keycode, RIGHT_ARROW):
				skey <- "RA"
			case reflect.DeepEqual(keycode, LEFT_ARROW):
				skey <- "LA"
			case reflect.DeepEqual(keycode, UP_ARROW):
				skey <- "UA"
			case reflect.DeepEqual(keycode, DOWN_ARROW):
				skey <- "DA"
			}
		}

	}
}

func ReadEntry(trange ...string) error {

	if len(trange) == 0 {
		// get all or the default range
	} else if len(trange) == 2 {
		// get between reange
	} else {
		return errors.New("Invalid number of range arguments")
	}

	return nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func FlagParser(f []string) {
	// disable input buffering so we can read arrow keys in term
	err := exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// TEMP
	// exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	check(err)
	RenderHeader()
	ReceiveEntry()

}

func main() {
	// TODO import using os.Getenv() for consistency
	// GOBIN := "/usr/bin"
	// STORE_FILES := "/var/lib/go-journal"
	// fmt.Println(os.Environ())

	// fmt.Println(reflect.TypeOf(os.Args))

	FlagParser(os.Args[1:])

}
