package main

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestFile() int64 {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatalf("failed to create test file")
	}
	defer file.Close()

	bw := bufio.NewWriter(file)
	for i := 0; i < 100; i++ {
		_, err = bw.Write([]byte("1 2 3 4 5 6 7 8 9 q w e r t y u i o p\n"))
		if err != nil {
			log.Fatalf("error writing to input testFile")
		}
	}
	bw.Flush()
	if err != nil {
		log.Fatalf("failed to write: %v", err)
	}
	inputInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("failed to get input file info")
	}
	return inputInfo.Size()
}

func removeTestFile(fileName string) {
	err := os.Remove(fileName)
	if err != nil {
		log.Fatalf("Failed to remove file: %s", err.Error())
	}
}

func TestCopy(t *testing.T) {
	inputSize := createTestFile()
	defer removeTestFile("test.txt")

	err := Copy("test.txt", "result.txt", 0, -1)
	if err != nil {
		log.Fatalf("failed to copy")
	}

	output, err := os.Open("result.txt")
	if err != nil {
		log.Fatalf("failed to open test output file")
	}
	defer removeTestFile("result.txt")
	defer output.Close()
	outputInfo, err := output.Stat()
	if err != nil {
		log.Fatalf("failed to get test output info")
	}
	assert.Equal(t, inputSize, outputInfo.Size(), "copy equal files")
}

func TestCopyLimitZero(t *testing.T) {
	createTestFile()
	defer removeTestFile("test.txt")

	err := Copy("test.txt", "result.txt", 0, 0)
	if err != nil {
		log.Fatalf("failed to copy")
	}

	output, err := os.Open("result.txt")
	if err != nil {
		log.Fatalf("failed to open test output file")
	}
	defer removeTestFile("result.txt")
	defer output.Close()
	outputInfo, err := output.Stat()
	if err != nil {
		log.Fatalf("failed to get test output info")
	}
	assert.Equal(t, int64(0), outputInfo.Size(), "copy files with limit 0")
}

func TestCopyLimit(t *testing.T) {
	inputSize := createTestFile()
	defer removeTestFile("test.txt")
	limit := int64(1024)
	err := Copy("test.txt", "result.txt", 0, limit)
	if err != nil {
		log.Fatalf("failed to copy")
	}

	output, err := os.Open("result.txt")
	if err != nil {
		log.Fatalf("failed to open test output file")
	}
	defer removeTestFile("result.txt")
	defer output.Close()
	outputInfo, err := output.Stat()
	if err != nil {
		log.Fatalf("failed to get test output info")
	}
	if outputInfo.Size() < limit {
		limit = inputSize
	}
	assert.Equal(t, limit, outputInfo.Size(), "copy files with limit 0")
}

func TestCopyOffset(t *testing.T) {
	inputSize := createTestFile()
	defer removeTestFile("test.txt")
	offsetSize := int64(420)
	err := Copy("test.txt", "result.txt", offsetSize, -1)
	if err != nil {
		log.Fatalf("failed to copy")
	}

	output, err := os.Open("result.txt")
	if err != nil {
		log.Fatalf("failed to open test output file")
	}
	defer removeTestFile("result.txt")
	defer output.Close()
	outputInfo, err := output.Stat()
	if err != nil {
		log.Fatalf("failed to get test output info")
	}
	assert.Equal(t, inputSize-offsetSize, outputInfo.Size(), "copy files with limit 0")
}
