#!/bin/sh

set -e

VERSION="1.0.3"
REPO="egeuysall/rest"
TMPDIR=$(mktemp -d)

cleanup() {
  rm -rf "$TMPDIR"
}
trap cleanup EXIT

# Detect OS
OS=$(uname | tr '[:upper:]' '[:lower:]')
case "$OS" in
  linux|darwin) ;;
  mingw*|msys*|cygwin*)
    OS="windows"
    ;;
  *)
    echo "Unsupported OS: $OS"
    exit 1
    ;;
esac

# Detect Architecture
ARCH=$(uname -m)
case "$ARCH" in
  x86_64|amd64) ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *)
    echo "Unsupported architecture: $ARCH"
    exit 1
    ;;
esac

FILE="rest_${VERSION}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/download/v${VERSION}/${FILE}"
CHECKSUM_URL="https://github.com/${REPO}/releases/download/v${VERSION}/rest_${VERSION}_checksums.txt"

echo "Downloading checksum file..."
curl -fsSL "$CHECKSUM_URL" -o "$TMPDIR/checksums.txt"

echo "Downloading $FILE ..."
curl -fsSL "$URL" -o "$TMPDIR/$FILE"

echo "Verifying checksum..."
grep "$FILE" "$TMPDIR/checksums.txt" | sha256sum -c -

echo "Extracting archive..."
tar -xzf "$TMPDIR/$FILE" -C "$TMPDIR"

echo "Installing binary to /usr/local/bin ..."
chmod +x "$TMPDIR/rest"
sudo mv "$TMPDIR/rest" /usr/local/bin/rest

echo "Installation complete. Version info:"
rest --version