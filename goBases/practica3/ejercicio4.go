const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
 )
  
 
 func operation

 
 minFunc, err := operation(minimum)
 averageFunc, err := operation(average)
 maxFunc, err := operation(maximum)
  
 
  
 minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)
 averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
 maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
 