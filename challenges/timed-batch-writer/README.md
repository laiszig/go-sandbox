### Go Concurrency Challenge - Timed Batch Writer

Build a system in Go that receives incoming data from an API and persists it to a database efficiently.

#### Input

* The API sends requests at random intervals between **0ms and 200ms**
* Each request contains a **registry**

#### Requirements

The solution must:

* Use **goroutines (concurrency)**
* Use **channels** for communication
* Use a **`select` statement**
* Run in an **infinite loop** to continuously process incoming data

#### Persistence Rules

You must buffer incoming registries and save them to the database based on the following conditions:

1. **Batch size condition**

    * When **100 registries** are accumulated → save all of them to the database

2. **Time condition**

    * If **10 seconds** pass since the last save → save whatever has been accumulated so far (even if < 1000)

#### Goal

Efficiently batch writes to the database by balancing:

* Throughput (saving in large batches)
* Latency (not waiting too long when traffic is low)