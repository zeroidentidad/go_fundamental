package main

import (
	"log"

	server "./gin-server"
	gui "./gui-lorca"
	core "./mod-asisten-core"
	extra "./mod-asisten-core/extras"
)

func main() {
	// Parse dirs
	GoAsistenDataDir = extra.ParseDirsEnv(GoAsistenDataDir)
	GoAsistenPreDataDir = extra.ParseDirsEnv(GoAsistenPreDataDir)
	// PrepareDataCore
	err := extra.DataPrepare(GoAsistenDataDir, GoAsistenPreDataDir)
	if err != nil {
		log.Panic(err)
	}
	// Create new Core Client
	coreClient, err := core.NewCoreClient(
		GoAsistenCoreDir,
		GoAsistenDataDir,
	)
	if err != nil {
		log.Panic(err)
	}

	// Create Server Http For View
	server.InitServer(ServerPort, coreClient.GetDB(), GoAsistenGUIDir)

	// Open gui HTTP
	ui, err := gui.NewLorcaHTTP("http://localhost:"+ServerPort+"/asisten", 800, 600)
	defer ui.Close()
	// Set Methods Speech
	ui.Bind("speech", speech(coreClient))
	ui.Bind("StopGUI", func() {
		ui.Close()
	})
	if err != nil {
		log.Panic(err)
	}

	// Waiting finish app
	<-ui.Done()
}

func speech(coreClient *core.ModelCore) func(string) []string {
	return func(words string) []string {
		// client Words convert in Asisten response
		responseAsisten, funcAsisten := coreClient.WordsToAsistenResponse(words)

		// Asisten Function run
		funcAsisten()

		// Convert AsisteWords in Speech Base64 for JavaScript
		bs64Speech := coreClient.WordsToSpeechBase64(responseAsisten)

		// Return Response
		return []string{responseAsisten, bs64Speech}
	}
}
