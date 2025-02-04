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

func GetInstanceByID(client *Client, instanceId string) error {
	instancelist, _, err := client.apiClient.PubliccloudAPI.GetInstance(client.ctx, instanceId).Execute()
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

func DeleteInstance(client *Client, instanceId string) error {
	instanceTerminationReq := client.apiClient.PubliccloudAPI.TerminateInstance(client.ctx, instanceId)
	_, err := client.apiClient.PubliccloudAPI.TerminateInstanceExecute(instanceTerminationReq)
	if err != nil {
		return fmt.Errorf("error terminating the instance: %w", err)
	}
	fmt.Printf("Instance terminated successfully. ID: %s \n", instanceId)
	return nil
}

func UpdateInstance(client *Client, instanceId string, updateInstanceOpts apiclient.UpdateInstanceOpts) error {
	instanceUpdateReq := client.apiClient.PubliccloudAPI.UpdateInstance(client.ctx, instanceId).UpdateInstanceOpts(updateInstanceOpts)
	instance, _, err := client.apiClient.PubliccloudAPI.UpdateInstanceExecute(instanceUpdateReq)
	fmt.Printf("Test. instace: %v \n", instance)
	if err != nil {
		return fmt.Errorf("error updating the instance: %w", err)
	}
	fmt.Printf("Test. instace: %v \n", instance)

	name := instance.GetReference()
	// image := instance.GetImage()  this give us image so I need to do the next line
	imageName := instance.Image.GetName()
	state := instance.GetState()

	fmt.Printf("Instance updated successfully successfully. ID: %s, Name: %s, Image: %s, state: %s \n", instanceId, name, imageName, state)
	return nil
}
