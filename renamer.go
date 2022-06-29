// This script can be used to insert docnos into the c4 no clean collection
// It takes a path argument that should point to the root of the c4 git repo
// It will insert a docno field into the training documents in c4 no clean
// and write them into a new directory called en.noclean.withdocnos
// The docno of a document will be "c4nc-<file_number>-<line_number>"
// line_number starts at zero for each file
// file_number is taken from the file's name
// e.g. the document on the second line of file c4-train.01234-of-07168.json.gz
// would have a docno of c4nc-1234-00001
// To run, place script in an empty directory then
// use `go mod init` and then `go run main.go -path <path to c4>` in your terminal

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	gzip "github.com/klauspost/pgzip"
	"github.com/schollz/progressbar/v3"
)

func insert_docno(file_number string, line_number int, document string) string {
	return fmt.Sprintf("{\"docno\":\"c4nc-%s-%06d\",%s", file_number[1:], line_number, document[1:])
}

func main() {
	path := flag.String("path", ".", "Path to C4 repo.")
	pattern := flag.String("pattern", "*****", "File pattern to match if you wish to insert docnos into a subset of the collection")
	blocks := flag.Int("blocks", 4, "# of gzip blocks being written in parallel.")
	flag.Parse()

	f, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger := log.New(f, "prefix", log.LstdFlags)

	file_names := fmt.Sprintf("%s/en.noclean/c4-train.%s-of-07168.json.gz", *path, *pattern)

	files, err := filepath.Glob(file_names)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	total_files := len(files)

	bar := progressbar.Default(int64(total_files))

	for _, file_name := range files {
		file_number := file_name[len(file_name)-22 : len(file_name)-22+5]

		fi, err := os.Open(file_name)
		if err != nil {
			logger.Println("Error reading: ", file_number, "\n", err)
			continue
		}

		gr, err := gzip.NewReader(fi)
		if err != nil {
			logger.Println("Error reading: ", file_number, "\n", err)
			continue
		}

		br := bufio.NewReader(gr)

		new_file_name := fmt.Sprintf("%s/en.noclean.withdocnos/%s", *path, file_name[len(file_name)-31:])

		fo, err := os.Create(new_file_name)
		if err != nil {
			log.Fatal(err)
		}

		gw := gzip.NewWriter(fo)
		gw.SetConcurrency(100000, *blocks)

		line_number := 0
		for {
			document, err := br.ReadString('\n')
			if err != nil && err != io.EOF {
				logger.Println("Error reading: ", file_number, "\n", err)
				break
			}
			if err != nil && err == io.EOF {
				break
			}
			new_line := insert_docno(file_number, line_number, document)
			_, err = gw.Write([]byte(new_line))
			if err != nil {
				logger.Println("Error writing:", file_number, "\n", err)
			}
			line_number += 1
		}
		gw.Close()
		fo.Close()
		gr.Close()
		fi.Close()
		bar.Add(1)
	}
	bar.Close()
}
