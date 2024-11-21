package main

import (
    "fmt"
    "testing"
)

type Config struct {
    Env   string
    Proxy string
}

func (conf Config) ConfInfo() string {
    fmt.Println("Env: ", conf.Env)
    fmt.Println("Proxy: ", conf.Proxy)
    return "----------------------"
}

func TestConfInfo_f9d1a7b626(t *testing.T) {
    // Test case 1: Valid config
    conf := Config{
        Env:   "production",
        Proxy: "http://proxy.example.com",
    }
    expected := "----------------------"
    actual := conf.ConfInfo()
    if actual != expected {
        t.Errorf("Expected %s, got %s", expected, actual)
    }

    // Test case 2: Empty config
    conf = Config{}
    expected = "----------------------"
    actual = conf.ConfInfo()
    if actual != expected {
        t.Errorf("Expected %s, got %s", expected, actual)
    }
}
