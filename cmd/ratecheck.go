/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/spf13/cobra"
)

// ratecheckCmd represents the ratecheck command
var ratecheckCmd = &cobra.Command{
	Use:   "ratecheck -r regexp filename",
	Aliases: []string{"r"},
	Short: "Scrape a log file to calculate rate of events",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// handle command line
		regex, _ := cmd.Flags().GetString("regexp")
		if regex == "" {
			os.Exit(1)
		}
		_, err := regexp.Compile(regex)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}

		scrapeFile(args[0], regex)

		fmt.Printf("ratecheck -r %s %s\n", regex, args[0])
	},
}

func init() {
	rootCmd.AddCommand(ratecheckCmd)

	ratecheckCmd.Flags().StringP("regexp", "r", "", "Regular expression to match")
}

func scrapeFile(filename string, regex string) {
	fmt.Printf("Scraping %s with %s\n", filename, regex)

	// open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// for each line in the file, check if it matches the regex
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(regex)
	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line: %s\n", line)
		if re.MatchString(line) {
			count++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Total matches: %d\n", count)
}
