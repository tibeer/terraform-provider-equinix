package main

import (
    "text/template"
    "os"
    "gopkg.in/yaml.v2"
    "io"
    "log"
)

func readConfig(filePath string) (*Config, error) {
    // Open the file
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    // Read the file contents
    fileContents, err := io.ReadAll(file)
    if err != nil {
        return nil, err
    }

    // Unmarshal the YAML
    var config Config
    err = yaml.Unmarshal(fileContents, &config)
    if err != nil {
    return nil, err
    }

    return &config, nil
}

func main() {
    // Check if a command line argument is provided
    if len(os.Args) < 2 {
        log.Fatal("Usage: <program> <path to config.yaml>")
    }
    configFilePath := os.Args[1]

    // Read and parse the configuration file
    config, err := readConfig(configFilePath)
    if err != nil {
        log.Fatalf("Error reading config: %v", err)
    }
    
    // Parse the template
    tmpl, err := template.New("model.go.tmpl").Funcs(template.FuncMap{
        "ToCamelCase": ToCamelCase,
        "SDKTypeName": SDKTypeName,
        "ToSnakeCase": ToSnakeCase,
        "ToPascalCase": ToPascalCase,
    }).ParseFiles("./templates/model.go.tmpl")
    if err != nil {
        log.Fatalf("Error parsing template: %v", err)
    }

    // Format directory name
    dirName := ToSnakeCase(config.Service) + "_" + ToSnakeCase(config.Resource)
    
    // Create the directory if it does not exist
    if err := os.MkdirAll(dirName, 0755); err != nil {
        log.Fatalf("Failed to create directory: %s, error: %v", dirName, err)
    }
    
    // Path for the new file
    filePath := dirName + "/framework_models.go"
    
    file, err := os.Create(filePath)
    if err != nil {
        log.Fatalf("Error creating framework_models template: %v", err)
    }
    defer file.Close()

    // Execute the template with config data
    err = tmpl.Execute(file, config)
    if err != nil {
        log.Fatalf("Error executing framework_models template: %v", err)
    }
}
