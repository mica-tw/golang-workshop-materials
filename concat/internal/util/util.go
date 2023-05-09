package util

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
	"log"
	"os"
	"strings"
)

const (
	letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers = "0123456789"
)

func ForEachLines(path string, fn func(string)) {
	f, err := os.Open(path)
	if err != nil {
		log.Println("ERROR: processing file.", err)
		return
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		fn(line)
	}

	if err := sc.Err(); err != nil {
		log.Println("ERROR: reading input file:", err)
	}
}

func GenerateID() string {
	id, err := getShortUUID(uuid.New().String(), 5)
	if err != nil {
		return ""
	}
	return id
}

func getShortUUID(uuidStr string, length int) (string, error) {
	uuidObj, err := uuid.Parse(uuidStr)
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256([]byte(uuidObj.String()))
	hashStr := fmt.Sprintf("%x", hash)
	hashStr = strings.ToLower(hashStr)
	return hashStr[:length], nil
}
