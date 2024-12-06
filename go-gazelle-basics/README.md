# barebones-bazel: Go Foundations - Go Level Up

This Bazel workspace demonstrates how Gazelle simplifies Go project setup within Bazel by automatically generating build files for a basic, dependency-free project. This eliminates much of the manual configuration typically required.

## Getting Started

Ensure you have Go (version 1.22 or higher) and Bazel installed.

## Project Structure

```
.
├── MODULE.bazel
├── MODULE.bazel.lock
├── README.md
├── WORKSPACE.bazel
├── go.mod
├── client
│   ├── BUILD.bazel
│   └── main.go
└── currency
    ├── BUILD.bazel
    ├── converter.go
    └── converter_test.go
```

## Generation

If the BUILD.bazel files are not found in the `client` and `currency` folders, run the command to use gazelle to generate the file.

```bash
bazel run gazelle
```

## Building and Running

**Build:**

```bash
bazel build ...
```

<details>
<summary>Commands to build target by target</summary>

```bash
bazel build //client:client
bazel build //client:client_lib
bazel build //currency:converter
bazel build //currency:converter_test
```

</details>

**Run:**

```bash
bazel run //client:service
```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.