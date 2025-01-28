package instances

import (
	"flag"
	"fmt"
	"log"

	"public-cloud-api/pkg/cloud"
	"public-cloud-api/pkg/util"

	apiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func CreateInstance(args []string) error {

	// Load environment variables
	apiKey, apiHost := util.LoadEnv()

	// Define flags for command-line arguments
	createInstaneCmd := flag.NewFlagSet("instance create", flag.ExitOnError)
	name := createInstaneCmd.String("name", "", "Refrence for the instance")
	region := createInstaneCmd.String("region", "", "Region for the instance")
	instanceType := createInstaneCmd.String("type", "", "Instance type")
	image := createInstaneCmd.String("image", "", "Image for the instance")
	contractType := createInstaneCmd.String("contract-type", "HOURLY", "Contract type")
	contractTerm := createInstaneCmd.Int("contract-term", 0, "Contract term")
	billingFrequency := createInstaneCmd.Int("billing-frequency", 1, "Billing frequency")
	rootDiskStorageType := createInstaneCmd.String("storage-type", "", "Root Disk Storage Type")
	rootDiskStorageSize := createInstaneCmd.Int("storage-size", 5, "Root Disk Size")
	sshKey := createInstaneCmd.String("ssh-key", "", "Setting ssh key")

	// Parse flags
	err := createInstaneCmd.Parse(args)
	if err != nil {
		return fmt.Errorf("error parsing flags: %w", err)
	}
	launchOpts := apiclient.NewLaunchInstanceOpts(
		apiclient.RegionName(*region),
		apiclient.TypeName(*instanceType),
		*image,
		apiclient.ContractType(*contractType),
		apiclient.ContractTerm(*contractTerm),
		apiclient.BillingFrequency(*billingFrequency),
		apiclient.StorageType(*rootDiskStorageType),
	)

	// Apply optional values if they exist
	if name != nil {
		launchOpts.SetReference(*name)
	}
	if rootDiskStorageSize != nil {
		rootDiskSize := int32(*rootDiskStorageSize)
		launchOpts.SetRootDiskSize(rootDiskSize)
	}
	if sshKey != nil {
		launchOpts.SetSshKey(*sshKey)
	}

	// Validate required flags
	if *region == "" || *instanceType == "" || *image == "" {
		missingFlags := []string{}
		if *region == "" {
			missingFlags = append(missingFlags, "-region")
		}
		if *instanceType == "" {
			missingFlags = append(missingFlags, "-type")
		}
		if *image == "" {
			missingFlags = append(missingFlags, "-image")
		}
		return fmt.Errorf("missing required flags: %s", missingFlags)
	}

	// Create the cloud client
	client := cloud.NewClient(apiKey, apiHost)

	// Create a new instance
	fmt.Printf("Starting creation of instance created successfully. Name:%s, Region:%s, IntanceType:%s, ContractType:%s, ContractTerm:%d, BillingFrequency:%d, StorageType:%s, sshKey:%s, storage-size:%d\n", *name, *region, *instanceType, *contractType, *contractTerm, *billingFrequency, *rootDiskStorageType, *sshKey, *rootDiskStorageSize)
	err = cloud.CreateInstance(client, *launchOpts)
	if err != nil {
		return fmt.Errorf("failed to create instance: %w", err)
	}

	log.Println("Instance created successfully!")
	return nil
}
