"""
Link: https://www.janestreet.com/mock-interview/
Problem: Unit Conversion
Description: Write a program to answer unit coversion programs, given a list of conversion facts about units

Examples

example facts:
    m = 3.28 ft
    f = 12 in
    hr = 60 min
    min = 60 sec

example queries:
    2m = ? in --> answer = 78.72
    13 in = ? m --> answer = 0.330 (roughly)
    13 in = ? hr --> "not convertible"

Requirements:
    A Fact is passed as a (String, Float, String) tuple
    A Query is passed as a (Float, String, String) tuple
"""

def convert_units(facts: tuple[str, float, str], queries: tuple[float, str, str]) --> list[tuple[float, str]]: 
    conversions = {}
    for fact in facts:
        larger_unit_name = fact[0]
        smaller_unit_amount = fact[1]
        smaller_unit_name = fact[2]

        conversions[larger_unit_name] = (smaller_unit_name, smaller_unit_amount)

    results = []
    for query in queries:

        while true:
            

