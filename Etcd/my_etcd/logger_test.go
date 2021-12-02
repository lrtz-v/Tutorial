package main

import "testing"


func TestInitLogger(t *testing.T) {
	InitLogger("./log")
	Logger.Info("Info Test")
	Logger.Warn("Warn Test")
	Logger.Debug("Debug Test")
	Logger.Error("Error Test")
}
