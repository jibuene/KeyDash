#!/bin/bash

GO_BINARY="keydash"

# Destination directory in the user's PATH (adjust if needed)
DEST_DIR="/usr/local/bin"

# Check if the destination directory is writable
if [ ! -w "$DEST_DIR" ]; then
    echo "Error: You do not have write permission to $DEST_DIR."
    exit 3
fi

# Check if the binary is present in destination directory
if [ -f "$DEST_DIR/$GO_BINARY" ]; then
    rm -f "$DEST_DIR/$GO_BINARY"
    echo "Existing $GO_BINARY binary has been removed."
fi

# Move the binary to the destination directory
sudo mv "$GO_BINARY" "$DEST_DIR"

# Ensure it's executable
# sudo chmod +x "$DEST_DIR/$(basename "$GO_BINARY")"

echo "Go binary has been successfully moved to $DEST_DIR and is now executable."

