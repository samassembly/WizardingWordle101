# WizardingWordle101: Prove your mastery of spells from Ravenwood Academy!

Inspired by Wordle, guess a spell from the world of Wizard101. Narrow down your guesses by paying attention to the Cost, School, Accuracy, and Type attributes. If you can guess the secret spell within 5 guesses, you win!

As an avid enjoyer of Wizard101, I struggle to introduce new friends to the 10+ year old MMO due to the overwhelming amount of content the game has to offer. I decided to ease them into the game with a series of companion applications to help learn the game and its mechanics. WizardingWordle101 is designed to provide a fun daily challenge that reinforces knowledge of spells across all classes to better a user's combat potential. Understanding what spells might be cast at you is key to becoming a masterful practitioner of magic!

<img src="https://github.com/samassembly/WizardingWordle101/images/WizWordDemo" width="100%" alt="WizardingWordle101 demo image">

## Quick Start

### 1. Installation

### 2. Initialization

### 3. Play!
```bash 
#Navigate to the intstall directory
./WizardingWordle101 play

WizWord101> start
```

## Usage

Available Dev commands:
```bash
#Register a new user
./WizardingWordle101 register <name>
#Change current user
./WizardingWordle101 login <name>
#List all users
./WizardingWordle101 users

#Add spell to database
./WizardingWordle101 spell <name> <cost> <school> <accuracy> <type>
#List all spells
./WizardingWordle101 spells

#Start the game shell
./WizardingWordle101 play

#Reset the database
./WizardingWordle101 reset
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