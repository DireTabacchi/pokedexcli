package main

import "os"

func commandExit(ds *dexState, args ...string) error {
    os.Exit(0)
    return nil
}
