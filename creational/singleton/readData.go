package singleton

import (
	"bufio"
	"os"
	"strconv"
)

func readData(path string) (map[string]int, error) {
  // ex, err := os.Executable()
  // if err != nil {
  //   panic(err)
  // }
  // exPath := filepath.Dir(ex)
	exPath := "/home/renatospaka/dev/cursos/udemy/design-patterns-go/my-code/singleton"
	
  // file, err := os.Open(exPath + path)
  file, err := os.Open(exPath + "/" + path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanLines)

  result := map[string]int{}
  
  for scanner.Scan() {
    k := scanner.Text()
    scanner.Scan()
    v, _ := strconv.Atoi(scanner.Text())
    result[k] = v
  }

  return result, nil
}