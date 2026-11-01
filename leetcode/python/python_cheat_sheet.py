# Data Types
stringg = "string"
integerr = 5
floatt = 5.0
listt = ["use", "square", "brackets"]
tuplee = ("use", "parenthesis")
dictionaryy = {"use": "squiggly", "brackets": 5}
sett = {"also", "use", "squiggly", "brackets", "no", "colons"}
booll = True
bytess = b"Hello"

# Type Casting
d = int(3)
e = float(3)
f = str(3)
print(type(d), type(e), type(f))

# Strings
## strings are arrays
for x in "banana":
    print(x)
## string length
print(len("hello"))
## check if sub-string exits
txt = "hello world"
if "world" in txt:
    print("helllllllo worlddddd")
## slicing
txt2 = "this is a sample string"
print(txt2[2:5])
## modify stringg
a = "Hello, World!"
print(a.lower())
print(a.upper())
print(a.strip())  # returns "Hello, World!"
print(a.replace("H", "J"))
print(a.split(","))  # returns ['Hello', ' World!']
## For Loops


for i in range(256):
    binary = format(i, "08b")  # 8-bit binary representation
    char = chr(i)  # Character from Unicode code point
    print(f"Decimal: {i:3} | Binary: {binary} | Char: {repr(char)}")
