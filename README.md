# Barebones Bazel: Streamlining Go Builds with a Practical Guide

This repository accompanies the "Barebones Bazel" series of articles, providing practical examples and a step-by-step approach to building Go projects with Bazel. Learn how to set up a workspace, manage dependencies, build, generate,  and test code, and organize multi-module projects. Explore the power and efficiency of Bazel, including the latest advancements introduced in Blzmod 5.0 for your Go development.

## Getting Started

* Go version 1.22 or higher. As of this writing, the latest version is 1.23. Keep in mind that each major Go release is supported until two newer major releases are available. For the most up-to-date information on Go's release policy, refer to the [official documentation](https://go.dev/doc/devel/release#policy).
* Bazel version 7.0 or higher

## Project Structure

The repository are broken into the respective directories:

* **go-basics:** This directory houses a simple Go project with a single `main.go` file that builds a basic "Hello, Bazel!" binary using Bazel.
* **go-level-up:** This directory contains a Go project showcasing the use of Go modules and testing within a Bazel build environment.

## Building and Running

Each Go project directory contains a `README.md` file with specific instructions. Please refer to the `README.md` within each project for guidance.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.