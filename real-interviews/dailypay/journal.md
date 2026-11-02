DailyPay Phone Screen
Setting
Phone Call
Thursday, April 4 @ 12:30 pm EST
Drew Messner
Agenda
Overview of DailyPay
Growth projecting this year
650 employees, looking to ad 350-400
Structured growth, not looking for layoffs
On demand paycheck is a major market
Role itself (SE2)
MId-level engineer, Money in and money out team
Publishing payment APIs, connecting with clients to get info and sending wherever it is needed internally
Building and consuming RESTful APIs
Kafka quite a bit
Tech stacks change, moving from Ruby on Rails → Golang
Speed performance benefits of Golang
Questions about me
Interview Process
David and Ryan at the same time
David eng manager
Ryan Senior manager bove Dave
2 technicals calls
DSAa
System design
APIs
With a senior engineer
Timeline

Q/A

DailyPay Hiring Manager Round (Behavioral)

Setting
Fri Apr 12, 2024 12pm – 12:45pm (EDT)
Dave Cunningham
Overview of DailyPay
Ruby monolith → golang
Ledger and payment
Ledger ruby → go
Payment new microservice
Questions
Senior Hiring Engineer
Rundown on your resume
Background
Go, APIs, Kafka
Hiring Engineer
Can you tell me about something you’re really proud of?
How do you go about researching?
Tell me about a time something didn’t go right?
Do you have any questions?
How do you guys go about onboarding engineer
Consuming APIS
So what’s next
Me and Ryan discuss
REcruiter reach out
Tech challenge
Some sort of design challenge
Questions

DailyPay LeetCode Round
Setting
Date/Time: Apr 25, 2024 3:00pm-4:00pm EDT
Interviewers: Sebon Koo, Carl Smith
Zoom call
Sebon was a staff
Carl was a senior

Resume interview
What do you do at matchbook
Elevator pitch for matchbook
My journey as a data engineer, complex world of sql server
Transition to software due to limitations of a database
Talked about Golang
Mentioned the router I use, why I like it, light weight
Mentioned that gorilla/mix is deprecated
Asked if I used frameworks in go
Used libraries
No frameworks, maybe go-Kafka would be a framework
Asked about Kafka,
High level research, have poc at work to implement
Not expert
Asked what they use kafka
Said a client recommended it and its a popular stack
Have you done research into other technologies instead of kafka

Whiteboarding
Hackerpad
Two sum question, return Boolean
Got this correct
Pretended I didn’t know, instead of naming complement I named it diff
Did have a mistake initially but fixed that
Gave time and space complexity
Solution passed
Agreed it was optimal
Expanded on the problem by asking me to return tuples of pairs in a list
Couldn’t run this one but we all agreed my second solution looked correct
The turned it into system design and asked me what to do if you wanted to make this a rest api
Listed endpoints
POST api/v1/twosum
Body with nums and target
Response with status code and json with transaction object status object and data object
GET api/v1/twosum/{transactionId}
Follow up on system design, asked what id do in a processing case
Went through the whole flow
On the initial POST I’d kick off the twosum method async
Store transactionId and metadata somewhere
Respond to the client with a processing status
When the async method finishes, store results with metadata in db
The client can keep requesting with the GET and transactionId
Then I broke down the more specific code
A server with a router in it
Router sets up the controllers that correspond to each route
Controller orchestrate business logic
Business later does the real code
Db layer handles db work  
Sebon said the System Design stuff looked “great”

Questions
What is your experience with mentorship and growth, mentioned the lack of mentorship at my startup
Carl said he thinks you learn best by watching good things happen around you and he thinks a lot of good stuff is currently happening at DailyPay to learn from
Sebon said he thinks good stuff is happening today, not so much a few years ago
Do they both work on ledgers project or are they building stuff from scratch
I didn’t fully listen here, but during this part he said the Ruby on Rails repository kinda sucks but it’s great experience to learn from that codebase (I think it’s a good sign he said this because it means he might envision me working on it)

DailyPay System Design Round
Setting
Date/Time: May 8, 2024 11:00am-12:00pm EDT
Derek Rushing, Annette McCullough (Principal Software Engineers) (Annette ended up bailing)
Brendan Ko (Senior Software Engineer) (replaced Annette)

DailyPay Offer Call

Setting
5/13/2024 @ 5:32 pm
Call lasted 10 minutes
Drew Messner
What happened
The NYC people team is on an offsite and he had to do some stuff to get them to respond
said he had good news
The team wants to move forward with me
Said I had a great scorecard
I did so well they opened up 2 more positions
SWE2
140K Base
10k sign on
10% bonus
5000 options, Strike price is $14 4 year vest 1/36th monthly
I said I was happy with the base and sign on
I asked if anything changes if I moved to NYC and he said we’d figure that out when/if it happens
I asked if the stock was options or RSUs, it is options (said that’s pretty commmon from private companies)
I asked about the bonus, and he said that was performance enhanced

    Said I was a front-runner candidate
    They had to move some things around to make it work
    Asked if I was going to sign, I said I’m definitely going to sign as long as the contract looks good
    June 3rd start date
