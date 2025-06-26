// internal/env/env_test.go
package env

import (
    "os"
    "testing"
)

func TestBGPDisabled(t *testing.T) {
    tests := []struct {
        name     string
        envValue string
        want     bool
    }{
        {"NotSet", "", false},
        {"SetToFalse", "false", false},
        {"SetToTrue", "true", true},
        {"SetToInvalid", "foo", false},
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            os.Setenv("METALLB_DISABLE_BGP", tt.envValue)
            if got := BGPDisabled(); got != tt.want {
                t.Errorf("BGPDisabled() = %v, want %v", got, tt.want)
            }
            os.Unsetenv("METALLB_DISABLE_BGP")
        })
    }
}