package utils

import (
	"testing"

	"github.com/spf13/viper"
)

func TestCheckMutuallyExclusiveSettings(t *testing.T) {
	config := viper.New()
	got := CheckMutuallyExclusiveSettings(config)
	if got != nil {
		t.Errorf("got: %s, want: %s", got, "")
	}
}

func TestCheckMutuallyExclusiveSettingsFail(t *testing.T) {
	config := viper.New()
	config.Set("image", "test")
	config.Set("dockerfile", "test")
	got := CheckMutuallyExclusiveSettings(config)
	if got == nil {
		t.Errorf("got: %s, want: %s", got, "error")
	}
}

func TestAliasesDockerfile_buildDockerfile(t *testing.T) {
	dockerfile := "../Dockerfile"
	config := viper.New()
	config.RegisterAlias("dockerfile", "build.dockerfile")
	config.Set("dockerfile", dockerfile)
	got := config.Get("dockerfile")
	if got != dockerfile {
		t.Errorf("with param \"dockerfile\", got: %s, want: %s", got, dockerfile)
	}
	got = config.Get("build.dockerfile")
	if got != dockerfile {
		t.Errorf("with param \"build.dockerfile\", got: %s, want: %s", got, dockerfile)
	}
}
