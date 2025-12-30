# Maven Plus

> A terminal-friendly Maven workflow for people who prefer CLIs over IDEs

Maven is great, and that's why I'm building on top of it.

But let's be honest: Maven shines when you abstract it behind an IDE. When it comes to using a terminal-centric workflow and a text editor like Neovim, Java tooling just isn't there.

**Maven Plus (mvnp)** bridges that gap with an intuitive command-line experience for Maven projects

## Features

**Maven Plus** streamlines your Java development workflow:

- **Interactive Project Creation** - Create new Maven projects with an intuitive TUI
- **Generate Classes** - Generate classes, iterface, enums and record with a single command

---

## Installation

**Via Go:**

```bash
go install github.com/maxbrt/mvnp@latest
```

**From source:**

```bash
git clone https://github.com/maxbrt/mvnp.git
cd mvnp
go build -o mvnp
```

**Prerequisites:**

- Maven installed and available in your PATH
- Java Development Kit (JDK)

## Usage

### Create a New Project

```bash
mvnp init
```

### Generate Classes

```bash
mvnp gen
```
---

## Built With

Maven Plus is powered by some excellent Go libraries:

| Library | Purpose |
|---------|---------|
| [Cobra](https://github.com/spf13/cobra) | CLI framework |
| [Bubble Tea](https://github.com/charmbracelet/bubbletea) | Terminal UI framework |
| [Bubbles](https://github.com/charmbracelet/bubbles) | TUI components |
| [Lipgloss](https://github.com/charmbracelet/lipgloss) | Terminal styling |
| [etree](https://github.com/beevik/etree) | XML parsing |

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
