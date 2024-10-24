package functions

import (
	"autodock-be/dto"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

func GenerateDockerComposeFile(application dto.DockerCompose) error {
	tmpl, err := template.ParseFiles("templates\\docker-compose\\docker-compose.tmp")
	buf := new(strings.Builder)
	if err != nil {
		return err
	}
	err = tmpl.Execute(buf, application)
	if err != nil {
		return err
	}
	file, err := os.Create("outs\\docker-compose.yml")
	if err != nil {
		return err
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
	return nil
}
