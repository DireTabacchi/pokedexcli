package main

import "os"

func commandExit(ds *dexState) error {
    os.Exit(0)
    return nil
}
