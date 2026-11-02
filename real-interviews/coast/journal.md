## Context

- Tamara Zaman (recruiter) reached out on November 7, 2025
- Anton Maximov (CTO) reached out on November 13, 2025
- Modern Commercial Charge Card and Expense Management Platform
- Never applied
- Series B, Venture Capital backed, Not Profitable
- Job Application: https://job-boards.greenhouse.io/coast/jobs/4000170004
- Podcast: https://fintechlayercake.buzzsprout.com/1985266/episodes/14997028-fleet-card-masterclass-with-coast-founder-daniel-simon

## Phone Screen (Phone Call) on Friday, November 17

- She explained what Coast did and what fleet management is
  Helping companies track their expenses with a card that has modern features for administrating over a company’s fleets’ expenses
- Tamara asked standard interview questions
  - Why are you interested in Coast?
  - Background, DailyPay focused
  - Biggest project at DailyPay
- Asked compensations
  - I said I currently make 180k base, and also other bonus and equity structure
  - Tech Stach is Python/Typescript, Postgres, AWS

## 1:1 Zoom w/ Anton Maximov on Tuesday, November 25

- In responding to the CTO’s email, he offered a chat so I scheduled a call with him
- Spoke with Anton who seemed to be at their office, quick 15 minute discussion
- He started by giving background on Coast
  - Modern software approach to cost management software
  - Provides specific features for managing a fleet
  - Tech Stack is…
    - Containers (microservices) and serverless (lambdas)
    - Python (old) and Typescript full stack (new)
    - React frontend
    - Postgres and DynamoDB
    - AWS, ECS, AppConfig and Feature Flags
  - 30 engineers currently, looking to hire 12 more in the next year
  - About 100 employees total and growing
  - Not profitable, Series B, Venture Capital backed
  - Looking to move into expense management and bill pay
- I asked a few questions
  - Does your deployment setup include ArgoCD or anything?
  - Value proposition of their card?
    - The software is the value proposition, modern features like mileage usage, etc.
  - How do your teams operate? Sprint structure? Many deployments?
    - Many deployment, up to 5 a day
    - Modern agile
  - Are you guys Microservices, Monolith?
    - Basically said that “containers” means domain separated so microservices
    - Work closely with product and design
  - Said to check out their CEO’s podcast at the bottom of the job application
  - He asked why I wanted to leave DailyPay as it is a large successful company
    - “I love DailyPay! Great people, great tech stack, good engineering.But, I moved
      to New York about 6 months ago to expand my opportunities, and will pursue any
      good opportunities that I see. When Tamara reached out, Coast looked like a
      cool mission and a great learning opportunity, no real reason to leave DailyPay
      but I am excited for other opportunities.
  - He asked what my biggest project has been
    - “So DailyPay works by giving earned wage access to our users. If you work a
      12-hour shift today, but pay day isn’t till Friday, you can take your money out
      right after your shift. To be able to do this, we need a line of credit backing
      the advances that we give to our users. This line of credit limits how much money
      we can give out, therefore how much we can profit. In January of 2025, we started
      a new initiative to increase our credit line. Similar to how mortgage backed securities
      are explained in the Big Short movie, we created a new securitization vehicle that uses
      the pay-checks that will be paid to DailyPay as a pool of assets. Investors are able to buy
      these securities, giving DailyPay immediate access to cash, which increases our credit line,
      and the investors get a larger pay-out in the future.
      I was the lead engineer in designing a service that could collect, analyze, summarize and
      report the credit rating of DailyPay’s outstanding advances, which powered that whole product I
      just discussed. This gave DailyPay an extra $200 million on our credit line.”

## 1:1 Technical Round w/ Abe Heidebrecht

- Abe introduced himself very quickly
- According to LinkedIn, he has only been at the company for 6 months and is a staff engineer
- Not a very welcoming vibe
- He said that I was supposed to have a hello world application setup already. I thought it was supposed to be setup on the call. Unfortunate mis-understanding of requirements
- I tried to build a new hello world API on the fly, but I accidentally did http instead of mux so it wasn’t working. Wasted about 5 ish minutes
- I switched to my existing mock app that I was practicing with
- Now that we had an operation app, he explained the prompt
- “Build a URL shortener. We want an API that can take in a URL and return a shortened version of it. Similar to what bit.ly does. An arbitrary string to short code”
- I asked some clarifying questions
  - What kind of character do we want?
    - “Alphanumeric, 12 characters”
  - Should the final URL have any relation from the previous URL inside of it?
    - “No”
- I started theorizing some solutions, immediately brought up hashing
  - He didn’t really respond to this at all, but then I started thinking through some issues with hashing
  - Mentioned hash collisions, and then he started steering me back to hashing by saying “good thought, but how likely is that?” So we didn’t worry about that anymore
- I started setting up the endpoint. Just wanted to setup what that interface looks like
  - I asked what kind of endpoints they want. POST, GET, etc.
  - He turned the question back around and asked what my opinions is on that. I responded “I like to stick pretty strictly to the guidelines of the various verbs. A GET should be just retrieving information, a POST is for creating/updating/deleting”. Stumbled a bit on this response with also going into PUTS but not a big deal
  - Then I said I was just going to go with a GET request cause we aren’t writing anything. To this he said, “Would it change your mind if eventually we wanted to store this?” OF COURSE that changed my mind… a POST it is
  - Built the API interface quickly and tested it without issue
- Switched to building out the shortener logic
  - Asked if I can google stuff cause I didn’t know the hashing library, he said that was fine but if I use AI he’ll ask a lot of questions
  - Wrote a hashing function real quick and plugged that into the POST request, it worked well
  - Had to google how to get the substring
- Once that worked, Prompt #1 was completed
- Now he asked (Prompt #2), what if we wanted to be able to get the original URL back from the shortened version
  - First thought was there’s no way to reverse the shortened URL, given we had to cut it up
  - Asked if he’s cool with me storing it in a database/cache, not that we had enough time to actually set one of those up but could mock one in memory with a map
  - Built that out rather quickly, but there was a bug
  - At this point he said “Ike, we only have 5 minutes left, are there any questions you want to ask me? I responded “Honest, I really just want to get this thing working”
  - Quickly found the bug, I wasn’t storing the shortened hash but instead the longer one as the key
  - Debugged that super fast and had 4 minutes for questions
    - “What’s your deployment cycle like? I’m personally a fan of quick deployment cycles.”
      - He said they are fast
    - “How’s the tech stack? I heard it’s Python for some old stuff, full-stack Typescript for new stuff”
      - He said that’s exactly it. He’s a fan of the type-ing in Typescript and I said I prefer static languages too
      - Next.js is the backend
      - They’re not worried about alerts, etc.
    - “What are you looking for in this position?”
      - “A solid engineer”
- Said goodbye, told him to have a nice thanksgiving and good luck in the interviewing process
- My initial thoughts are..
  - It went well
  - Could’ve been faster so he could’ve asked more questions about my solution
  - Start was a bummer and a bit slow
  - Got the solutions for both prompts
  - Never felt like I stopped coding so I’m not fully sure where the time went
  - I guess I was too slow up-front
  - The two arguments I see are:
    - “A senior engineer would’ve accomplished this quicker”
    - “A mid-level would’ve needed hand-holding”
  - I didn’t get any hand-holding

## System Design + Behavioral, December 10th, 2025, 1:00 - 3:15 pm EST

- System Design with Santiago and Kevin
- Behavioral with Danny

## CEO Interview, December 15th, 2025

Questions for Daniel Simon

Can you explain your journey to me please? Seems like you were a software engineer, then went to law school, then founded coast?
What was the motivation to go to law school?
Were you aiming to start a company through all of this?

I was told you all are hiring several engineers. What plans does coast have in the next year causing the hiring wave?
Are you currently outsourcing any work to consultants? Do you plan to?
I was checking the people I met at Coast on LinkedIn, I noticed a lot of Lab49. Can you describe that relationship?

Comparing Coast to Bread, Bread seems to have been acquired right at that 5.5 year mark. With Coast you are also at that 5.5 year mark.
Is this where you wanted to be with Coast now?
Is the end goal acquisition or establishment?

What team do you envision me having the biggest impact on?
Considering your background is in software engineering, how do you expect the person filling this role to contribute?

On the IO Podcast you mention that you should have some humility as a startup should have some humility when entering the market.
Interesting point
I’d never think of that
How do you put “humility” into place practically while trying to beat those giants?

On the Primary Ventures Podcast you mention that you felt like there was a slowness in the velocity of building the user facing product despite the talent of the company. You mentioned a re-org that fixed this to get rid of inter-team dependencies. Can you talk more about the details of this re-org?
Hired very senior engineers initially

My answers to his questions
Why do you want to work at Coast?
I moved to NYC 6 months ago with the goal of expanding my experience in the Fintech space. I currently work in Fintech and feel like I know that space quite well; Coast is an entirely different sector for me to learn and experience.
Strong engineering culture
Vision

What other companies are you interviewing with?
3 other companies, 1 in fintech but all startups
Equal ponit in interview process with one of them

Why are you drawn to fintech?
Money never goes out of style
The management of money and assets is vast and complex
Love the idea of math, algorithms, etc.
Reliable industry
So many different realms to explore
