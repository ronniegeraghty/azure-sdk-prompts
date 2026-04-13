package main

import (
"os"

"github.com/ronniegeraghty/hyoka/hyoka/cmd"
)

func main() {
if err := cmd.Execute(); err != nil {
os.Exit(1)
}
}
