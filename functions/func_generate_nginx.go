package functions

import (
	"autodock-be/dto"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

func GenerateNginxFile(application dto.NginxConf) error {
	tmpl, err := template.ParseFiles("templates/nginx/nginx-conf.tmp")
	if err != nil {
		return err
	}

	buf := new(strings.Builder)

	err = tmpl.Execute(buf, application)
	if err != nil {
		return err
	}

	err = os.MkdirAll("nginx.conf", os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	filePath := fmt.Sprintf("nginx.conf/%s.conf", application.ServerName)

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

	return nil
}
