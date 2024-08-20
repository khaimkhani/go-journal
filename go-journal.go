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
			// Use Scanner here maybe
			os.Stdin.Read(b)
			bslice := b[:]
			fmt.Println(reflect.TypeOf(bslice))
			fmt.Println(bslice)
			keycode := bslice
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
