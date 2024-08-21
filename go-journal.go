package main

import (
	//"bufio"
	"errors"
	"fmt"
	//"io"
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

var (
	RIGHT_ARROW = [3]byte{27, 91, 67}
	LEFT_ARROW  = [3]byte{27, 91, 68}
	UP_ARROW    = [3]byte{27, 91, 65}
	DOWN_ARROW  = [3]byte{27, 91, 66}
	BACK_SPACE  = [3]byte{27, 91, 67}
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

	// stdout channel
	stdoutc := make(chan byte, 1)
	defer close(stdoutc)

	// listening routines
	go readStdIn(quit, stdoutc)
	go specialKeyListener(quit, stdoutc, skey)

	// display routines
	// figure out how termbox does this

	select {
	case <-quit:
		fmt.Println("qutitting")
		return
	}
}

func readStdIn(quit chan<- bool, stdoutc <-chan byte) {

	entname := "test-entry"
	filename := fmt.Sprintf("%s/%s.txt", STOREPATH, entname)

	// might not need to reopen
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	check(err)

	for {
		select {
		case char, ok := <-stdoutc:
			fmt.Println(char)
			_, err := f.WriteString(string(char))
			if !ok || err != nil {
				check(err)
			}
		}
	}

}

func specialKeyListener(quit <-chan bool, stdoutc chan<- byte, skey chan<- string) {
	akbuf := make([]byte, 3)
	specialbuf := make([]byte, 1)
	for {
		select {
		case <-quit:
			fmt.Println("quitting")
			return
		default:
			// can do something like this
			// os.Stdout.Write(outbuf)
			// to display only writter characters, ignoring ^[[C type shit
			// remember to hide char input in term
			// also switch off after exiting program

			// read 3 inputs. reset after each read
			os.Stdin.Read(akbuf)
			akbuf := akbuf[:]

			// read singular special keys
			os.Stdin.Read(specialbuf)
			specialbuf := specialbuf[:]

			fmt.Println(akbuf)
			fmt.Println(specialbuf)
			keycode := akbuf
			// keycode = strings.Join(bslice, "")
			// keycode := string(bslice)
			fmt.Println(keycode)

			// switch {
			// case keycode == RIGHT_ARROW:
			// 	fmt.Println("ra")
			// 	skey <- "RA"
			// case keycode == LEFT_ARROW:
			// 	fmt.Println("la")
			// 	skey <- "LA"
			// case keycode == UP_ARROW:
			// 	fmt.Println("ua")
			// 	skey <- "UA"
			// case keycode == DOWN_ARROW:
			// 	fmt.Println("da")
			// 	skey <- "DA"
			// }
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
	// -echo to turn off, echo to turn on
	exec.Command("stty", "-F", "/dev/tty", "echo").Run()
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
	fmt.Println(reflect.TypeOf("1"))

	defer os.Exit(1)

	FlagParser(os.Args[1:])

}
