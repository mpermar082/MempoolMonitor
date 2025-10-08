// cmd/mempoolmonitor/main.go
package main

import (
    "flag"
    "log"
    "os"

    "mempoolmonitor/internal/mempoolmonitor"
)

func main() {
    // Define command line flags
    var (
        verbose = flag.Bool("verbose", false, "Enable verbose logging")
        input   = flag.String("input", "", "Input file path")
        output  = flag.String("output", "", "Output file path")
    )

    // Parse command line flags
    flag.Parse()

    // Create a new app instance with verbose logging
    app := mempoolmonitor.NewApp(*verbose)
    if err := app.Run(*input, *output); err != nil {
        // Log and exit on error
        log.Printf("Error: %v", err)
        os.Exit(1)
    }
}