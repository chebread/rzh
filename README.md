# rzh
`rzh` is a program that manages Zsh history.

## Table of Contents
- [Features](#features)
- [Note](#note)
- [How to use](#how-to-use)
- [Installation](#installation)
- [License](#license)

## Features
- Add commands to Zsh history
- Remove commands from Zsh history
- Deduplicate Zsh history entries
- Back up Zsh history

## Note
This program operates on the `.zsh_history` file located directly in your home directory (e.g. `~/.zsh_history`). If the file is in a different location, the program will not function correctly.

## How to use
### Add commands
```shell
rzh add "<command1>" ["<command2>"...] [--backup/-b] [--force/-f]
```
Adds one or more new commands to your .zsh_history file.

- The `--backup` flag creates a full backup of your history file before adding.
- The `--force` flag skips the confirmation message and adds the command immediately.

```shell
# Add a single command "ls -la" to history
> rzh add "ls -la"

# Add multiple commands at once
> rzh add "go build" "go test" "git push"

# Add multiple commands after creating a backup
> rzh add "npm install" "npm run dev" --backup
```

### Remove commands
```shell
rzh remove "<command1>" ["<command2>"...] [--prefix/-p] [--backup/-b] [--force/-f]
```
Removes all commands from your ~/.zsh_history file that contain any of the specified search terms. You can provide one or more terms.

By default, the command will show you all the matching lines that will be deleted and ask for confirmation. The matching is case-sensitive.

- The `--backup` flag creates a full backup of your history file before removal.
- The `--force` flag skips the confirmation prompt and deletes the matching lines immediately.

### Deduplicate history
```shell
rzh dedup [--force/-f]
```
Cleans your `.zsh_history` by removing duplicate command entries.

It keeps the **most recent entry** for each duplicated command based on its timestamp and removes all older ones.

- The `--force` flag skips the confirmation prompt.

```shell
# Deduplicate history with a confirmation prompt
> rzh dedup

# Deduplicate history without confirmation
> rzh dedup --force
```

### Manage backups
```shell
rzh backup [create | list | remove | restore] [--force/-f]
```
Manages backups of your history file, stored in `~/rzh/backup/`. The default action is `create`.

- `create`: Creates a new backup with a timestamped filename.
- `list`: Shows all available backups.
- `remove`: Launches an interactive interface to select and delete backups.
- `restore`: Presents a numbered list of available backups and prompts you to select one to restore. The contents of the selected backup will completely overwrite your current `.zsh_history` file.

- The `--force` flag is applicable to `remove` and `restore` to skip interactive prompts.

```shell
# Create a new backup (default action)
> rzh backup
> rzh backup create

# List all backups
> rzh backup list

# Interactively select a backup to restore and merge with current history
> rzh backup restore
```

## Installation
1. Visit [the GitHub Releases page](https://github.com/chebread/rzh/releases) for `rzh`.
2. Download the appropriate file for your operating system and architecture.
3. Unachive the downloaded file.
4. Execute the `rzh` executable file.
5. For easier access, consider adding `rzh` executable file to your system's PATH environment variable.

## License
MIT LICENSE &copy; 2025 Cha Haneum