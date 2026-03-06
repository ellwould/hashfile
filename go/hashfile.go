/*
MIT License

Copyright (c) 2026 Elliot Michael Keavney

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Clear Screen Function
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

// American National Standards Institute (ANSI) reset colour code
const resetColour = "\033[0m"

// American National Standards Institute (ANSI) text colour code
const textBoldBlack = "\033[1;30m"
const textBoldWhite = "\033[1;37m"

// American National Standards Institute (ANSI) background colour codes
const bgRed = "\033[41m"
const bgGreen = "\033[42m"
const bgYellow = "\033[43m"
const bgBlue = "\033[46m"
const bgMagenta = "\033[45m"

// Function to draw box with stars around a message
func messageBox(message string, bgColour string) {
	clearScreen()
	topBottomStar := strings.Repeat(" ✸", (len(message)/2)+6)
	inbetweenSpace := strings.Repeat(" ", len(message)+8)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("     " + bgColour + textBoldWhite + topBottomStar + " " + resetColour)
	fmt.Println("     " + bgColour + textBoldWhite + " ✸" + inbetweenSpace + "✸ " + resetColour)
	fmt.Println("     " + bgColour + textBoldWhite + " ✸    " + message + "    ✸ " + resetColour)
	fmt.Println("     " + bgColour + textBoldWhite + " ✸" + inbetweenSpace + "✸ " + resetColour)
	fmt.Println("     " + bgColour + textBoldWhite + topBottomStar + " " + resetColour)
	fmt.Println("")
	fmt.Println("")
	os.Exit(0)
}

// Function to compute hash of a file
func hashTheFile(rootDirPath string, fileName string, algorithmVersion string) string {

	// Go introduced OpenRoot in version 1.24, it restricts file operations to a single directory
	rootDir, err := os.OpenRoot(rootDirPath)
	if err != nil {
		messageBox("Directory path cannot be opened or does not exist", bgYellow)
	}

	defer rootDir.Close()

	// Open the file
	file, err := rootDir.Open(fileName)
	if err != nil {
		messageBox("File cannot be opened or does not exist", bgYellow)
	}

	defer file.Close()

	if algorithmVersion == "sha256" || algorithmVersion == "sha-256" || algorithmVersion == "SHA256" || algorithmVersion == "SHA-256" {
		hash := sha256.New()
		if _, err := io.Copy(hash, file); err != nil {
			log.Fatal(err)
		}
		return hex.EncodeToString(hash.Sum(nil))
	} else if algorithmVersion == "sha512" || algorithmVersion == "sha-512" || algorithmVersion == "SHA512" || algorithmVersion == "SHA-512" {
		hash := sha512.New()
		if _, err := io.Copy(hash, file); err != nil {
			log.Fatal(err)
		}
		return hex.EncodeToString(hash.Sum(nil))
	} else {
		messageBox("Algorithm Required!", bgYellow)
		return ""
	}
}

// Function to terminate the running program if a user types "exit" or "quit"
func exit(value string) {
	if value == "exit" || value == "quit" {
		fmt.Print(resetColour)
		clearScreen()
		os.Exit(0)
	}
}

func main() {
	var (
		expectedHashValue string
		algorithmVersion  string
		rootDirPath       string
		fileName          string
	)
	clearScreen()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("     " + bgMagenta + textBoldWhite + " ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓ " + resetColour)
	fmt.Println("     " + bgMagenta + textBoldWhite + " ┃                                                                       ┃ " + resetColour)
	fmt.Println("     " + bgMagenta + textBoldWhite + " ┃   ██╗  ██╗  █████╗  ███████╗ ██╗  ██╗ ███████╗ ██╗ ██╗      ███████╗  ┃ " + resetColour)
	fmt.Println("     " + bgMagenta + textBoldWhite + " ┃   ██║  ██║ ██╔══██╗ ██╔════╝ ██║  ██║ ██╔════╝ ██║ ██║      ██╔════╝  ┃ " + resetColour)
	fmt.Println("     " + bgMagenta + textBoldWhite + " ┃   ███████║ ███████║ ███████╗ ███████║ █████╗   ██║ ██║      █████╗    ┃ " + resetColour)
	fmt.Println("     " + bgMagenta + textBoldWhite + " ┃   ██╔══██║ ██╔══██║ ╚════██║ ██╔══██║ ██╔══╝   ██║ ██║      ██╔══╝    ┃ " + resetColour)
	fmt.Println("     " + bgMagenta + textBoldWhite + " ┃   ██║  ██║ ██║  ██║ ███████║ ██║  ██║ ██║      ██║ ███████╗ ███████╗  ┃ " + resetColour)
	fmt.Println("     " + bgMagenta + textBoldWhite + " ┃   ╚═╝  ╚═╝ ╚═╝  ╚═╝ ╚══════╝ ╚═╝  ╚═╝ ╚═╝      ╚═╝ ╚══════╝ ╚══════╝  ┃ " + resetColour)
	fmt.Println("     " + bgMagenta + textBoldWhite + " ┃                                                                       ┃ " + resetColour)
	fmt.Println("     " + bgMagenta + textBoldWhite + " ┃     Source code available at https://github.com/ellwould/hashfile     ┃ " + resetColour)
	fmt.Println("     " + bgMagenta + textBoldWhite + " ┃                                                                       ┃ " + resetColour)
	fmt.Println("     " + bgMagenta + textBoldWhite + " ┃            " + resetColour + bgBlue + textBoldWhite + "                                               " + resetColour + bgMagenta + textBoldWhite + "            ┃ " + resetColour)
	fmt.Println("     " + bgMagenta + textBoldWhite + " ┃            " + resetColour + bgBlue + textBoldWhite + "  Type \"exit\" or \"quit\" to terminate hashfile  " + resetColour + bgMagenta + textBoldWhite + "            ┃ " + resetColour)
	fmt.Println("     " + bgMagenta + textBoldWhite + " ┃            " + resetColour + bgBlue + textBoldWhite + "                                               " + resetColour + bgMagenta + textBoldWhite + "            ┃ " + resetColour)
	fmt.Println("     " + bgMagenta + textBoldWhite + " ┃                                                                       ┃ " + resetColour)
	fmt.Println("     " + bgMagenta + textBoldWhite + " ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛ " + resetColour)

	fmt.Print(textBoldBlack)
	fmt.Println("")
	fmt.Print("         Enter expected hash value for the file\n         [usally found inside a checksum file]: ")
	fmt.Scan(&expectedHashValue)
	exit(expectedHashValue)
	fmt.Println("")
	fmt.Print("         Enter the algorithm version [sha256/sha512]: ")
	fmt.Scan(&algorithmVersion)
	exit(algorithmVersion)
	fmt.Println("")
	fmt.Print("         Enter the directory path where the file to hash is located\n         or enter ./ to use current working directory: ")
	fmt.Scan(&rootDirPath)
	exit(rootDirPath)
	fmt.Println("")
	fmt.Print("         Enter the name of the file name: ")
	fmt.Scan(&fileName)
	exit(fileName)
	fmt.Println("")
	fmt.Print(resetColour)
	result := hashTheFile(rootDirPath, fileName, algorithmVersion)
	if expectedHashValue == result {
		messageBox("Hashes match!", bgGreen)
	} else {
		messageBox("WARNING: Hashes do not match!", bgRed)
	}
}

// Contributor(s):
// Elliot Michael Keavney
