package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "alchemist",
	Short: "Alchemist is a base's project generator for React-apps, nest-servers and Go-servers",
	Long: `
	
   _____   .__           .__                      .__           __   
  /  _  \  |  |    ____  |  |__    ____    _____  |__|  _______/  |_ 
 /  /_\  \ |  |  _/ ___\ |  |  \ _/ __ \  /     \ |  | /  ___/\   __\
/    |    \|  |__\  \___ |   Y  \\  ___/ |  Y Y  \|  | \___ \  |  |  
\____|__  /|____/ \___  >|___|  / \___  >|__|_|  /|__|/____  > |__|  
        \/            \/      \/      \/       \/          \/
   `,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.alchemist.yaml)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".alchemist" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".alchemist")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
