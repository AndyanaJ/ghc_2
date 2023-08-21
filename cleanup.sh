#!/bin/bash

# Define the directory where cleanup will be performed
cleanup_dir="./garbage/"

# Define the file extension to be deleted
file_extension=".txt"

# Move to the cleanup directory
cd "$cleanup_dir" || exit

# Delete files with the specified extension
find . -type f -name "*$file_extension" -delete

echo "Cleanup completed: Deleted $file_extension files in $cleanup_dir"
