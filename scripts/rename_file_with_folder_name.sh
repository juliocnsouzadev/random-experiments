#!/bin/bash

# Function to recursively rename files
function rename_files {
    local parent="$1"
    local files=("$parent"/*)

    for file in "${files[@]}"; do
        if [[ -d "$file" ]]; then
            rename_files "$file"
        else
            local filename="$(basename "$file")"
            mv -- "$file" "$parent/$parent-$filename"
        fi
    done
}

# Get parent directory
parent="$1"

# Rename files recursively
rename_files "$parent"
