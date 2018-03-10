# diff
Go package to get differences of two files

# bench
    BenchmarkCompare-8          3000            346911 ns/op
    ok      github.com/sj14/diff    1.088s
    
    BenchmarkCompare-8          3000            380519 ns/op
    ok      github.com/sj14/diff    1.190s

    BenchmarkCompare-8          3000            363884 ns/op
    ok      github.com/sj14/diff    1.138s

# memory profile
    Type: inuse_space
    Time: Mar 10, 2018 at 1:19pm (CET)
    Showing nodes accounting for 512.03kB, 100% of 512.03kB total
          flat  flat%   sum%        cum   cum%
      512.03kB   100%   100%   512.03kB   100%  flag.(*FlagSet).Var
             0     0%   100%   512.03kB   100%  flag.(*FlagSet).String
             0     0%   100%   512.03kB   100%  flag.(*FlagSet).StringVar
             0     0%   100%   512.03kB   100%  flag.String
             0     0%   100%   512.03kB   100%  main.init
             0     0%   100%   512.03kB   100%  runtime.main
             0     0%   100%   512.03kB   100%  testing.init


# cpu profile
    Type: cpu
    Time: Mar 10, 2018 at 1:15pm (CET)
    Duration: 1.40s, Total samples = 1.40s (99.89%)
    Showing nodes accounting for 1.40s, 100% of 1.40s total
          flat  flat%   sum%        cum   cum%
         0.39s 27.86% 27.86%      0.41s 29.29%  syscall.Syscall
         0.26s 18.57% 46.43%      0.28s 20.00%  syscall.Syscall6
         0.18s 12.86% 59.29%      0.18s 12.86%  runtime.epollctl
         0.05s  3.57% 62.86%      0.05s  3.57%  runtime.memclrNoHeapPointers
         0.04s  2.86% 65.71%      0.04s  2.86%  runtime.unlock
         0.03s  2.14% 67.86%      0.03s  2.14%  runtime.scanobject
         0.02s  1.43% 69.29%      0.39s 27.86%  bufio.(*Scanner).Scan
         0.02s  1.43% 70.71%      0.02s  1.43%  bufio.(*Scanner).advance (inline)
         0.02s  1.43% 72.14%      0.02s  1.43%  runtime.(*mheap).freeSpanLocked
         0.02s  1.43% 73.57%      0.05s  3.57%  runtime.(*mspan).sweep
         0.02s  1.43% 75.00%      0.02s  1.43%  runtime.findObject
         0.02s  1.43% 76.43%      0.02s  1.43%  runtime.reentersyscall
         0.02s  1.43% 77.86%      0.02s  1.43%  runtime.scanblock
         0.01s  0.71% 78.57%      0.01s  0.71%  bufio.(*Scanner).Split
         0.01s  0.71% 79.29%      0.01s  0.71%  bytes.IndexByte
         0.01s  0.71% 80.00%      1.26s 90.00%  github.com/sj14/diff.getLine
         0.01s  0.71% 80.71%      0.23s 16.43%  internal/poll.(*pollDesc).init
         0.01s  0.71% 81.43%      0.01s  0.71%  runtime.(*gcBitsArena).tryAlloc (inline)
         0.01s  0.71% 82.14%      0.01s  0.71%  runtime.(*gcSweepBuf).pop (inline)
         0.01s  0.71% 82.86%      0.03s  2.14%  runtime.(*mheap).allocSpanLocked
         0.01s  0.71% 83.57%      0.04s  2.86%  runtime.(*mheap).alloc_m
         0.01s  0.71% 84.29%      0.01s  0.71%  runtime.(*mheap).lookupMaybe (inline)
         0.01s  0.71% 85.00%      0.01s  0.71%  runtime.(*mspan).countAlloc
         0.01s  0.71% 85.71%      0.01s  0.71%  runtime.(*mspan).ensureSwept
         0.01s  0.71% 86.43%      0.02s  1.43%  runtime.(*pollCache).free
         0.01s  0.71% 87.14%      0.07s  5.00%  runtime.SetFinalizer
         0.01s  0.71% 87.86%      0.01s  0.71%  runtime._ExternalCode
         0.01s  0.71% 88.57%      0.01s  0.71%  runtime.atomicstorep
         0.01s  0.71% 89.29%      0.02s  1.43%  runtime.exitsyscall
         0.01s  0.71% 90.00%      0.01s  0.71%  runtime.exitsyscallfast
         0.01s  0.71% 90.71%      0.01s  0.71%  runtime.futex
         0.01s  0.71% 91.43%      0.01s  0.71%  runtime.gomcache (inline)
         0.01s  0.71% 92.14%      0.07s  5.00%  runtime.gosweepone
         0.01s  0.71% 92.86%      0.06s  4.29%  runtime.gosweepone.func1
         0.01s  0.71% 93.57%      0.01s  0.71%  runtime.indexbytebody
         0.01s  0.71% 94.29%      0.01s  0.71%  runtime.lock
         0.01s  0.71% 95.00%      0.01s  0.71%  runtime.mProf_Free
         0.01s  0.71% 95.71%      0.17s 12.14%  runtime.mallocgc
         0.01s  0.71% 96.43%      0.01s  0.71%  runtime.memmove
         0.01s  0.71% 97.14%      0.01s  0.71%  runtime.nextFreeFast (inline)
         0.01s  0.71% 97.86%      0.26s 18.57%  runtime.systemstack
         0.01s  0.71% 98.57%      0.01s  0.71%  runtime.usleep
         0.01s  0.71% 99.29%      0.01s  0.71%  runtime.wbBufFlush1
         0.01s  0.71%   100%      0.02s  1.43%  syscall.ByteSliceFromString
             0     0%   100%      0.04s  2.86%  bufio.(*Scanner).Text (inline)
             0     0%   100%      0.02s  1.43%  bufio.ScanLines
             0     0%   100%      1.26s 90.00%  github.com/sj14/diff.BenchmarkCompare
             0     0%   100%      1.26s 90.00%  github.com/sj14/diff.Compare
             0     0%   100%      0.18s 12.86%  internal/poll.(*FD).Close
             0     0%   100%      0.23s 16.43%  internal/poll.(*FD).Init
             0     0%   100%      0.23s 16.43%  internal/poll.(*FD).Read
             0     0%   100%      0.18s 12.86%  internal/poll.(*FD).decref
             0     0%   100%      0.18s 12.86%  internal/poll.(*FD).destroy
             0     0%   100%      0.11s  7.86%  internal/poll.runtime_pollClose
             0     0%   100%      0.09s  6.43%  internal/poll.runtime_pollOpen
             0     0%   100%      0.02s  1.43%  internal/poll.runtime_pollUnblock
             0     0%   100%      0.20s 14.29%  os.(*File).Close
             0     0%   100%      0.23s 16.43%  os.(*File).Read
             0     0%   100%      0.23s 16.43%  os.(*File).read
             0     0%   100%      0.20s 14.29%  os.(*file).close
             0     0%   100%      0.60s 42.86%  os.Open
             0     0%   100%      0.60s 42.86%  os.OpenFile
             0     0%   100%      0.30s 21.43%  os.newFile
             0     0%   100%      0.60s 42.86%  os.openFileNolog
             0     0%   100%      0.11s  7.86%  runtime.(*mcache).nextFree
             0     0%   100%      0.11s  7.86%  runtime.(*mcache).nextFree.func1
             0     0%   100%      0.11s  7.86%  runtime.(*mcache).refill
             0     0%   100%      0.11s  7.86%  runtime.(*mcentral).cacheSpan
             0     0%   100%      0.01s  0.71%  runtime.(*mcentral).freeSpan
             0     0%   100%      0.09s  6.43%  runtime.(*mcentral).grow
             0     0%   100%      0.08s  5.71%  runtime.(*mheap).alloc
             0     0%   100%      0.04s  2.86%  runtime.(*mheap).alloc.func1
             0     0%   100%      0.01s  0.71%  runtime.SetFinalizer.func1
             0     0%   100%      0.03s  2.14%  runtime.SetFinalizer.func2
             0     0%   100%      0.02s  1.43%  runtime._System
             0     0%   100%      0.03s  2.14%  runtime.addfinalizer
             0     0%   100%      0.01s  0.71%  runtime.addspecial
             0     0%   100%      0.06s  4.29%  runtime.bgsweep
             0     0%   100%      0.01s  0.71%  runtime.deductSweepCredit
             0     0%   100%      0.01s  0.71%  runtime.deferreturn
             0     0%   100%      0.02s  1.43%  runtime.entersyscall
             0     0%   100%      0.01s  0.71%  runtime.execute
             0     0%   100%      0.01s  0.71%  runtime.findrunnable
             0     0%   100%      0.01s  0.71%  runtime.freespecial
             0     0%   100%      0.01s  0.71%  runtime.futexsleep
             0     0%   100%      0.02s  1.43%  runtime.gcAssistAlloc
             0     0%   100%      0.02s  1.43%  runtime.gcAssistAlloc.func1
             0     0%   100%      0.02s  1.43%  runtime.gcAssistAlloc1
             0     0%   100%      0.03s  2.14%  runtime.gcBgMarkWorker
             0     0%   100%      0.03s  2.14%  runtime.gcBgMarkWorker.func2
             0     0%   100%      0.03s  2.14%  runtime.gcDrain
             0     0%   100%      0.02s  1.43%  runtime.gcDrainN
             0     0%   100%      0.01s  0.71%  runtime.gcWriteBarrier
             0     0%   100%      0.01s  0.71%  runtime.gopreempt_m
             0     0%   100%      0.01s  0.71%  runtime.goschedImpl
             0     0%   100%      0.01s  0.71%  runtime.heapBits.initSpan
             0     0%   100%      0.11s  7.86%  runtime.makeslice
             0     0%   100%      0.02s  1.43%  runtime.markroot
             0     0%   100%      0.02s  1.43%  runtime.markrootBlock
             0     0%   100%      0.01s  0.71%  runtime.mcall
             0     0%   100%      0.01s  0.71%  runtime.morestack
             0     0%   100%      0.01s  0.71%  runtime.mstart
             0     0%   100%      0.01s  0.71%  runtime.mstart1
             0     0%   100%      0.09s  6.43%  runtime.netpollclose
             0     0%   100%      0.09s  6.43%  runtime.netpollopen
             0     0%   100%      0.01s  0.71%  runtime.newMarkBits
             0     0%   100%      0.02s  1.43%  runtime.newobject
             0     0%   100%      0.01s  0.71%  runtime.newstack
             0     0%   100%      0.01s  0.71%  runtime.notesleep
             0     0%   100%      0.01s  0.71%  runtime.park_m
             0     0%   100%      0.01s  0.71%  runtime.removefinalizer
             0     0%   100%      0.01s  0.71%  runtime.removespecial
             0     0%   100%      0.02s  1.43%  runtime.schedule
             0     0%   100%      0.04s  2.86%  runtime.slicebytetostring
             0     0%   100%      0.01s  0.71%  runtime.stopm
             0     0%   100%      0.05s  3.57%  runtime.sweepone
             0     0%   100%      0.01s  0.71%  runtime.sysmon
             0     0%   100%      0.01s  0.71%  runtime.wbBufFlush
             0     0%   100%      0.01s  0.71%  runtime.wbBufFlush.func1
             0     0%   100%      0.02s  1.43%  syscall.BytePtrFromString
             0     0%   100%      0.18s 12.86%  syscall.Close
             0     0%   100%      0.30s 21.43%  syscall.Open
             0     0%   100%      0.23s 16.43%  syscall.Read
             0     0%   100%      0.30s 21.43%  syscall.openat
             0     0%   100%      0.23s 16.43%  syscall.read
             0     0%   100%      1.26s 90.00%  testing.(*B).launch
             0     0%   100%      1.26s 90.00%  testing.(*B).runN
