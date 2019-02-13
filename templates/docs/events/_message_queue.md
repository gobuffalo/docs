## Integrating a Messaging Queue

It is often desirable to take events emitted and send them to a message queue, such as Kafka or Redis, to be processed externally. The <%= doclink("github.com/gobuffalo/events") %> package does not have a directhook for this sort of functionality, the most direct way of enabling this behavior is to register a <%= doclink("github.com/gobuffalo/events#Listener") %> that can then hand the event over to the appropriate message queue.

```go
events.Listen(func(e events.Event) {
  myMessageQ.DoWork(e)
})
```
