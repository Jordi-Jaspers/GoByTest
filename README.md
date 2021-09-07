# GoByTest
## Biography  

**Authors:**  
Jordi Jaspers [[Github](https://github.com/Jordi-Jaspers "Github Page"), [Linkedin](https://www.linkedin.com/in/jordi-jaspers/ "Linkedin Page")]  
  
**Date of initial commit:**  
10/09/2021

**Description:**  
Learning Go (Made by Google) by test-driven examples to promote test-driven development (=TTD). TDD emphasizes the importance of an effective and sustainable testing approach. TDD also contributes directly to the overall quality of software. It’s a truism for small or large system development that often goes missing in the day-to-day hustle to get new functionality into production. Quality software gets built when there’s an acknowledgment that quality test code should receive the same amount of attention and resources as quality production code, as they are equally essential in development.
  
### What We Learned (so far..)
* Test driven design with GO

**Chapter 1**
  
* Writing tests
* Declaring functions, with arguments and return types
* "if", "const" and "switch"
* Declaring variables and constants

**Chapter 2**
* More practice of the TDD workflow
* "Integers", "addition"
* Writing better documentation so users of our code can understand its usage quickly
* Examples of how to use our code, which are checked as part of our tests

**Chapter 3**
* More TDD practice
* Learned "for"
* Learned how to write benchmarks

**Chapter 4**
* Arrays
* Slices
* The various ways to make them
* How they have a fixed capacity but you can create new slices from old ones
* using append
* How to slice, slices
* len to get the length of an array or slice
* Test coverage tool
* reflect.DeepEqual and why it's useful but can reduce the type-safety of your code

**Chapter 5**
* Declaring structs to create your own data types which lets you bundle related data together and make the intent of your code clearer
* Declaring interfaces so you can define functions that can be used by different types (parametric polymorphism)
* Adding methods so you can add functionality to your data types and so you can implement interfaces
* Table driven tests to make your assertions clearer and your test suites easier to extend & maintain

**Chapter 6**
**Pointers**
* Go copies values when you pass them to functions/methods, so if you're writing a function that needs to mutate state you'll need it to take a pointer to the thing you want to change.
* The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make a copy of something, in which case you need to pass a reference. Examples include referencing very large data structures or things where only one instance is necessary (like database connection pools).
**nil**
* Pointers can be nil
* When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime exception - the compiler won't help you here.
* Useful for when you want to describe a value that could be missing
**Errors**
* Errors are the way to signify failure when calling a function/method.
* By listening to our tests we concluded that checking for a string in an error would result in a flaky test. So we refactored our implementation to use a meaningful value instead and this resulted in easier to test code and concluded this would be easier for users of our API too.
* This is not the end of the story with error handling, you can do more sophisticated things but this is just an intro. Later sections will cover more strategies.
* Don’t just check errors, handle them gracefully
**Create new types from existing ones**
* Useful for adding more domain specific meaning to values
* Can let you implement interfaces


**Code Wars**  
All these challenges are from <https://www.codewars.com/users/Jordi-Jaspers>.
### References

* Troubleshooting: <https://stackoverflow.com/>
* The TDD Go course: <https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration>
