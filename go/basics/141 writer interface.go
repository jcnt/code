package main

import (
	"log"
	"os"
)

//type person struct {
//	first string
//}
//
//func (p person) writeOut(w io.Writer) {
//	_, err := w.Write([]byte(p.first))
//	return err
//}

func main() {
	f, err := os.Create("141output.txt")
	if err != nil {
		log.Fatalf("error %s\n", err)
	}
	defer f.Close()

	s := []byte("Hello Gophers!")

	_, err = f.Write(s)
	if err != nil {
		log.Fatalf("error %s\n", err)
	}

}
