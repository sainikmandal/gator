# gator 🐊
A CLI tool RSS feed aggregator in Go

## No one uses RSS bro, why even do this shit?
A great man once said, we do pointless things because they are hard and we can do it. Well that was not what he said, neither he was a great man.

---

## Prerequisites
Now lets pretend that someone is reading this and (pretend 10 times that they are thinking of installing this)
Make sure you have Go and Postgres installed!

---

## Installation
Run:
```bash
go install github.com/sainikmandal/gator@latest
```

Before running gator make a config file in your root dir. Run:
```bash
touch ~/gatorconfig.json
```

Please do not get creative with the config file name, not here.
gator expects a config file with format:
```json
{
  "db_url": "postgres://username:password@localhost:5432/gatordb?sslmode=disable",
  "current_user_name": "demo"
}
```

okay now sing a nice prayer and run:
```bash
gator register [time to be creative]
```

## Commands

### Authentication & User Management

* **Register a new user**:
```bash
gator register <username>
```

* **Login as existing user**:
```bash
gator login <username>
```

* **Reset database** (clears all data):
```bash
gator reset
```

* **List all users**:
```bash
gator users
```

### Feed Management

* **Add a feed**:
```bash
gator addfeed <feed_url> <feed_name>
```

* **List all feeds**:
```bash
gator feeds
```

* **Follow a feed**:
```bash
gator follow <feed_url>
```

* **List feeds you're following**:
```bash
gator following
```

* **Unfollow a feed**:
```bash
gator unfollow <feed_url>
```

### Content & Aggregation

* **Aggregate feeds** (fetch new posts):
```bash
gator agg <time_between_requests>
```
Example: `gator agg 1m` (fetches every minute)

* **Browse posts** (show latest posts from followed feeds):
```bash
gator browse <limit>
```
Example: `gator browse 10` (shows latest 10 posts)
