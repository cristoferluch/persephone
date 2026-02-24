# 🐾 Persephone

A terminal-based UI for exploring PostgreSQL database schemas. Browse your tables, columns, and indexes without leaving the command line.

![Go](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/license-MIT-blue)

---

## ✨ Features

- 🔍 **Live search** — filter tables by name as you type
- 📋 **Column inspector** — view name, type, length, precision, nullability, and primary key info
- 🗂️ **Index viewer** — see all indexes and their keys for any table
- ⚡ **In-memory caching** — column and index data is fetched once and cached for fast navigation
- 🖱️ **Mouse support** — navigate with keyboard or mouse

---

## 📸 Preview

![Screenshot](https://github.com/cristoferluch/assets/blob/main/persephone.gif)

---

## 🚀 Getting Started

### Option 1 — Download the binary

Grab the latest release from the [Releases](https://github.com/cristoferluch/persephone/releases) page.

**Linux / macOS**
```bash
curl -L https://github.com/cristoferluch/persephone/releases/latest/download/persephone -o persephone
chmod +x persephone
mv persephone /usr/local/bin/
```

**Windows**

Download `persephone.exe` from the releases page and add it to your `PATH`.

### Option 2 — Build from source

```bash
git clone https://github.com/cristoferluch/persephone.git
cd persephone
go build -o persephone ./cmd/main.go
```

### Configuration

Create a `settings.yaml` file in the project root:

```yaml
host: localhost
port: 5432
user: your_user
password: your_password
database: your_database
sslmode: disable
```

### Run

```bash
go run .\cmd\app\main.go
```

---

## ⌨️ Keybindings

| Key       | Action                               |
|-----------|--------------------------------------|
| `Ctrl+K`  | Focus the search input               |
| `Tab`     | Move focus from search to table list |
| `↑` / `↓` | Navigate the table list              |
| `Ctrl+C`  | Quit                                 |

---

## 🛠️ Tech Stack

| Library                                   | Purpose                  |
|-------------------------------------------|--------------------------|
| [tview](https://github.com/rivo/tview)    | Terminal UI framework    |
| [tcell](https://github.com/gdamore/tcell) | Terminal cell rendering  |
| [Viper](https://github.com/spf13/viper)   | Configuration management |
| [lib/pq](https://github.com/lib/pq)       | PostgreSQL driver        |

---

## 🤝 Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

1. Fork the repository
2. Create a feature branch: `git checkout -b feat/my-feature`
3. Commit your changes: `git commit -m 'feat: add my feature'`
4. Push and open a Pull Request

---

## 📄 License

MIT License — see [LICENSE](LICENSE) for details.
