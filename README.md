<h3 align="center">GoByTest</h3>

<div align="center">

  [![Status](https://img.shields.io/badge/status-active-success.svg)]() 
  [![GitHub Issues](https://img.shields.io/github/issues/Jordi-Jaspers/GoByTest.svg)](https://github.com/kylelobo/The-Documentation-Compendium/issues)
  [![GitHub Pull Requests](https://img.shields.io/github/issues-pr/Jordi-Jaspers/GoByTest.svg)](https://github.com/kylelobo/The-Documentation-Compendium/pulls)
  [![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

### Biography  

**Authors:**  
Jordi Jaspers [[Github](https://github.com/Jordi-Jaspers "Github Page"), [Linkedin](https://www.linkedin.com/in/jordi-jaspers/ "Linkedin Page")]  
  
**Date of initial commit:**  
10/09/2021

**Description:**  
Learning Go (Made by Google) by test-driven examples to promote test-driven development (=TTD). TDD emphasizes the importance of an effective and sustainable testing approach. TDD also contributes directly to the overall quality of software. It’s a truism for small or large system development that often goes missing in the day-to-day hustle to get new functionality into production. Quality software gets built when there’s an acknowledgment that quality test code should receive the same amount of attention and resources as quality production code, as they are equally essential in development.  

---

## 📝 Table of Contents
- [Go Fundamentals](#fundamentals)
- [Application](#application)
- [Meta](#meta)
- [TODO](#todo)
- [Code Wars](#code_wars)
- [Built Using](#built_using)
- [References](#references)

## Go Fundamentals <a name = "fundamentals"></a>
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

##### Mock Definitions by Martin Fowler
* **Dummies:** objects are passed around but never actually used. Usually they are just used to fill parameter lists.
* **Fakes:** objects actually have working implementations, but usually take some shortcut which makes them not suitable for production (an InMemoryTestDatabase is a good example).
* **Stubs:** provide canned answers to calls made during the test, usually not responding at all to anything outside what's programmed in for the test.
* **Spies:** are stubs that also record some information based on how they were called. One form of this might be an email service that records how many messages it was sent.
* **Mocks:** are pre-programmed with expectations which form a specification of the calls they are expected to receive. They can throw an exception if they receive a call they don't expect and are checked during verification to ensure they got all the calls they were expecting.
  
**Chapter 10**  
* goroutines, the basic unit of concurrency in Go, which let us check more than one website at the same time.
* anonymous functions, which we used to start each of the concurrent processes that check websites.
* channels, to help organize and control the communication between the different processes, allowing us to avoid a race condition bug.
* the race detector which helped us debug problems with concurrent code

**Chapter 11**  
**select**

* Helps you wait on multiple channels.
* Sometimes you'll want to include time.After in one of your cases to prevent your system blocking forever.

**httptest**

* A convenient way of creating test servers so you can have reliable and controllable tests.
* Using the same interfaces as the "real" net/http servers which is consistent and less for you to learn.

**Chapter 12**
* Introduced some of the concepts from the reflect package.
* Used recursion to traverse arbitrary data structures.
* Avoid reflection: <https://go.dev/blog/laws-of-reflection>

**Chapter 13**
* Mutex allows us to add locks to our data.
* Waitgroup is a means of waiting for goroutines to finish jobs.
* Do not expose structs with Mutex or other sensitive data. Try to create a func that passes a reference.
* use __'Go vet'__ to find bugs or subtle mistakes.
* Don't use embedding because it's convenient, this is a very bad practice because of exposing your API.

#### Mutex OR Channels
"Share memory by communicating, don't communicate by sharing memory." - Golang 2020  
The general rule is not to overuse both and consider the effectiveness in your use case. Commenly used practices consider the following.  
**Channel** -> passing ownership of data, distributing units of work, communicating async results  
**Mutex** ->  caches, state  

**Chapter 14**
* How to test a HTTP handler that has had the request cancelled by the client.
* How to use context to manage cancellation.
* How to write a function that accepts context and uses it to cancel itself by using goroutines, select and channels.
* Follow Google's guidelines as to how to manage cancellation by propagating request scoped context through your call-stack.
* How to roll your own spy for http.ResponseWriter if you need it.

#### Google guidelines about canceling request  
At Google, we require that Go programmers pass a Context parameter as the first argument to every function on the call path between incoming and outgoing requests. This allows Go code developed by many different teams to interoperate well. It provides simple control over timeouts and cancelation and ensures that critical values like security credentials transit Go programs properly. source: <https://go.dev/blog/context>  

```Go
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

//Http.Handler creates server that fetches the store data.
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log error however you like
		}

		fmt.Fprint(w, data)
	}
}
```

**Chapter 15**
* Built into the standard library
* If you can think of ways to describe your domain rules in code, they are an excellent tool for giving you more confidence
* Force you to think about your domain deeply
* Potentially a nice complement to your test suite

## Application <a name = "application"></a>
The application must have the following endpoints:

* `GET /players/{name}`  
should return a number indicating the total number of wins
  
* `POST /players/{name}`  
should record a win for that name, incrementing for every subsequent POST

## Meta <a name = "meta"></a>

## 🔧 TODO <a name = "todo"></a>

## 👨‍💻 Code Wars <a name = "code_wars"></a>

All the challenge are directly from the [CodeWars](https://www.codewars.com/users/Jordi-Jaspers) website.

## ⛏️ Built Using <a name = "built_using"></a>
- [Golang](https://golang.org/) - Backend Framework

## ✍️ References <a name = "references"></a>
* Troubleshooting: <https://stackoverflow.com/>
* The TDD Go course: <https://quii.gitbook.io/learn-go-with-tests>
