// Code generated by pkg/commands/internal/migrate/cloner/cloner.go. DO NOT EDIT.

package versiontwo

type Output struct {
	Formats    Formats  `yaml:"formats,omitempty" toml:"formats,multiline,omitempty"`
	SortOrder  []string `yaml:"sort-order,omitempty" toml:"sort-order,multiline,omitempty"`
	ShowStats  *bool    `yaml:"show-stats,omitempty" toml:"show-stats,multiline,omitempty"`
	PathPrefix *string  `yaml:"path-prefix,omitempty" toml:"path-prefix,multiline,omitempty"`
	PathMode   *string  `yaml:"path-mode,omitempty" toml:"path-mode,multiline,omitempty"`
}
