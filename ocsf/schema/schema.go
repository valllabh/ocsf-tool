package schema

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"github.com/valllabh/ocsf-tool/commons"
)

var ocsfSchema *OCSFSchema
var schemaLoaders map[string]SchemaLoader = map[string]SchemaLoader{}

func LoadOCSFSchema() *OCSFSchema {

	// if ocsfSchema is already loaded, return it
	if ocsfSchema != nil {
		return ocsfSchema
	}

	// load schema.loading.strategy from config
	strategy := viper.GetString("schema.loading.strategy")

	println("Loading schema using strategy:", strategy)

	// get schema loader for the strategy
	schemaLoader, schemaLoaderExists := GetSchemaLoader(strategy)

	if !schemaLoaderExists {
		fmt.Println("Invalid schema loading strategy:", strategy)
		os.Exit(1)
	}

	// initialize the schema loader
	schemaLoader.Init()

	// load the schema
	ocsfSchema, ocsfSchemaLoadingError := schemaLoader.Load()
	if ocsfSchemaLoadingError != nil {
		fmt.Println("Error loading schema:", ocsfSchemaLoadingError)
		os.Exit(1)
	}

	return ocsfSchema
}

// InitOCSFSchemaLoader initializes all the schema loaders
func init() {

	// set schema path
	viper.SetDefault("schema.path", "$CWD/schema")

	// set default schema loading strategy
	viper.SetDefault("schema.loading.strategy", "repository")

	// set default extensions
	viper.SetDefault("extensions.discovery.paths", []string{"$CWD/extensions"})
	viper.SetDefault("extensions.selected", []string{})

	// set default profiles
	viper.SetDefault("profiles.selected", []string{})

}

// get schema loader by name
func GetSchemaLoader(name string) (SchemaLoader, bool) {
	schemaLoader, schemaLoaderExists := schemaLoaders[name]
	return schemaLoader, schemaLoaderExists
}

func RegisterSchemaLoader(name string, schemaLoader SchemaLoader) {
	println("Registering schema loader:", name)
	schemaLoaders[name] = schemaLoader

	schemaLoaders[name].Config()
}

func LoadCommonOptions(sl SchemaLoader) {
	sl.SetExtensions(viper.GetStringSlice("extensions.selected"))
	sl.SetProfiles(viper.GetStringSlice("profiles.selected"))
}

func GetSchemaJsonFilePath(sl SchemaLoader) string {

	fileNameHash := sl.GetSchemaHash()

	// Load directory from config
	path := viper.GetString("schema.path")

	// Prepare file name
	path += fmt.Sprintf("/ocsf-schema-%s.json", fileNameHash)

	// Prepare Directory Path
	path = commons.PathPrepare(path)

	return path
}
