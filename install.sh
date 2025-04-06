#!/bin/bash

set -e

INSTALL_DIR="/usr/local/bin"
BINARY_URL="https://github.com/hoseinmontazer/stickynotes/releases/download/v1.0.0/stickynotes-linux-amd64.tar.gz"
TMP_DIR=$(mktemp -d)
ARCHIVE="$TMP_DIR/stickynotes.tar.gz"

echo "📥 Downloading StickyNotes..."
curl -L "$BINARY_URL" -o "$ARCHIVE"

echo "📦 Extracting..."
tar -xzf "$ARCHIVE" -C "$TMP_DIR"

echo "🚀 Installing to $INSTALL_DIR..."
sudo mv "$TMP_DIR/stickynotes" "$INSTALL_DIR"

echo "🧹 Cleaning up..."
rm -rf "$TMP_DIR"

echo "✅ Installation complete!"
echo "Run with: stickynotes"
