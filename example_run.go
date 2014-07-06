import eperf

func ExampleRun() {
  perftest := eperf.New("code test", 1000)  //Create a new Perftest
  for perftest.Run() {                      //Use Run() in a loop
    code()                                  //code() will be executed 1000 times. 
  }
  
  //The fastest execution time will be logged. Example output:
  //2014/07/06 17:26:59 eperf: code test runtime: 133.806us

}
