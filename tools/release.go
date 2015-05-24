package main

import (
  "bufio"
  "fmt"
  "flag"
  // "io"
  "os"
  "strings"
  "strconv"
  "errors"
)

var versionFile = flag.String("v", ".version", "version file")
var releaseType = flag.String("r", "PATCH", "type of release MAJOR.MINOR.PATCH")
var user1File = flag.String("i1", ".firmware/user1.bin", "image user1.bin path")
var user2File = flag.String("i2", ".firmware/user1.bin", "image user2.bin path")

func readVersion(pathname *string) (version int, err error) {
  inputFile, inputError := os.Open(*pathname)
  if inputError != nil {
    fmt.Printf("Uhm, file %s should be created already, shouldn't it?", *pathname)
    return version, errors.New("Cannot open version file")
  }
  inputReader := bufio.NewReader(inputFile)
  inputString, readerError := inputReader.ReadString('\n')
  if readerError != nil {
    fmt.Printf("Error in reading file %s?", readerError)
    return version, errors.New("Cannot read version file");
  }
  tmp := strings.Split(inputString, "\n")
  inputString = tmp[0]

  versionParts := strings.Split(inputString, ".")

  for _, val := range versionParts {
    v, e := strconv.Atoi(val)
    if (e != nil) {
      fmt.Printf("Invalid version file\n");
      return version, errors.New("Invalide version file")
    }
    version = version*256+ v
  }
  return version, nil
}

func main() {
  flag.PrintDefaults()
  flag.Parse()

  version, err := readVersion(versionFile);
  if (err != nil) {
    return;
  }
  fmt.Printf("version %d\n", version)
}