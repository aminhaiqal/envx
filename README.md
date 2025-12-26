# ⚡ envx

<div align="center">

![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go)
![License](https://img.shields.io/badge/license-MIT-blue?style=for-the-badge)
![Build Status](https://img.shields.io/badge/build-passing-brightgreen?style=for-the-badge)
![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen?style=for-the-badge)

**Lightning-fast environment variable management for modern developers.**

A blazingly fast, local-first CLI tool to manage environment variables across all your projects. Zero setup, zero servers, zero hassle.

[Features](#-features) • [Installation](#-installation) • [Quick Start](#-quick-start) • [Documentation](#-documentation) • [Contributing](#-contributing)

</div>

---

## Why envx?

Managing environment variables shouldn't require a PhD. **envx** is built for solo developers who want:

- **Instant access** to env vars across projects
- **Zero configuration** - just install and go
- **Local-first** - no servers, no cloud, no complexity
- **One command** to switch between dev/staging/prod

Stop juggling `.env` files. Stop losing configurations. Stop the chaos.

**envx makes environment management invisible.**

## Features

- **Blazingly Fast** - Written in Go, optimized for speed
- **Multi-Project** - Manage unlimited projects from one place
- **Multi-Environment** - Dev, staging, production profiles
- **Import/Export** - Works seamlessly with `.env` files
- **Secure Storage** - Encrypted values for sensitive data (Phase 4)
- **Smart Templates** - Share env structure without secrets
- **Cross-Platform** - macOS, Linux, Windows support
- **Beautiful CLI** - Intuitive interface, clear output
- **Git-Friendly** - Optional version control integration
- **Zero Dependencies** - Single binary, no runtime required

## Installation

### Using Go

```bash
go install github.com/aminhaiqal/envx@latest
```

### Using Homebrew (macOS/Linux)

```bash
# Coming soon
brew install envx
```

### Download Binary

Grab the latest release for your platform from the [releases page](https://github.com/aminhaiqal/envx/releases).

### Build from Source

```bash
git clone https://github.com/aminhaiqal/envx.git
cd envx
go build -o envx cmd/envx/main.go
```

## Quick Start

### Initialize your first project

```bash
envx init myapp
```

### Add your environment variables

```bash
envx set myapp DATABASE_URL="postgresql://localhost:5432/mydb"
envx set myapp API_KEY="sk-your-secret-key-here"
envx set myapp PORT=3000
envx set myapp NODE_ENV="development"
```

### View your variables

```bash
envx list myapp
```

Output:
```
📦 myapp (development)

DATABASE_URL  postgresql://localhost:5432/mydb
API_KEY       sk-*********************here
PORT          3000
NODE_ENV      development

4 variables
```

### Export to .env file

```bash
envx export myapp
# ✓ Exported to .env
```

### Use in your workflow

```bash
# Generate .env and start your app
envx export myapp && npm run dev

# Quick export for different environments
envx export myapp --env production -o .env.prod
```

## Documentation

### Core Commands

#### `envx init <project> [flags]`

Initialize a new project profile.

```bash
# Create a new project
envx init myapp

# Create with specific environment
envx init myapp --env production

# Create with description
envx init myapp --desc "My awesome API"
```

**Flags:**
- `--env, -e` - Environment name (default: development)
- `--desc, -d` - Project description

---

#### `envx set <project> <KEY=value> [flags]`

Set an environment variable.

```bash
# Set a variable
envx set myapp DATABASE_URL="postgresql://localhost:5432/db"

# Set for specific environment
envx set myapp API_URL="https://api.prod.com" --env production

# Set with description (helpful for team templates)
envx set myapp SECRET_KEY="xxx" --desc "OAuth secret key"

# Set multiple at once
envx set myapp KEY1=value1 KEY2=value2 KEY3=value3
```

**Flags:**
- `--env, -e` - Target environment
- `--desc, -d` - Variable description
- `--secret, -s` - Mark as secret (masked in list)

---

#### `envx get <project> <KEY> [flags]`

Retrieve a specific variable value.

```bash
# Get a variable
envx get myapp DATABASE_URL

# Get from specific environment
envx get myapp API_URL --env production

# Copy to clipboard
envx get myapp API_KEY --copy
```

**Flags:**
- `--env, -e` - Source environment
- `--copy, -c` - Copy value to clipboard

---

#### `envx list <project> [flags]`

List all environment variables for a project.

```bash
# List all variables (secrets masked)
envx list myapp

# Show all values including secrets
envx list myapp --show-secrets

# List specific environment
envx list myapp --env production

# Export as JSON
envx list myapp --json
```

**Flags:**
- `--env, -e` - Target environment
- `--show-secrets` - Display secret values
- `--json` - Output as JSON

---

#### `envx rm <project> <KEY> [flags]`

Delete an environment variable.

```bash
# Delete a variable
envx rm myapp OLD_CONFIG

# Delete from specific environment
envx rm myapp TEMP_KEY --env staging

# Delete multiple
envx rm myapp KEY1 KEY2 KEY3
```

**Flags:**
- `--env, -e` - Target environment
- `--force, -f` - Skip confirmation

---

#### `envx export <project> [flags]`

Export variables to a .env file.

```bash
# Export to .env in current directory
envx export myapp

# Export to custom path
envx export myapp --output ./config/.env.local

# Export specific environment
envx export myapp --env production --output .env.prod

# Export with comments
envx export myapp --with-comments
```

**Flags:**
- `--env, -e` - Source environment
- `--output, -o` - Output file path (default: .env)
- `--with-comments` - Include variable descriptions as comments
- `--overwrite` - Overwrite existing file without prompt

---

#### `envx import <project> <file> [flags]`

Import variables from an existing .env file.

```bash
# Import from .env file
envx import myapp .env

# Import to specific environment
envx import myapp .env.production --env production

# Merge without overwriting existing values
envx import myapp .env --merge

# Dry run (preview changes)
envx import myapp .env --dry-run
```

**Flags:**
- `--env, -e` - Target environment
- `--merge` - Don't overwrite existing variables
- `--dry-run` - Preview import without applying

---

#### `envx clone <source> <destination> [flags]`

Clone a project profile.

```bash
# Clone entire project
envx clone myapp myapp-backup

# Clone to different environment
envx clone myapp myapp-staging --env staging

# Clone specific environment only
envx clone myapp myapp-prod --from production --to production
```

**Flags:**
- `--env, -e` - Target environment for destination
- `--from` - Source environment
- `--to` - Destination environment

---

#### `envx template <project> [flags]`

Generate a template file (variable names without values).

```bash
# Output to stdout
envx template myapp

# Save to file
envx template myapp --output .env.template

# Include descriptions
envx template myapp --with-desc > .env.example
```

**Flags:**
- `--output, -o` - Output file path
- `--with-desc` - Include variable descriptions
- `--env, -e` - Source environment

---

#### `envx projects [flags]`

List all managed projects.

```bash
# List all projects
envx projects

# Show detailed info
envx projects --detailed

# Filter by environment
envx projects --env production
```

**Flags:**
- `--detailed, -d` - Show variable counts and descriptions
- `--env, -e` - Filter by environment

---

#### `envx switch <project> <environment>`

Set the default environment for a project.

```bash
# Switch default environment
envx switch myapp production

# Now all commands use production by default
envx list myapp  # Lists production env
```

---

#### `envx envs <project>`

List all environments for a project.

```bash
envx envs myapp

# Output:
# development (default)
# staging
# production
```

---

### Configuration

envx stores data locally in a secure location:

**Storage locations:**
- **macOS/Linux**: `~/.config/envx/`
- **Windows**: `%APPDATA%\envx\`

**Configuration file:**
- `~/.config/envx/config.json`

**Project data:**
- `~/.config/envx/projects/<project-name>.json`

## Project Structure

```
envx/
├── cmd/
│   └── envx/
│       └── main.go              # CLI entry point
├── internal/
│   ├── storage/
│   │   ├── storage.go           # Storage interface
│   │   ├── json.go              # JSON storage
│   │   └── encrypted.go         # Encrypted storage (Phase 4)
│   ├── config/
│   │   └── config.go            # Config management
│   ├── profile/
│   │   ├── profile.go           # Profile operations
│   │   └── environment.go       # Environment handling
│   ├── exporter/
│   │   ├── dotenv.go            # .env export
│   │   └── template.go          # Template generation
│   └── importer/
│       └── dotenv.go            # .env import
├── pkg/
│   └── envx/
│       ├── vault.go             # Core operations
│       ├── types.go             # Data structures
│       └── validator.go         # Input validation
├── tests/
│   ├── unit/
│   └── integration/
├── docs/
│   ├── getting-started.md
│   ├── best-practices.md
│   └── faq.md
├── go.mod
├── go.sum
├── README.md
├── LICENSE
├── .goreleaser.yml
└── .github/
    └── workflows/
        ├── test.yml
        ├── release.yml
        └── lint.yml
```

## Development

### Prerequisites

- Go 1.21 or higher
- Git
- Make (optional)

### Local Setup

```bash
# Clone the repository
git clone https://github.com/aminhaiqal/envx.git
cd envx

# Install dependencies
go mod download

# Build
make build
# OR
go build -o envx cmd/envx/main.go

# Run
./envx --help
```

### Running Tests

```bash
# Run all tests
make test
# OR
go test ./...

# Run with coverage
make test-coverage
# OR
go test -cover ./...

# Run integration tests
go test ./tests/integration/... -v

# Run benchmarks
go test -bench=. ./...
```

### Code Quality

```bash
# Run linter
golangci-lint run

# Format code
gofmt -s -w .

# Vet code
go vet ./...
```

## 🗺️ Roadmap

### Phase 1: MVP ✅ (Weeks 1-2)
- [x] Project initialization
- [x] Basic CRUD operations (set, get, list, rm)
- [x] Export to .env
- [x] Import from .env
- [x] Cross-platform file handling

### Phase 2: Multi-Environment 🚧 (Weeks 3-4)
- [ ] Multiple environment profiles per project
- [ ] Environment switching
- [ ] Default environment management
- [ ] Environment listing

### Phase 3: Polish & UX 📋 (Week 5)
- [ ] Beautiful CLI output with colors
- [ ] Template generation
- [ ] Profile cloning
- [ ] Bulk operations
- [ ] Better error messages

### Phase 4: Security 🔒 (Week 6-7)
- [ ] Encrypted value storage
- [ ] Master password protection
- [ ] Secure value masking
- [ ] Audit logging

### Phase 5: Advanced Features 🚀 (Future)
- [ ] Git integration (commit env changes)
- [ ] Variable validation and schemas
- [ ] Shell completions (bash, zsh, fish)
- [ ] Interactive TUI mode
- [ ] Cloud backup (optional)
- [ ] Team sharing (encrypted exports)
- [ ] Docker integration
- [ ] Kubernetes secrets export
- [ ] Plugin system

## 🎓 What You'll Learn

This project covers essential Go concepts:

**Go Fundamentals:**
- Project structure and package organization
- Error handling patterns
- File I/O and path management
- JSON encoding/decoding
- Command-line interfaces with Cobra
- Cross-platform development

**Software Engineering:**
- Clean architecture principles
- Test-driven development
- API design
- State management
- Security best practices

**DevOps & Distribution:**
- CI/CD with GitHub Actions
- Automated releases with GoReleaser
- Cross-platform builds
- Package distribution

## Contributing

Contributions make open source amazing! Whether you're fixing bugs, adding features, or improving docs—**all contributions are welcome**.

### How to Contribute

1. **Fork** the repository
2. **Create** your feature branch (`git checkout -b feature/AmazingFeature`)
3. **Commit** your changes (`git commit -m 'Add some AmazingFeature'`)
4. **Push** to the branch (`git push origin feature/AmazingFeature`)
5. **Open** a Pull Request

### Development Guidelines

- Write tests for new features
- Follow Go conventions and best practices
- Update documentation for user-facing changes
- Keep commits atomic and descriptive
- Run tests before submitting PR
- Add examples for new commands

### Good First Issues

Look for issues labeled `good-first-issue` to get started!

## Community & Support

- **Bug Reports**: [Open an issue](https://github.com/aminhaiqal/envx/issues/new?template=bug_report.md)
- **Feature Requests**: [Open an issue](https://github.com/aminhaiqal/envx/issues/new?template=feature_request.md)
- **Discussions**: [GitHub Discussions](https://github.com/aminhaiqal/envx/discussions)
- **Docs**: [Wiki](https://github.com/aminhaiqal/envx/wiki)

## License

Distributed under the MIT License. See [`LICENSE`](LICENSE) for more information.

## 🌟 Show Your Support

If envx makes your dev life easier, consider:
- **Starring** this repo
- **Sharing** on Twitter
- **Writing** about your experience
- **Contributing** code or docs

## Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for CLI framework
- Inspired by the need for simpler env management
- Thanks to all contributors who make this project better

## Contact

**Amin Haiqal**
- GitHub: [@aminhaiqal](https://github.com/aminhaiqal)
- LinkedIn: [linkedin.com/in/amin-haiqal](https://www.linkedin.com/in/amin-haiqal/)

**Project Link:** [https://github.com/aminhaiqal/envx](https://github.com/aminhaiqal/envx)

---

<div align="center">

**Built with ❤️ and Go by a developer, for developers**

[⭐ Star this repo](https://github.com/aminhaiqal/envx/stargazers) • [Report Bug](https://github.com/aminhaiqal/envx/issues) • [Request Feature](https://github.com/aminhaiqal/envx/issues)

</div>
