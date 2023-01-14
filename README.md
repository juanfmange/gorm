# gorm
My gorm config to general use to mysql, sqlserver and postgres.
You must need to add the config toml for the env. For example

**mysql.toml

with:

[default]
- host
- port
- database
- user
- password
- charset
- sql_log 
- time_zone 
- MaxIdLeConns
- MaxOpenConns

