# Lab 6 Dining Philosophers

For running the program and the license see the [main README](../README.md).

The 2 Go folders contain 2 different solutions to the dining philosophers problem.
 - 1st solution uses a channel that is 1 size smaller than the amount of forks and philosophers
  so that at any time there are enough forks available.
 - 2nd solution makes all of the philosophers pick up their right fork first,
  except for the last one which is left handed, so that again their are enough forks available to eat.