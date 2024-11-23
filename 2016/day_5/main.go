package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func GetMd5Hash(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func HasEmpty(arr [8]string) bool {
	for _, str := range arr {
		if str == "" {
			return true
		}
	}
	return false
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func main() {
	doorId := "reyedfim"

	counter := 1
	var password1 []string
	var password2 [8]string

	for {
		computable := doorId + strconv.Itoa(counter)
		counter++
		hashed := GetMd5Hash(computable)
		zeroes := hashed[:5]
		possibleChar := string(hashed[5])
		possibleChar2 := string(hashed[6])

		if zeroes == "00000" && possibleChar != string('0') && len(password1) < 8 {
			password1 = append(password1, possibleChar)
		}

		if zeroes == "00000" && HasEmpty(password2) && isNumeric(string(possibleChar)) {
			fmt.Println(computable, hashed, possibleChar, possibleChar2)
			index, _ := strconv.Atoi(possibleChar)
			if index < 8 {
				atLocation := password2[index]
				if atLocation == "" {
					password2[index] = string(possibleChar2)
				}
			}

		}

		if len(password1) >= 8 && !HasEmpty(password2) {
			break
		}
	}

	fmt.Println("Problem 1:", strings.Join(password1, ""))
	fmt.Println("Problem 2:", strings.Join(password2[:], ""))
}
