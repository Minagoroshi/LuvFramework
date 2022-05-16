package luvcheckerlib

import (
	"bufio"
	"fmt"
	"github.com/gookit/color"
	"os"
	"strconv"
	"strings"
)

type Account struct {
	Username string
	Password string
	Capture  []string
}

//AddCaptureString adds a string to the capture list
func (a *Account) AddCaptureString(name string, data string) {
	a.Capture = append(a.Capture, fmt.Sprintf("%s: %s", name, data))
}

//AddCaptureInt adds an int to the capture list
func (a *Account) AddCaptureInt(name string, data int) {
	a.Capture = append(a.Capture, fmt.Sprintf("%s: %s", name, strconv.Itoa(data)))
}

//AddCaptureBool adds a bool to the capture list
func (a *Account) AddCaptureBool(name string, data bool) {
	a.Capture = append(a.Capture, fmt.Sprintf("%s: %t", name, data))
}

//AddHit Concatenates the username and password into a single string for use in the Hit-list
//Prints the hit to console using the default hits format if verbose is true
func (a *Account) AddHit(verbose bool) {
	HitList = append(HitList, a)
	if verbose {
		color.Success.Println(fmt.Sprintf("[+] [Hit: %s:%s] {%s}", a.Username, a.Password, strings.Join(a.Capture, " | ")))
	}
	OnCheck <- 1
}

//AddFail Concatenates the username and password into a single string
func (a *Account) AddFail(verbose bool) {
	if verbose {
		color.Error.Println(fmt.Sprintf("[-] [Fail: %s:%s]", a.Username, a.Password))
	}
	OnCheck <- 2
}

//LoadCombos is a function to load each line of a file into a map of usernames
//and passwords defined by the Combo struct
func LoadCombos(file string) (int, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != AccountRegex {
			continue
		}
		parts := strings.Split(line, ":")
		AccountList = append(AccountList, Account{parts[0], parts[1], []string{}})
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return len(AccountList), err
}

//GetUserPass returns the username and password of an account in the list
func GetUserPass(i int) string {
	return AccountList[i].Username + ":" + AccountList[i].Password
}
