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

	// Now we want to migrate the `ultra-confidential` secret into the `confidential-path` path
	destinationPath := "/confidential-path"
	updatedSecret, _ := api.UpdateSecret(&secret_manager.UpdateSecretRequest{
		SecretID: secret.ID,
		Path:     &destinationPath,
	})

	fmt.Println("updated secret : ")
	fmt.Println("ID: ", updatedSecret.ID, " Name: ", updatedSecret.Name, " Path: ", updatedSecret.Path)
}
