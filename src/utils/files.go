package utils

import ("os"
        "strings"
)

// Reads file content.  
// Removes \n or \r\n at the end of file
func ReadKeyFile(fname string) ([]byte, error) {
        key, err := os.ReadFile(fname)

        if err != nil {
            return nil, err
        }

        key = []byte(strings.TrimSuffix(string(key), "\n"))
        key = []byte(strings.TrimSuffix(string(key), "\r"))

        return key, nil
}


