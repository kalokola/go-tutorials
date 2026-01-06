#!/bin/bash

# Prompt for commit message
read -p "Enter commit message: " msg

# Check if message is empty
if [ -z "$msg" ]; then
    echo "Error: Commit message cannot be empty"
    exit 1
fi

# Git commands
git add .
git commit -m "$msg"
git push -u origin main