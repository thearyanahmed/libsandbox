# Learning PostgreSQL
This is a sort of like a journal for postgres SQL, jotting down the common and basic quries/functions/features of postgres SQL.

So, **What is PostgreSQL?** PostgresSQL is an open-source relational database maangement system. 

Table of content
- 
- [Installing postgress](#installation)
- [Creating a user](#creating-a-user)
- [Database](#creating-a-database)
- [INSERT](#inserting-data)
- [UPDATE](#updating-a-record)
- [DELETE](#deleting-a-record)
- [Comparison operators](#comparison-operators)
- [Arithmatic operators](#arithmetic-operators)
- [ORDER BY](#sorting-or-ordering-data)
- [DISTINCT](#distinct-or-selecting-the-unique)
- [WHERE](#where)
- [OR](#or)
- [IN](#in)
- [BETWEEN](#between)
- [LIKE](#like)
- [ILIKE](#ilike)
- [LIMIT](#limit)
- [GROUP BY](#group-by)
- [HAVING](#having)
- [OFFSET](#offset)
- [MAX](#max)
- [MIN](#minx)
- [AVG](#avg)
- [SUM](#sum)
- [Alias](#alias)
- [COALESCE](#Coalesce)
- [DATE](#date)
- [Primary Keys](#primary-key)
- [UNIQUE](#unique)
- [Check constraints](#check-constraint)
- [Data types](#data-types)
- [Relations](#relations)
	- [INNER JOIN or JOIN](#inner-join)
	- [LEFT JOIN](#left-join)


### Installation
You can download the installer for your distribution [from here](https://www.postgresql.org/download/).

### Starting postgresSQL
To start/stop/control the server I'll be using **pg_ctl**. 
```
pg_ctl -D /usr/local/var/postgres start
```
I'm using a mac, so the directory **/usr/local/var/postgres** might vary from your machine.


### Using the CLI
psql is the PostgreSQL interactive terminal. To run the interactive shell, type
```
psql postgres 
```
It will promt a new cli. The **postgres** is a default database that came with the installation.

To check if everything is alright, we can check the version by typing
```
select version();
```

and the output should be similar to 

```sql
PostgreSQL 12.3 on x86_64-apple-darwin19.5.0, compiled by Apple clang version 11.0.3 (clang-1103.0.32.62), 64-bit
(1 row)
```
### Creating a user
syntax
```sql
CREATE USER name [ [ WITH ] option [ ... ] ]
```
options are 
```sql
    SYSID uid 
    | CREATEDB | NOCREATEDB
    | CREATEUSER | NOCREATEUSER
    | IN GROUP groupname [, ...]
    | [ ENCRYPTED | UNENCRYPTED ] PASSWORD 'password'
    | VALID UNTIL 'abstime'
```

Lets create a user who can login and create database and interact with them.
```sql
CREATE ROLE $username WITH LOGIN PASSWORD = 'password-goes-here';
```
example
```sql
CREATE ROLE plaban WITH LOGIN PASSWORD = 'secret';
```

### Creating a database

```
create DATABASE $dbName;

# example 
create DATABASE test_db;
```

the output should be
```
CREATE DATABASE
```

### Creating a database with specific owner
syntax
```sql
CREATE DATABASE $DB_NAME WITH OWNER $username;
```
example
```sql
CREATE USER project_db WITH OWNER plaban;
```

### Listing all databases
To list all database in the interactive shell, type ```\l```, it is not a SQL command so no need to use semi colons.

![Listing all databases](https://i.ibb.co/wrwdkDP/image.png)

### Switching databases
To swithc to database, type ```\c $dbname``` to start using the database. For example
```\c test_db```.
The output should be similar to 
```
You are now connected to database "test_db" as user "prophecy".
```
Again, this is not a SQL command, so no need to add semi-colons.

So now we are in our practice database, let's create some tables. We'll start with ```colours``` table.

### Deleting a database
```sql
DROP DATABASE $database_name;
```

### Connecting to database
The command to connect to a database is 
```bash
psql -h $host_ip_or_domain -p $port -U $username -W $dbname
# using -W (capital W) will ask for a password, if you don't have any password use -w 
```
So in my case, (password is ommited as my default password blank/I don't have a default password.)
```bash
psql -h localhost -p 5432 -U username dbname
```

### Creating a table
To create a table, our syntax is 
```sql
CREATE TABLE $tableName (
	$column_1 $data_type $optional_meta_information, 
	$column_2 $data_type $optional_meta_infromation,
	$column_N $data_type $optional_meta_information,
	$table_constraints
);
```
Lets start simple, 
So our command will be 
```sql
CREATE TABLE colours (
	id int PRIMARY KEY NOT NULL, 
	name varchar(20) NOT NULL, 
	hex varchar(7)
);
```
To know if it worked, we can check tables,  by typeing ```\dt``` while we are in the database.
![List of all tables](https://i.ibb.co/L5Gwg2t/image.png)

### Inserting data
As of now, we have our table and we can store some data in it. To store data, we use a **insert** query. The insert query is structured as follows 
```sql
INSERT INTO $table_name ($column_1, $column_2, ... $column_N) VALUES ($value_of_column_1, $value_of_column_2, ... $value_of_column_N);
```
This will insert a single row. If we want to insert multiple rows at once, we can append to the values.

```sql 
INSERT INTO $table_name ($column_1, $column_2, ... $column_N) VALUES ($row_1_column_1, $row_1_column_2, ... $row_1_column_N),($row_2_column_1, $row_2_column_2, ... $row_2_column_N),($row_N_column_1, $row_N_column_2, ... $row_N_column_N);
```
#### Example 
So for creating a single row 

```sql
INSERT INTO colours (id, name) VALUES (1, 'red');
```
For inserting multiple rows 
```sql
INSERT INTO colours (id, name, hex) VALUES (2, 'green',''), (3, 'blue','3399ff');
```

The output should be ```INSERT 0 2```  or similar.

To see if everything worked correctly, we can list all our records from our table by using the ```SELECT``` command.
```sql
SELECT * FROM $table_name;
```
So our command will be,
```sql
SELECT ALL FROM colours;
```
The result should be a list data.
![Selecting all records from our colours table.](https://i.ibb.co/VJgLCGB/image.png)
Now, with the ``*`` , we have selected all the columns. To select a specific column, we can specificy them in place of asterisk.
```sql
SELECT column_a, column_b FROM $table_name;
``` 
So our example will be
```sql
SELECT name, hex from colours;
```
![Returning specific columns](https://i.ibb.co/qMTRsHQ/image.png)


### Selecting record/s
syntax
```sql
SELECT * FROM $table_name;
# or, if you want to be specific with column names.
SELECT $column_name_a, $column_name_b,... $column_name_n FROM $table_name;
```
example
```sql
SELECT * FROM cars;
```
result
![enter image description here](https://i.ibb.co/x7L7vfK/image.png)

### Updating a record
syntax 
```sql
UPDATE $table_name SET $column_name = $value, $another_column = $another value, ... $nth_column_name = $some_value WHERE $condition;
```
example
```sql
UPDATE users SET name = 'Aryan Ahmed Anik' WHERE email='thearyanahmed@gmail.com' OR email='theprophecy@gmail.com';
```
**Warning** If you don't provide a where clause or a always true where clause, the operation will be performed against all the data on the table.

### Deleting a record
syntax
```sql
DELETE FROM $table_name WHERE $condition;
```
example
```sql
DELETE FROM users WHERE id > 10;
```

**Warning** If you don't provide a where clause or a always true where clause, the operation will be performed against all the data on the table.
```sql 
DELETE FROM users WHERE 1;
# or 
DELETE FROM users; 
```
All the records will be deleted.


### Comparison operators
|syntax| what it means  | example |
|--|--|--|
| = | equals to | id = 2 |
| > | greater than | id > 2 |
| < | less than | id < 10 |
| >= | greater than or equals to | id >= 10 |
| <= | less than or equals to | id <= 100 |
| <> | not equals to | id <> 102 |
| \|\| | contact strings | 'hello ' \|\| 'world' |
| !!= | not in | 3 !! = i |
| ~~ | like | 'scrappy,marc,hermit' ~~ '%scrappy%' |
| !~~ | not like | 'bruce' !~~ '%al%' |
| ~ | match (regex), case sensitive | 'thomas' ~ '\*.thomas\*.' |
| ~* | match (regex), case insensitive | 'thomas' ~\* '\*.thomas\*.' |
| !~ | Does not match (regex), case sensitive | 'thomas' !~ '\*.Thomas\*.'|
| !~*| Does not match (regex), case insensitive | 'thomas' !~ '\*.vadim\*.' |

### Arithmetic Operators
|syntax| what it means  | example |
|--|--|--|
| + | addition | SELECT 2 + 5 |
| - | subtraction | SELECT 5 -2 |
| / | division | 10 / 2 |
| * | multiplication | 5 * 3 |
| % | modulo | 12 % 5 |
| % | truncate | % 4.5 |
| ! | factorial | 3! |
| !! | factorial (left operation) | !! 5 |
| : | natural Exponentiation | : 3.0 |
| ; | natural Logarithm | (; 5.0) |
| @ | absolute value | @ -5.0 |
| \|/ | square root | \|/ 25.0 |
| \|\|/ | cubic root | \|\|/ 27.0 |


### Sorting or Ordering data
To order our data by a column in a ascending order,
```sql
SELECT * FROM $table_name ORDER BY $column_name $ASC_OR_DESC;
```
so to order our data by name, 
```sql
SELECT * FROM users ORDER BY name ASC;
```
Our result looks like this,
![user's list](https://i.ibb.co/sw5ZScQ/image.png)
As you see the id doesn't add up to the next row's id.  If we try it in descending order, 
![enter image description here](https://i.ibb.co/JrPsxRh/image.png)

#### Combining multiple columns in sorting

```sql
SELECT * FROM $table_name ORDER BY $column_1, $column_2, ... $column_N $ASC_OR_DESC;
```
for example 
```sql
SELECT * FROM users ORDER BY id, name ASC;
```
**Note:** If you do not specify the order (ascending or descending), the default will be Ascending.
### Distinct or selecting the unique
```sql
SELECT DISTINCT $column_name from $table_name;
```
### WHERE 
```sql
SELECT * FROM users WHERE $column = $value;
```
For adding more than 1 condition, we use **AND**.
```sql
SELECT * FROM users WHERE $column = $value AND $another_column = $some_value AND ... $some_other_column_n = $some_value;
```
for example 
```sql
SELECT * FROM users WHERE id > 30 and gender = 'Male' ORDER BY name ASC;
```
And we have our results.
![enter image description here](https://i.ibb.co/NjM1VHj/image.png)

### OR 
syntax 
```sql
SELECT * FROM users WHERE $column = $value OR $column = $value; # columns can be same or different.
```
example.
```sql
SELECT * FROM users WHERE country = 'Bangladesh' OR country = 'Singapore' OR country = 'France';
```

### IN 
If you are quering in a single column, you can shorten the code by using in keyword.
syntax 
```sql
SELECT * FROM users WHERE country IN ('Bangladesh', 'Singapore','France'); 
```
**Note :** You can obviously use multiple *OR*, even with *IN* clause.

### BETWEEN
The between keyword is used to select data from a range.
syntax 
```sql
SELECT * FROM $table_name WHERE $column_name BETWEEN $starting_point AND $ending_point;
```
example
```sql
SELECT * FROM users WHERE id BETWEEN 1 AND 10;
```

### LIKE
Like is used to match patterns.
syntax
```sql
SELECT * FROM $table_name WHERE $column_name LIKE $pattern;
```
So, lets query out all the users who has 'United' in their country name.
query
```sql
SELECT name, country FROM users WHERE country LIKE '%United%' order by name;
```
![enter image description here](https://i.ibb.co/GWJgmrP/image.png)

#### patterns 

|like| description  |
|--|--|--|
| WHERE CustomerName LIKE 'a%' | Finds any values that start with "a"|
| WHERE CustomerName LIKE '%a'| Finds any values that end with "a"|
| WHERE CustomerName LIKE '%or%' | Finds any values that have "or" in any position|
|WHERE CustomerName LIKE '_r%' |Finds any values that have "r" in the second position|
|WHERE CustomerName LIKE 'a_%'|Finds any values that start with "a" and are at least 2 characters in length|
|WHERE CustomerName LIKE 'a__%'|Finds any values that start with "a" and are at least 3 characters in length|
|WHERE ContactName LIKE 'a%o'|Finds any values that start with "a" and ends with "o"|

**Note to self** underscores are used to match single characters.
### ILIKE
ILIKE is used to search with case insensitivity.

### LIMIT
syntax 
```sql
SELECT * FROM users LIMIT $limit;
```
example
```sql
SELECT * FROM users LIMIT 10;
```

![limiting](https://i.ibb.co/m6m8fp8/image.png)

**Note to self** Limit is not a SQL standard. We can achieve the same result by using **FETCH**.
syntax 
```sql
... FETCH FIRST N ROWS; # n is positive integer, can be ommited if n = 1.
```

### OFFSET
We often need to select records after a specific number of results. We might need to select 10 items from 30th. Meaning 30th to 40th elements of the results.Pagination is a good example of it.
If such scenarios, *OFFSET* comes into play. 
syntax 
```sql
SELECT * FROM $table_name OFFSET $n; # n = positive integer.
```
example
```sql
SELECT * FROM users OFFSET 5, LIMIT 10 ORDER BY id;
```

![enter image description here](https://i.ibb.co/Bj30Fzx/image.png)


### GROUP BY
syntax 
```sql
... GROUP BY $column_name;
```

example 
```sql
SELECT country, COUNT(*) FROM users GROUP BY country;
```
![enter image description here](https://i.ibb.co/YXCpp3V/image.png)

### HAVING
Having works with group by, adds up an extra layer of filter.
example
```sql
SELECT country, count(*) FROM users GROUP BY country HAVING COUNT(*) > 5;
```
This will return all the country's where the count of people are greater than 5.
**Note** having must come before *order by*.

### MAX 
syntax 
```sql
SELECT MAX($column_name) FROM $table_name;
```
example
```sql
SELECT MAX(price) FROM cars;
```
It will return a single row, with the max priced value.

### MIN 
syntax 
```sql
SELECT MIN($column_name) FROM $table_name;
```
example
```sql
SELECT MIN(price) FROM cars;
```
It will return a single row, with the minimum priced value.

### AVG
syntax 
```sql
SELECT AVG($column_name) FROM $table_name;
```
example
```sql
SELECT AVG(price) FROM cars;
```
It will return a single row, with the average priced value.

### SUM 
syntax 
```sql
SELECT SUM($column_name) FROM $table_name;
```
example
```sql
SELECT SUM(price) FROM cars;
```
It will return a single row, with the sum of the price column.

### Alias
You can use the `as` keyword followed by a name or alias you want to use.
syntax
```sql
SELECT $column_name AS $alias FROM $table_name;
```
example
```sql
SELECT name as car_name FROM cars;
```
It will output the name column as car_name, the values will remain the same.

### Coalesce 
Coalesce function gives us the benefit to add a default value if the first value does not exists.
syntax
```sql
SELECT COALESCE($first_value, $default_value);
```
you can have multiple values to check, and if it does not exist, it will fall back to the next default value.
For example,
```sql
SELECT COALESCE($first_value, $second_value_if_first_value_is_null, $third_value_if_second_value_is_null,...,$nth_value_if_(n-1)th_value_is_null);
```
### Date
Get the current time 
```sql
SELECT NOW();
```
Get current time
```sql
SELECT NOW()::TIME;
```
Get current date
```sql
SELECT NOW()::DATE;
```

#### calcuating dates
Subtracting date
```sql
SELECT NOW() - INTERVAL '10 YEARS'; # should return the date from 10 years ago.
```
Adding days
```sql
SELECT NOW() + INTERVAL '10 MONTHS'; # should return the date after 10 months.
```
#### extracting data
Extracting can be done by the EXTRACT function.
```sql
SELECT EXTRACT(DAY FROM NOW());
```

### PRIMARY KEY
Primary keys are unique identifiers accross all the data of any table. Primary key is not a must have column. This does not hold any value outside of the database. It is unique everytime its generated. Even if we delete a key, the next one will not be the same.
#### Adding primary key
```sql
ALTER TABLE $table_name ADD PRIMARY KEY ($key1,$key2,...$keyN);
```
example
```sql
ALTER TABLE cars ADD PRIMARY KEY (id);
```
So, why did our syntax say **KEY (\$key1,\$key2,.. \$keyN);**.  Well because, we often need a unique key based on two or multiple columns. Those are called **composite keys**.

#### Dropping primary key by altering table
syntax
```sql
ALTER TABLE $table_name DROP CONSTRAINT $primary_key_constraint;
```
so if we want to drop cars table's primary key constraint,
```sql
ALTER TABLE cars DROP CONSTRAINT cars_pkey;
```
If you are using numeric column as primary key, it will auto increment by 1 by default.

### UNIQUE
we have unique constraint to make a record unique accross the whole table. For example, we can add unique constraint to the email to make sure there are no duplicate emails.
syntax
```sql
ALTER TABLE $table_name ADD CONSTRAINT $identifier UNIQUE ($column_name);
```
example 
```sql
ALTER TABLE cars ADD CONSTRAINT unique_email_address UNIQUE (email);
```
### CHECK CONSTRAINT
syntax
```sql
ALTER TABLE $table_name ADD CONSTRAINT $constrant_name CHECK ($logic);
```
example
```sql
ALTER TABLE users ADD CONSTRAINT gender_constraint CHECK (gender = 'Male');
```

### Data types
PostgresSQL supports a various number of data types. You can find them [here in the PostgresSQL supported data types document.](https://www.postgresql.org/docs/9.5/datatype.html)

## Relations

What are relations and what are joins? Well, in RDMSs, we work with structured relation data. Meaning, One data might be related to another. In real life scenario, in a blogging platform, an author can have multiple articles. So, author and aritcle are related. Authors have `1 to many` relation with articles, as a author may have multiple articles. But an article is realted to an author in `1 to 1` relation. Cause A article can belong to only one author (in our scenarion, if your application has the feature where an article can be collarboratedly writted by mulitple authors, like google docs, the relation will be 1 to many as well). 

We primarily have 
- 1 to 1 
- 1 to many
- many to many 
- polymorphic

#### one to one
One to one is when the entity A can have only 1 association of entity B. Sort of like mother-child relation. A child can have one mother.

#### one to many
In this case, an entity might have multiple related entities. Like a mother can have multiple children.

#### many to many 
In this case, a any entity can have multiple assoicated entites, from both end. For example, a unversity student have many courses and each course can be enrolled by many students. 
#### polymorphic 
Till now, we have talked about realted entities in two tables or 3 (in case of many to many) but the number of tables are fixed, but what if, we have our related model but the number of tables are not fixed and is dynamic. 
For example, in any social media now days,

you can lets say comment on a post, a photo, album, life event etc. These `posts`, `photos`, `albums`, `life events` etc are different tables, but every comment shares the same attributes, a comment body itself, an author / user. In these scenraios, we don't make post_comments, photo_comments, album_comments multiple `$_comments` table with the same structure, we create a single comments table with the reference to the table. Our schema will have the following attributes

```
id
comment_body
source_id
source_type // posts, photos,albums etc
```

So a row with `soruce_id` 1 `source_type` posts is a comment to a post where the id of the post is 1. A row with `soruce_id` 1 `source_type` albums is a comment to an `album` with the id of 1.

### Inner join

Inner join takes the values that are common/present to both table.
![inner join diagram.](https://images.squarespace-cdn.com/content/v1/5732253c8a65e244fd589e4c/1464122775537-YVL7LO1L7DU54X1MC2CI/ke17ZwdGBToddI8pDm48kMjn7pTzw5xRQ4HUMBCurC5Zw-zPPgdn4jUwVcJE1ZvWMv8jMPmozsPbkt2JQVr8L3VwxMIOEK7mu3DMnwqv-Nsp2ryTI0HqTOaaUohrI8PIvqemgO4J3VrkuBnQHKRCXIkZ0MkTG3f7luW22zTUABU/image-asset.png)

syntax 
```sql
SELECT $columns FROM $first_table JOIN $second_table ON $first_table.joining_column = $second_table.reference_column;
```
Suppose we have a 1 to 1 relationship, where each user can have a car. A column `car_id` in users table points to the `id` column in the cars table.  So our query would be,
```sql
SELECT * FROM users JOIN cars ON users.car_id = cars.id;
```
We have 4 users, 1 user doesn't have a car assocaited with them.
![enter image description here](https://i.ibb.co/DQXrzrS/image.png)
### Left join
Left join selects all the records from the left table even if they don't have a related record in the right table. 
![enter image description here](https://images.squarespace-cdn.com/content/v1/5732253c8a65e244fd589e4c/1464122797709-C2CDMVSK7P4V0FNNX60B/ke17ZwdGBToddI8pDm48kMjn7pTzw5xRQ4HUMBCurC5Zw-zPPgdn4jUwVcJE1ZvWEV3Z0iVQKU6nVSfbxuXl2c1HrCktJw7NiLqI-m1RSK4p2ryTI0HqTOaaUohrI8PIO5TUUNB3eG_Kh3ocGD53-KZS67ndDu8zKC7HnauYqqk/image-asset.png)

syntax 
```sql
SELECT $columns FROM $first_table LEFT JOIN $second_table ON $first_table.joining_column = $second_table.reference_column;
```
```sql
SELECT * FROM users LEFT JOIN cars ON users.car_id = cars.id;
```
![enter image description here](https://i.ibb.co/Ykdn93Y/image.png)

### Selecting specific columns
We can prepend the table name with a dot to be more speicifc query. Instead of `SELECT name from users` we can write `SELECT users.name FROM users`, being more specific.
**Why is this necessary?** Often we have same named columns in multiple tables, for example, a user can have name, a company table can have name, a cars table can have name. You get the idea. And also, when **JOINING** , we can be more specific instead of selecting all columns.  
So, a inner join would be

```sql
SELECT users.name, cars.name FROM users JOIN cars ON users.car_id = cars.id; # we are only selecting user's name and the name of the car they own.
```
The doc will be updated with time.

**Note to self** We don't need to memorize every bit but build up the queries from a simple statement and adding up more and more. 


