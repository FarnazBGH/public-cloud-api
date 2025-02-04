package instances

import (
	"flag"
	"fmt"
	"log"

	"public-cloud-api/pkg/cloud"
	"public-cloud-api/pkg/util"

	apiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func UpdateInstance(args []string) error {

	// Load environment variables
	apiKey, apiHost := util.LoadEnv()

	// Define flags for command-line arguments
	updateInstaneCmd := flag.NewFlagSet("instance update", flag.ExitOnError)
	instanceId := updateInstaneCmd.String("id", "", "Instance ID to fetch details")
	name := updateInstaneCmd.String("name", "", "Refrence for the instance")
	instanceType := updateInstaneCmd.String("type", "", "Instance type")
	rootDiskStorageSize := updateInstaneCmd.Int("storage-size", 5, "Root Disk Size")

	// Parse flags
	err := updateInstaneCmd.Parse(args)
	if err != nil {
		return fmt.Errorf("error parsing flags: %w", err)
	}

	// Validate required flags
	if *instanceId == "" {
		missingFlags := []string{}
		if *instanceId == "" {
			missingFlags = append(missingFlags, "-id")
		}
		return fmt.Errorf("missing required flags: %s", missingFlags)
	}

	// Apply  values if they exist for the ones need set
	updateLaunchOpts := NewUpdateInstanceOptsFromValues(*name, *instanceType, *rootDiskStorageSize)

	// Update the cloud client
	client := cloud.NewClient(apiKey, apiHost)

	// Update a new instance
	fmt.Printf("Starting updating of instance created successfully. id:%s\n", *instanceId)
	err = cloud.UpdateInstance(client, *instanceId, *updateLaunchOpts)
	if err != nil {
		return fmt.Errorf("failed to update instance: %w", err)
	}

	log.Println("Instance updated successfully!")
	return nil
}

func NewUpdateInstanceOptsFromValues(name string, instanceType string, rootDiskSize int) *apiclient.UpdateInstanceOpts {
	opts := apiclient.NewUpdateInstanceOpts()
	if name != "" {
		opts.SetReference(name)
	}
	if instanceType != "" {
		typeName := apiclient.TypeName(instanceType)
		opts.Type = &typeName
	}
	if rootDiskSize > 0 {
		diskSize := int32(rootDiskSize)
		opts.SetRootDiskSize(diskSize)
	}
	return opts
}
