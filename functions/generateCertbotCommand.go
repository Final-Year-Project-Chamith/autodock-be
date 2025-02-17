package functions

import (
	"autodock-be/dto"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

func GenerateCertbotSHFile(domain string) error {
	tmpl, err := template.ParseFiles("templates\\certbot\\sh.tmp")
	if err != nil {
		return err
	}
	model := dto.Certbot{
		Domain: domain,
	}

	buf := new(strings.Builder)

	err = tmpl.Execute(buf, model)
	if err != nil {
		return err
	}

	err = os.MkdirAll("certbot", os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	filePath := fmt.Sprintf("certbot/certbot_%s.sh", domain)

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
