// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package repository

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	testlib "github.com/uber/zanzibar/test/lib"
)

const (
	exampleGateway                  = "../../examples/example-gateway"
	httpClientUpdateRequestFile     = "data/client/contacts_client_update.json"
	tchannelClientUpdateRequestFile = "data/client/baz_client_update.json"
)

func TestUpdateHTTPClient(t *testing.T) {
	testUpdateClientConfig(t, httpClientUpdateRequestFile, "contacts")
}

func TestUpdateTchannelClient(t *testing.T) {
	testUpdateClientConfig(t, tchannelClientUpdateRequestFile, "baz")
}

func testUpdateClientConfig(t *testing.T, requestFile string, clientName string) {
	req := &ClientConfig{}
	err := readJSONFile(requestFile, req)
	assert.NoError(t, err, "Failed to unmarshal client config.")
	tempDir, err := copyExample(t)
	if !assert.NoError(t, err, "Failed to copy example.") {
		return
	}
	r := &Repository{
		localDir: tempDir,
	}
	clientCfgDir := "clients"
	err = r.UpdateClientConfigs(req, clientCfgDir, "{{placeholder}}")
	if !assert.NoError(t, err, "Failed to write client config.") {
		return
	}

	clientConfigActFile := filepath.Join(tempDir, clientCfgDir, clientName, clientConfigFileName)
	actualClientCfg, err := ioutil.ReadFile(clientConfigActFile)
	if !assert.NoError(t, err, "Failed to read client config file.") {
		return
	}
	clientConfigExpFile := filepath.Join(exampleGateway, clientCfgDir, clientName, clientConfigFileName)
	testlib.CompareGoldenFile(t, clientConfigExpFile, actualClientCfg)

	clientModuleCfg, err := ioutil.ReadFile(filepath.Join(tempDir, clientCfgDir, clientModuleFileName))
	if !assert.NoError(t, err, "Failed to read client module config file.") {
		return
	}
	clientModuleExpFile := filepath.Join(exampleGateway, clientCfgDir, clientModuleFileName)
	testlib.CompareGoldenFile(t, clientModuleExpFile, clientModuleCfg)

	productionJSON, err := ioutil.ReadFile(filepath.Join(tempDir, productionCfgJSONPath))
	if !assert.NoError(t, err, "Failed to read client production JSON config file.") {
		return
	}
	productionJSONExpFile := filepath.Join(exampleGateway, productionCfgJSONPath)
	testlib.CompareGoldenFile(t, productionJSONExpFile, productionJSON)
}

func copyExample(t *testing.T) (string, error) {
	tempDir, err := ioutil.TempDir("", "example-gateway")
	if err != nil {
		return "", err
	}
	err = copyDir(exampleGateway, tempDir, []string{
		filepath.Join(exampleGateway, "build"),
		filepath.Join(exampleGateway, "middlewares"),
	})
	if err != nil {
		return "", err
	}
	t.Logf("Temp dir is created at %s\n", tempDir)
	return tempDir, nil
}

func copyDir(src, dest string, ignoredPrefixes []string) error {
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil || src == path {
			return err
		}
		for _, prefix := range ignoredPrefixes {
			if strings.HasPrefix(path, prefix) {
				return nil
			}
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		destPath := filepath.Join(dest, rel)
		if info.IsDir() {
			return os.Mkdir(destPath, info.Mode())
		}
		srcFile, err := os.Open(path)
		defer closeFile(srcFile)
		if err != nil {
			return err
		}
		destFile, err := os.Create(destPath)
		defer closeFile(destFile)
		if err != nil {
			return err
		}
		_, err = io.Copy(destFile, srcFile)
		return err
	}
	return filepath.Walk(src, walkFn)
}

func closeFile(file *os.File) {
	if err := file.Close(); err != nil {
		fmt.Errorf("Failed to close file %q: %+v", file.Name(), err)
	}
}