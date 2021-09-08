# FindSmallestNumberOfFactorialContainingNZeroes <img alt="Go" src="https://img.shields.io/badge/Go-v1.17-blue"/>
This is the GO implementation of 'FindSmallestNumberOfFactorialContainingNZeros' I had previously written in C#

Given ![\Large N](https://latex.codecogs.com/svg.latex?\Large&space;N) number, the task is to determine the smallest number ![\Large N](https://latex.codecogs.com/svg.latex?\Large&space;N), so that ![\Large N!](https://latex.codecogs.com/svg.latex?\Large&space;N!) has at least ![\Large k](https://latex.codecogs.com/svg.latex?\Large&space;k) trailing zeros.

In more simple terms, the number of trailing zeros ![\Large k](https://latex.codecogs.com/svg.latex?\Large&space;k) must match the input ![\Large N](https://latex.codecogs.com/svg.latex?\Large&space;N) by <i>finding the smallest factorial with trailing zeros matching the input</i>.

## Overview ##
[This "Hour of Code 2017" challenge](https://www.hackerrank.com/contests/hour-of-code-2017/challenges/number-of-zeros/problem "HackerRank's Hour of Code 2017") was part of a contest HackerRank held on Dec 4 2017, 01:00 am CDT to Dec 11 2017, 01:00 am CDT for primary school and high school students, and everywhone else who wanted to learn programming. The contest was made for the Computer Science Education Week and organized by CTK Rijeka. 

Everyone knows that ![\Large 10^6](https://latex.codecogs.com/svg.latex?\Large&space;10^6) is called a "million", and most of us know that the number ![\Large 10^100](https://latex.codecogs.com/svg.latex?\Large&space;10^100) is called "googol" (not Google, that's a tech company).

Even though one googol is a really large number, there are <i>many</i> greater numbers. It will be your task to find the <b>smallest number</b>, so that ![\Large N!](https://latex.codecogs.com/svg.latex?\Large&space;N!) has at least ![\Large k](https://latex.codecogs.com/svg.latex?\Large&space;k) trailing zeros.

## Input Format ##
In the first and only line of input there is a single integer number ![\Large k](https://latex.codecogs.com/svg.latex?\Large&space;k), representing the number of trailing zeros in ![\Large N!](https://latex.codecogs.com/svg.latex?\Large&space;N!).

## Constraints ##
![\Large o<= k < 10^9](https://latex.codecogs.com/svg.latex?\Large&space;0<k<10^9) 

## Output Format ##
In a single line of output print the number ![\Large N](https://latex.codecogs.com/svg.latex?\Large&space;N) from the problem statement.

## Sample Input 1 ##
```
1
```

## Sample Output 1 ##
```
5
```

## Explanation 1 ##
We can check this by writing the first factorials:
```
1! = 1
2! = 2
3! = 6
4! = 24
5! = 120
```
Notice that the smallest number to have 1 trailing zero in it's factorial is 5.

## Sample Input 2 ##
```
6
```

## Sample Output 2 ##
```
25
```

## Explanation 2 ##
We can check this by writing the first factorials:
```
1! = 1
2! = 2
3! = 6
4! = 24
5! = 120
6! = 720
7! = 5040
8! = 40320
9! = 362880
10! = 3628800
11! = 39916800
12! = 479001600
13! = 6227020800
14! = 87178291200
15! = 1307674368000
16! = 20922789888000
17! = 355687428096000
18! = 6402373705728000
19! = 121645100408832000
20! = 2432902008176640000
21! = 51090942171709440000
22! = 1124000727777607680000
23! = 25852016738884976640000
24! = 620448401733239439360000
25! = 15511210043330985984000000
```
Notice that the smallest number to have 6 trailing zero in it's factorial is 25. Notice the difference between sample 1 vs sample 2 factorials.

## Remarks ##
* I found [this HackerRank "Hour of Code 2017" challenge](https://www.hackerrank.com/contests/hour-of-code-2017/challenges/number-of-zeros/problem "HackerRank's Hour of Code 2017") to be fascinating, fun and quite enlightening!!
* I added this to my GitHub repos for reference and explanation into the challenges of math and factorials, as a brain teaser.
* I originally wrote this using integers, for loops and several conditions which was not optimal. I was able to get it to pass the HackerRank tests but not fully (it was failing at 50% due to timeout issues). The secret is that you must use factorials and calculations to project the numbers. I scoured the internet looking for helpful code and noticed there wasn't very much in Go (and also C#), so I wanted to contribute a better solution and better explanation, than what was found on GitHub at this time of writing.
* Below you can begin to visually see the pattern of the factorials. 
* When you add the input to the expected output as the next row's input, watch the expected output multiple by 5!
* For example: 1 + 5 = 6, 6 + 25 = 31, 31 + 125 = 156, 156 + 625 = 781, 781 + 3125 = 3906, etc. You can noticed the expected output multiples by a factor of 5!

```
func main() {
	fmt.Println(findSmallestFactorialForNZeros(1))    // expected output 5
	fmt.Println(findSmallestFactorialForNZeros(6))    // expected output 25
	fmt.Println(findSmallestFactorialForNZeros(31))   // expected output 125
	fmt.Println(findSmallestFactorialForNZeros(156))  // expected output 625
	fmt.Println(findSmallestFactorialForNZeros(781))  // expected output 3125
	fmt.Println(findSmallestFactorialForNZeros(3906)) // expected output 15625
}
```

## Contact ##

You can contact me <a href="mailto:richardsmailbox@gmail.com?subject=Getting in touch!&body=Hi, I would like to contact you, Devtr0n.">via gmail</a> or at my website http://www.richardhollon.com or find me on <a href="https://www.linkedin.com/in/richardhollon/" target="_blank"><img alt="LinkedIn" src="https://img.shields.io/badge/linkedin-%230077B5.svg?style=for-the-badge&logo=linkedin&logoColor=white"/></a>
