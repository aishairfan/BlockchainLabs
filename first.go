package main

import (
    "crypto/sha256"
    "fmt"
    "strconv"
    "strings"
    "time"
)

func mineBlock(targetPrefix string, blockData string) (string, int, int64) {
    nonce := 0
    start := time.Now()

    for {
        dataToHash := blockData + strconv.Itoa(nonce)
        hash := sha256.Sum256([]byte(dataToHash))
        hashString := fmt.Sprintf("%x", hash)

        if strings.HasPrefix(hashString, targetPrefix) {
            end := time.Now()
            elapsed := end.Sub(start).Milliseconds()
            return hashString, nonce, elapsed
        }

        nonce++
    }
}

func main() {
    targetPrefix := "00"
    blockData := "This is a secret message"

    hash, nonce, elapsed := mineBlock(targetPrefix, blockData)

    fmt.Printf("Nonce found: %d\n", nonce)
    fmt.Printf("Hash: %s\n", hash)
    fmt.Printf("Attempts: %d\n", nonce+1)
    fmt.Printf("Time taken (ms): %d\n", elapsed)
}