package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	for year := 2015; year < 2024; year++ {
		dirPath := filepath.Join("..", fmt.Sprintf("%d", year))

		err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() && strings.Contains(info.Name(), "solution") {
				switch {
				case strings.HasSuffix(info.Name(), ".js"):
					fmt.Printf("Running %s\n", path)
					cmd := exec.Command("node", path)
					runCommand(cmd)
				case strings.HasSuffix(info.Name(), ".go"):
					fmt.Printf("Running %s\n", path)
					cmd := exec.Command("go", "run", path)
					runCommand(cmd)
					// case strings.HasSuffix(info.Name(), ".py"):
					// 	fmt.Printf("Running %s\n", path)
					// 	cmd := exec.Command("python3", path)
					// 	runCommand(cmd)
				}
			}
			return nil
		})

		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func runCommand(cmd *exec.Cmd) {
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("An error occurred: %s\n", err)
	}
	if len(output) > 0 {
		fmt.Println("Output:")
		fmt.Println(string(output))
	}
}
