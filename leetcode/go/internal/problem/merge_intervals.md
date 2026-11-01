# MERGE OVERLAPPING INTERVALS

Given an array of `intervals` (interval = [start, end]; start will always be less <= end), merge intervals where their times overlap

## CONTEXT

- What is an overlapping interval?
  [ [1, 3], [2, 5] ] --> these two overlap with 2 and 3

- How do we merge intervals?
  Following the problem set given from the first context bullet, the merge results would be:
  [ [1, 5] ] --> The lower `start` and the greater `end` will be used as the new interval.
  At the end of a merge, two intervals have become one.

## QUESTIONS

- Is there any order to the intervals within the array? Like ascending order from one interval to the next?
  I don't think there is, which means we'd have to compare each interval to all other intervals, can't just compare side by side

## SOLUTION IDEATING

- Comparing two intervals, how do we know they overlap? Define overlapping as an algorithm:

---

[ [start1, end1], [start2, end2] ]

- if `start1` >= `start2` AND `start1` <= `end2` then THERE IS OVERLAP
- if `start1` <=`start2` AND `end1` >= `start2` then THERE IS OVERLAP

- if `end1` >= `start2` OR `end1` >= `end2` then THERE IS OVERLAP

Ex. [ [10, 70], [6, 13] ]

- is 10 >

---

### BRUTE FORCE - Compare each interval to every other interval

### SORT - Sort intervals by start then compare end to start
