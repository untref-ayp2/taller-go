package main

import (
	"os"
	"testing"
)

func TestLeerConfig(t *testing.T) {
	contenido := []byte("puerto=8080\nhost=localhost\ndebug=true\n")
	os.WriteFile("_config.txt", contenido, 0644)
	defer os.Remove("_config.txt")

	cfg, err := LeerConfig("_config.txt")
	if err != nil {
		t.Fatal("no deberia haber error")
	}
	if cfg.Puerto != 8080 || cfg.Host != "localhost" || !cfg.Debug {
		t.Errorf("config incorrecta: %+v", cfg)
	}
}

func TestLeerConfigNoExiste(t *testing.T) {
	_, err := LeerConfig("_no_existe.txt")
	if err == nil {
		t.Error("archivo inexistente deberia dar error")
	}
}
