# WoW Query update tool

This tool was created to maintain wow-query.dev databse synchronized with the API.

This only works with Postgres and will create its own tables and schemas (requires advanced privileges)

# How it works

This will iterate over ALL API endpoints and break data into a relational database.
It only fetch **static** data.  

There are 3 types of tasks:

#### 1. Index task
Used to fetch data from endpoints that expose a `/index` path.

This will fetch the index, iterate over its result and call a function for each item.

#### 2. Range task
Used to fetch data from endpoints such as items and spell.

This will iterate over a large range to cover all possible URIs.

#### 3. Media endpoints

Used to fetch media information for previously imported data.

This will iterate over records already on the database like items and spell.

# Warning

This will perform a VERY LARGE AMOUNT OF REQUESTS.

That means it will eventually get a 429 (too many requests) response and halt execution for 1 hour.

Current API limitation is 36k requests per hour, but it might execute more than this limit on some cases.

Released for educational purposes only. Use at your own risk.