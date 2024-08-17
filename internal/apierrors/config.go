package apierrors

import (
	"errors"
	"fmt"
)

var ErrInvalidConfig error = errors.New("invalid or missing configuration")

func NewInvalidConfigError(configName string) error {
	return fmt.Errorf("%s: %w", configName, ErrInvalidConfig)
}
