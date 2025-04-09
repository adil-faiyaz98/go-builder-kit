package builders

import (
	"fmt"
	
	"github.com/adil-faiyaz98/go-builder-kit/models"
	
)

// DeviceBuilder builds a Device model
type DeviceBuilder struct {
	device *models.Device
	// Custom validation functions
	validationFuncs []func(*models.Device) error
}

// NewDeviceBuilder creates a new DeviceBuilder
func NewDeviceBuilder() *DeviceBuilder {
	return &DeviceBuilder{
		device: &models.Device{
			Type: "",
			Model: "",
			SerialNumber: "",
			PurchaseDate: "",
			OS: "",
			LastUsed: "",
			Apps: []*models.Application{},
			Settings: map[string]string{},
		},
		validationFuncs: []func(*models.Device) error{},
	}
}

// NewDeviceBuilderWithDefaults creates a new DeviceBuilder with sensible defaults
func NewDeviceBuilderWithDefaults() *DeviceBuilder {
	builder := NewDeviceBuilder()
	// Add default values here if needed
	return builder
}
// WithType sets the Type
func (b *DeviceBuilder) WithType(value string) *DeviceBuilder {
	b.device.Type = value
	return b
}

// WithModel sets the Model
func (b *DeviceBuilder) WithModel(model string) *DeviceBuilder {
	b.device.Model = model
	return b
}

// WithSerialNumber sets the SerialNumber
func (b *DeviceBuilder) WithSerialNumber(serialNumber string) *DeviceBuilder {
	b.device.SerialNumber = serialNumber
	return b
}

// WithPurchaseDate sets the PurchaseDate
func (b *DeviceBuilder) WithPurchaseDate(purchaseDate string) *DeviceBuilder {
	b.device.PurchaseDate = purchaseDate
	return b
}

// WithOS sets the OS
func (b *DeviceBuilder) WithOS(oS string) *DeviceBuilder {
	b.device.OS = oS
	return b
}

// WithLastUsed sets the LastUsed
func (b *DeviceBuilder) WithLastUsed(lastUsed string) *DeviceBuilder {
	b.device.LastUsed = lastUsed
	return b
}

// WithApps sets the Apps
func (b *DeviceBuilder) WithApps(apps *ApplicationBuilder) *DeviceBuilder {
	// Ensure the slice is initialized
	if b.device.Apps == nil {
		b.device.Apps = []*models.Application{}
	}
	// Handle nested slice element
	builtValue := apps.Build().(*models.Application)
	b.device.Apps = append(b.device.Apps, builtValue)
	return b
}

// WithSettings sets the Settings
func (b *DeviceBuilder) WithSettings(key string, val string) *DeviceBuilder {
	if b.device.Settings == nil {
		b.device.Settings = make(map[string]string)
	}
	b.device.Settings[key] = val
	return b
}


// WithValidation adds a custom validation function
func (b *DeviceBuilder) WithValidation(validationFunc func(*models.Device) error) *DeviceBuilder {
	b.validationFuncs = append(b.validationFuncs, validationFunc)
	return b
}

// Build builds the Device
func (b *DeviceBuilder) Build() interface{} {
	return b.device
}

// BuildPtr builds the Device and returns a pointer
func (b *DeviceBuilder) BuildPtr() *models.Device {
	return b.device
}

// BuildAndValidate builds the Device and validates it
func (b *DeviceBuilder) BuildAndValidate() (*models.Device, error) {
	device := b.device

	// Run custom validation functions
	for _, validationFunc := range b.validationFuncs {
		if err := validationFunc(device); err != nil {
			return nil, fmt.Errorf("custom validation failed: %w", err)
		}
	}

	// Run model's Validate method if it exists
	if v, ok := interface{}(device).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return device, err
		}
	}

	return device, nil
}

// MustBuild builds the Device and panics if validation fails
func (b *DeviceBuilder) MustBuild() *models.Device {
	model, err := b.BuildAndValidate()
	if err != nil {
		panic(err)
	}
	return model
}

// Clone creates a deep copy of the builder
func (b *DeviceBuilder) Clone() *DeviceBuilder {
	clonedDevice := *b.device
	return &DeviceBuilder{
		device: &clonedDevice,
		validationFuncs: append([]func(*models.Device) error{}, b.validationFuncs...),
	}
}
