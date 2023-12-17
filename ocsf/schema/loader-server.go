package schema

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/viper"
	"github.com/valllabh/ocsf-tool/commons"
)

func init() {
	// Discontinue server strategy
	// RegisterSchemaLoader("server", &SchemaServerSchemaLoader{})
}

func (sl *SchemaServerSchemaLoader) Init() {

	// print extensions
	fmt.Println("Extensions:", sl.extensions)

	// print profiles
	fmt.Println("Profiles:", sl.profiles)

}

func (sl *SchemaServerSchemaLoader) Config() {

	// Schema Server Strategy Default Options
	viper.SetDefault("schema.loading.strategies.server.url", "https://schema.ocsf.io")
	viper.SetDefault("schema.loading.strategies.server.api.version", "1.0.0")

}

// Downloads ocsf schema from and write schema to file
func (sl *SchemaServerSchemaLoader) Load() (*OCSFSchema, error) {

	// Load common options
	LoadCommonOptions(sl)

	schemaJSON := GetSchemaJsonFilePath(sl)

	// Check if the schema JSON file doe not exists
	if !commons.PathExists(schemaJSON) {
		// Download the schema and save it to disk
		err := sl.downloadSchemaAndSave()
		if err != nil {
			fmt.Println("Error downloading schema:", err)
			os.Exit(1)
		}
	}

	// Read the JSON file
	data, err := os.ReadFile(schemaJSON)
	if err != nil {
		log.Fatalf("Error reading JSON file: %v", err)
	}

	// Define a variable of the struct type
	var schema OCSFSchema

	// Unmarshal the JSON data into the struct
	if err := json.Unmarshal(data, &schema); err != nil {
		log.Fatalf("Error unmarshalling JSON data: %v", err)
	}

	return &schema, err
}

// GetSchemaHash returns the hash of the schema
func (sl *SchemaServerSchemaLoader) GetSchemaHash() string {
	return commons.Hash(
		strings.Join(sl.GetExtensions(), " "),
		strings.Join(sl.GetProfiles(), " "),
		viper.GetString("schema.loading.strategies.server.api.version"),
	)
}

// SetExtensions sets extensions
func (sl *SchemaServerSchemaLoader) SetExtensions(extensions []string) {
	sl.extensions = extensions
}

// GetExtensions returns extensions
func (sl *SchemaServerSchemaLoader) GetExtensions() []string {
	return sl.extensions
}

// SetProfiles sets profiles
func (sl *SchemaServerSchemaLoader) SetProfiles(profiles []string) {
	sl.profiles = profiles
}

// GetProfiles returns profiles
func (sl *SchemaServerSchemaLoader) GetProfiles() []string {
	return sl.profiles
}

// ProfileExists returns true if profile exists
func (sl *SchemaServerSchemaLoader) ProfileExists(profile string) bool {
	return commons.Contains(sl.GetProfiles(), profile)
}

// ExtensionExists returns true if extension exists
func (sl *SchemaServerSchemaLoader) ExtensionExists(extension string) bool {
	return commons.Contains(sl.GetExtensions(), extension)
}

// func to get schema api base url from config
func getSchemaServerURL(path string, queryParams map[string]string) (string, error) {

	url, urlError := url.Parse(viper.GetString("schema.loading.strategies.server.url"))

	url.Path = commons.CleanPath(url.Path + "/" + path)

	// add query param to url from queryParams
	q := url.Query()
	for key, value := range queryParams {
		q.Add(key, value)
	}
	url.RawQuery = q.Encode()

	return url.String(), urlError
}

// downloadSchemaAndSave downloads the schema JSON from the API and saves it to a file.
func (sl *SchemaServerSchemaLoader) downloadSchemaAndSave() error {

	// TODO: version specific API calls

	// Build the query string for extensions and profiles
	queryParams := map[string]string{}
	if len(sl.extensions) > 0 {
		queryParams["extensions"] = strings.Join(sl.extensions, ",")
	}
	if len(sl.profiles) > 0 {
		queryParams["profiles"] = strings.Join(sl.profiles, ",")
	}

	// Construct the full API URL with query parameters
	url, urlError := getSchemaServerURL("export/schema", queryParams)

	if urlError != nil {
		fmt.Println("Error constructing API URL:", urlError)
		return urlError
	}

	// Send a GET request to the API endpoint
	fmt.Println("Sending GET request to API endpoint...")
	fmt.Printf("API URL: %s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return err
	}
	defer resp.Body.Close()

	// Check if the response status code is 200 (OK)
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("API returned status code: %d\n", resp.StatusCode)
		return fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	// Read the response body
	fmt.Println("Reading response body...")
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return err
	}

	schemaJSON := GetSchemaJsonFilePath(sl)

	// Save the response body as a JSON file
	fmt.Printf("Saving schema to file: %s\n", schemaJSON)

	// Create the directory if it does not exist
	commons.EnsureDirExists(schemaJSON)

	err = os.WriteFile(schemaJSON, body, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return err
	}

	fmt.Printf("Schema saved to %s\n", schemaJSON)
	return nil
}
