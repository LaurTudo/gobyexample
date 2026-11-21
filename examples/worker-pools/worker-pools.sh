# La rulare, programul arată cum cele 5 sarcini sunt
# executate de diferiți workeri. Programul durează doar
# aproximativ 2 secunde, deși face în total aproximativ
# 5 secunde de muncă, pentru că avem 3 workeri care
# lucrează concurent.
$ time go run worker-pools.go 
worker 1 started  job 1
worker 2 started  job 2
worker 3 started  job 3
worker 1 finished job 1
worker 1 started  job 4
worker 2 finished job 2
worker 2 started  job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5

real	0m2.358s
