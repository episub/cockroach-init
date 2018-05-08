package main

// Tool to import scripts from scripts folder

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	scriptsFolder := os.Getenv("SCRIPTS_FOLDER")
	start := time.Now()

	for time.Now().Before(start.Add(time.Second * 120)) {
		ready := isDBReady()

		if ready {
			err := importScripts(scriptsFolder)

			if err != nil {
				log.Fatal(err)
			}
			break
		}
		time.Sleep(time.Second)
	}
}

func importScripts(scriptsFolder string) error {
	files, err := ioutil.ReadDir(scriptsFolder)

	if err != nil {
		return err
	}

	for _, file := range files {
		if !file.IsDir() {
			log.Printf("Importing %s/%s", scriptsFolder, file.Name())
			out, err := exec.Command("bash", "-c", fmt.Sprintf("/cockroach/cockroach sql --port=26257 --insecure < %s/%s", scriptsFolder, file.Name())).Output()

			if err != nil {
				return fmt.Errorf("foo command failed: %v: %s", err, err.(*exec.ExitError).Stderr)
			}

			log.Printf(string(out))
		}
	}

	return nil
}

// isDBReady Returns true if DB is ready, false otherwise.  Will fail if node name has word 'true' in it!
func isDBReady() bool {
	out, err := exec.Command("bash", "-c", "/cockroach/cockroach node status --insecure | grep -i true | wc -l").Output()

	if err != nil {
		log.Printf("Error executing query. Error: %s\n. Output: %s", err, string(out))
		return false
	}

	if strings.TrimSpace(string(out)) == "1" {
		return true
	}

	return false
}
