# diff
Go package to get differences of two files

# bench  
    BenchmarkCompare-8        100000             18906 ns/op
    ok      github.com/sj14/diff    2.094s

# memory profile
    File: diff.test
    Type: inuse_space
    Time: Mar 10, 2018 at 2:34pm (CET)
    Showing nodes accounting for 512.19kB, 100% of 512.19kB total
          flat  flat%   sum%        cum   cum%
      512.19kB   100%   100%   512.19kB   100%  runtime.malg
             0     0%   100%   512.19kB   100%  runtime.allocm
             0     0%   100%   512.19kB   100%  runtime.mstart
             0     0%   100%   512.19kB   100%  runtime.mstart1
             0     0%   100%   512.19kB   100%  runtime.newm
             0     0%   100%   512.19kB   100%  runtime.resetspinning
             0     0%   100%   512.19kB   100%  runtime.schedule
             0     0%   100%   512.19kB   100%  runtime.startm
             0     0%   100%   512.19kB   100%  runtime.wakep

# cpu profile
    File: diff.test
    Type: cpu
    Time: Mar 10, 2018 at 2:35pm (CET)
    Duration: 2.10s, Total samples = 2.13s (101.30%)
    Showing nodes accounting for 1.96s, 92.02% of 2.13s total
    Dropped 53 nodes (cum <= 0.01s)
          flat  flat%   sum%        cum   cum%
         0.53s 24.88% 24.88%      0.57s 26.76%  syscall.Syscall
         0.24s 11.27% 36.15%      0.24s 11.27%  runtime.epollctl
         0.24s 11.27% 47.42%      0.24s 11.27%  syscall.Syscall6
         0.15s  7.04% 54.46%      0.37s 17.37%  runtime.mallocgc
         0.08s  3.76% 58.22%      0.08s  3.76%  runtime.lock
         0.08s  3.76% 61.97%      0.08s  3.76%  runtime.memclrNoHeapPointers
         0.06s  2.82% 64.79%      0.06s  2.82%  runtime.futex
         0.05s  2.35% 67.14%      0.80s 37.56%  bufio.(*Scanner).Scan
         0.04s  1.88% 69.01%      1.81s 84.98%  github.com/sj14/diff.getContent
         0.04s  1.88% 70.89%      0.04s  1.88%  runtime.heapBitsSetType
         0.03s  1.41% 72.30%      0.14s  6.57%  runtime.growslice
         0.03s  1.41% 73.71%      0.03s  1.41%  runtime.memmove
         0.03s  1.41% 75.12%      0.03s  1.41%  runtime.newdefer
         0.03s  1.41% 76.53%      0.04s  1.88%  runtime.scanobject
         0.03s  1.41% 77.93%      0.12s  5.63%  runtime.slicebytetostring
         0.03s  1.41% 79.34%      0.06s  2.82%  runtime.sweepone
         0.02s  0.94% 80.28%      0.03s  1.41%  bufio.ScanLines
         0.02s  0.94% 81.22%      1.91s 89.67%  github.com/sj14/diff.Compare
         0.02s  0.94% 82.16%      0.04s  1.88%  internal/poll.runtime_pollUnblock
         0.02s  0.94% 83.10%      0.56s 26.29%  os.(*File).Read
         0.02s  0.94% 84.04%      0.53s 24.88%  os.(*File).read
         0.02s  0.94% 84.98%      0.02s  0.94%  runtime._ExternalCode
         0.02s  0.94% 85.92%      0.02s  0.94%  runtime.findObject
         0.02s  0.94% 86.85%      0.03s  1.41%  runtime.newMarkBits
         0.02s  0.94% 87.79%      0.02s  0.94%  runtime.unlock
         0.01s  0.47% 88.26%      0.33s 15.49%  internal/poll.(*FD).Init
         0.01s  0.47% 88.73%      0.51s 23.94%  internal/poll.(*FD).Read
         0.01s  0.47% 89.20%      0.32s 15.02%  internal/poll.(*pollDesc).init
         0.01s  0.47% 89.67%      0.65s 30.52%  os.openFileNolog
         0.01s  0.47% 90.14%      0.02s  0.94%  runtime.(*mcentral).freeSpan
         0.01s  0.47% 90.61%      0.04s  1.88%  runtime.(*mspan).sweep
         0.01s  0.47% 91.08%      0.03s  1.41%  runtime.goschedImpl
         0.01s  0.47% 91.55%      0.03s  1.41%  runtime.reentersyscall
         0.01s  0.47% 92.02%      0.32s 15.02%  runtime.systemstack
             0     0% 92.02%      1.91s 89.67%  github.com/sj14/diff.BenchmarkCompare
             0     0% 92.02%      0.11s  5.16%  internal/poll.(*FD).Close
             0     0% 92.02%      0.11s  5.16%  internal/poll.(*FD).decref
             0     0% 92.02%      0.11s  5.16%  internal/poll.(*FD).destroy
             0     0% 92.02%      0.12s  5.63%  internal/poll.runtime_pollClose
             0     0% 92.02%      0.14s  6.57%  internal/poll.runtime_pollOpen
             0     0% 92.02%      0.13s  6.10%  os.(*File).Close
             0     0% 92.02%      0.13s  6.10%  os.(*file).close
             0     0% 92.02%      0.65s 30.52%  os.Open
             0     0% 92.02%      0.65s 30.52%  os.OpenFile
             0     0% 92.02%      0.39s 18.31%  os.newFile
             0     0% 92.02%      0.14s  6.57%  runtime.(*mcache).nextFree
             0     0% 92.02%      0.14s  6.57%  runtime.(*mcache).nextFree.func1
             0     0% 92.02%      0.14s  6.57%  runtime.(*mcache).refill
             0     0% 92.02%      0.14s  6.57%  runtime.(*mcentral).cacheSpan
             0     0% 92.02%      0.10s  4.69%  runtime.(*mcentral).grow
             0     0% 92.02%      0.07s  3.29%  runtime.(*mheap).alloc
             0     0% 92.02%      0.05s  2.35%  runtime.SetFinalizer
             0     0% 92.02%      0.02s  0.94%  runtime.SetFinalizer.func1
             0     0% 92.02%      0.03s  1.41%  runtime._System
             0     0% 92.02%      0.06s  2.82%  runtime.bgsweep
             0     0% 92.02%      0.03s  1.41%  runtime.deferproc
             0     0% 92.02%      0.03s  1.41%  runtime.entersyscall
             0     0% 92.02%      0.03s  1.41%  runtime.findrunnable
             0     0% 92.02%      0.03s  1.41%  runtime.futexsleep
             0     0% 92.02%      0.03s  1.41%  runtime.futexwakeup
             0     0% 92.02%      0.02s  0.94%  runtime.gcAssistAlloc
             0     0% 92.02%      0.02s  0.94%  runtime.gcAssistAlloc.func1
             0     0% 92.02%      0.02s  0.94%  runtime.gcAssistAlloc1
             0     0% 92.02%      0.05s  2.35%  runtime.gcBgMarkWorker
             0     0% 92.02%      0.03s  1.41%  runtime.gcBgMarkWorker.func2
             0     0% 92.02%      0.04s  1.88%  runtime.gcDrain
             0     0% 92.02%      0.02s  0.94%  runtime.gcDrainN
             0     0% 92.02%      0.02s  0.94%  runtime.gcMarkDone
             0     0% 92.02%      0.02s  0.94%  runtime.gcMarkTermination
             0     0% 92.02%      0.02s  0.94%  runtime.gentraceback
             0     0% 92.02%      0.02s  0.94%  runtime.gosched_m
             0     0% 92.02%      0.06s  2.82%  runtime.gosweepone
             0     0% 92.02%      0.06s  2.82%  runtime.gosweepone.func1
             0     0% 92.02%      0.03s  1.41%  runtime.heapBits.initSpan
             0     0% 92.02%      0.17s  7.98%  runtime.makeslice
             0     0% 92.02%      0.02s  0.94%  runtime.markroot
             0     0% 92.02%      0.06s  2.82%  runtime.mcall
             0     0% 92.02%      0.12s  5.63%  runtime.netpollclose
             0     0% 92.02%      0.12s  5.63%  runtime.netpollopen
             0     0% 92.02%      0.03s  1.41%  runtime.newobject
             0     0% 92.02%      0.02s  0.94%  runtime.notesleep
             0     0% 92.02%      0.03s  1.41%  runtime.notewakeup
             0     0% 92.02%      0.04s  1.88%  runtime.park_m
             0     0% 92.02%      0.02s  0.94%  runtime.removefinalizer
             0     0% 92.02%      0.02s  0.94%  runtime.resetspinning
             0     0% 92.02%      0.06s  2.82%  runtime.schedule
             0     0% 92.02%      0.02s  0.94%  runtime.startm
             0     0% 92.02%      0.03s  1.41%  runtime.stopm
             0     0% 92.02%      0.02s  0.94%  runtime.wakep
             0     0% 92.02%      0.11s  5.16%  syscall.Close
             0     0% 92.02%      0.25s 11.74%  syscall.Open
             0     0% 92.02%      0.46s 21.60%  syscall.Read
             0     0% 92.02%      0.25s 11.74%  syscall.openat
             0     0% 92.02%      0.46s 21.60%  syscall.read
             0     0% 92.02%      1.91s 89.67%  testing.(*B).launch
             0     0% 92.02%      1.91s 89.67%  testing.(*B).runN
