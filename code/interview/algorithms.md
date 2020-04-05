## 1. Split numbers 
### title
Given an array of ints, for example `[6, 4, -3, 0, 5, -2, -1, 0, 1, -9]`, implement in one of the following languages
to move all positive integers to the left, all negative integers to the right, and all zeros to the middle.

You are not allowed to consume extra O(n) storage space, aka. your algorithm space complexity should be O(1).

Your answer should be compilable and runnable, in one of the following function signatures:

### solution


## 2. Serialize reversed list
### title
Given a “reversed list”, whose first tuple value is the id of string type, and second tuple value is the path to locate the id in a structured document. 

For example, `{"1": "bar", "2": "foo.bar", "3": "foo.foo", "4": "baz.cloudmall.com", "5": "baz.cloudmall.ai"}`
Your mission is to transform the list back to a document in JSON format.

For example, a legit JSON for the above list would be:
```
{
  "bar": "1",
  "foo": {
    "bar": "2",
    "foo": "3"
  },
  "baz": {
    "cloudmall": {
      "com": "4",
      "ai": "5"
    }
  }
}
```
Your answer should be compilable and runnable, in one of the following function signatures:

### solution
