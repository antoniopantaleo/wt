<p align="center">
<img height="320px" alt="image" src="https://github.com/user-attachments/assets/438d0a25-e9c1-4e11-8b2c-2b8309b58056" />
</p>

`wt` is a small CLI that helps you manage and view git worktrees across ALL yout repositories you want to track.

## Features

- Track all of your repositories in a single local config file
- List worktrees for all tracked repositories
- Show tracked repositories
- Add/remove tracked repositories

## Config

`wt` reads config from:

```text
$XDG_CONFIG_HOME/.wt/config.json
```

## Usage

### Add a repository

```bash
wt add /path/to/repo
```

### List worktrees from all managed repos

```bash
wt ls
```

Example output:

```text
PATH                                     BRANCH         HEAD
---------------------------------------  -------------  ----------------------------------------
/Users/me/src/my-api                     main           3e92a8e6d8c1b2d5f8d6f0f4a6d9f1f2a1b3c4d5
/Users/me/worktrees/my-api/feature/auth  feature/auth   29c4b7f50a8c49322b7f39e88030e3e4f2aa1ccd
```

### Show managed repositories

```bash
wt managed
```

Example output:

```text
• /my/repo/one
• /my/repo/two
```

### Remove a managed repository

```bash
wt rm /path/to/repo
```
