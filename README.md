# Gator A CLI to manage rss feeds

## How to install
```bash
go install github.com/saniak-hub/gator
```
Requiremets 
1. Go and Postgresql installed in your working environment

## Build the program
```bash
go build
```
This will generate a executable called gator


## To register
```bash
./gator register username
```
Register username

## To login
```bash
./gator login username
```
Login to start using the cli



## To reset
```bash
./gator reset
```
Remove all the data in the databases

## To users
```bash
./gator users
```
List all the register users

## To agg
```bash
./gator agg <time between requests>
```
This command recursively stores the items in a rss feed to a poststable. 
It requres one positional argument <time between requests>, which is the time it will pause before scraping a new site

## To addfeed
```bash
./gator addfeed url
```
Adds the feed to the database for persistence  


## To feeds
```bash
./gator feeds
```
Lists all the feeds stored.  
It prints The name of the feed and the url to the feed site  

## To follow
```bash
./gator follow url
```
Requires one positional argument which is the url to the site you want to follow

## To unfollow
```bash
./gator unfollow url
```
Use this to unfollow a site

## To following
```bash
./gator following
```
This command lists the sites you follow

## To browse
```bash
./gator browse [limit]
```
Prints Title and the link of the posts to the screen 
the limit is optional, which defaults to 2


