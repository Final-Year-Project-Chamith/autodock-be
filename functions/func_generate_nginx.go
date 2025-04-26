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
func GenerateNginxFileCMD(config dto.NginxConf) error {
	filePath := fmt.Sprintf("/etc/nginx/conf.d/%s.conf", config.ServerName)

	
	nginxTemplate := `
server {
    listen 80;
  
    server_name {{.ServerName}};
    
    location / {
        proxy_pass       http://localhost:{{.Port}};
        proxy_redirect   off;
        proxy_set_header Host $host;
        proxy_set_header Proxy "";

        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Host $server_name;
        proxy_set_header X-Forwarded-Proto $scheme;

        # log files
        access_log /var/log/nginx/{{.ServerName}}.access.log;
        error_log /var/log/nginx/{{.ServerName}}.error.log;
    }

}
`
	tmpl, err := template.New("nginxConf").Parse(nginxTemplate)
	if err != nil {
		return err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("could not create Nginx config file: %v", err)
	}
	defer file.Close()

	err = tmpl.Execute(file, config)
	if err != nil {
		return fmt.Errorf("could not write to Nginx config file: %v", err)
	}

	return nil
}