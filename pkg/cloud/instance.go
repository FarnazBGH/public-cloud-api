package cloud

import (
	"fmt"

	apiclient "github.com/leaseweb/leaseweb-go-sdk/publiccloud"
)

func ListInstances(client *Client) error {
	instancelist, _, err := client.apiClient.PubliccloudAPI.GetInstanceList(client.ctx).Execute()
	if err != nil {
		return fmt.Errorf("error when calling PubliccloudAPI.GetInstanceList: %w", err)
	}

	fmt.Println("Instances:")
	for _, instance := range instancelist.Instances {
		fmt.Printf("- ID: %s, RegionName: %s, State: %s\n", instance.Id, instance.Region, instance.State)
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

func CreateInstance(client *Client, instanceLaunchOpts apiclient.LaunchInstanceOpts) error {
	instanceLaunchReq := client.apiClient.PubliccloudAPI.LaunchInstance(client.ctx).LaunchInstanceOpts(instanceLaunchOpts)
	instance, _, err := client.apiClient.PubliccloudAPI.LaunchInstanceExecute(instanceLaunchReq)
	instanceName := instance.GetReference()
	if err != nil {
		return fmt.Errorf("error launching the instance: %w", err)
	}
	fmt.Printf("Instance created successfully. ID: %s and Name:%v and Region: %s   \n", instance.Id, instanceName, string(instance.Region))
	return nil
}
