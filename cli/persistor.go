package cli

import "code.cloudfoundry.org/cli/cf/configuration"

// noopPersistor
type noopPersistor struct {
}

// Delete -
func (p *noopPersistor) Delete() {
}

// Exists -
func (p *noopPersistor) Exists() bool {
	return true
}

// Load -
func (p *noopPersistor) Load(configuration.DataInterface) error {
	return nil
}

// Save -
func (p *noopPersistor) Save(configuration.DataInterface) error {
	return nil
}
