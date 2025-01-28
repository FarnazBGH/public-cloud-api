package cloud

import (
	"fmt"
)

func ListInstances(client *Client) error {
	instancelist, _, err := client.apiClient.PubliccloudAPI.GetInstanceList(client.ctx).Execute()
	if err != nil {
		return fmt.Errorf("error when calling PubliccloudAPI.GetInstanceList: %w", err)
	}

	fmt.Println("Instances:")
	for _, instance := range instancelist.Instances {
		fmt.Printf("- ID: %s, Name: %s, State: %s\n", instance.Id, instance.Region, instance.State)
	}
	return nil
}

func GetInstanceByID(client *Client, instanceID string) error {
	instancelist, _, err := client.apiClient.PubliccloudAPI.GetInstance(client.ctx, instanceID).Execute()
	if err != nil {
		return fmt.Errorf("error when calling PubliccloudAPI.GetInstance: %w", err)
	}

	fmt.Printf("Instance Details:\nID: %s\nRegion: %s\nOfferingType: %s\nProductType: %s \nState: %s\n",
		instancelist.Id, instancelist.Region, instancelist.Type, instancelist.ProductType, instancelist.State)

	return nil
}
