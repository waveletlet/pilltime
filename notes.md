# pilltime
Questions to answer:
- Given a drug and dose schedule, how much is left at time X 
  - And graph it over time 
- Tapers: If I decrease/increase the dose by X% per time unit, when will I be
  down/up to Y dosage?
  - Options: steady change or stepwise change
    - Ex: For a decrease of 25% per week, do I want to drop by 25% at the
      beginning of each week or 3.57% per day?

Features:
- Scheduling: import/export schedule to text file
  - Simple schedule config file format
  - Export as ical, etc for import into calendar program
- Interactive reminder timer
  - Option to ask what time pill was actually taken, vs what time
    reminder was dismissed, vs assuming if the reminder was dismissed it was
    taken on time

Advanced would be nice stuff:
- Half-life range vs average
  - Graph display like cone of probability? Or three lines?

## Exponential decay
y(t) = a × e^kt

Where y(t) = value at time "t"
a = value at the start
k = rate of growth (when >0) or decay (when <0)
t = time

If I know the half life and starting quantity, first need to solve for k:

````
y / a = e^kt
ln(y / a) = kt
ln(y / a) / t = k
a = starting quantity
y = starting quantity / 2
t = half life
````

In go terms: `k := math.Log(y/a) / t`


Need to decide at what value of y is it ~0? Or use rule of thumb: after 5
half-lives considered gone

## Models
- Dose
  - Dosage
  - ScheduledTime time.DateTime
    - ScheduledTime = TakenTime until interactive reminder timer functional
  - TakenTime time.DateTime
  - ZeroTime time.DateTime

- Dosage
  - Qty float? bignum?
    - Probably can store as float and do calcs as bignum
  - QtyUnits string
  - Chemical struct

- Chemical 
  - HalfLife time.Duration

- Schedule
  - Dose
  - BeginTime time.DateTime
  - EndTime time.DateTime
  - Interval/schedule type
    - (see Calendula schedule options)
