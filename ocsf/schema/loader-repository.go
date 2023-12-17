package schema

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/spf13/viper"
	"github.com/valllabh/ocsf-tool/commons"
	"gopkg.in/src-d/go-git.v4"
)

func init() {
	RegisterSchemaLoader("repository", &SchemaRepositorySchemaLoader{})
}

func (sl *SchemaRepositorySchemaLoader) Init() {

}

func (sl *SchemaRepositorySchemaLoader) Config() {

	// Schema Repository Strategy Default Options

	// config for git repository url
	viper.SetDefault("schema.loading.strategies.repository.url", "https://github.com/ocsf/ocsf-schema")

	// config for directory to clone git repository
	viper.SetDefault("schema.loading.strategies.repository.directory.path", "$CWD/schema/git")

	// config for branch to checkout
	viper.SetDefault("schema.loading.strategies.repository.branch.name", "main")
}

// Downloads ocsf schema from and write schema to file
func (sl *SchemaRepositorySchemaLoader) Load() (*OCSFSchema, error) {

	// Load common options
	LoadCommonOptions(sl)

	// Download git repository
	downloadRepoError := sl.downloadRepo()

	// Check for error while downloading git repository
	if downloadRepoError != nil {
		return nil, downloadRepoError
	}

	// Process git repository
	println("Processing git repository")
	return sl.processRepo()

}

// SetExtensions sets extensions
func (sl *SchemaRepositorySchemaLoader) SetExtensions(extensions []string) {
	sl.extensions = extensions
}

// GetExtensions returns extensions
func (sl *SchemaRepositorySchemaLoader) GetExtensions() []string {
	return sl.extensions
}

// SetProfiles sets profiles
func (sl *SchemaRepositorySchemaLoader) SetProfiles(profiles []string) {
	sl.profiles = profiles
}

// GetProfiles returns profiles
func (sl *SchemaRepositorySchemaLoader) GetProfiles() []string {
	return sl.profiles
}

// ProfileExists returns true if profile exists
func (sl *SchemaRepositorySchemaLoader) ProfileExists(profile string) bool {

	// return true if no profile is configured
	if len(sl.GetProfiles()) == 0 {
		return true
	}

	return commons.Contains(sl.GetProfiles(), profile)
}

// ExtensionExists returns true if extension exists
func (sl *SchemaRepositorySchemaLoader) ExtensionExists(extension string) bool {

	// return true if no extension is configured
	if len(sl.GetExtensions()) == 0 {
		return true
	}

	return commons.Contains(sl.GetExtensions(), extension)
}

func repoPath(path string) string {

	directory := viper.GetString("schema.loading.strategies.repository.directory.path")

	directory = commons.PathPrepare(directory)

	return commons.CleanPath(directory + "/" + path)
}

// function to load repository path from config
func (sl *SchemaRepositorySchemaLoader) loadRepoPath() string {
	// Load directory from config
	directory := viper.GetString("schema.loading.strategies.repository.directory.path")

	// Prepare Directory Path
	directory = commons.PathPrepare(directory)

	return directory
}

// function to download git repository for given URL and save it to disk
func (sl *SchemaRepositorySchemaLoader) downloadRepo() error {

	// Load directory from config
	directory := sl.loadRepoPath()

	// Load URL from config
	url := viper.GetString("schema.loading.strategies.repository.url")

	// Check if directory exists using commons
	if commons.PathExists(directory) {
		// check if directory is a git repository
		if commons.GitIsValidGitRepository(directory) {

			// reset uncommitted changes
			println("Resetting uncommitted changes")
			errGitResetUncommittedChanges := commons.GitResetUncommittedChanges(directory)
			if errGitResetUncommittedChanges != nil {
				return errGitResetUncommittedChanges
			}

			// checkout branch
			branch := viper.GetString("schema.loading.strategies.repository.branch.name")
			println("Checking out branch: " + branch)
			errGitCheckoutBranch := commons.GitCheckoutBranch(directory, branch)

			if errGitCheckoutBranch != nil {
				return errGitCheckoutBranch
			}

			// pull latest changes
			println("Pulling latest changes")
			errGitPullRepository := commons.GitPullRepository(directory)

			// if schema is already up to date
			if errGitPullRepository == git.NoErrAlreadyUpToDate {
				println("Schema is already up to date")
				return nil
			}

		}
	} else {

		// Clone git repository
		println("Cloning git repository from " + url + " to " + directory)
		errGitCloneRepository := commons.GitCloneRepository(url, directory)
		if errGitCloneRepository != nil {
			return errGitCloneRepository
		}

	}

	return nil
}

// func to process repository
func (sl *SchemaRepositorySchemaLoader) processRepo() (*OCSFSchema, error) {

	// Load directory from config
	directory := sl.loadRepoPath()

	// Load schema from directory
	println("Loading schema from " + directory)
	schema, loadSchemaError := sl.loadSchemaFromDirectory(directory)

	// Check for error while loading schema
	if loadSchemaError != nil {
		return nil, loadSchemaError
	}

	// Return schema
	return schema, nil
}

// function to load schema from directory
func (sl *SchemaRepositorySchemaLoader) loadSchemaFromDirectory(directory string) (*OCSFSchema, error) {

	// Declare schema
	var schema OCSFSchema

	// Load version from directory
	schemaVersion, schemaVersionError := sl.loadVersionFromDirectory()

	// Check for error while loading version from directory
	if schemaVersionError != nil {
		return nil, schemaVersionError
	}

	dictionary := Dictionary{}
	objectMap := make(map[string]Object)
	eventMap := make(map[string]Event)

	// Load dictionary from dictionary.json file
	dictionaryFile := repoPath("/dictionary.json")
	dictionaryLoadingError := sl.loadDictionary(dictionaryFile, &dictionary)

	if dictionaryLoadingError != nil {
		return nil, dictionaryLoadingError
	}

	// Load objects defined in json files in the map from /objects directory using schema.Object struct
	objectsDirectory := repoPath("/objects")
	objectLoadingError := sl.loadObjects(objectsDirectory, &objectMap, &dictionary)

	if objectLoadingError != nil {
		return nil, objectLoadingError
	}

	// extend object attributes
	for _, object := range objectMap {
		sl.extendAttribute(&object.Attributes, &dictionary, &objectMap)
	}

	// Load events defined in json files in the map from /events directory using schema.Event struct
	eventsDirectory := repoPath("/events")

	eventLoadingError := sl.loadEvents(eventsDirectory, &eventMap, &dictionary)

	if eventLoadingError != nil {
		return nil, eventLoadingError
	}

	// extend event attributes
	for _, event := range eventMap {
		sl.extendAttribute(&event.Attributes, &dictionary, &objectMap)
	}

	// Build schema
	schema = OCSFSchema{
		Objects:    objectMap,
		Classes:    eventMap,
		Version:    schemaVersion.Version,
		Types:      dictionary.Types.Attributes,
		Dictionary: dictionary,
	}

	// Load extensions defined in json files in the map from /extensions directory using schema.Extension struct
	extensionsDirectory := repoPath("/extensions")
	extensionsLoadingError := sl.loadExtensions(extensionsDirectory, &schema)

	if extensionsLoadingError != nil {
		return nil, extensionsLoadingError
	}

	return &schema, nil
}

// GetSchemaHash returns the hash of the schema
func (sl *SchemaRepositorySchemaLoader) GetSchemaHash() string {
	return commons.Hash(
		strings.Join(sl.GetExtensions(), " "),
		strings.Join(sl.GetProfiles(), " "),
	)
}

// function to load extensions from directory
func (sl *SchemaRepositorySchemaLoader) loadExtensions(directory string, ocsfSchema *OCSFSchema) error {

	// recursively load each file in extensions directory
	err := commons.Walk(directory, func(path string, info os.FileInfo, err error) error {

		// Check if file is a json file
		if strings.HasSuffix(path, "extension.json") {

			// Load extension from file
			err := sl.loadExtensionFromDirectory(commons.Dir(path), ocsfSchema)

			// Check for error while loading extension from file
			if err != nil {
				return err
			}

		}

		return nil
	})

	return err
}

func (sl *SchemaRepositorySchemaLoader) loadExtension(path string) (Extension, error) {

	// Declare extension
	var extension Extension

	// load data from file []byte
	data, loadDataError := os.ReadFile(path)

	if loadDataError != nil {
		return extension, loadDataError
	}

	// unmarshal data into extension
	if err := json.Unmarshal(data, &extension); err != nil {
		return extension, err
	}

	return extension, nil
}

// function to load extension from file
func (sl *SchemaRepositorySchemaLoader) loadExtensionFromDirectory(path string, ocsfSchema *OCSFSchema) error {

	println("Loading extension from " + path)

	extensionFile := path + "/" + "extension.json"
	extension, extensionLoadingError := sl.loadExtension(extensionFile)

	if extensionLoadingError != nil {
		return extensionLoadingError
	}

	// check if extension is configured to be loaded else ignore
	if !sl.ExtensionExists(extension.Name) {
		println("Ignoring extension " + extension.Caption + " as it is not configured to be loaded.")
		return nil
	}

	// Load dictionary from dictionary.json file
	dictionaryFile := commons.CleanPath(path + "/" + "dictionary.json")
	dictionaryLoadingError := sl.loadDictionary(dictionaryFile, &ocsfSchema.Dictionary)

	if dictionaryLoadingError != nil {
		return dictionaryLoadingError
	}

	// Load objects defined in json files in the map from /objects directory using schema.Object struct
	objectsDirectory := commons.CleanPath(path + "/" + "objects")
	objectLoadingError := sl.loadObjects(objectsDirectory, &ocsfSchema.Objects, &ocsfSchema.Dictionary)

	if objectLoadingError != nil {
		return objectLoadingError
	}

	// extend object attributes
	for _, object := range ocsfSchema.Objects {
		sl.extendAttribute(&object.Attributes, &ocsfSchema.Dictionary, &ocsfSchema.Objects)
	}

	// Load events defined in json files in the map from /events directory using schema.Event struct
	eventsDirectory := commons.CleanPath(path + "/" + "events")
	eventLoadingError := sl.loadEvents(eventsDirectory, &ocsfSchema.Classes, &ocsfSchema.Dictionary)

	if eventLoadingError != nil {
		return eventLoadingError
	}

	return nil
}

// func to load version from directory returns version and error. version.json from repo directory is loaded and version is returned
func (sl *SchemaRepositorySchemaLoader) loadVersionFromDirectory() (Version, error) {

	// Declare version
	var version Version

	// Version file path
	versionFile := repoPath("/version.json")

	// Load data from file []byte
	data, loadDataError := os.ReadFile(versionFile)

	if loadDataError != nil {
		return version, loadDataError
	}

	// Unmarshal data into version
	if err := json.Unmarshal(data, &version); err != nil {
		return version, err
	}

	return version, nil
}

// function to load objects from directory
func (sl *SchemaRepositorySchemaLoader) loadObjects(directory string, objects *(map[string]Object), dictionary *Dictionary) error {

	rootDir := commons.Dir(directory)

	// recursively load each file in objects directory
	err := commons.Walk(directory, func(path string, info os.FileInfo, err error) error {

		// Check if file is a json file
		if strings.HasSuffix(path, ".json") {

			println("Loading object from " + path)

			// Load object from file
			object, err := sl.loadObjectFromFile(path, rootDir)

			// Check for error while loading object from file
			if err != nil {
				return err
			}

			// if object.name is blank and object.extend exists and objects contain object.extend then merge object.profiles and merge object.attributes
			if object.Name == "" && object.Extends != "" {

				println("Extending object " + object.Extends)

				originalObject, objectExists := (*objects)[object.Extends]

				if objectExists {
					// if originalObject.Profiles does not exist then create it
					if originalObject.Profiles == nil {
						originalObject.Profiles = make([]string, 0)
					}

					originalObject.Profiles = append(originalObject.Profiles, object.Profiles...)

					// iterate over attributes and add them to originalObject
					for key, value := range object.Attributes {
						originalObject.Attributes[key] = value
					}

					(*objects)[object.Extends] = originalObject
				} else {
					println("Object " + object.Extends + " does not exist")
				}
			} else {

				// Add object to schema
				(*objects)[object.Name] = object

			}
		}

		return nil
	})

	// iterate over objects and resolve extends
	for _, object := range *objects {

		// extend object
		sl.extendObject(&object, objects)

	}

	return err
}

// function to load events from directory
func (sl *SchemaRepositorySchemaLoader) loadEvents(directory string, events *(map[string]Event), dictionary *Dictionary) error {

	rootDir := commons.Dir(directory)

	// recursively load each file in events directory
	err := commons.Walk(directory, func(path string, info os.FileInfo, err error) error {

		// Check if file is a json file
		if strings.HasSuffix(path, ".json") {

			println("Loading event from " + path)

			// Load event from file
			event, err := sl.loadEventFromFile(path, rootDir)

			// Check for error while loading event from file
			if err != nil {
				return err
			}

			// Add event to schema
			(*events)[event.Name] = event
		}

		return nil
	})

	// iterate over events and resolve extends
	for _, event := range *events {

		// extend event
		sl.extendEvent(&event, events)

	}

	return err
}

// function to load Dictionary from dictionary.json file
func (sl *SchemaRepositorySchemaLoader) loadDictionary(path string, dictionary *Dictionary) error {

	// Load data from file []byte
	data, loadDataError := os.ReadFile(path)

	if loadDataError != nil {
		return loadDataError
	}

	_dictionary := Dictionary{}

	// Unmarshal data into dictionary
	if err := json.Unmarshal(data, &_dictionary); err != nil {
		return err
	}

	// set map if nil
	if (*dictionary).Attributes == nil {
		(*dictionary).Attributes = make(map[string]Attribute)
	}

	// merge _dictionary attributes into dictionary attributes
	for key, value := range _dictionary.Attributes {
		(*dictionary).Attributes[key] = value
	}

	return nil
}

// function to get all parent attributes recursively of Object using Extends from given map of items
func (sl *SchemaRepositorySchemaLoader) extendObject(item *Object, items *map[string]Object) {

	// check if item has parent
	if item.Extends != "" {
		// get parent item
		parentItem := (*items)[item.Extends]

		// iterate over parent attributes and add them to attributes
		for key, value := range parentItem.Attributes {
			item.Attributes[key] = value
		}

		// extend parent
		sl.extendObject(&parentItem, items)

		(*items)[item.Name] = *item
	}

}

// function to get all parent attributes recursively of Event using Extends from given map of items
func (sl *SchemaRepositorySchemaLoader) extendEvent(item *Event, items *map[string]Event) {

	// check if item has parent
	if item.Extends != "" {
		// get parent item
		parentItem := (*items)[item.Extends]

		// set attributes
		for key, value := range parentItem.Attributes {
			item.Attributes[key] = value
		}

		// if item category is not null then use parent
		if item.Category == "" {
			item.Category = parentItem.Category
		}

		// get parent attributes
		sl.extendEvent(&parentItem, items)

		(*items)[item.Name] = *item
	}

}

// function to extend attribute from dictionary
func (sl *SchemaRepositorySchemaLoader) extendAttribute(attributes *(map[string]Attribute), dictionary *Dictionary, objects *(map[string]Object)) {

	// iterate over attributes
	for key, attribute := range *attributes {
		dictionaryAttribute, dictionaryAttributeExists := dictionary.Attributes[key]
		// if attribute exists in dictionary then copy it to attribute
		if dictionaryAttributeExists {

			copyError := copier.CopyWithOption(&attribute, &dictionaryAttribute, copier.Option{IgnoreEmpty: true, DeepCopy: true})

			// if attribute.Type does not end with _t
			if !strings.HasSuffix(attribute.Type, "_t") {
				// check if attribute.Type exists in objects
				object, objectExists := (*objects)[attribute.Type]
				if objectExists {
					attribute.Type = "object_t"
					attribute.ObjectType = object.Name
					attribute.ObjectName = object.Caption
				}
			}

			(*attributes)[key] = attribute

			if copyError != nil {
				println("Error while extending attribute " + key)
				println(copyError.Error())
			}
		}
	}
}

// function to load object from file
func (sl *SchemaRepositorySchemaLoader) loadObjectFromFile(path string, includeRootPath string) (Object, error) {
	// Declare object
	var object Object

	// load data from file []byte
	data, loadDataError := os.ReadFile(path)

	if loadDataError != nil {
		return object, loadDataError
	}

	type Alias RepositoryObject
	o := RepositoryObject{}
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(&o),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return object, err
	}

	for key, value := range o.Attributes {
		if key == "$include" {
			includes := value.([]interface{})
			for _, include := range includes {

				// parse include file path
				includeFilePath := commons.CleanPath(includeRootPath + "/" + include.(string))

				// include file path
				attributes, includeError := sl.includeFile(includeFilePath)

				if includeError != nil {
					return object, includeError
				}

				if len(attributes) != 0 {
					// append attributes to event
					for key, value := range attributes {
						o.Attributes[key] = value
					}
				}

			}
			delete(o.Attributes, "$include")
		} else {
			attr := Attribute{}
			jsonAttr, _ := json.Marshal(value)
			json.Unmarshal(jsonAttr, &attr)
			o.Attributes[key] = attr
		}
	}

	// copy each field from RepositoryObject to Object
	object.Caption = o.Caption
	object.Constraints = o.Constraints
	object.Description = o.Description
	object.Extends = o.Extends
	object.Name = o.Name

	// iterate over attributes and copy each attribute to object
	object.Attributes = make(map[string]Attribute)
	for key, value := range o.Attributes {
		object.Attributes[key] = value.(Attribute)
	}

	return object, nil
}

// function to load event from file
func (sl *SchemaRepositorySchemaLoader) loadEventFromFile(path string, includeRootPath string) (Event, error) {
	// Declare event
	var event Event

	// load data from file []byte
	data, loadDataError := os.ReadFile(path)

	if loadDataError != nil {
		return event, loadDataError
	}

	type Alias RepositoryEvent
	e := RepositoryEvent{}
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(&e),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return event, err
	}

	for key, value := range e.Attributes {
		if key == "$include" {
			includes := value.([]interface{})
			for _, include := range includes {

				// parse include file path
				includeFilePath := commons.CleanPath(includeRootPath + "/" + include.(string))

				// include file path
				attributes, includeError := sl.includeFile(includeFilePath)

				if includeError != nil {
					return event, includeError
				}

				if len(attributes) != 0 {
					// append attributes to event
					for key, value := range attributes {
						e.Attributes[key] = value
					}
				}

			}
			delete(e.Attributes, "$include")
		} else {
			attr := Attribute{}
			jsonAttr, _ := json.Marshal(value)
			json.Unmarshal(jsonAttr, &attr)
			e.Attributes[key] = attr
		}
	}

	// copy each field from RepositoryEvent to Event
	event.Caption = e.Caption
	event.Description = e.Description
	event.Extends = e.Extends
	event.Name = e.Name
	event.Category = e.Category
	event.CategoryName = e.CategoryName
	event.Profiles = e.Profiles
	event.Uid = e.Uid

	// iterate over attributes and copy each attribute to object
	event.Attributes = make(map[string]Attribute)
	for key, value := range e.Attributes {
		event.Attributes[key] = value.(Attribute)
	}

	return event, nil

}

// func detectAndLoadInclude which accepts include file path and loads it using diffrent function based on path. Possible paths are /includes, /profiles
func (sl *SchemaRepositorySchemaLoader) includeFile(path string) (map[string]Attribute, error) {

	// switch based on path prefix
	switch {
	case strings.Contains(path, "includes/"):
		data, err := sl.loadIncludeFromFile(path)
		return data.Attributes, err
	case strings.Contains(path, "profiles/"):
		data, err := sl.loadProfileFromFile(path)
		return data.Attributes, err
	default:
		return nil, errors.New("Invalid include path " + path)
	}

}

// function to load include from file returns Include struct and error. It accepts include file path as string and loads json file from repo Directory and returns Include struct and error
func (sl *SchemaRepositorySchemaLoader) loadIncludeFromFile(path string) (Include, error) {

	// declare include
	var include Include

	// load data from file []byte
	data, loadDataError := os.ReadFile(path)

	if loadDataError != nil {
		return include, loadDataError
	}

	// unmarshal data into include
	if err := json.Unmarshal(data, &include); err != nil {
		return include, err
	}

	return include, nil
}

// function to load profile from file returns Profile struct and error. It accepts profile file path as string and loads json file from repo Directory and returns Profile struct and error
func (sl *SchemaRepositorySchemaLoader) loadProfileFromFile(path string) (Profile, error) {

	println("Loading profile from " + path)

	extensionFile := commons.CleanPath(commons.Dir(commons.Dir(path)) + "/" + "extension.json")
	extension, extensionLoadingError := sl.loadExtension(extensionFile)

	// declare profile
	var profile Profile

	// load data from file []byte
	data, loadDataError := os.ReadFile(path)

	if loadDataError != nil {
		return profile, loadDataError
	}

	// unmarshal data into profile
	profileUnmarshalError := json.Unmarshal(data, &profile)

	if profileUnmarshalError != nil {
		return profile, profileUnmarshalError
	}

	profileName := profile.Name

	// check if profile belongs to extension
	if extensionLoadingError == nil {
		profileName = extension.Name + "/" + profile.Name
	}

	if !sl.ProfileExists(profileName) {
		println("Ignoring profile " + profileName + " as it is not configured to be loaded.")
		return Profile{}, nil
	}

	// iterate over attributes and add attribute.profile = profileName
	for key, attribute := range profile.Attributes {
		attribute.Profile = profileName
		profile.Attributes[key] = attribute
	}

	return profile, nil
}
