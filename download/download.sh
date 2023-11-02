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
echo "👍 OS and Architecture detected"

# Tar File Name Pattern
TAR_FILE="${FILE_BASENAME}_${OS}_${ARCH}.tar.gz"

# Get the latest release URL
LATEST_VERSION_URL=$(curl -sSLI -o /dev/null -w %{url_effective} "$RELEASES_URL/latest")

# Extract the latest release version from the URL
VERSION=$(echo "$LATEST_VERSION_URL" | awk 'match($0, /v[0-9]+\.[0-9]+\.[0-9]+/) { print substr($0, RSTART, RLENGTH) }')

# Error if Version is not there
if [ -z "$VERSION" ]; then
    echo
    echo "😭 Could not detect latest version of OCSF-Tool."
    echo "Please use Release URL to download the files manually."
    echo $RELEASES_URL 
    echo
    exit 1
fi

echo "👍 Detected latest version of OCSF-Tool"

# Download the tar file and checksums
curl -sfLO "$RELEASES_URL/download/$VERSION/$TAR_FILE"
curl -sfLO "$RELEASES_URL/download/$VERSION/checksums.txt"
echo "👍 Downloaded OCSF-Tool $VERSION (latest)"

# Exit if Tar file is not donwloaded
if [ ! -f "$TAR_FILE" ]; then
    echo
    echo "😭 Could not download \"$TAR_FILE\"."
    echo "Please use Release URL to download the binary manually."
    echo $RELEASES_URL 
    echo
    exit 1
fi

# Verify Checksums
# if sha256sum command exists
if command -v sha256sum >/dev/null 2>&1; then
    sha256sum --ignore-missing --quiet --check checksums.txt
    echo "👍 Verified downloaded files"
else
    echo "😦 \"sha256sum\" command not available to verify downloaded file"
fi

# Installation #

# Create ocsf-tool directory if does not exists
mkdir -p "./ocsf-tool"

# Clear ocsf-tool directory
rm -rf "./ocsf-tool/*"

# Extract tar file in ocsf-tool directory
tar -xf "$TAR_FILE" -C "./ocsf-tool"
echo "👍 Extracted the Tar in ./ocsf-tool directory"

# Remove tar file and checksums
rm -rf "$TAR_FILE" "checksums.txt"
echo "👍 Tar and Checksums removed"


# get current directory


# Done
echo
echo "🎉 Dwonload Complete!"
echo
echo "Go to \"$(pwd)/ocsf-tool\""
echo "And Run \"./ocsf-tool\""
echo

# Usage
./ocsf-tool/ocsf-tool --help