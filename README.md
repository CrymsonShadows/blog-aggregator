# blog-aggregator

This is an aggregator for RSS feeds to store new updates of content locally into a database using a cli.

To use this program you need Postgres 15, Go, and Goose installed.

## Installing Dependencies

### 1. Getting postgres macOS with brew

`brew install postgresql@15`

The psql command-line utility is the default client for Postgres. Use it to make sure you’re on version 15+ of Postgres:

`psql --version`

#### To Start Postgres Server In Background

`brew services start postgresql@15`

Enter the psql shell and create a new database:

`CREATE DATABASE gator;`

Get your connection string and save it for later in the config. A connection string is just a URL with all of the information needed to connect to a database. The format is:

```
protocol://username:password@host:port/database
```

For macOS (no password, your username):

`postgres://username:@localhost:5432/gator`

### 2. Install Goose

`go install github.com/pressly/goose/v3/cmd/goose@latest`

Then go to the sql/schema directory and do:
`goose postgres <connection_string> up`

### 3. Get uuid Package

`go get github.com/google/uuid`

### 4. Get a Postgres Driver

`go get github.com/lib/pq`

## Setup Gator app

Manually create a config file in your home directory, ~/.gatorconfig.json, with the connection string followed by `?sslmode=disable`:

```go
{
  "db_url": "protocol://username:password@host:port/database?sslmode=disable"
}
```

Then to install gator do.
`go install -o gator`

## Commands

### 1. Register

A username has to be registered to futher use any commands. More than one user can be registered. Each user follows their own feeds.
`gator register <username>`

### 2. Login

Login to a registered username to be able to use gator with that user's followed feeds.
`gator login <username>`

### 3. Reset

Reset all data stored in database.
`gator reset`

### 4. Users

Lists all regisetered users.
`gator users`

### 5. Feeds

List all followed feeds.
`gator feeds`

### 6. Add Feed

Add a feed to the current user's followed feeds by giving a name for the feed and then the url. You can't add a feed already followed by another user.
`gator addfeed "Feed Name" "https://url.com"`

### 7. Follow

Give the url of a feed to follow for the current user. The feed must have been added by a user before with the addfeed command.
`gator follow "https://url.com"`

### 8. Following

Show all the feeds the current user is following.
`gator following`

### 9. Unfollow

Unfollow a feed for the current user by giving the feed's url.
`gator unfollow "https://url.com"`

### 10. Aggregate

To have gator scrape feeds at a given interval by oldest scrapped feed first. This runs forrever unless inturrupted.
`gator agg <time between reqs (10s or 10m)>`

### 11. Browse

Browse the most recent posts from followed feeds for the given user. Default number of posts given is 2 if no number is given.
`gator browse <number of posts>`
