name: Go Build and Test

on:
  push:
    branches:
      - main  # Change to your default branch name if needed

jobs:
  build-and-test:
    runs-on: ubuntu-latest  # You can change the runner if needed

    steps:
    - name: Checkout Repository
      uses: actions/checkout@v2  # This action checks out your repository

    - name: Set Up Go
      uses: actions/setup-go@v2
      with:
        go-version: 1.16  # Change this to your preferred Go version

    - name: Build Application
      run: go build -o myapp main.go  # Adjust the binary name and source file

    - name: Run Tests
      run: go test ./...  # Run tests for all packages

    - name: Archive Artifact
      uses: actions/upload-artifact@v2
      with:
        name: myapp-binary
        path: myapp  # Name of the binary file created in the build step
