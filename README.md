# go-tdd
TDD is currently a software development method that is widely used considering that the results of products developed using TDD will have superior quality and minimal bugs because it breaks a problem into smaller and more manageable parts, thus helping the problem solving process. here i want to try using TTD on queue data structure.

## queue requirement
tdd will be very difficult to do if we don't have clear requirements in our system (even though tdd also helps us imagine what our requirements are), so let's describe what is happening in our queue system
- an queue has 0 length on construction
- an queue is empty on construction
- after n prepend to an empty queue, the queue is non empty and its length equal n
- after we prepend a value, the value must be in frist order
- if we take a data fromq queue, the data must be in the last order

above are the requirements of a queue, even though the requirements above do not fully describe what is happening in the queue, but it is enough for us to start writing tests, and refactor