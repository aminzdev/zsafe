# 🚀 ZSafe

![GitHub release (latest by date)](https://img.shields.io/github/v/release/aminzdev/zsafe?style=for-the-badge)
![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/aminzdev/zsafe/codecov.yml?style=for-the-badge)
![GitHub issues](https://img.shields.io/github/issues/aminzdev/zsafe?style=for-the-badge)
![GitHub stars](https://img.shields.io/github/stars/aminzdev/zsafe?style=for-the-badge)
![Code Coverage](https://img.shields.io/codecov/c/github/aminzdev/zsafe?style=for-the-badge)
![GitHub license](https://img.shields.io/github/license/aminzdev/zsafe?style=for-the-badge)
![GitHub forks](https://img.shields.io/github/forks/aminzdev/zsafe?style=for-the-badge)
![GitHub watchers](https://img.shields.io/github/watchers/aminzdev/zsafe?style=for-the-badge)
![GitHub contributors](https://img.shields.io/github/contributors/aminzdev/zsafe?style=for-the-badge)
![Supported OS](https://img.shields.io/badge/OS-Linux%20%26%20Windows-blue?style=for-the-badge&logo=linux&logoColor=white)
![Docs](https://img.shields.io/badge/docs-up_to_date-brightgreen?style=for-the-badge)

> **A simple and efficient Backup tool.**

## ✨ Features

- [x] Create and restore backups of files & directories 📂
- [x] Compress backups using `zstd` for high performance 🗜️
- [x] Encrypt backups with `AES-256` for security 🔒
- [ ] Upload backups to **Telegram** 📤
- [ ] Upload backups to **Google Drive** ☁️

---

## 📥 Installation

### Install via Go

```bash
go install github.com/aminzdev/zsafe@latest
```

### Build from Source

```bash
git clone https://github.com/aminzdev/zsafe.git
cd zsafe
go build -o zsafe
```

After installation, you can run:

```bash
zsafe --help
```

---

## 🚀 Usage

### 🔹 Create a Backup

```bash
zsafe backup /path/to/directory
```

### 🔹 Create a Backup with password

```bash
zsafe backup /path/to/directory --password "secret"
```

### 🔹 Restore a Backup

```bash
zsafe restore backup.zst /path/to/directory
```

### 🔹 Restore an Encrypted Backup

```bash
zsafe restore backup.zst.enc /path/to/directory --password "secret"
```

---

## 🛠️ Development & Contribution

#### 1. Clone the repository:

```bash
git clone https://github.com/aminzdev/zsafe.git
cd zsafe
```

#### 2. Install dependencies:

```bash
go mod tidy
```

#### 3. Run the tool locally:

```bash
go run main.go --help
```

🚀 Contributions are welcome! Feel free to fork the repo and submit PRs. 😊

---

## 📜 License

This project is licensed under the **GPL-2.0** - see the [LICENSE](LICENSE) file for details.

---

## ⭐ Support & Feedback

If you like this project, consider giving it a **⭐ star**! 😊  
For issues or suggestions, open an [issue](https://github.com/aminzdev/zsafe/issues).
