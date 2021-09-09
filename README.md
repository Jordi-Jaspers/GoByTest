# GoByTest
### Biography  

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
  
**Chapter 7**  
* Full CRUD dictionary API
* Create maps
* Search for items in maps
* Add new items to maps
* Update items in maps
* Delete items from a map
* Learned more about errors
* How to create errors that are constants
* Writing error wrappers
  
**Chapter 8**  
Motivated by our tests we refactored the code so we could control where the data was written by injecting a dependency which allowed us to:
* Test our code If you can't test a function easily, it's usually because of dependencies hard-wired into a function or global state. If you have a global database connection pool for instance that is used by some kind of service layer, it is likely going to be difficult to test and they will be slow to run. DI will motivate you to inject in a database dependency (via an interface) which you can then mock out with something you can control in your tests.  
* Separate our concerns, decoupling where the data goes from how to generate it. If you ever feel like a method/function has too many responsibilities (generating data and writing to a db? handling HTTP requests and doing domain level logic?) DI is probably going to be the tool you need.  
* Allow our code to be re-used in different contexts The first "new" context our code can be used in is inside tests. But further on if someone wants to try something new with your function they can inject their own dependencies.  
  
**Chapter 9**  
* Mocking
* Without mocking important areas of your code will be untested.
* Without mocks you may have to set up databases and other third parties things just to test simple business rules. You're likely to have slow tests, resulting in slow feedback loops.
* By having to spin up a database or a webservice to test something you're likely to have fragile tests due to the unreliability of such services.
* Reference to Martin Fowler: <https://martinfowler.com/bliki/TestDouble.html>

#### Mock Definitions by Martin Fowler
* **Dummies:** objects are passed around but never actually used. Usually they are just used to fill parameter lists.
* **Fakes:** objects actually have working implementations, but usually take some shortcut which makes them not suitable for production (an InMemoryTestDatabase is a good example).
* **Stubs:** provide canned answers to calls made during the test, usually not responding at all to anything outside what's programmed in for the test.
* **Spies:** are stubs that also record some information based on how they were called. One form of this might be an email service that records how many messages it was sent.
* **Mocks:** are pre-programmed with expectations which form a specification of the calls they are expected to receive. They can throw an exception if they receive a call they don't expect and are checked during verification to ensure they got all the calls they were expecting.
  
### Code Wars
All these challenges are from <https://www.codewars.com/users/Jordi-Jaspers>.
### References
  
* Troubleshooting: <https://stackoverflow.com/>
* The TDD Go course: <https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration>
