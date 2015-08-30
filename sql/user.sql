-- Create the user table 

begin transaction
DROP TABLE IF EXISTS Users

CREATE TABLE Users(Id INTEGER PRIMARY KEY, Email Varchar(50), Hashward TEXT, Rank Integer, Date )




