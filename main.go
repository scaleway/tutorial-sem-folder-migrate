package main

import (
	"fmt"

	secret_manager "github.com/scaleway/scaleway-sdk-go/api/secret/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func main() {
	// Create a Scaleway client
	client, _ := scw.NewClient(scw.WithEnv())
	api := secret_manager.NewAPI(client)
	projectID, _ := client.GetDefaultProjectID()

	// We create a secret located on the root path
	path := "/"
	secret, _ := api.CreateSecret(&secret_manager.CreateSecretRequest{
		ProjectID: projectID,
		Name:      "ultra-confidential",
		Path:      &path,
	})
	fmt.Println("Created secret : ")
	fmt.Println("ID: ", secret.ID, " Name: ", secret.Name, " Path: ", secret.Path)

	// We create a folder called `confidential-folder` on the root path
	folderPath := "/"
	folder, _ := api.CreateFolder(&secret_manager.CreateFolderRequest{
		ProjectID: projectID,
		Name:      "confidential-folder",
		Path:      &folderPath,
	})
	fmt.Println("New folder : ")
	fmt.Println("Name: ", folder.Name, " Path: ", folder.Path)

	// We migrate the `ultra-confidential` secret into the `confidential-folder` folder
	destinationPath := "/confidential-folder"
	updatedSecret, _ := api.UpdateSecret(&secret_manager.UpdateSecretRequest{
		SecretID: secret.ID,
		Path:     &destinationPath,
	})

	fmt.Println("updated secret : ")
	fmt.Println("ID: ", updatedSecret.ID, " Name: ", updatedSecret.Name, " Path: ", updatedSecret.Path)
}
