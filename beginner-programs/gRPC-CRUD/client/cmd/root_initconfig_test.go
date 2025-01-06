package cmd

import (
	"testing"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

type T struct {
	common
	isEnvSet bool
	context  *testContext // For running tests and subtests.
}
func TestinitConfig(t *testing.T) {

	tests := []struct {
		name     string
		cfgFile  string
		wantFile string
		wantErr  bool
	}{
		{
			name:     "Configuration File Provided",
			cfgFile:  "testconfig.toml",
			wantFile: "testconfig.toml",
			wantErr:  false,
		},
		{
			name:     "No Configuration File Provided",
			cfgFile:  "",
			wantFile: ".cobra",
			wantErr:  false,
		},
		{
			name:     "Error Reading Configuration File",
			cfgFile:  "nonexistent.toml",
			wantFile: "nonexistent.toml",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			cfgFile = tt.cfgFile

			defer func() {
				if r := recover(); r != nil {
					t.Log("Recovered in initConfig", r)
				}
			}()

			initConfig()

			if viper.ConfigFileUsed() != tt.wantFile && !tt.wantErr {
				t.Errorf("initConfig() = %v, want %v", viper.ConfigFileUsed(), tt.wantFile)
			}

			home, _ := homedir.Dir()
			if tt.cfgFile == "" && viper.ConfigFileUsed() != home+"/"+tt.wantFile {
				t.Errorf("initConfig() = %v, want %v", viper.ConfigFileUsed(), home+"/"+tt.wantFile)
			}
		})
	}
}
