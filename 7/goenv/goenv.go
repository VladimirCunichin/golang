package goenv

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func ReadDir(path string) error {
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	for _, fi := range dir {
		fp := filepath.Join(path, fi.Name())
		if fi.IsDir() {
			if err := ReadDir(fp); err != nil {
				log.Fatal(err)
			}
		} else {
			value, err := ioutil.ReadFile(fp)
			if err != nil {
				return err
			}
			if err := os.Setenv(fi.Name(), string(value)); err != nil {
				return err
			}
		}
	}
	return nil
}

func RunCmd(name string) error {
	fmt.Println(name)
	cmd := exec.Command(name)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
