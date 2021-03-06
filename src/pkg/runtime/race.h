// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Definitions related to data race detection.

#ifdef RACE
enum { raceenabled = 1 };
#else
enum { raceenabled = 0 };
#endif

// Initialize race detection subsystem.
void	runtime·raceinit(void);
// Finalize race detection subsystem, does not return.
void	runtime·racefini(void);

void	runtime·racemalloc(void *p, uintptr sz, void *pc);
void	runtime·racefree(void *p);
void	runtime·racegostart(int32 goid, void *pc);
void	runtime·racegoend(int32 goid);
void	runtime·racewritepc(void *addr, void *pc);
void	runtime·racereadpc(void *addr, void *pc);
void	runtime·racefingo(void);
void	runtime·raceacquire(void *addr);
void	runtime·raceacquireg(G *gp, void *addr);
void	runtime·racerelease(void *addr);
void	runtime·racereleaseg(G *gp, void *addr);
void	runtime·racereleasemerge(void *addr);
void	runtime·racereleasemergeg(G *gp, void *addr);
