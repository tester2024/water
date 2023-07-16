//go:build !linux && !darwin && !windows
// +build !linux,!darwin,!windows

package water

// PlatformSpecificParams
type PlatformSpecificParams struct {
}

func defaultPlatformSpecificParams() PlatformSpecificParams {
	return PlatformSpecificParams{}
}
