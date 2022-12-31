package cmd

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "replace \"names/to/use.txt\" \"../Astrox Imperium/Astrox Imperium_Data/MOD/sectors/\"",
	Short: "Replaces sector names in Astrox Imperium sector_x.txt files.",
	Long:  `Replace is a name replacement tool for Astrox Imperium sector_x.txt files.`,
	Run: func(cmd *cobra.Command, args []string) {
		var names []string

		nameFile, err := os.Open(args[0])
		if err != nil {
			log.Fatal(err)
		}

		for scanner := bufio.NewScanner(nameFile); scanner.Scan(); {
			names = append(names, scanner.Text())
		}
		nameFile.Close()

		//shuffle names if random flag is set
		if cmd.Flag("random").Changed {
			rand.Seed(time.Hour.Milliseconds())
			rand.Shuffle(len(names), func(i, j int) {
				names[i], names[j] = names[j], names[i]
			})
		}

		sectorFiles, err := os.ReadDir(args[1])
		if err != nil {
			log.Fatal(err)
		}

		for i, f := range sectorFiles {
			sectorFile, err := os.OpenFile(filepath.Join(args[1], f.Name()), os.O_RDWR, 0755)
			if err != nil {
				log.Fatal(err)
			}

			// read in file into buffer
			scanner := bufio.NewScanner(sectorFile)
			var buffer []string
			for scanner.Scan() {
				buffer = append(buffer, scanner.Text())
			}
			// done with file, rename for backup purposes
			sectorFile.Close()
			os.Rename(filepath.Join(args[1], f.Name()), filepath.Join(args[1], "\\OLD_"+f.Name()))

			// find string to change
			for j, line := range buffer {
				if strings.HasPrefix(line, "SECTOR;name;") {
					fields := strings.Split(line, ";")
					fields[2] = names[i]
					buffer[j] = strings.Join(fields, ";")
					break
				}
			}

			// write modified file
			newSectorFile, err := os.OpenFile(filepath.Join(args[1], f.Name()), os.O_WRONLY|os.O_CREATE, 0755)
			if err != nil {
				log.Fatal(err)
			}
			defer newSectorFile.Close()

			writer := bufio.NewWriter(newSectorFile)
			for _, line := range buffer {
				_, e := writer.WriteString(line + "\n")
				if e != nil {
					log.Print(e)
				}
			}

			e := writer.Flush()
			if e != nil {
				log.Print(e)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goapp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("random", "r", false, "Shuffles names randomly if set.")
}
