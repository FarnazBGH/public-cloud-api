package instances

import (
	"flag"
	"fmt"
	"log"

	"public-cloud-api/pkg/cloud"
	"public-cloud-api/pkg/util"
)

func DeleteInstance(args []string) error {

	// Load environment variables
	apiKey, apiHost := util.LoadEnv()

	// Define flags for command-line arguments
	deleteInstaneCmd := flag.NewFlagSet("instance delete", flag.ExitOnError)
	id := deleteInstaneCmd.String("id", "", "Id of the instance")

	// Parse flags
	err := deleteInstaneCmd.Parse(args)
	if err != nil {
		return fmt.Errorf("error parsing flags: %w", err)
	}

	// Validate required flags
	if *id == "" {
		missingFlags := []string{}
		if *id == "" {
			missingFlags = append(missingFlags, "-id")
		}
		return fmt.Errorf("missing required flags: %s", missingFlags)
	}

	// Delete the cloud client
	client := cloud.NewClient(apiKey, apiHost)

	// Delete a new instance
	err = cloud.DeleteInstance(client, *id)
	if err != nil {
		return fmt.Errorf("failed to delete instance: %w", err)
	}

	log.Println("Instance deleted successfully!")
	return nil
}
