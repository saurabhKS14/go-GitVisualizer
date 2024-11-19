# Git Visualizer CLI

`git-visualizer` is a command-line interface (CLI) tool built in Go to visualize Git repository data. It provides several interactive features, such as commit heatmaps, contributor statistics, and more, to help developers understand their Git history and collaboration trends.

## Features

- **Commit Heatmap**: View a heatmap showing commit activity over time (daily, weekly, monthly).
- **Contributor Statistics**: See top contributors for a selected time range (week/month).
- **Interactive Navigation**: Navigate through commit data using arrow keys.
- **Export Options**: Export commit heatmap and contributor stats to CSV, JSON, or text files.
- **Easy Setup**: Install and run the tool with minimal configuration.

## Installation

### Prerequisites

- **Go**: Make sure you have Go installed. You can download it from [here](https://golang.org/doc/install).
- **Git**: Ensure that Git is installed and configured on your machine.

### Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/<your-username>/git-visualizer.git
   ```

2. Navigate to the project directory:

   ```bash
   cd git-visualizer
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

4. Build the project:

   ```bash
   go build
   ```

5. Run the CLI tool:

   ```bash
   ./gitviz
   ```

## Usage

### Basic Commands

- **Visualize Branches**

  To list all branches in your Git repository and highlight current branch:

  ```bash
  gitviz .
  ```

## Development

If you'd like to contribute or learn more about how this project is built, check out the development section.


### Building the Project

To build the CLI tool from source, run the following:

```bash
go build
```

### Adding Features

Feel free to fork the repository and create new features. Open an issue to discuss your ideas or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

### Key Technologies Used

- **Go**: The core programming language for this project.
- **go-git**: A Git library in Go for interacting with Git repositories.
- **Bubble Tea**: A Go package for building interactive terminal UIs.
- **Cobra**: A Go library for building command-line applications.

## Contributing

If you would like to contribute to this project, feel free to fork the repository and submit pull requests. If you find any bugs or have feature requests, please open an issue.

---
