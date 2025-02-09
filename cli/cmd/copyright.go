// Copyright 2025 The DocWiz Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
package cmd

import (
	"bytes"
	"docwiz/internal/log"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
)

// copyrightCmdParameter stores parameters for the "copyright" command.
type copyrightCmdParameter struct {
	// pattern defines the file search pattern to locate target files.
	// Example patterns: "*.go", "src/*.js", "**/*.md".
	pattern string

	// tailInsert determines whether to insert the copyright notice at the end of the file.
	// If false, the notice is inserted at the beginning of the file.
	tailInsert bool

	// file specifies the path to a file containing the copyright notice.
	// If provided, its content will be used instead of the --content flag.
	file string

	// content holds the copyright notice text to be inserted.
	// If both --file and --content are provided, --file takes precedence.
	content string

	// repeat controls whether the copyright notice can be inserted multiple times.
	// If false, it ensures that the notice is added only once per file.
	repeat bool
}

var (
	copyrightParameter copyrightCmdParameter
	copyrightCmd       = &cobra.Command{
		Use:   "copyright",
		Short: "Insert copyright notices into files matching a pattern.",
		Long: `The 'copyright' command searches for files matching a specified pattern 
and inserts a copyright notice at the beginning or end of each file. 
It supports inserting content from a file or directly from a string.`,
		Example: `  docwiz copyright -p "*.go" -c "Copyright 2025 The DocWiz Authors."
  docwiz copyright -p "src/*.js" -f LICENSE_HEADER.txt
  docwiz copyright -p "*.md" --tail --repeat -c "© 2025 Open Source Project"`,
		Run: func(cmd *cobra.Command, args []string) {
			files, err := filepath.Glob(copyrightParameter.pattern)
			if err != nil {
				log.Fata(err)
			}

			if len(copyrightParameter.file) != 0 {
				data, err := os.ReadFile(copyrightParameter.file)
				if err != nil {
					log.Fata(err)
				}
				copyrightParameter.content = string(data)
			}

			if !copyrightParameter.repeat {
				for _, file := range files {
					if copyrightParameter.tailInsert {
						uniqueAppendToFile(file, []byte(copyrightParameter.content))
					} else {
						uniquePrependToFile(file, []byte(copyrightParameter.content))
					}
				}
			} else {
				for _, file := range files {
					if copyrightParameter.tailInsert {
						appendToTail(file, copyrightParameter.content)
					} else {
						insertAtHead(file, copyrightParameter.content)
					}
				}
			}

		},
	}
)

func init() {
	docwizCmd.AddCommand(copyrightCmd)
	copyrightCmd.PersistentFlags().StringVarP(&copyrightParameter.file, "file", "f", "", "Path to a file containing the copyright notice")
	copyrightCmd.PersistentFlags().StringVarP(&copyrightParameter.content, "content", "c", "", "Direct copyright content to be inserted")
	copyrightCmd.PersistentFlags().StringVarP(&copyrightParameter.pattern, "pattern", "p", "", "File search pattern (e.g., '*.go', 'src/*.js', '**/*.go')")
	copyrightCmd.PersistentFlags().BoolVarP(&copyrightParameter.tailInsert, "tail", "t", false, "Insert the copyright notice at the end of the file")
	copyrightCmd.PersistentFlags().BoolVarP(&copyrightParameter.repeat, "repeat", "r", false, "Allow repeated insertion of the copyright notice")
}

// uniquePrependToFile inserts data at the beginning of a file if it doesn't already start with it.
func uniquePrependToFile(filePath string, prefix []byte) error {
	prefix = bytes.ReplaceAll(prefix, []byte(`\n`), []byte("\n"))

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the beginning part of the file to check if it already starts with the prefix
	buffer := make([]byte, len(prefix))
	_, err = file.Read(buffer)
	if err != nil && err != io.EOF {
		return err
	}

	// If the file already starts with the prefix, return
	if bytes.HasPrefix(buffer, prefix) {
		log.Warnf("%s %s File already starts with the given prefix, no need to insert.\n", filePath)
		return nil
	}

	// Read the rest of the file content
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Create new content by prepending the prefix
	newContent := append(prefix, fileContent...)

	// Write the modified content back to the file
	err = os.WriteFile(filePath, newContent, 0644)
	if err != nil {
		return err
	}

	log.Infof("%s %s Prefix inserted successfully.\n", filePath)
	return nil
}

// uniqueAppendToFile appends data to the end of a file if it doesn't already end with it.
func uniqueAppendToFile(filePath string, suffix []byte) error {
	suffix = bytes.ReplaceAll(suffix, []byte(`\n`), []byte("\n"))

	// Read the entire file
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Check if the file already ends with the suffix
	if bytes.HasSuffix(data, suffix) {
		log.Warnf("%s %s The file already ends with the specified suffix, skipping append.\n", filePath)
		return nil
	}

	// Open the file in append mode
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Append the suffix
	_, err = file.Write(suffix)
	if err != nil {
		return err
	}

	log.Infof("%s %s Data successfully appended to the file.\n", filePath)
	return nil
}

// insertAtHead inserts the specified content at the beginning of the file.
func insertAtHead(filePath, content string) error {
	content, err := strconv.Unquote(`"` + content + `"`)
	if err != nil {
		return fmt.Errorf("error parsing content: %v", err)
	}
	// Temporary file path
	tempFilePath := "temp_" + filePath

	// Open the source file
	sourceFile, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening source file: %v", err)
	}
	defer sourceFile.Close()

	// Create a temporary file to store the new content
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		return fmt.Errorf("error creating temporary file: %v", err)
	}
	defer tempFile.Close()

	// Write the head content into the temporary file
	_, err = tempFile.WriteString(content)
	if err != nil {
		return fmt.Errorf("error writing to temp file: %v", err)
	}

	// Copy the content of the original file into the temporary file
	_, err = tempFile.ReadFrom(sourceFile)
	if err != nil {
		return fmt.Errorf("error reading from source file: %v", err)
	}

	// Close both files
	tempFile.Close()
	sourceFile.Close()

	// Remove the original file
	err = os.Remove(filePath)
	if err != nil {
		return fmt.Errorf("error removing original file: %v", err)
	}

	// Rename the temporary file to the original file name
	err = os.Rename(tempFilePath, filePath)
	if err != nil {
		return fmt.Errorf("error renaming temp file: %v", err)
	}

	return nil
}

// appendToTail appends the specified content at the end of the file.
func appendToTail(filePath, content string) error {
	content, err := strconv.Unquote(`"` + content + `"`)
	if err != nil {
		return fmt.Errorf("error parsing content: %v", err)
	}
	// Open the file in append mode
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file for append: %v", err)
	}
	defer file.Close()

	// Append the content at the end of the file
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return nil
}
