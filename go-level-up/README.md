# barebones-bazel: Go Foundations - Go Level Up

This Bazel workspace showcases a more intricate Go project. While still utilizing core Go components, it incorporates dependencies from a locally-defined library within the workspace.

## Getting Started

Ensure you have Go (version 1.22 or higher) and Bazel installed.

## Project Structure

```
.
├── MODULE.bazel
├── MODULE.bazel.lock
├── README.md
├── WORKSPACE.bazel
├── client
│   ├── BUILD.bazel
│   ├── go.mod
│   └── main.go
└── currency
    ├── BUILD.bazel
    ├── converter.go
    ├── converter_test.go
    └── go.mod
```

## Building and Running

**Build:**

```bash
bazel build ...
```

<details>
<summary>Commands to build target by target</summary>

```bash
bazel build //currency:converter
bazel build //currency:converter_test
bazel build //client:service
```

</details>

**Run:**

```bash
bazel run //client:service
```

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.