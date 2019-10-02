package main

import "flag"

// ServerPort config
var ServerPort = ""

// GoAsistenCoreDir dir core
var GoAsistenCoreDir = ""

// GoAsistenPreDataDir dir
var GoAsistenPreDataDir = ""

// GoAsistenDataDir dir
var GoAsistenDataDir = ""

// GoAsistenGUIDir dir
var GoAsistenGUIDir = ""

func init() {
	flag.StringVar(&ServerPort, "server-port", "1325", "Server port")
	flag.StringVar(&GoAsistenCoreDir, "go-asisten-core", "/opt/go-asisten/go-asisten-core/go-asisten-core.so", "Core file")
	flag.StringVar(&GoAsistenPreDataDir, "go-predata-dir", "/opt/go-asisten/go-asisten-core/data/", "Pre Data Dir")
	flag.StringVar(&GoAsistenDataDir, "go-data-dir", "$HOME/.config/go-asisten/", "Data Dir")
	flag.StringVar(&GoAsistenGUIDir, "go-gui-dir", "/opt/go-asisten/go-asisten-gui/", "Gui Dir")

	flag.Parse()
}
