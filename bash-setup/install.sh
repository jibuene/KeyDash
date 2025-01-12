#!/bin/bash

GO_BINARY="keydash"

# Destination directory in the user's PATH (adjust if needed)
DEST_DIR="/usr/local/bin"

# Check if the destination directory is writable
if [ ! -w "$DEST_DIR" ]; then
    echo "Error: You do not have write permission to $DEST_DIR."
    exit 3
fi

# Move the binary to the destination directory
sudo mv "$GO_BINARY" "$DEST_DIR"

# Ensure it's executable
# sudo chmod +x "$DEST_DIR/$(basename "$GO_BINARY")"

echo "Go binary has been successfully moved to $DEST_DIR and is now executable."

