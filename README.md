# Gas station simulation
## Original assignment
1. Cars arrive at the gas station and wait in the queue for the free station
2. Total number of cars and their arrival time is configurable
3. There are 4 types of stations: gas, diesel, LPG, electric
4. Count of stations and their serve time is configurable as interval (e.g. 2-5s) and can be different for each type
5. Each station can serve only one car at a time, serving time is chosen randomly from station's interval
6. After the car is served, it goes to the cash register.
7. Count of cash registers and their handle time is configurable
8. After the car is handled (random time from register handle time range) by cast register, it leaves the station.
9. Program collects statistics about the time spent in the queue, time spent at the station and time spent at the cash register for every car
10. Program prints the aggregate statistics at the end of the simulation

## Input `config.yaml`
```yaml
cars:
  count: 1000
  arrival_time_min: 5
  arrival_time_max: 10
  main_queue_length_max: 10
stations:
  1:
    count: 4
    serve_time_min: 20
    serve_time_max: 50
    queue_length_max: 3
  2:
    count: 3
    serve_time_min: 30
    serve_time_max: 60
    queue_length_max: 3
  3:
    count: 2
    serve_time_min: 70
    serve_time_max: 100
    queue_length_max: 3
  4:
    count: 1
    serve_time_min: 10
    serve_time_max: 40
    queue_length_max: 3
registers:
  count: 2
  handle_time_min: 10
  handle_time_max: 30
  queue_length_max: 1

```
## Formatted output
```
Pumps:
  Gas:
  	Cars Served: 282
  	Total Service Time: 12.1046903s
  	Total Queue Time: 420µs
  	Average Time: 1.489µs
  LPG:
  	Cars Served: 255
  	Total Service Time: 13.6946459s
  	Total Queue Time: 92.0978ms
  	Average Time: 361.167µs
  Electric:
  	Cars Served: 231
  	Total Service Time: 26.4145919s
  	Total Queue Time: 4.9233493s
  	Average Time: 21.3132ms
  Diesel:
  	Cars Served: 232
  	Total Service Time: 10.3350512s
  	Total Queue Time: 2.6811344s
  	Average Time: 11.556613ms

Registers:
	Cars Served: 1000
	Total Service Time: 37.393076s
	Total Queue Time: 10.8836648s
	Average Time: 10.883664ms
Total elapsed time: 15.633313s
```
