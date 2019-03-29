package main

import (
	"bufio"
	"flag"
	"github.com/paradoxxl/protoJsonTagStripper/lib"
	"log"
	"os"
	"strings"
)

var folder = flag.String("folder", "", "Specify the folder with the proto files. Only .pb.go-Files are processed")
var recursive = flag.Bool("Recursive", false, "Recursive folder traversal")

var file = flag.String("file", "", "Specify only when folder not set. Processes the specified file only")

func main() {
	flag.Parse()

	var err error
	if *folder == "" && *file == "" {
		log.Fatal("A folder or file needs to be specified")
	}

	var files []string
	if *folder != "" {
		files, err = lib.SearchFiles(*folder, *recursive)
		if err != nil {
			log.Fatal("Error Traversing folder ", err)
		}
	} else {
		log.Println(file)
		files = append(files, *file)
	}

	log.Println("Continue replacing files? [Y]/N")
	reader := bufio.NewReader(os.Stdin)
	ans, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("Error reading user input ", err)
	}

	if !strings.ContainsAny(ans, "yY") {
		log.Println("Abort mission")
		return
	}

	for _, v := range files {
		err = lib.ReplaceOmits(v)
		if err != nil {
			log.Println(err)
		}
	}

	log.Println("Done. Good bye.")

}
