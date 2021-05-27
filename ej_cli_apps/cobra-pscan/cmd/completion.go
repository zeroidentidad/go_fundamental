package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
)

// completionCmd representa el comando de completado
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generar completado de bash para comandos",
	Long: `Para cargar completado ejecutar
	
	source <(pScan completion)

	Para cargar completado automáticamente al abrir terminal, agregar esta línea a su archivo .bashrc
	
	$ ~/.bashrc
	
	source <(pScan completion)
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return completionAction(os.Stdout)
	},
}

func completionAction(out io.Writer) error {
	return rootCmd.GenBashCompletion(out)
}

func init() {
	rootCmd.AddCommand(completionCmd)

	// Aquí definir flags y ajustes de configuración.

	// Cobra admite flags persistentes que funcionarán para
	// el comando y todos los subcomandos, por ejemplo:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra admite flags locales que solo se ejecutarán cuando
	// se llame directamente al comando, por ejemplo:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
