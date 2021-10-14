package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var input string
var output string
var offset int64
var limit int64

func init() {
	flag.StringVar(&input, "from", "", "file to read from")
	flag.StringVar(&output, "to", "", "file to write to")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
	flag.Int64Var(&limit, "limit", -1, "limit of bytes to copy")
}

func Copy(input, output string, offset, limit int64) error {
	file, err := os.Open(input)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return fmt.Errorf("can't read file information: %s", err.Error())
	}
	bytesToCopy := info.Size()
	inputFile := io.LimitReader(file, bytesToCopy)
	if limit != -1 && limit < bytesToCopy {
		inputFile = io.LimitReader(file, limit)
		bytesToCopy = limit
	}
	if offset > info.Size() {
		return errors.New("offset bigger than filesize")
	}
	if offset > 0 {
		pos, err := file.Seek(offset, 0)
		if err != nil || pos != offset {
			return errors.New("error while trying to offset reading")
		}
		if bytesToCopy+offset >= info.Size() {
			bytesToCopy = info.Size() - offset
		}
	}

	outputFile, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("failed to create file %v", err)
	}

	defer outputFile.Close()
	var current int64 = 0
	bar := pb.StartNew(int(bytesToCopy))
	for current < bytesToCopy {
		bytesWritten, err := io.CopyN(outputFile, inputFile, 2048)
		current += bytesWritten
		if err == io.EOF {
			bar.Add64(bytesWritten)
			break
		} else if err != nil {
			return err
		}
		bar.Add64(bytesWritten)
	}
	bar.Finish()

	return nil
}

func main() {
	flag.Parse()

	err := Copy(input, output, offset, limit)
	if err != nil {
		fmt.Print(err)
	}
}
