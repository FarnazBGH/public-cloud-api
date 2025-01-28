package instances

import (
	"flag"
	"fmt"

	"public-cloud-api/pkg/cloud"
	"public-cloud-api/pkg/util"
)

func ListInstances(args []string) error {

	// Load environment variables
	apiKey, apiHost := util.LoadEnv()

	// Define flags for command-line arguments
	listInstaneCmd := flag.NewFlagSet("instance list", flag.ExitOnError)
	instanceID := listInstaneCmd.String("id", "", "Instance ID to fetch details for (optional)")

	// Parse the flags for the "listinstances" command
	err := listInstaneCmd.Parse(args)
	if err != nil {
		return fmt.Errorf("error parsing flags: %w", err)
	}

	// Create the cloud client
	client := cloud.NewClient(apiKey, apiHost)

	// If an instance ID is provided, fetch its details
	if *instanceID != "" {
		err := cloud.GetInstanceByID(client, *instanceID)
		if err != nil {
			return fmt.Errorf("error fetching instance details: %w", err)
		}

	} else {
		// Otherwise, list all instances
		err := cloud.ListInstances(client)
		if err != nil {
			return fmt.Errorf("error listing instances: %w", err)
		}
	}
	return nil
}
