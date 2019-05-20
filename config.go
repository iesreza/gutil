package main

// ConfigVersion returns the version, this should be incremented every time the config changes
var ConfigVersion = "1.0.0"

type Config struct {
	Version          string
	IntegerValue     int
	StringValue      string
	StringArrayValue []string
}
