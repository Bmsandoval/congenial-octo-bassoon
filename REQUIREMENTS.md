# Scenario

You are given a very simple log of events(`server_log.csv`) for a server that allows users to upload or download files.


Using this log file as input, write a program that will output the following metrics:
1. How many users accessed the server?
2. How many uploads were larger than `50kB`?
3. How many times did `jeff22` upload to the server on April 15th, 2020?


If you are interested in a programming position, please complete this challenge in Go. Otherwise, use a programming or scripting language of your choice.


# Details
The log is represented by a CSV formatted file where column data is ordered: `timestamp`, `username`, `operation`, `size`.

- The `timestamp` is recorded in the UNIX date format.
- The `username` is a unique identifier for the user.
- The `operation` indicates if an `upload` or `download` occurred, no other values will appear in this column.
- The `size` is an integer reflecting file size in `kB`.


# Important Note
The intention of this coding challenge is to gauge a candidate's ability to program, write clean code, and discuss their work with others. While we would like to see your best work, we don't want candidates to spend a large amount of their free time optimizing their solution to this challenge. So, our expectation is not necessarily to receive an optimal or even completely correct solution. With this in mind, we recommend spending about an hour on this challenge.