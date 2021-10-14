package goenv

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func generateEnvFilesInTempFolder() (string, error) {
	tmpDir, err := ioutil.TempDir("", "goenv")
	if err != nil {
		return "", err
	}

	one, err := ioutil.TempFile(tmpDir, "goenv-")
	if err != nil {
		log.Fatal(err)
	}
	defer one.Close()
	one.WriteString(one.Name())

	two, err := ioutil.TempFile(tmpDir, "goenv-")
	if err != nil {
		log.Fatal(err)
	}
	defer two.Close()
	two.WriteString(two.Name())

	return tmpDir, nil
}

func TestAddEnvsFromFolder(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name:    "Test add envs from folder 1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpDir, err := generateEnvFilesInTempFolder()
			if err != nil {
				log.Fatal(err)
			}
			tt.path = tmpDir
			if err := ReadDir(tt.path); (err != nil) != tt.wantErr {
				t.Errorf("AddEnvsFromFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
			count := 0
			for _, env := range os.Environ() {
				if strings.Contains(env, "goenv") {
					count++
				}
			}
			if count != 3 {
				t.Errorf("count files must be 3, not %v", count)
			}
		})
	}
}
