package npm

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

var RunError = errors.New("npm run error")

func Run(script, root string) error {
	scripts := getScripts(root)
	if _, exists := scripts[script]; !exists {
		return fmt.Errorf("script does not exist: %s", script)
	}

	c := exec.Command("npm", "run", script)
	out, err := c.Output()
	if err != nil {
		return RunError
	}

	log.Println(string(out))

	return nil
}

func getScripts(root string) map[string]string {
	var obj map[string]json.RawMessage

	fp := filepath.Join(root, "package.json")

	if _, err := os.Stat(fp); os.IsNotExist(err) {
		log.Printf("no package.json.\ndid you npm init?\n")
		return nil
	}

	b, err := os.ReadFile(fp)
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	if err := json.Unmarshal(b, &obj); err != nil {
		log.Println(err.Error())
		return nil
	}

	scripts := make(map[string]string)
	err = json.Unmarshal(obj["scripts"], &scripts)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return scripts
}
