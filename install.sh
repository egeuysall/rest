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

if command -v sha256sum >/dev/null 2>&1; then
  CHECKSUM_TOOL="sha256sum"
elif command -v shasum >/dev/null 2>&1; then
  CHECKSUM_TOOL="shasum -a 256"
else
  echo "Error: neither sha256sum nor shasum found."
  exit 1
fi

echo "Verifying checksum..."

EXPECTED_SUM=$(grep "$FILE" "$TMPDIR/checksums.txt" | awk '{print $1}')
ACTUAL_SUM=$($CHECKSUM_TOOL "$TMPDIR/$FILE" | awk '{print $1}')

if [ "$EXPECTED_SUM" != "$ACTUAL_SUM" ]; then
  echo "❌ Checksum verification failed!"
  echo "Expected: $EXPECTED_SUM"
  echo "Actual:   $ACTUAL_SUM"
  exit 1
fi

echo "✅ Checksum verified."

echo "Extracting archive..."
tar -xzf "$TMPDIR/$FILE" -C "$TMPDIR"

# Get the first directory name from the archive
EXTRACTED_DIR=$(tar -tzf "$TMPDIR/$FILE" | head -1 | cut -f1 -d"/")

# Verify the binary exists before proceeding
if [ ! -f "$TMPDIR/$EXTRACTED_DIR/rest" ]; then
  echo "❌ Error: Binary not found in extracted archive"
  echo "Contents of $TMPDIR/$EXTRACTED_DIR:"
  ls -la "$TMPDIR/$EXTRACTED_DIR"
  exit 1
fi

echo "Installing binary to /usr/local/bin ..."
chmod +x "$TMPDIR/$EXTRACTED_DIR/rest"
sudo mv "$TMPDIR/$EXTRACTED_DIR/rest" /usr/local/bin/rest

echo "✅ Installation complete. Version info:"
rest --version