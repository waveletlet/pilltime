# pilltime

A drug halflife calculator and eventually dose scheduler. 

Questions to answer:
- [x]Given a drug and dose schedule, how much is left at time X 
  - [ ] And graph it over time 
- [ ] Tapers: If I decrease/increase the dose by X% per time unit, when will I be
  down/up to Y dosage?
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


