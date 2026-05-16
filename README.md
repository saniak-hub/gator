# Gator - A CLI Tool to Manage RSS Feeds

## Overview

Gator is a command-line tool that helps you aggregate and manage RSS feeds from multiple sources. It allows you to subscribe to feeds, fetch new items at regular intervals, and browse posts from all your subscriptions in one convenient place.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Commands](#commands)
  - [User Management](#user-management)
  - [Feed Management](#feed-management)
  - [Content Management](#content-management)
- [Typical Workflow](#typical-workflow)
- [Troubleshooting](#troubleshooting)

## Prerequisites

- **Go** 1.16 or higher
- **PostgreSQL** - Must be installed and running in your working environment
- .gatorconfig.json file in the home directory
  ```json
  {
  "db_url": "<postgres connection string>"
  example: postgres://user:password@localhost:port/db_name?sslmode=disable
  }
  ```
  

## Installation

### Using `go install`

```bash
go install github.com/saniak-hub/gator
```

### Build from Source

```bash
go build
```

This will generate an executable called `gator`.

## Quick Start

1. **Register a new user:**
   ```bash
   ./gator register myusername
   ```

2. **Login:**
   ```bash
   ./gator login myusername
   ```

3. **Add a feed to follow:**
   ```bash
   ./gator addfeed https://example.com/feed.xml
   ```

4. **Start aggregating (fetches feeds periodically):**
   ```bash
   ./gator agg 30s
   ```

5. **Browse your posts:**
   ```bash
   ./gator browse
   ```

## Commands

### User Management

#### `register`

```bash
./gator register <username>
```

Register a new user account with the specified username.

**Example:**
```bash
./gator register john_doe
```

---

#### `login`

```bash
./gator login <username>
```

Login to start using the CLI. You must be logged in to use feed management and browsing commands.

**Example:**
```bash
./gator login john_doe
```

---

#### `users`

```bash
./gator users
```

List all registered users in the system.

---

#### `reset`

```bash
./gator reset
```

⚠️ **Warning:** Remove all data from the database. This action cannot be undone.

---

### Feed Management

#### `addfeed`

```bash
./gator addfeed <url>
```

Add a new RSS feed to the database for persistence. You must include the full URL to the RSS feed.

**Example:**
```bash
./gator addfeed https://example.com/rss.xml
```

---

#### `feeds`

```bash
./gator feeds
```

List all feeds you've added to the system. Displays the feed name and URL.

**Example Output:**
```
Feed Name: TechNews
URL: https://technews.example.com/feed.xml
```

---

#### `follow`

```bash
./gator follow <url>
```

Subscribe to a feed. Requires the URL of the RSS feed you want to follow.

**Example:**
```bash
./gator follow https://blog.example.com/feed.xml
```

---

#### `unfollow`

```bash
./gator unfollow <url>
```

Unsubscribe from a feed.

**Example:**
```bash
./gator unfollow https://blog.example.com/feed.xml
```

---

#### `following`

```bash
./gator following
```

List all the feeds you are currently subscribed to.

---

### Content Management

#### `agg`

```bash
./gator agg <time_between_requests>
```

Recursively fetch and store new items from all subscribed feeds to the database.

**Parameters:**
- `<time_between_requests>`: The time interval between feed requests (e.g., `30s`, `1m`, `5m`)

**Example:**
```bash
./gator agg 30s    # Fetch feeds every 30 seconds
./gator agg 5m     # Fetch feeds every 5 minutes
```

---

#### `browse`

```bash
./gator browse [limit]
```

Display posts from your subscribed feeds in the terminal. Shows the title and link for each post.

**Parameters:**
- `[limit]` (optional): Number of posts to display. Defaults to 2 if not specified.

**Example:**
```bash
./gator browse      # Show 2 latest posts
./gator browse 10   # Show 10 latest posts
```

---

## Typical Workflow

1. **Setup** - Register an account: `./gator register username`
2. **Login** - Start a session: `./gator login username`
3. **Subscribe** - Add feeds you want to follow: `./gator addfeed <url>`
4. **Aggregate** - Keep feeds up to date: `./gator agg 5m`
5. **Browse** - Read new posts: `./gator browse`
6. **Manage** - Follow/unfollow feeds as needed: `./gator follow <url>` or `./gator unfollow <url>`

## Troubleshooting

**PostgreSQL Connection Error**
- Ensure PostgreSQL is installed and running on your system
- Check that your database credentials are properly configured

**Feed Not Updating**
- Verify the URL is a valid RSS feed
- Check your internet connection
- Ensure the feed server is accessible

**Permission Denied When Running gator**
- Make sure the executable has execute permissions: `chmod +x gator`

---

For more information or to report issues, visit the [GitHub repository](https://github.com/saniak-hub/gator).
