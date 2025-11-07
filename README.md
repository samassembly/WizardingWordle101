# WizardingWordle101: Prove your mastery of spells from Ravenwood Academy!

Inspired by Wordle, guess a spell from the world of Wizard101. Narrow down your guesses by paying attention to the Cost, School, Accuracy, and Type attributes. If you can guess the secret spell within 5 guesses, you win!

## 🧙‍♂️ Motivation

As an avid enjoyer of Wizard101, I struggle to introduce new friends to the 10+ year old MMO due to the overwhelming amount of content the game has to offer. I decided to ease them into the game with a series of companion applications to help learn the game and its mechanics. WizardingWordle101 is designed to provide a fun daily challenge that reinforces knowledge of spells across all classes to better a user's combat potential. Understanding what spells might be cast at you is key to becoming a masterful practitioner of magic!

<img src="https://github.com/samassembly/WizardingWordle101/blob/main/demos/WizardingWordleDemo.gif" width="100%" alt="WizardingWordle101 demo gif">

## 🪄 Quick Start

### 1. Installation
This project requires **Go**, **PostgreSQL**, and **Goose** for database migrations. Follow the steps below to set up your environment:
--
### Install Go
   - **Linux**:  
     ```bash
     sudo apt-get update
     sudo apt-get install golang
     ```
   
Verify installation:  
    ```bash
    go version
    ```

### Install Postgres
- **Linux**:  
     ```bash
     sudo apt-get install postgresql postgresql-contrib
     ```

Verify installation:  
    ```bash
    psql --version
    ```

### Postgres Setup
```bash
#Update your postgres password
sudo passwd postgres

#Start Postgres Server
sudo service postgresql start

#Enter psql shell
sudo -u postgres psql

#Create & Connect to database
CREATE DATABASE wizword101;
\c wizword101

#Set Password and Verify
ALTER USER postgres PASSWORD 'password';
SELECT version();

#Leave psql
exit
```
### Goose Database Migrations
Install Goose:
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Determine your Connection String
```bash
'postgres://<username>:<password>@localhost:5432/wizword101?sslmode=disable'

#Test connection string
psql 'postgres://<username>:<password>@localhost:5432/wizword101'
```

Run Up Migrations
```bash
cd wizardingwordle101/sql/schema
goose postgres <connection_string> up

#Verify User and Spell tables exist
psql wizword101
\dt
```

### Create Configuration file
Create a config file in your root directory called `.w101config.json` with the following content:
`{"db_url":"postgres://<username>:<password>@localhost:5432/wizword101?sslmode=disable","current_user_name":""}` 

### 2. Initialization
Register yourself as a player to start playing:
```bash
#Navigate to the install directory
./WizardingWordle101 register <name>
```

As you discover and learn new spells in the spiral, add them to the database to increase the WizardingWordle101 pool.
```bash
./WizardingWordle101 spell <name> <cost> <school> <accuracy> <type>

#Some starter spells to get you started
./WizardingWordle101 spell "Fire Cat" 1 Fire 75 Attack
./WizardingWordle101 spell "Frost Beetle" 1 Ice 80 Attack
./WizardingWordle101 spell "Thunder Snake" 1 Storm 70 Attack
./WizardingWordle101 spell "Imp" 1 Life 90 Attack
./WizardingWordle101 spell "Dark Sprite" 1 Death 85 Attack
./WizardingWordle101 spell "Blood Bat" 1 Myth 80 Attack
./WizardingWordle101 spell "Scarab" 1 Balance 85 Attack
```

### 3. Play!
```bash 
./WizardingWordle101 play

WizWord101> start
```

## 📜 Usage

Available Dev commands:
```bash
#Register a new user
./WizardingWordle101 register <name>

User created successfully:
 * ID:      c5c9c70d-721c-42e8-82df-1df5ef80fe9b
 * Name:    S@m

#Change current user
./WizardingWordle101 login <name>

User switched successfully!

#List all users
./WizardingWordle101 users

* Sam 
* S@m (current)

#Add spell to database
./WizardingWordle101 spell <name> <cost> <school> <accuracy> <type>

./WizardingWordle101 spell "Dark Sprite" 1 Death 85 Attack
Spell saved successfully:
 * Name:    Dark Sprite
 * Cost:    1
 * SpellSchool:    Death
 * Accuracy:    85
 * SpellType:    Attack

#List all spells
./WizardingWordle101 spells

Name            | Cost   | School     | Accuracy | Type      
-------------------------------------------------------------
DeathBlade      | 0      | Death      | 100      | Charm     
LifeBlade       | 0      | Life       | 100      | Charm     
Vampire         | 4      | Death      | 85       | Drain     
Dark Sprite     | 1      | Death      | 85       | Attack  
...

#Start the game shell
./WizardingWordle101 play

 __      __.__                         ._____      __                .___.__         ___________  ____ 
/  \    /  \__|____________ _______  __| _/  \    /  \___________  __| _/|  |   ____/_   \   _  \/_   |
\   \/\/   /  \___   /\__  \_   __ \/ __ |\   \/\/   /  _ \_  __ \/ __ | |  | _/ __ \|   /  /_\  \|   |
 \        /|  |/    /  / __ \|  | \/ /_/ | \        (  <_> )  | \/ /_/ | |  |_\  ___/|   \  \_/   \   |
  \__/\  / |__/_____ \(____  /__|  \____ |  \__/\  / \____/|__|  \____ | |____/\___  >___|\_____  /___|
       \/           \/     \/           \/       \/                   \/           \/           \/     
WizWord101 > 

#Reset the database
./WizardingWordle101 reset

Database Reset Successfully!
```

Available Play commands:
```bash
#start a game of WizardingWordle101
WizWord> start

#Print description of Play Commands
WizWord> help

#Exit the Play shell
WizWord> exit
```

## 🪶 Contributing

### Clone the repo

```bash
git clone https://github.com/samassembly/WizardingWordle101.git
cd WizardingWordle101
```

### Build the compiled binary

```bash
go build
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.

Some ideas for future implemenation:
Setup/Initialization process script
Database initialization of spells for each Story Arc
Scoreboard for users to compete with each other
Server to generate daily word, clients on endpoints (Wordle-like)