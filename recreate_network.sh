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

# Create posts
./bin/client post 1 1 "Just finished reading an amazing book about distributed systems!"
./bin/client post 1 2 "Working on a new project. Excited to share more soon!"
./bin/client post 2 3 "Beautiful sunset today 🌅"
./bin/client post 2 4 "Coffee and code - the perfect morning combination"
./bin/client post 3 5 "Attended an interesting conference today. Learned so much about microservices architecture and how they can scale horizontally. The key takeaways were about service discovery, load balancing, and maintaining consistency across distributed systems."
./bin/client post 4 6 "Weekend vibes ✨"
./bin/client post 4 7 "New recipe experiment: homemade pasta from scratch. It turned out better than expected!"
./bin/client post 5 8 "Just discovered a new hiking trail nearby. The views were absolutely breathtaking!"
./bin/client post 5 9 "Reflecting on the importance of work-life balance. It's something we all struggle with, but finding that equilibrium is crucial for long-term happiness and productivity."
./bin/client post 6 10 "Movie night recommendations?"
./bin/client post 6 11 "Finished watching an incredible documentary series. Highly recommend!"
./bin/client post 7 12 "Today I learned about graph databases and their applications in social networks. The way relationships are modeled is fascinating - nodes represent entities, edges represent relationships, and you can traverse connections efficiently. This is particularly powerful for recommendation systems and analyzing network structures."
./bin/client post 8 13 "Morning run complete! 🏃"
./bin/client post 8 14 "Planning a trip to the mountains next month"
./bin/client post 9 15 "New music discovery: found an amazing indie artist"
./bin/client post 9 16 "The evolution of programming languages over the decades has been remarkable. From assembly to high-level languages, each generation brought new abstractions and paradigms. Functional programming, object-oriented design, and now reactive programming - each approach solves different classes of problems."
./bin/client post 10 17 "Weekend project: building a small garden"
./bin/client post 10 18 "Fresh vegetables from the garden taste so much better!"
./bin/client post 11 19 "Book recommendation: just finished an excellent novel"
./bin/client post 11 20 "The art of storytelling in modern literature has evolved significantly. Contemporary authors are experimenting with narrative structures, unreliable narrators, and non-linear timelines. This creates more engaging and thought-provoking reading experiences."
./bin/client post 12 21 "Tech stack decisions are always interesting. Choosing between different frameworks, databases, and architectures requires balancing performance, maintainability, team expertise, and long-term scalability. There's rarely a one-size-fits-all solution."
./bin/client post 12 22 "Working on improving my photography skills"
./bin/client post 13 23 "Beautiful day for a walk in the park"
./bin/client post 13 24 "Cooking experiment: trying a new cuisine tonight"
./bin/client post 14 25 "The importance of mental health awareness cannot be overstated. Creating safe spaces for conversations, reducing stigma, and providing accessible resources are all crucial steps toward a healthier society."
./bin/client post 14 26 "Weekend reading session"
./bin/client post 15 27 "New workout routine is showing results!"
./bin/client post 15 28 "Sustainable living practices are becoming increasingly important. From reducing waste to choosing renewable energy sources, every small action contributes to a larger impact on our planet's future."
./bin/client post 16 29 "Art gallery visit was inspiring"
./bin/client post 16 30 "Working on a creative writing project"
./bin/client post 17 31 "The future of remote work is fascinating. Companies are rethinking office spaces, collaboration tools, and team dynamics. The hybrid model seems to be emerging as a popular solution, balancing flexibility with in-person connection."
./bin/client post 17 32 "Learning a new language - challenging but rewarding"
./bin/client post 18 33 "Weekend hiking adventure planned"
./bin/client post 18 34 "Nature photography is so peaceful"
./bin/client post 19 35 "The intersection of technology and art creates incredible possibilities. Digital art, interactive installations, and AI-generated content are pushing boundaries and redefining what we consider artistic expression."
./bin/client post 19 36 "Attending a local music festival this weekend"
./bin/client post 20 37 "Morning meditation routine is life-changing"
./bin/client post 20 38 "Reflecting on personal growth and development. It's important to regularly assess our goals, values, and progress. Self-awareness is the foundation of meaningful change."
./bin/client post 21 39 "New restaurant discovery - amazing food!"
./bin/client post 21 40 "Weekend project: home improvement"
./bin/client post 22 41 "The science of habit formation is fascinating. Understanding how neural pathways are created and strengthened helps explain why some behaviors become automatic while others require constant effort. Building good habits is about creating the right systems and environment."
./bin/client post 22 42 "Book club discussion was thought-provoking"
./bin/client post 23 43 "Exploring new neighborhoods in the city"
./bin/client post 23 44 "Weekend farmers market finds"
./bin/client post 24 45 "The evolution of social media platforms has dramatically changed how we communicate and share information. From simple status updates to complex content ecosystems, these platforms continue to shape our social interactions and information consumption patterns."
./bin/client post 24 46 "Working on a side project - excited about the progress"
./bin/client post 25 47 "Morning coffee and journaling - perfect start to the day"
./bin/client post 25 48 "The concept of deep work and focused attention is becoming increasingly valuable in our distraction-filled world. Creating boundaries, eliminating interruptions, and dedicating blocks of time to meaningful work can significantly improve both productivity and satisfaction."
./bin/client post 1 49 "Update on the project: making great progress!"
./bin/client post 3 50 "Sharing some thoughts on software architecture patterns. The choice between monolithic and microservices architectures depends on many factors including team size, complexity, deployment frequency, and organizational structure. Each approach has its place."
