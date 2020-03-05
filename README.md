## prospace-test

# GALAXY MERCHANT TRADING GUIDE

Problem Description

You decided to give up on earth after the latest financial collapse left 99.99% of the earth's population with 0.01% of the wealth. Luckily, with the scant sum of money that is left in your account, you are able to afford to rent a spaceship, leave earth, and fly all over the galaxy to sell common metals and dirt (which apparently is worth a lot). Buying and selling over the galaxy requires you to convert numbers and units, and you decided to write a program to help you. The numbers used for intergalactic transactions follows similar convention to the roman numerals and you have painstakingly collected the appropriate ranslation between them. Roman numerals are based on seven symbols:
Symbol Value: 
I 1
V 5
X 10
L 50
C 100
D 500
M 1,000

![alt text](https://github.com/johanesarnold/prospace-test/blob/master/diagram-solution.png)

## Input
```
glob is I 
prok is V 
pish is X 
tegj is L 
glob glob Silver is 34 Credits 
glob prok Gold is 57800 Credits 
pish pish Iron is 3910 Credits 
how much is pish tegj glob glob ? 
how many Credits is glob prok Silver ? 
how many Credits is glob prok Gold ? 
how many Credits is glob prok Iron ? 
how much wood could a woodchuck chuck if a woodchuck could chuck wood ?
```
## Expected
```
pish tegj glob glob is 42 
glob prok Silver is 68 Credits 
glob prok Gold is 57800 Credits 
glob prok Iron is 782 Credits 
I have no idea what you are talking about
```
## How to run solution
```
go run main.go
```