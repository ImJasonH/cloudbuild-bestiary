package main

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"testing"
)

func TestYAMLs(t *testing.T) {
	b, err := ioutil.ReadAll(io.LimitReader(rand.Reader, 30))
	if err != nil {
		t.Fatal(err)
	}
	tag := "run-" + base64.URLEncoding.EncodeToString(b)[:5]
	t.Logf("Using build tag %q", tag)

	fis, err := ioutil.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	for _, fi := range fis {
		if !strings.HasSuffix(fi.Name(), ".yaml") {
			continue
		}
		t.Run(fi.Name(), func(t *testing.T) {
			b, err := submit(fi.Name(), tag)
			t.Log(string(b))
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func submit(config, tag string) ([]byte, error) {
	log.Println("tag", tag)
	return exec.Command("gcloud", "builds", "submit", "--config", config, "--no-source", "--substitutions=_TAG="+tag).CombinedOutput()
}
