# Cuvva take home challenge

Thanks for applying to become a Cuvva Backend Engineer 🎉. Below is a challenge we'd like you complete and send back to us.

## Rules & time limits

Before we start lets set some rules on how much you should feel you need to do:

- The challenge shouldn't take you more than **four hours** (this is the recommended time but we expect it to take less)
    - If you reach the maximum allowed time and haven't completed the task in the way you want, you can use a README
    to document what you would have liked to do. If you do so, please be as detailed as possible.
- You should write code as if it were going to go into production. This means using best practices where possible.
    - We favour well-documented, simple code over code that is complex, terse and hard to read.
    - We favour well tested code.
    - We favour making proper use of concurrency to speed up processes where possible.
    - We believe its good practice to make sure inputs are validated and sanitized before being used.
- Once you're finished we'd like you to submit your code back to us as a zip. Please do not use a public repository.

## The task

Cuvva sells a lot of motor insurance, you can purchase cover for a range of durations, from 30 minutes all the way up to
28 days. To do this we need a system that prices a user accurately and quickly.

Pricing means taking a set of information (usually details about a user) and ultimately generating a single price.
This is the 'quote price' that we will ultimately end up charging a customer.

In this scenario, the Product Manager has asked you to come up with a system to perform pricing calculations on a single
user for a range of durations.
The PM mentioned that the system needs to potentially be able to reprice thousands of users within a single minute.
They also mentioned that the price needs to be deterministic, given the same input information it should always give
you the same output.

### Resources

You will find attached

- A code stub in Go implementing an RPC endpoint
- A model, this is an excel sheet detailing how we translate risks about the user to a price. [Link here](https://docs.google.com/spreadsheets/d/10z5KCDmTXJQiCGIXDgPH7pmqRffbObUCdORXZLWmFVU/edit?usp=sharing).

### The Challenge

#### Task 1 - Familiarise yourself with the resources

Take a look at the code stub and familiarise yourself with the structure.
This implements an RPC endpoint (`generate_pricing`) which is called with a single users details.

Take a look at [the provided model](https://docs.google.com/spreadsheets/d/10z5KCDmTXJQiCGIXDgPH7pmqRffbObUCdORXZLWmFVU/edit?usp=sharing). The model gives you details on how to take in a risk i.e. an insurance group
and translate it to a "factor". These factors get multiplied with the "Base Rate" price (found in the "Base Rates"
sheet) to form a total price the user should pay. The base rate changes depending on the duration of cover selected.
We typically give the user a price for every duration of cover we offer, letting the user pick the duration that makes
sense for them.
Take a look at the "Calculator" sheet and familiarise yourself with how the total price a user pays is calculated.

A few notes worthwhile mentioning:
- We do not intend you to read, parse and use the Excel spreadsheet we provided. This is only provided to communicate
the model. Feel free to copy the factors into a structure you find appropriate for the exercise.
- We expect you to think about scalability of the model and factor it into the design of the application.

#### Task 2  - Accept details about a user

Update the `generate_pricing` endpoint to accept the following example input.

```json
{
    "date_of_birth": "1970-12-04",
    "insurance_group": 12,
    "license_held_since": "1988-08-01"
}
```

This represents details about a single user.

#### Task 3 -  Produce a users price for a range of durations

Update the `generate_pricing` endpoint to respond with the total price a user should pay for each duration of cover
found in the base rates sheet.

The output must be in JSON, the schema of the output is left to you.
