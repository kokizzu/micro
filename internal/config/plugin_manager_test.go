package config

import "testing"

func TestNewPluginInfoAllowsJSON5Comments(t *testing.T) {
	data := []byte(`[
  // repo.json files are decoded as JSON5 by the plugin installer.
  {
    "Name": "example",
    "Description": "Example plugin",
    "Website": "https://example.com"
  }
]`)

	info, err := NewPluginInfo(data)
	if err != nil {
		t.Fatal(err)
	}

	if info.Name != "example" {
		t.Fatalf("expected plugin name %q, got %q", "example", info.Name)
	}
}

func TestNewPluginInfoRejectsMissingName(t *testing.T) {
	for _, data := range []string{`[]`, `null`, `[{}]`, `[{"Description": "no name"}]`} {
		if info, err := NewPluginInfo([]byte(data)); err != ErrMissingName {
			t.Errorf("NewPluginInfo(%s) = %v, %v; want ErrMissingName", data, info, err)
		}
	}
}
