# barebones-bazel: Go Foundations - Go Basics

This Bazel workspace demonstrates building and running a basic Go project. The project uses fundamental Go elements and has no external dependencies.

## Getting Started

Ensure you have Go (version 1.22 or higher) and Bazel installed.

## Project Structure

```
.
├── BUILD.bazel 
├── MODULE.bazel
├── README.md
├── WORKSPACE.bazel
└── main.go
```

## Building and Running

**Build:**

```bash
bazel build //:main
```

**Run:**

```bash
bazel run //:main
```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.