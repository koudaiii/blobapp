package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/labstack/echo/v4"
)

// handleError is a helper function that handles errors in this quickstart.
func handleError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// getClientFromContext is a helper function that retrieves the ServiceURL from the context.
func getClientFromContext(c echo.Context) *azblob.Client {
	return c.Get("client").(*azblob.Client)
}

// Get all routes
func getAllRoutes(c echo.Context) error {
	routes := c.Echo().Routes()
	var routeList []string
	for _, route := range routes {
		routeList = append(routeList, fmt.Sprintf("%s %s", route.Method, route.Path))
	}
	return c.JSON(http.StatusOK, routeList)
}

// Get list of containers.
func getContainers(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get list of containers")
}

// Get a container.
func getContainer(c echo.Context) error {
	return c.JSON(http.StatusOK, "Get a container")
}

// Create the container.
func createContainer(c echo.Context) error {
	return c.JSON(http.StatusOK, "Create the container")
}

// Delete the container.
func deleteContainer(c echo.Context) error {
	return c.String(http.StatusOK, "Deleted container successfully")
}

// Upload a file in the container
func uploadFile(c echo.Context) error {
	return c.String(http.StatusOK, "Created file successfully")
}

// Download a file in the container.
func downloadFile(c echo.Context) error {
	return c.String(http.StatusOK, "Downloaded file successfully")
}

// Delete the blob.
func deleteFile(c echo.Context) error {
	return c.String(http.StatusOK, "Deleted file successfully")
}

func main() {
	// Get Azure Storage Account Name
	accountName, ok := os.LookupEnv("AZURE_STORAGE_ACCOUNT_NAME")
	if !ok {
		panic("AZURE_STORAGE_ACCOUNT_NAME could not be found")
	}

	// Get Azure Storage Account Key
	serviceURL := fmt.Sprintf("https://%s.blob.core.windows.net/", accountName)

	// create a credential by using system identity managed or user identity managed
	// if you want to use service principal, you can use azidentity.NewClientSecretCredential
	// https://pkg.go.dev/github.com/Azure/azure-sdk-for-go/sdk/azidentity@v1.4.0#ManagedIDKind
	// ManagedIdentityCredential authenticates an Azure managed identity in any hosting environment supporting managed identities. This credential authenticates a *system-assigned identity by default*.
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	handleError(err)

	// create a ServiceURL to call the API
	client, err := azblob.NewClient(serviceURL, cred, nil)
	handleError(err)

	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("client", client)
			return next(c)
		}
	})

	// Get all routes
	e.GET("/", getAllRoutes)

	// Get list blob storage container
	e.GET("containers", getContainers)

	// Get blob storage container
	e.GET("containers/:containerName", getContainer)

	// Create blob storage container
	e.POST("containers/:containerName", createContainer)

	// Delete blob storage container
	e.DELETE("containers/:containerName", deleteContainer)

	// Upload blob storage file
	e.POST("containers/:containerName/:fileName", uploadFile)

	// Download blob storage file
	e.GET("containers/:containerName/:fileName", downloadFile)

	// Delete blob storage file
	e.DELETE("containers/:containerName/:fileName", deleteFile)

	e.Logger.Fatal(e.Start(":1323"))
}
