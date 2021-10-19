package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'designerPdfViewer' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY h
 *  2. STRING word
 */

func designerPdfViewer(h []int32, word string) int32 {
    var max_h int32 = -1
    for _, letter := range []rune(word) {
        if h[toNum(letter)] > max_h {
            max_h = h[toNum(letter)]
        }
    }
    return max_h*int32(len(word))

}

// toNum function.
// It returns the alphabetical index of a rune.
func toNum(i rune) int32 {
        var NUM int = 31
        return int32(int(i)&NUM - 1)
    }

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    hTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var h []int32

    for i := 0; i < 26; i++ {
        hItemTemp, err := strconv.ParseInt(hTemp[i], 10, 64)
        checkError(err)
        hItem := int32(hItemTemp)
        h = append(h, hItem)
    }

    word := readLine(reader)

    result := designerPdfViewer(h, word)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}