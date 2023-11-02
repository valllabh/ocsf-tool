package schema

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	apiURL     = "https://schema.ocsf.io/export/schema"
	schemaJSON = "./ocsf-schema.json"
)

func LoadOCSFSchema() (OCSFSchema, error) {

	// Download the schema and save it to disk
	err := DownloadSchemaAndSave(nil, nil)
	if err != nil {
		fmt.Println("Error downloading schema:", err)
		os.Exit(1)
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

	return schema, err
}

// DownloadSchemaAndSave downloads the schema JSON from the API and saves it to a file.
func DownloadSchemaAndSave(extensions []string, profiles []string) error {

	// Return if the schema JSON file already exists
	if _, err := os.Stat(schemaJSON); err == nil {
		fmt.Printf("Using schema file: %s\n", schemaJSON)
		return nil
	}

	// Build the query string for extensions and profiles
	queryParams := []string{}
	if len(extensions) > 0 {
		queryParams = append(queryParams, "extensions="+strings.Join(extensions, ","))
	}
	if len(profiles) > 0 {
		queryParams = append(queryParams, "profiles="+strings.Join(profiles, ","))
	}
	queryString := ""
	if len(queryParams) > 0 {
		queryString = "?" + strings.Join(queryParams, "&")
	}

	// Construct the full API URL with query parameters
	fullURL := apiURL + queryString

	// Send a GET request to the API endpoint
	fmt.Println("Sending GET request to API endpoint...")
	fmt.Printf("API URL: %s\n", fullURL)
	resp, err := http.Get(fullURL)
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

	// Save the response body as a JSON file
	fmt.Printf("Saving schema to file: %s\n", schemaJSON)
	err = os.WriteFile(schemaJSON, body, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return err
	}

	fmt.Printf("Schema saved to %s\n", schemaJSON)
	return nil
}
