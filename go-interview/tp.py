print('Hello, world!')
print('What is your name?')
myName = input()
print('It is nice to met you, ' + myName)
print('The length of your name is:')
print(len(myName))
print('What is your age?')
myAge = input()
if myAge.isdigit():
    print('You will be ' + str(int(myAge) + 1) + ' in "one" year')
else:
    print('that is not a number. Lets try again, what is your age?')
C = 0
while (C <= 3):
    myAge2 = input()
    C += 1
    if myAge2.isdigit():
        print('You will be ' + str(int(myAge2) + 1) + ' in "one" year')
        break
    else:
        print('Try again')
