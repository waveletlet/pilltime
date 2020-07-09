# pilltime
Questions to answer:
- Given a drug and dose schedule, how much is left at time X 
  - And graph it over time 
- Tapers: If I decrease/increase the dose by X% per time unit, when will I be
  down/up to Y dosage?
  - Options: steady change or stepwise change
    - Ex: For a decrease of 25% per week, do I want to drop by 25% at the
      beginning of each week or 3.57% per day?
- If I want a target level (with max and mins) of drug in my blood stream, what
  is my:
  - Ideal schedule (i.e. every 7.5 hours)
  - Realistic schedule (i.e. 3x a day)
    - Should be logical way of constraining realistic schedule (sleep times,
      etc)

Features:
- Scheduling: import/export schedule to text file
  - Simple schedule config file format
  - Export as ical, etc for import into calendar program
- CSV output
- Graph output
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
y / a = e^kh
ln(y / a) = kh
ln(y / a) / h = k
a = starting quantity
y = half starting quantity = a/2
h = half life
````

In go terms: `k := math.Log(a/2/a) / h`
Plug that into: `Y := a * math.Exp(k * t)`

But found a much more straightforward formula for half life:
Y = a * 2^-(t/h)
Y := a * math.Exp2(-t/h)



Need to decide at what value of y is it ~0? Or use rule of thumb: after 5
half-lives considered gone
- Test with 25 mg w/ 24h half life: at 5 half lives it's ~0.78 mg. Some people
  say a dose that small will still affect them, which will also depend on the
  drug in question and person in question. 10 half-lives is about 1/1000
  original dose, I'll arbitrarily go with that as a default "0" for cumulative
  calculations, but have it changeable.
  - Still have to play with if I want the variable to be in terms of drug
    quantity (and if that, absolute or relative to original dose), or half lifes

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
  - EmptyStomachOK yes/no/doesn't matter
    - Don't need to implement until scheduling system differentiates between
      meal times and nonmeal times

- DrugSchedule
  - Dose
  - BeginTime time.DateTime
  - EndTime time.DateTime
  - Interval/schedule type
    - (see Calendula schedule options)

- PatientRoutine
  - WakeTime
  - SleepTime
  - Meals (includes snacks)
    - MealTimes
  - Unavailable
    - Like sleep times but more flexible for misc other reasons someone
      predictably can't take a dose (ex. fasting, commute, class)

