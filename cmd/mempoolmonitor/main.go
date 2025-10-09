// cmd/mempoolmonitor/main.go
package main

import (
    "flag"
    "log"
    "os"

    "mempoolmonitor/internal/mempoolmonitor"
)

func main() {
    // Define and parse command line flags
    var (
        verbose = flag.Bool("verbose", false, "Enable verbose logging")
        input   = flag.String("input", "", "Input file path")
        output  = flag.String("output", "", "Output file path")
    )
    flag.Parse()

    // Validate input and output flags
    if *input == "" || *output == "" {
        log.Fatal("Input and output file paths are required")
    }

    // Create a new app instance with verbose logging
    app := mempoolmonitor.NewApp(*verbose)
    if err := app.Run(*input, *output); err != nil {
        // Log and exit on error
        log.Printf("Error: %v", err)
        os.Exit(1)
    }
}