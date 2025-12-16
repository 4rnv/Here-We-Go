package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var fileExtensions = map[string]string{
	".html": "HTML", ".htm": "HTML",
	".css": "CSS",
	".js":  "JavaScript", ".jsx": "JavaScript",
	".ts": "TypeScript", ".tsx": "TypeScript",
	".xml":  "XML",
	".yaml": "YAML", ".yml": "YAML",
	// ".md":    "Markdown",
	// ".json":  "JSON",
	".go":    "Go",
	".py":    "Python",
	".ipynb": "Jupyter Notebook",
	".java":  "Java",
	".c":     "C",
	".h":     "C Header",
	".cpp":   "C++", ".cxx": "C++", ".cc": "C++",
	".hpp":   "C++ Header",
	".cs":    "C#",
	".sh":    "Shell Script",
	".txt":   "Text",
	".php":   "PHP",
	".swift": "Swift",
	".rb":    "Ruby",
	".kt":    "Kotlin",
	".pl":    "Perl",
	".r":     "R",
	".sql":   "SQL",
	".ino":   "Arduino",
}

var ignoredDirectories = map[string]bool{
	"venv":              true,
	".venv":             true,
	"env":               true,
	".env":              true,
	"__pycache__":       true,
	"pyproject.toml":    true,
	"node_modules":      true,
	".pnp":              true,
	".yarn":             true,
	".git":              true,
	"playwright-report": true,
}

func getSize(filepath string) (int64, error) {
	fi, err := os.Stat(filepath)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	size := fi.Size()
	return size, nil
}

func traverseDirectory(directory string, langSizes map[string]int64, totalSizeDirectory *int64) {
	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fullPath := filepath.Join(directory, file.Name())
		if file.IsDir() && !ignoredDirectories[file.Name()] {
			traverseDirectory(fullPath, langSizes, totalSizeDirectory)
		} else {
			filesize, err := getSize(filepath.Join(directory, file.Name()))
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}
			ext := strings.ToLower(filepath.Ext(fullPath))
			langName, found := fileExtensions[ext]
			if !found {
				langName = "Other"
				continue // Skipping files with unknown extensions (for now)
			}
			langSizes[langName] += filesize
			*totalSizeDirectory += filesize
		}
	}
}

func main() {
	folder := flag.String("folder", "", "Please pass a valid path to a directory/folder.")
	flag.Parse()
	directory := filepath.Clean(*folder)
	langSizes := make(map[string]int64)
	var totalSizeDirectory int64 = 0
	fmt.Println("~Starting analysis in: ", directory)
	traverseDirectory(directory, langSizes, &totalSizeDirectory)
	languages := make([]string, 0, len(langSizes))

	for lang := range langSizes {
		languages = append(languages, lang)
	}
	sort.Strings(languages)
	fmt.Printf("%-18s %12s %16s\n", "Language", "Size (Bytes)", "Percentage")
	fmt.Println(strings.Repeat("-", 50))

	for _, language := range languages {
		size := langSizes[language]
		percentage := (float64(size) / float64(totalSizeDirectory)) * 100.0
		fmt.Printf("%-18s %12d %16.2f%%\n", language, size, percentage)
	}

	fmt.Println(strings.Repeat("-", 50))
	fmt.Printf("%-18s %12d %16.2f%%\n", "Total Identified", totalSizeDirectory, 100.0)
}
