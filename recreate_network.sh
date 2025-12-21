#!/bin/bash
set -e
go build -o ./bin/client ./client
./bin/client reset

# Create people
./bin/client create-person "John Doe" 30 1
./bin/client create-person "Jane Doe" 25 2
./bin/client create-person "Alice Smith" 28 3
./bin/client create-person "Bob Johnson" 35 4
./bin/client create-person "Charlie Brown" 22 5
./bin/client create-person "Diana Prince" 29 6
./bin/client create-person "Eve Wilson" 27 7
./bin/client create-person "Frank Miller" 32 8
./bin/client create-person "Grace Lee" 26 9
./bin/client create-person "Henry Davis" 31 10
./bin/client create-person "Ivy Chen" 24 11
./bin/client create-person "Jack Taylor" 33 12
./bin/client create-person "Kate Anderson" 28 13
./bin/client create-person "Liam O'Brien" 29 14
./bin/client create-person "Mia Rodriguez" 25 15
./bin/client create-person "Noah Kim" 27 16
./bin/client create-person "Olivia Martinez" 30 17
./bin/client create-person "Paul Thompson" 34 18
./bin/client create-person "Quinn White" 23 19
./bin/client create-person "Rachel Green" 26 20
./bin/client create-person "Sam Wilson" 31 21
./bin/client create-person "Tina Park" 28 22
./bin/client create-person "Uma Patel" 29 23
./bin/client create-person "Victor Chen" 32 24
./bin/client create-person "Wendy Zhang" 27 25

# Create follow relationships
./bin/client follow-person 1 2
./bin/client follow-person 2 1
./bin/client follow-person 1 3
./bin/client follow-person 2 3
./bin/client follow-person 3 1
./bin/client follow-person 3 2
./bin/client follow-person 1 4
./bin/client follow-person 4 1
./bin/client follow-person 2 5
./bin/client follow-person 5 2
./bin/client follow-person 3 6
./bin/client follow-person 6 3
./bin/client follow-person 4 7
./bin/client follow-person 7 4
./bin/client follow-person 5 8
./bin/client follow-person 8 5
./bin/client follow-person 6 9
./bin/client follow-person 9 6
./bin/client follow-person 7 10
./bin/client follow-person 10 7
./bin/client follow-person 8 11
./bin/client follow-person 11 8
./bin/client follow-person 9 12
./bin/client follow-person 12 9
./bin/client follow-person 10 13
./bin/client follow-person 13 10
./bin/client follow-person 11 14
./bin/client follow-person 14 11
./bin/client follow-person 12 15
./bin/client follow-person 15 12
./bin/client follow-person 13 16
./bin/client follow-person 16 13
./bin/client follow-person 14 17
./bin/client follow-person 17 14
./bin/client follow-person 15 18
./bin/client follow-person 18 15
./bin/client follow-person 16 19
./bin/client follow-person 19 16
./bin/client follow-person 17 20
./bin/client follow-person 20 17
./bin/client follow-person 18 21
./bin/client follow-person 21 18
./bin/client follow-person 19 22
./bin/client follow-person 22 19
./bin/client follow-person 20 23
./bin/client follow-person 23 20
./bin/client follow-person 21 24
./bin/client follow-person 24 21
./bin/client follow-person 22 25
./bin/client follow-person 25 22

# Additional cross-connections
./bin/client follow-person 1 10
./bin/client follow-person 1 15
./bin/client follow-person 1 20
./bin/client follow-person 2 12
./bin/client follow-person 2 18
./bin/client follow-person 3 7
./bin/client follow-person 3 14
./bin/client follow-person 4 9
./bin/client follow-person 4 16
./bin/client follow-person 5 11
./bin/client follow-person 5 19
./bin/client follow-person 6 13
./bin/client follow-person 6 21
./bin/client follow-person 7 15
./bin/client follow-person 7 22
./bin/client follow-person 8 17
./bin/client follow-person 8 23
./bin/client follow-person 9 20
./bin/client follow-person 9 24
./bin/client follow-person 10 25
./bin/client follow-person 11 1
./bin/client follow-person 12 2
./bin/client follow-person 13 3
./bin/client follow-person 14 4
./bin/client follow-person 15 5
./bin/client follow-person 16 6
./bin/client follow-person 17 7
./bin/client follow-person 18 8
./bin/client follow-person 19 9
./bin/client follow-person 20 10
./bin/client follow-person 21 11
./bin/client follow-person 22 12
./bin/client follow-person 23 13
./bin/client follow-person 24 14
./bin/client follow-person 25 15