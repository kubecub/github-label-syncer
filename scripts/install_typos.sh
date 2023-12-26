#!/bin/bash

# Download the typos tool
curl -L https://github.com/crate-ci/typos/releases/download/1.0.14/typos-x86_64-unknown-linux-gnu.tar.gz -o typos.tar.gz

# Extract the downloaded file
tar -xvf typos.tar.gz

# Move the typos tool to the tools directory
mv typos ./tools/

# Clean up the downloaded file
rm typos.tar.gz
