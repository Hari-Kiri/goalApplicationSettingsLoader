package goalApplicationSettingsLoader

import (
	"encoding/json"
	"io/ioutil"
)

// Application settings
type ApplicationSettings struct {
	Settings ApplicationSettingsData
}
type ApplicationSettingsData struct {
	Name         string
	Port         int
	Organisation string
	Version      string
}

// Database settings
type ApplicationDatabaseConfiguration struct {
	DatabaseConfiguration ApplicationDatabaseConfigurationData
}
type ApplicationDatabaseConfigurationData struct {
	User           string
	Password       string
	ConnectionType string
	Hostname       string
	DatabaseName   string
}

// Static server setting
type StaticServerConfiguration struct {
	StaticServer StaticServerConfigurationData
}
type StaticServerConfigurationData struct {
	Uri string
}

// Load application settings data
func LoadSettings() (*ApplicationSettings, error) {
	// Open our jsonFile
	openSettingsFile, error := ioutil.ReadFile("settings.json")
	// If ioutil.ReadFile returns an error then handle it
	if error != nil {
		return nil, error
	}
	// JSON data
	var appData ApplicationSettings
	// we unmarshal our JSON data which contains our
	// json file's content into 'appData' which we defined above
	error = json.Unmarshal(openSettingsFile, &appData)
	// If json.Unmarshal returns an error then handle it
	if error != nil {
		return nil, error
	}
	// Return data
	return &appData, error
}

// Load database configuration data
func LoadDatabaseConfiguration() (*ApplicationDatabaseConfiguration, error) {
	// Open our jsonFile
	openSettingsFile, error := ioutil.ReadFile("settings.json")
	// If ioutil.ReadFile returns an error then handle it
	if error != nil {
		return nil, error
	}
	// JSON data
	var dbConfig ApplicationDatabaseConfiguration
	// we unmarshal our JSON data which contains our
	// json file's content into 'appData' which we defined above
	error = json.Unmarshal(openSettingsFile, &dbConfig)
	// If json.Unmarshal returns an error then handle it
	if error != nil {
		return nil, error
	}
	// Return data
	return &dbConfig, error
}

// Load static server configuration data
func LoadStaticServerConfiguration() (*StaticServerConfiguration, error) {
	// Open our jsonFile
	openSettingsFile, error := ioutil.ReadFile("settings.json")
	// If ioutil.ReadFile returns an error then handle it
	if error != nil {
		return nil, error
	}
	// JSON data
	var staticServerConfigurationData StaticServerConfiguration
	// we unmarshal our JSON data which contains our
	// json file's content into 'appData' which we defined above
	error = json.Unmarshal(openSettingsFile, &staticServerConfigurationData)
	// If json.Unmarshal returns an error then handle it
	if error != nil {
		return nil, error
	}
	// Return data
	return &staticServerConfigurationData, error
}
