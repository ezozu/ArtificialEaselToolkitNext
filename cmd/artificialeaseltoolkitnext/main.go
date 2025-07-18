// cmd/artificialeaseltoolkitnext/main.go
package main

import (
"flag"
"log"
"os"

"artificialeaseltoolkitnext/internal/artificialeaseltoolkitnext"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialeaseltoolkitnext.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
