package main

import (
	"bytes"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"gopkg.in/yaml.v3"
)

type Settings struct {
	CalDAV CalDAVConfig     `yaml:"caldav"`
	ToDos  []CalendarConfig `yaml:"todos"`

	Templates map[string]string `yaml:"templates"`
}

type CalDAVConfig struct {
	URL string `yaml:"url"`

	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type CalendarConfig struct {
	Name  string `yaml:"name"`
	Label string `yaml:"label,omitempty"`
	Desc  string `yaml:"description,omitempty"`
}

func (c *CalendarConfig) GetLabel() string {
	if c.Label != "" {
		return c.Label
	}
	return c.Name
}

// Loads settings from default location.
func GetSettings() *Settings {
	path, err := xdg.SearchConfigFile(filepath.Join("daily", "settings.yaml"))
	if err != nil {
		log.Fatal(err)
	}

	settings, err := ReadSettings(path)
	if err != nil {
		log.Fatalf("Error while reading settings: %v", err)
	}

	return settings
}

func ReadSettings(filePath string) (*Settings, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read settings file: %w", err)
	}

	var settings Settings
	if err := yaml.Unmarshal(data, &settings); err != nil {
		return nil, fmt.Errorf("failed to parse YAML settings: %w", err)
	}

	return &settings, nil
}

func (s *Settings) WriteSettings(filePath string) error {
	buf := bytes.Buffer{}
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	enc.Encode(s)

	data := buf.Bytes()
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write settings file: %w", err)
	}

	return nil
}

func (s *Settings) GetAbsoluteURL(URI string) string {
	u, err := url.Parse(s.CalDAV.URL)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s://%s%s", u.Scheme, u.Host, URI)
}
