package cmd

import (
	"fmt"

	"github.com/bzcorn/security-check-repo/checker"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "security-check-repo",
	Short: "security-check-repo checks if a project has had a commit in the last year",
	Long: `A Fast and Flexible checker for security teams to quickly determine if 
	a project has had a commit in the last year. This can be useful to detect an abandoned
	project that is no longer maintaining itself for security vulnerabilities.`,
	Run: func(cmd *cobra.Command, args []string) {
		repos := viper.GetStringSlice("repos")
		repos = append(repos, args...)
		if len(repos) < 1 {
			fmt.Println("Please provide at least one repository.")
			return
		}

		token := viper.GetString("token")
		if token == "" {
			fmt.Println("Please provide a GitHub Personal Access Token.")
			return
		}

		checker.CheckRepos(repos, token)
	},
}

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // path to look for the config file in
	viper.AutomaticEnv()          // read in environment variables that match

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
