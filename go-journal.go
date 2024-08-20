package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	// "time"
	// "strings"
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
	// concat buffer value. repeat for above for better accuracy
	BACK_SPACE = 12700
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
	// control channels
	quit := make(chan bool)
	skey := make(chan string)
	defer close(quit)
	defer close(skey)

	// listening routines
	go readStdIn(quit)
	go specialKeyListener(quit, skey)

	// display routines
	// figure out how termbox does this

	select {
	case <-quit:
		fmt.Println("qutitting")
		return
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
	return

}

func specialKeyListener(quit <-chan bool, skey chan<- string) {
	b := make([]byte, 3)
	for {
		select {
		case <-quit:
			fmt.Println("quitting")
			return
		default:
			os.Stdin.Read(b)
			bslice := string(b[:])
			fmt.Println(bslice)
			fmt.Println(b)
			//keycode = strings.Join(bslice, "")
			keycode := ""
			fmt.Println(bslice)
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
	err := exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// might be worth experimenting with this and implementing your own buffer
	//exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	check(err)
	RenderHeader()
	ReceiveEntry()
}

func main() {
	// TODO import using os.Getenv() for consistecy
	// GOBIN := "/usr/bin"
	// STORE_FILES := "/var/lib/go-journal"
	// fmt.Println(os.Environ())

	// fmt.Println(reflect.TypeOf(os.Args))

	defer os.Exit(1)

	FlagParser(os.Args[1:])

}
