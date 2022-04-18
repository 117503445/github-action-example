package test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/117503445/github-action-example/src"
	"github.com/stretchr/testify/assert"
)

// changeWorkDir change work dir to project root, in order to load config and other file
func changeWorkDir() error {
	const PROJECT_NAME = "github-action-example"

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	for {
		if filepath.Base(dir) == PROJECT_NAME {
			break
		}
		dir = filepath.Dir(dir)
		if dir == "" {
			return fmt.Errorf("Not Found '%v' in pwd", PROJECT_NAME)
		}
	}

	err = os.Chdir(dir)
	if err != nil {
		panic(err)
	}
	return nil
}

func TestHello(t *testing.T) {
	ast := assert.New(t)
	changeWorkDir()
	go src.Start()
	time.Sleep(time.Millisecond * 100) // wait for server start

	resp, err := http.Get("http://localhost:8080/")
	ast.Nil(err)

	respBytes, err := ioutil.ReadAll(resp.Body)
	ast.Nil(err)

	var respMap map[string]interface{}
	err = json.Unmarshal(respBytes, &respMap)
	ast.Nil(err)

	ast.Equal(0, int(respMap["code"].(float64)))
}
