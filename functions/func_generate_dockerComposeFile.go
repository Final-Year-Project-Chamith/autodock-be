package functions

import (
	"autodock-be/dto"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

func GenerateDockerComposeFile(application dto.DockerCompose, repoName string) error {
	tmpl, err := template.ParseFiles("templates/docker-compose/docker-compose.tmp")
	if err != nil {
		return err
	}

	buf := new(strings.Builder)

	err = tmpl.Execute(buf, application)
	if err != nil {
		return err
	}

	dirPath := fmt.Sprintf("docker-compose/%s", repoName)

	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	filePath := fmt.Sprintf("%s/docker-compose.yml", dirPath)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("failed to close output file: %v", err)
		}
	}(file)

	_, err = file.WriteString(buf.String())
	if err != nil {
		return fmt.Errorf("failed to write to output file: %w", err)
	}

	if err := testNginxConfig(); err != nil {
		return err
	}

	return nil
}
func testNginxConfig() error {
	cmd := exec.Command("nginx", "-t")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Nginx test failed: %s\n", string(output))
		return err
	}

	fmt.Printf("Nginx test succeeded: %s\n", string(output))
	return nil
}
