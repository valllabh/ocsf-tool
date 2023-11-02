# Set -e to exit on error
set -e

echo
echo "📦 OCSF-Tool Downloading"
echo

# Set Constants
FILE_BASENAME="ocsf-tool"
RELEASES_URL="https://github.com/valllabh/ocsf-tool/releases"


# Detect OS and Architecture
OS="$(uname -s)"
ARCH="$(uname -m)"
test "$ARCH" = "aarch64" && ARCH="arm64"

echo "✅ OS and Architecture detected"


# Tar File Name Pattern
TAR_FILE="${FILE_BASENAME}_${OS}_${ARCH}.tar.gz"

# Get the latest release URL

LATEST_VERSION_URL=$(curl -sSLI -o /dev/null -w %{url_effective} "$RELEASES_URL/latest")

# extract the latest release version from the URL
VERSION=$(echo $LATEST_VERSION_URL | grep -oP '(?<=/tag/).*$')
echo "✅ Detected latest version of OCSF-Tool"



# Download the tar file and checksums
curl -sfLO "$RELEASES_URL/download/$VERSION/$TAR_FILE"
curl -sfLO "$RELEASES_URL/download/$VERSION/checksums.txt"

echo "✅ Downloaded OCSF-Tool $VERSION (latest)"


# Exit if Tar file is not donwloaded
if [ ! -f "$TAR_FILE" ]; then
    echo "❌ Could not download files."
    echo "Please use Release URL to download the files manually."
    echo $RELEASES_URL 
    exit 1
fi

# Verify Checksums
sha256sum --ignore-missing --quiet --check checksums.txt
echo "✅ Verified downloaded files"

# Installation

# Create ocsf-tool directory if does not exists
mkdir -p "./ocsf-tool"

# Clear ocsf-tool directory
rm -rf "./ocsf-tool/*"

tar -xf "$TAR_FILE" -C "./ocsf-tool"
echo "✅ Extracted the tar file in ./ocsf-tool directory"



# remove tar file and checksums
rm -rf "$TAR_FILE" "checksums.txt"
echo "✅ Deleted tar file and checksums"


echo "✅ Dwonload Complete!"
echo


./ocsf-tool/ocsf-tool --help


# #!/bin/bash
# set -e


# if test "$DISTRIBUTION" = "pro"; then
# 	echo "Using Pro distribution..."
# 	RELEASES_URL="https://github.com/goreleaser/goreleaser-pro/releases"
# 	FILE_BASENAME="goreleaser-pro"
# 	LATEST="$(curl -sf https://goreleaser.com/static/latest-pro)"
# else
# 	echo "Using the OSS distribution..."
# 	RELEASES_URL="https://github.com/goreleaser/goreleaser/releases"
# 	FILE_BASENAME="goreleaser"
# 	LATEST="$(curl -sf https://goreleaser.com/static/latest)"
# fi

# test -z "$VERSION" && VERSION="$LATEST"

# test -z "$VERSION" && {
# 	echo "Unable to get goreleaser version." >&2
# 	exit 1
# }

# if test "$DISTRIBUTION" = "pro" && [[ "$VERSION" != *-pro ]]; then
# 	VERSION="$VERSION-pro"
# fi

# TMP_DIR="$(mktemp -d)"
# # shellcheck disable=SC2064 # intentionally expands here
# trap "rm -rf \"$TMP_DIR\"" EXIT INT TERM

# OS="$(uname -s)"
# ARCH="$(uname -m)"
# test "$ARCH" = "aarch64" && ARCH="arm64"
# TAR_FILE="${FILE_BASENAME}_${OS}_${ARCH}.tar.gz"

# (
# 	cd "$TMP_DIR"
# 	echo "Downloading GoReleaser $VERSION..."
# 	curl -sfLO "$RELEASES_URL/download/$VERSION/$TAR_FILE"
# 	curl -sfLO "$RELEASES_URL/download/$VERSION/checksums.txt"
# 	echo "Verifying checksums..."
# 	sha256sum --ignore-missing --quiet --check checksums.txt
# 	if command -v cosign >/dev/null 2>&1; then
# 		echo "Verifying signatures..."
# 		cosign verify-blob \
# 			--certificate-identity-regexp "https://github.com/goreleaser/goreleaser.*/.github/workflows/.*.yml@refs/tags/$VERSION" \
# 			--certificate-oidc-issuer 'https://token.actions.githubusercontent.com' \
# 			--cert "$RELEASES_URL/download/$VERSION/checksums.txt.pem" \
# 			--signature "$RELEASES_URL/download/$VERSION/checksums.txt.sig" \
# 			checksums.txt
# 	else
# 		echo "Could not verify signatures, cosign is not installed."
# 	fi
# )

# tar -xf "$TMP_DIR/$TAR_FILE" -C "$TMP_DIR"
# "$TMP_DIR/goreleaser" "$@"
