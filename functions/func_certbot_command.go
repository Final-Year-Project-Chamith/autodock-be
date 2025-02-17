package functions

import (
	"fmt"
	"os"
	"os/exec"
)

// GenerateCert writes a certbot script inside the mounted folder and executes it on the VM
func GenerateCert(domain string) error {
	// Define the script content
	scriptContent := fmt.Sprintf(`#!/bin/bash
certbot --nginx -d %s --non-interactive --agree-tos -m chamith.eos@gmail.com --redirect
`, domain)

	// Define paths
	containerScriptPath := "/certbot/certbot_script.sh" // Inside container (mounted)
	vmScriptPath := "/home/admin/certbot/certbot_script.sh" // Inside VM

	// Step 1: Create the script inside the container (mounted in VM)
	if err := os.WriteFile(containerScriptPath, []byte(scriptContent), 0700); err != nil {
		return fmt.Errorf("failed to create script file in container: %v", err)
	}

	// Step 2: Execute the script on the VM
	cmd := exec.Command("bash", "-c", vmScriptPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error executing script on VM: %v\nOutput: %s", err, string(output))
	}

	// Step 3: Delete the script after execution from VM
	if err := os.Remove(containerScriptPath); err != nil {
		return fmt.Errorf("failed to delete script file on VM: %v", err)
	}

	fmt.Println("✅ SSL Certificate generated successfully for:", domain)
	fmt.Println("📜 Output:", string(output))
	return nil
}
