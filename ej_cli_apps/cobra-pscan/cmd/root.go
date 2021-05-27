package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pScan",
	Short: "Escáner de puerto TCP",
	Long: `pScan - abreviatura de Port Scanner: ejecuta un escaneo de puertos TCP en una lista de hosts.
	
	pScan permite agregar, enumerar y eliminar hosts de la lista
	
	pScan ejecuta un escaneo de puertos en puertos TCP especificados. 
	Puede personalizar los puertos de destino utilizando un flag de línea de comando.`,
	Version: "0.1",
	// Eliminar el comentario de la siguiente línea si
	// la aplicación tiene una acción asociada:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute agrega todos los comandos secundarios al
// comando raíz y establece las flags de manera apropiada.
// Esto es llamado por main.main(). Solo necesario que suceda una vez en rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Aquí definir flags y ajustes de configuración.
	// Cobra admite flags persistentes que, si se definen aquí, serán globales en la aplicación.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
		"config file (default $HOME/.pScan.yaml)")

	rootCmd.PersistentFlags().StringP("hosts-file", "f", "pScan.hosts",
		"pScan hosts file")

	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("PSCAN")

	viper.BindPFlag("hosts-file", rootCmd.PersistentFlags().Lookup("hosts-file"))

	versionTemplate := `{{printf "%s: %s - version %s\n" .Name .Short .Version}}`
	rootCmd.SetVersionTemplate(versionTemplate)
}

// initConfig lee en archivo de configuración y variables ENV si están configuradas.
func initConfig() {
	if cfgFile != "" {
		// Usar archivo de configuración del flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Encontrar directorio de inicio.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Buscar configuración en directorio de inicio con el nombre ".pScan" (sin extensión).
		viper.AddConfigPath(home)
		viper.SetConfigName(".pScan")
	}

	viper.AutomaticEnv() // leer en las variables de entorno que coinciden

	// Si se encuentra un archivo de configuración, leerlo
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
