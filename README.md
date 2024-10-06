# GOGator
This is a CLI application capable of storing all the RSS feeds you want from the internet.

The collected Post is stored in PostgreSQL database allowing you to follow and unfollow the feeds added by other users

You can also see the summary of the aggregated posts in the terminal, with a link to the full post


## Requirements
You will need to download and install the GO programming language, Postgres Database manangement system and Git Version control from the link given below:-

- [Download Go](https://go.dev/)
- [Download PostgreSQL](https://www.postgresql.org/download/)
- [Download GIT](https://git-scm.com/downloads)
## Installation

After you have downloaded the everything you need you need to download from the [Requirements Section](##Requirements). You just need to follow the Steps Given Below.

### Step 1 
Clone the repository in your local machine by the navigating to the desired directory and typing the following command in the terminal.
```bash
  git clone https://github.com/SplinterSword/GOGator.git
  cd GOGator
```

### Step 2
To install it throught your system you need to run the following commands
```bash
  go build
  go install
```

### Step 3 (if the first 2 steps are not enough (they should be enough))
The last thing you need to is setup your database. Type the following commands in the terminal one by one.

```bash
  sudo -u postgres psql
  CREATE DATABASE gator;
  \c gator
  ALTER USER postgres PASSWORD 'postgres';
```

After these to steps GOGator will be installed in your local machine.
    
## How to Setup the config file

Before you can use it. You also need to Setup the gatorconfig.json file. Just Copy the text below in the gatorconfig.json file and make the appropriate changes

```
{
    "db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
    "current_user_name":""
}
```




## How to use GOGator

To Register a user
```
  GOGator register [name]
```

To Login as any other existing users
```
  GOGator login [Registered User]
```

To Get All the users
```
  GOGator users
```

To Add a RSS Feed
```
  GOGator addfeed [feedname] [feedURL]
```

To follow feed added by another user
```
  GOGator follow [feedurl]
```

To List ALL the available/added feeds
```
  GOGator feeds
```

To Get the feeds followed by the current user
```
  GOGator following
```

To Unfollow a particular feed
```
  GOGator unfollow [url]
```

To Get All the post from all the feeds(should be left running continously)

```
  GOGator agg [time_between_req(s/min/hrs)]
```

To Browse the posts from the aggregated feeds
```
  GOGator browse
```
To Reset Your Database
```
  GOGator reset
```