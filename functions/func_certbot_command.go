package functions

import (
	"fmt"
	"os"
	"os/exec"
)

// GenerateCert writes a certbot script inside the mounted folder and executes it on the VM
func GenerateCert(domain string) error {

	// Define paths
	containerScriptPath := "/certbot/certbot_" + domain + ".sh"     // Inside container (mounted)
	vmScriptPath := "/home/admin/certbot/certbot_" + domain + ".sh" // Inside VM

	if err := GenerateCertbotSHFile(domain); err != nil {
		return err
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
