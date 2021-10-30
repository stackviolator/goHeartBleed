package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

//TODO make the scanPort() func its own module?
//TEST command:
//scan -h google.com -p 80

// Global vars for super cool colors
var colorGreen = "\033[32m"
var colorReset = "\033[0m"
var colorRed = "\033[31m"

// Takes protocol, hostname and port to scan port, returns boolean
func scanPort(protocol, hostname string, port int,) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		fmt.Print("[" + colorRed + "*" + colorReset + "] Ran into error: ")
		log.Fatalln(err)
		return false
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(conn)
	return true
}

// Gets and handles input for commands
//TODO process arrow keys, add support for port ranges :)
//		Thinking of doing -p 1, 100 for syntax, split cmd[portNumber] by the comma and covert to int? something like that
//TODO need to wait to take multiple commands instead of one
// Need to incorporate output to a file, as well as performing multiple scans - while loop in main, add a quit command
func getCommand() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	var hostname string
	var portNumber string

	displayWelcomeMessage()
	fmt.Println("[" + string(colorGreen) + "*" + colorReset + "] Enter a command:")

	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}

	cmd := strings.Fields(input)
	switch cmd[0] {
	case "help":
		printHelp()
	case "scan":
		hostInArray, hostIndex := checkContains(cmd, "-h")
		portInArray, portIndex := checkContains(cmd, "-p")
		if hostInArray && portInArray{
			hostname = cmd[hostIndex +1]

			if cmd[portIndex + 1] == "ALL" {
				portNumber = "65536"
			} else {
				portNumber = cmd[portIndex + 1]
			}
		} else if portInArray {
			fmt.Println("Missing host")
		} else if hostInArray {
			fmt.Println("Missing port")
		}

		fmt.Println(portNumber)
		return hostname, portNumber
	default:
		printHelp()
	}
	return "", ""
}

func printHelp() {
	fmt.Println("Thank you for using my tool it make me happy thinking people are looking at this :) <3\nContact me via email: jpm7050@psu.edu")
	fmt.Println("Usage [Command] [Options]")
	fmt.Println("----HELP----:\n\tWill display this message, have fun, go crazy")
	fmt.Println("----SCAN----:\n\t-h: Used to specify the host name (full domain or IP)\n\t-p: Used to specify the port(s), a range can be specified with two ports separated by commas (-p 1,100), or ALL for all ports ")
}

// loops through the user inputted command and checks for a flag, returning a boolean and index in the array
func checkContains(arr []string, str string) (bool, int) {
	for k, a := range arr {
		if a == str {
			return true, k
		}
	}
	return false, -1
}

func main() {
	hostname, port := getCommand()
	portNumber, _ := strconv.Atoi(port)

	fmt.Println("Scanning host...")
	open := scanPort("tcp", hostname, portNumber)

	if open {
		fmt.Println("Open port found at " + colorGreen + hostname + ":" + port, colorReset)
	}
}

// Welcome message
func displayWelcomeMessage () {
	fmt.Println("   ____               ___ ___                           __       __________ .__                      .___    _________  .____     .___  \n  / ___\\  ____       /   |   \\   ____  _____  _______ _/  |_     \\______   \\|  |    ____   ____    __| _/    \\_   ___ \\ |    |    |   | \n / /_/  >/  _ \\     /    ~    \\_/ __ \\ \\__  \\ \\_  __ \\\\   __\\     |    |  _/|  |  _/ __ \\_/ __ \\  / __ |     /    \\  \\/ |    |    |   | \n \\___  /(  <_> )    \\    Y    /\\  ___/  / __ \\_|  | \\/ |  |       |    |   \\|  |__\\  ___/\\  ___/ / /_/ |     \\     \\____|    |___ |   | \n/_____/  \\____/      \\___|_  /  \\___  >(____  /|__|    |__|       |______  /|____/ \\___  >\\___  >\\____ |      \\______  /|_______ \\|___| \n                           \\/       \\/      \\/                           \\/            \\/     \\/      \\/             \\/         \\/     ")
}