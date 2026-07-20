//go:build !js
// +build !js

package model

import (
	"io"
)

type configuration struct {
	CreationDate                    string      `yaml:"created"`
	Version                         string      `yaml:"version"`
	CheckFileNameExt                bool        `yaml:"checkFileNameExt"`
	Reader15                        bool        `yaml:"reader15"`
	DecodeAllStreams                bool        `yaml:"decodeAllStreams"`
	ValidationMode                  string      `yaml:"validationMode"`
	PostProcessValidate             bool        `yaml:"postProcessValidate"`
	Eol                             string      `yaml:"eol"`
	WriteObjectStream               bool        `yaml:"writeObjectStream"`
	WriteXRefStream                 bool        `yaml:"writeXRefStream"`
	EncryptUsingAES                 bool        `yaml:"encryptUsingAES"`
	EncryptKeyLength                int         `yaml:"encryptKeyLength"`
	Permissions                     int         `yaml:"permissions"`
	Unit                            string      `yaml:"unit"`
	TimestampFormat                 string      `yaml:"timestampFormat"`
	DateFormat                      string      `yaml:"dateFormat"`
	Optimize                        bool        `yaml:"optimize"`
	OptimizeBeforeWriting           bool        `yaml:"optimizeBeforeWriting"`
	OptimizeResourceDicts           bool        `yaml:"optimizeResourceDicts"`
	OptimizeDuplicateContentStreams bool        `yaml:"optimizeDuplicateContentStreams"`
	CreateBookmarks                 bool        `yaml:"createBookmarks"`
	NeedAppearances                 bool        `yaml:"needAppearances"`
	Offline                         bool        `yaml:"offline"`
	Timeout                         int         `yaml:"timeout"`
	TimeoutCRL                      int         `yaml:"timeoutCRL"`
	TimeoutOCSP                     int         `yaml:"timeoutOCSP"`
	PreferredCertRevocationChecker  string      `yaml:"preferredCertRevocationChecker"`
	FormFieldListMaxColWidth        int         `yaml:"formFieldListMaxColWidth"`
	MaxStreamBytes                  *int64Value `yaml:"maxStreamBytes"`
	MaxDecodeBytes                  *int64Value `yaml:"maxDecodeBytes"`
	MaxImagePixels                  *int64Value `yaml:"maxImagePixels"`
	MaxImageBytes                   *int64Value `yaml:"maxImageBytes"`
}

type int64Value int64

func (i *int64Value) UnmarshalYAML(unmarshal func(interface{}) error) error {
	_ = "STUB: not implemented"
	return nil
}

func parseReadableInt64(s string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func loadValidationMode(c configuration, conf *Configuration) { _ = "STUB: not implemented"; return }

func loadedConfig(c configuration, configPath string) *Configuration {
	_ = "STUB: not implemented"
	return nil
}

func parseConfigFile(r io.Reader, configPath string) error { _ = "STUB: not implemented"; return nil }
