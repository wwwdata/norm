# norm
[Neoism](github.com/jmcvetta/neoism) is agreat client for neo4j. However if you are doing some basic CRUD stuff
with neo4j you end up writing code that is repeating itself a lot. I think this sucks, So I will try to implement
a very basic ORM.

## Planned features
- Save and Load a basic node which represents a struct without relationship and a label
- Define Relationships of a node
- Load Relationships and related nodes.
- Limits, Offsets
- Somehow an ability to write a custom cypher query to do crazy shit, but serialize the results into the previously defined structs
