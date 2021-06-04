package main

import (

	// "compress/gzip"
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

func main() {
	path := flag.String("path", ".", "Path to C4 repo.")
	pattern := flag.String("pattern", "*****", "File pattern to match.")
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

	count := len(files)

	bar := progressbar.Default(int64(count))

	for _, file := range files {
		file_number := file[len(file)-22 : len(file)-22+5]

		fi, err := os.Open(file)
		if err != nil {
			logger.Println(os.Stderr, "Error reading: ", file_number, "\n", err)
			continue
		}

		gr, err := gzip.NewReader(fi)
		if err != nil {
			logger.Println(os.Stderr, "Error reading: ", file_number, "\n", err)
			continue
		}

		br := bufio.NewReader(gr)

		new_file := fmt.Sprintf("%s/en.noclean.withdocnos/%s", *path, file[len(file)-31:])

		fo, err := os.Create(new_file)
		if err != nil {
			log.Fatal(err)
		}

		gw := gzip.NewWriter(fo)
		gw.SetConcurrency(100000, *blocks)

		i := 0
		for {
			line, err := br.ReadString('\n')
			if err != nil && err != io.EOF {
				logger.Println(os.Stderr, "Error reading: ", file_number, "\n", err)
				break
			}
			if err != nil && err == io.EOF {
				break
			}
			new_line := fmt.Sprintf("{\"docno\":\"%s.%05d\",%s", file_number, i, line[1:])
			_, err = gw.Write([]byte(new_line))
			if err != nil {
				logger.Println(os.Stderr, "Error writing:", file_number, "\n", err)
			}
			i += 1
		}
		gw.Close()
		fo.Close()
		gr.Close()
		fi.Close()
		bar.Add(1)
	}
	bar.Close()
}

