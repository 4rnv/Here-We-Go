A lightweight Go tool that analyzes your codebase's language composition from CLI. Designed to mimic GitHub's language percentage statistics, providing info about your project’s code composition locally.

## Usage
- Requires Go (1.16 or above)
- Clone this repo with `git clone https://github.com/4rnv/Holo.git`
- `cd Holo`
- `go run main.go --folder "path/to/directory"`
  - or `go build main.go"` followed by `./main --folder "path/to/directory"`
  
Run the tool by specifying the directory (folder) you want to analyze using the --folder flag. For example,

`./main --folder "C:/Users/XYZ/my-app"`

## Configuration
### File Extension Mapping
The file-to-language mappings are defined in the fileExtensions map within the source code. You can customize this mapping to support additional file types or adjust existing associations.
### Ignored Directories
Common package or environment directories (e.g., node_modules, .git, venv, etc.) are ignored. The list of directories to ignore is also defined in the source code via the ignoredDirectories map. Add or remove directories as needed for your projects.