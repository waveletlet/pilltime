# pilltime

A drug halflife calculator and eventually dose scheduler. 

Questions to answer:
- Given some drug, how much is left at time T 
  - [x] After one dose?
  - [ ] After a series of doses?
  - [ ] And graph it over time 
- Tapers: if I decrease/increase the dose by X% per time unit 
  - [ ] When will I be down/up to Y dosage?
  - [ ] What will my dosage be at time T?
  - [ ] Generate dose schedule
  - Options: 
    - [ ] steady change 
    - [ ] stepwise change
    - Ex: For a decrease of 25% per week, do I want to drop by 25% at the
      beginning of each week or 3.57% per day?
- If I want a target level (with max and mins) of drug in my blood stream, what
  is my:
  - [ ] Ideal schedule (i.e. every 7.5 hours)
  - [ ] Realistic schedule (i.e. 3x a day)
    - Should be logical way of constraining realistic schedule (sleep times,
      etc)
- If I have half a bottle of pills left,
  - What happens to by blood level if I stretch it out over X time?
  - When do I need to refill by if I don't want to drop below X blood level?

Note: references to blood levels are massive oversimplifications of how
pharmacokinetics works and I am neither a pharmacist nor doctor.

Features:
- [ ] Scheduling: import/export schedule to text file
  - [ ] Simple schedule config file format
  - [ ] Export as ical, etc for import into calendar program
- [ ] CSV output
- [ ] Graph output
- [ ] Interactive reminder timer
  - [ ] Option to ask what time pill was actually taken, vs what time
    reminder was dismissed, vs assuming if the reminder was dismissed it was
    taken on time

Advanced would be nice stuff:
- [ ] Half-life range vs average
  - [ ] Graph display like cone of probability? Or three lines?


