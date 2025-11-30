### Instructions
In this exercise you'll be organizing races between various types of remote controlled cars. Each car has its own speed and battery drain characteristics. <br>

Cars start with full (100%) batteries. Each time you drive the car using the remote control, it covers the car's speed in meters and decreases the remaining battery percentage by its battery drain. <br>

If a car's battery is below its battery drain percentage, you can't drive the car anymore. <br>

Each race track has its own distance. Cars are tested by checking if they can finish the track without running out of battery. <br>

Define a Car struct with the following int type fields:
- battery
- batteryDrain
- speed
- distance

Allow creating a remote controlled car by defining a function NewCar that takes the speed of the car in meters, and the battery drain percentage as its two parameters (both of type int) and returns a Car instance:
```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)
// => Car{speed: 5, batteryDrain: 2, battery:100, distance: 0}
```

Define another struct type called Track with the field distance of type integer. Allow creating a race track by defining a function NewTrack that takes the track's distance in meters as its sole parameter (which is of type int):

```go
distance := 800
track := NewTrack(distance)
// => Track{distance: 800}
```

Implement the Drive function that updates the number of meters driven based on the car's speed, and reduces the battery according to the battery drainage. If there is not enough battery to drive one more time the car will not move:
```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)
car = Drive(car)
// => Car{speed: 5, batteryDrain: 2, battery: 98, distance: 5}
```

To finish a race, a car has to be able to drive the race's distance. This means not 
draining its battery before having crossed the finish line. Implement the `CanFinish` 
function that takes a `Car` and a `Track` instance as its parameter and returns `true` if
 the car can finish the race; otherwise, return `false`.<br>

Assume that you are currently at the starting line of the race and start the engine of the car for the race. Take into account that the car's battery might not necessarily be fully charged when starting the race:
```go
speed := 5
batteryDrain := 2
car := NewCar(speed, batteryDrain)

distance := 100
track := NewTrack(distance)

CanFinish(car, track)
// => true
```