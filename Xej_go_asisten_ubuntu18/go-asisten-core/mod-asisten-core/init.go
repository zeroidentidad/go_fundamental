package core

import (
	"plugin"
)

// NewCoreClient new model
func NewCoreClient(CoreDir string, DataDir string) (*ModelCore, error) {
	// Open Shared Object
	pl, err := plugin.Open(CoreDir)
	if err != nil {
		return nil, err
	}

	// NewGoAsistenCore funtion opent to parse
	NewGoAsistenCore, err := pl.Lookup("NewGoAsistenCore")
	if err != nil {
		return nil, err
	}

	// AsistenCore methods
	AsistenCore, err := NewGoAsistenCore.(func(string) (interface{}, error))(DataDir)
	if err != nil {
		return nil, err
	}

	// Return data
	return &ModelCore{
		AsistenCore: AsistenCore,
		pl:          pl,
	}, nil
}
