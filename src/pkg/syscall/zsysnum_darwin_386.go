// mksysnum_darwin.sh /home/rsc/pub/xnu-1228/bsd/kern/syscalls.master
// MACHINE GENERATED BY THE ABOVE COMMAND; DO NOT EDIT

package syscall

const (
	// SYS_NOSYS = 0;  // { int nosys(void); }   { indirect syscall }
	SYS_EXIT  = 1 // { void exit(int rval); }
	SYS_FORK  = 2 // { int fork(void); }
	SYS_READ  = 3 // { user_ssize_t read(int fd, user_addr_t cbuf, user_size_t nbyte); }
	SYS_WRITE = 4 // { user_ssize_t write(int fd, user_addr_t cbuf, user_size_t nbyte); }
	SYS_OPEN  = 5 // { int open(user_addr_t path, int flags, int mode); }
	SYS_CLOSE = 6 // { int close(int fd); }
	SYS_WAIT4 = 7 // { int wait4(int pid, user_addr_t status, int options, user_addr_t rusage); }
	// SYS_NOSYS = 8;  // { int nosys(void); }   { old creat }
	SYS_LINK   = 9  // { int link(user_addr_t path, user_addr_t link); }
	SYS_UNLINK = 10 // { int unlink(user_addr_t path); }
	// SYS_NOSYS = 11;  // { int nosys(void); }   { old execv }
	SYS_CHDIR      = 12 // { int chdir(user_addr_t path); }
	SYS_FCHDIR     = 13 // { int fchdir(int fd); }
	SYS_MKNOD      = 14 // { int mknod(user_addr_t path, int mode, int dev); }
	SYS_CHMOD      = 15 // { int chmod(user_addr_t path, int mode); }
	SYS_CHOWN      = 16 // { int chown(user_addr_t path, int uid, int gid); }
	SYS_OGETFSSTAT = 18 // { int ogetfsstat(user_addr_t buf, int bufsize, int flags); }
	SYS_GETFSSTAT  = 18 // { int getfsstat(user_addr_t buf, int bufsize, int flags); }
	// SYS_NOSYS = 19;  // { int nosys(void); }   { old lseek }
	SYS_GETPID = 20 // { int getpid(void); }
	// SYS_NOSYS = 21;  // { int nosys(void); }   { old mount }
	// SYS_NOSYS = 22;  // { int nosys(void); }   { old umount }
	SYS_SETUID      = 23 // { int setuid(uid_t uid); }
	SYS_GETUID      = 24 // { int getuid(void); }
	SYS_GETEUID     = 25 // { int geteuid(void); }
	SYS_PTRACE      = 26 // { int ptrace(int req, pid_t pid, caddr_t addr, int data); }
	SYS_RECVMSG     = 27 // { int recvmsg(int s, struct msghdr *msg, int flags); }
	SYS_SENDMSG     = 28 // { int sendmsg(int s, caddr_t msg, int flags); }
	SYS_RECVFROM    = 29 // { int recvfrom(int s, void *buf, size_t len, int flags, struct sockaddr *from, int *fromlenaddr); }
	SYS_ACCEPT      = 30 // { int accept(int s, caddr_t name, socklen_t	*anamelen); }
	SYS_GETPEERNAME = 31 // { int getpeername(int fdes, caddr_t asa, socklen_t *alen); }
	SYS_GETSOCKNAME = 32 // { int getsockname(int fdes, caddr_t asa, socklen_t *alen); }
	// SYS_NOSYS = 27;  // { int nosys(void); }
	// SYS_NOSYS = 28;  // { int nosys(void); }
	// SYS_NOSYS = 29;  // { int nosys(void); }
	// SYS_NOSYS = 30;  // { int nosys(void); }
	// SYS_NOSYS = 31;  // { int nosys(void); }
	// SYS_NOSYS = 32;  // { int nosys(void); }
	SYS_ACCESS   = 33 // { int access(user_addr_t path, int flags); }
	SYS_CHFLAGS  = 34 // { int chflags(char *path, int flags); }
	SYS_FCHFLAGS = 35 // { int fchflags(int fd, int flags); }
	SYS_SYNC     = 36 // { int sync(void); }
	SYS_KILL     = 37 // { int kill(int pid, int signum, int posix); }
	// SYS_NOSYS = 38;  // { int nosys(void); }   { old stat  }
	SYS_GETPPID = 39 // { int getppid(void); }
	// SYS_NOSYS = 40;  // { int nosys(void); }   { old lstat }
	SYS_DUP     = 41 // { int dup(u_int fd); }
	SYS_PIPE    = 42 // { int pipe(void); }
	SYS_GETEGID = 43 // { int getegid(void); }
	SYS_PROFIL  = 44 // { int profil(short *bufbase, size_t bufsize, u_long pcoffset, u_int pcscale); }
	// SYS_NOSYS = 45;  // { int nosys(void); } { old ktrace }
	SYS_SIGACTION   = 46 // { int sigaction(int signum, struct __sigaction *nsa, struct sigaction *osa); }
	SYS_GETGID      = 47 // { int getgid(void); }
	SYS_SIGPROCMASK = 48 // { int sigprocmask(int how, user_addr_t mask, user_addr_t omask); }
	SYS_GETLOGIN    = 49 // { int getlogin(char *namebuf, u_int namelen); }
	SYS_SETLOGIN    = 50 // { int setlogin(char *namebuf); }
	SYS_ACCT        = 51 // { int acct(char *path); }
	SYS_SIGPENDING  = 52 // { int sigpending(struct sigvec *osv); }
	SYS_SIGALTSTACK = 53 // { int sigaltstack(struct sigaltstack *nss, struct sigaltstack *oss); }
	SYS_IOCTL       = 54 // { int ioctl(int fd, u_long com, caddr_t data); }
	SYS_REBOOT      = 55 // { int reboot(int opt, char *command); }
	SYS_REVOKE      = 56 // { int revoke(char *path); }
	SYS_SYMLINK     = 57 // { int symlink(char *path, char *link); }
	SYS_READLINK    = 58 // { int readlink(char *path, char *buf, int count); }
	SYS_EXECVE      = 59 // { int execve(char *fname, char **argp, char **envp); }
	SYS_UMASK       = 60 // { int umask(int newmask); }
	SYS_CHROOT      = 61 // { int chroot(user_addr_t path); }
	// SYS_NOSYS = 62;  // { int nosys(void); }   { old fstat }
	// SYS_NOSYS = 63;  // { int nosys(void); }   { used internally, reserved }
	// SYS_NOSYS = 64;  // { int nosys(void); }   { old getpagesize }
	SYS_MSYNC = 65 // { int msync(caddr_t addr, size_t len, int flags); }
	SYS_VFORK = 66 // { int vfork(void); }
	// SYS_NOSYS = 67;  // { int nosys(void); }   { old vread }
	// SYS_NOSYS = 68;  // { int nosys(void); }   { old vwrite }
	SYS_SBRK = 69 // { int sbrk(int incr) NO_SYSCALL_STUB; }
	SYS_SSTK = 70 // { int sstk(int incr) NO_SYSCALL_STUB; }
	// SYS_NOSYS = 71;  // { int nosys(void); }   { old mmap }
	SYS_OVADVISE = 72 // { int ovadvise(void) NO_SYSCALL_STUB; }   { old vadvise }
	SYS_MUNMAP   = 73 // { int munmap(caddr_t addr, size_t len); }
	SYS_MPROTECT = 74 // { int mprotect(caddr_t addr, size_t len, int prot); }
	SYS_MADVISE  = 75 // { int madvise(caddr_t addr, size_t len, int behav); }
	// SYS_NOSYS = 76;  // { int nosys(void); }   { old vhangup }
	// SYS_NOSYS = 77;  // { int nosys(void); }   { old vlimit }
	SYS_MINCORE   = 78 // { int mincore(user_addr_t addr, user_size_t len, user_addr_t vec); }
	SYS_GETGROUPS = 79 // { int getgroups(u_int gidsetsize, gid_t *gidset); }
	SYS_SETGROUPS = 80 // { int setgroups(u_int gidsetsize, gid_t *gidset); }
	SYS_GETPGRP   = 81 // { int getpgrp(void); }
	SYS_SETPGID   = 82 // { int setpgid(int pid, int pgid); }
	SYS_SETITIMER = 83 // { int setitimer(u_int which, struct itimerval *itv, struct itimerval *oitv); }
	// SYS_NOSYS = 84;  // { int nosys(void); }   { old wait }
	SYS_SWAPON    = 85 // { int swapon(void); }
	SYS_GETITIMER = 86 // { int getitimer(u_int which, struct itimerval *itv); }
	// SYS_NOSYS = 87;  // { int nosys(void); }   { old gethostname }
	// SYS_NOSYS = 88;  // { int nosys(void); }   { old sethostname }
	SYS_GETDTABLESIZE = 89 // { int getdtablesize(void); }
	SYS_DUP2          = 90 // { int dup2(u_int from, u_int to); }
	// SYS_NOSYS = 91;  // { int nosys(void); }   { old getdopt }
	SYS_FCNTL  = 92 // { int fcntl(int fd, int cmd, long arg); }
	SYS_SELECT = 93 // { int select(int nd, u_int32_t *in, u_int32_t *ou, u_int32_t *ex, struct timeval *tv); }
	// SYS_NOSYS = 94;  // { int nosys(void); }   { old setdopt }
	SYS_FSYNC       = 95 // { int fsync(int fd); }
	SYS_SETPRIORITY = 96 // { int setpriority(int which, id_t who, int prio); }
	SYS_SOCKET      = 97 // { int socket(int domain, int type, int protocol); }
	SYS_CONNECT     = 98 // { int connect(int s, caddr_t name, socklen_t namelen); }
	// SYS_NOSYS = 97;  // { int nosys(void); }
	// SYS_NOSYS = 98;  // { int nosys(void); }
	// SYS_NOSYS = 99;  // { int nosys(void); }   { old accept }
	SYS_GETPRIORITY = 100 // { int getpriority(int which, id_t who); }
	// SYS_NOSYS = 101;  // { int nosys(void); }   { old send }
	// SYS_NOSYS = 102;  // { int nosys(void); }   { old recv }
	// SYS_NOSYS = 103;  // { int nosys(void); }   { old sigreturn }
	SYS_BIND       = 104 // { int bind(int s, caddr_t name, socklen_t namelen); }
	SYS_SETSOCKOPT = 105 // { int setsockopt(int s, int level, int name, caddr_t val, socklen_t valsize); }
	SYS_LISTEN     = 106 // { int listen(int s, int backlog); }
	// SYS_NOSYS = 104;  // { int nosys(void); }
	// SYS_NOSYS = 105;  // { int nosys(void); }
	// SYS_NOSYS = 106;  // { int nosys(void); }
	// SYS_NOSYS = 107;  // { int nosys(void); }   { old vtimes }
	// SYS_NOSYS = 108;  // { int nosys(void); }   { old sigvec }
	// SYS_NOSYS = 109;  // { int nosys(void); }   { old sigblock }
	// SYS_NOSYS = 110;  // { int nosys(void); }   { old sigsetmask }
	SYS_SIGSUSPEND = 111 // { int sigsuspend(sigset_t mask); }
	// SYS_NOSYS = 112;  // { int nosys(void); }   { old sigstack }
	// SYS_NOSYS = 113;  // { int nosys(void); }   { old recvmsg }
	// SYS_NOSYS = 114;  // { int nosys(void); }   { old sendmsg }
	// SYS_NOSYS = 113;  // { int nosys(void); }
	// SYS_NOSYS = 114;  // { int nosys(void); }
	// SYS_NOSYS = 115;  // { int nosys(void); }   { old vtrace }
	SYS_GETTIMEOFDAY = 116 // { int gettimeofday(struct timeval *tp, struct timezone *tzp); }
	SYS_GETRUSAGE    = 117 // { int getrusage(int who, struct rusage *rusage); }
	SYS_GETSOCKOPT   = 118 // { int getsockopt(int s, int level, int name, caddr_t val, socklen_t *avalsize); }
	// SYS_NOSYS = 118;  // { int nosys(void); }
	// SYS_NOSYS = 119;  // { int nosys(void); }   { old resuba }
	SYS_READV        = 120 // { user_ssize_t readv(int fd, struct iovec *iovp, u_int iovcnt); }
	SYS_WRITEV       = 121 // { user_ssize_t writev(int fd, struct iovec *iovp, u_int iovcnt); }
	SYS_SETTIMEOFDAY = 122 // { int settimeofday(struct timeval *tv, struct timezone *tzp); }
	SYS_FCHOWN       = 123 // { int fchown(int fd, int uid, int gid); }
	SYS_FCHMOD       = 124 // { int fchmod(int fd, int mode); }
	// SYS_NOSYS = 125;  // { int nosys(void); }   { old recvfrom }
	SYS_SETREUID = 126 // { int setreuid(uid_t ruid, uid_t euid); }
	SYS_SETREGID = 127 // { int setregid(gid_t rgid, gid_t egid); }
	SYS_RENAME   = 128 // { int rename(char *from, char *to); }
	// SYS_NOSYS = 129;  // { int nosys(void); }   { old truncate }
	// SYS_NOSYS = 130;  // { int nosys(void); }   { old ftruncate }
	SYS_FLOCK      = 131 // { int flock(int fd, int how); }
	SYS_MKFIFO     = 132 // { int mkfifo(user_addr_t path, int mode); }
	SYS_SENDTO     = 133 // { int sendto(int s, caddr_t buf, size_t len, int flags, caddr_t to, socklen_t tolen); }
	SYS_SHUTDOWN   = 134 // { int shutdown(int s, int how); }
	SYS_SOCKETPAIR = 135 // { int socketpair(int domain, int type, int protocol, int *rsv); }
	// SYS_NOSYS = 133;  // { int nosys(void); }
	// SYS_NOSYS = 134;  // { int nosys(void); }
	// SYS_NOSYS = 135;  // { int nosys(void); }
	SYS_MKDIR   = 136 // { int mkdir(user_addr_t path, int mode); }
	SYS_RMDIR   = 137 // { int rmdir(char *path); }
	SYS_UTIMES  = 138 // { int utimes(char *path, struct timeval *tptr); }
	SYS_FUTIMES = 139 // { int futimes(int fd, struct timeval *tptr); }
	SYS_ADJTIME = 140 // { int adjtime(struct timeval *delta, struct timeval *olddelta); }
	// SYS_NOSYS = 141;  // { int nosys(void); }   { old getpeername }
	SYS_GETHOSTUUID = 142 // { int gethostuuid(unsigned char *uuid_buf, const struct timespec *timeoutp); }
	// SYS_NOSYS = 143;  // { int nosys(void); }   { old sethostid 	}
	// SYS_NOSYS = 144;  // { int nosys(void); }   { old getrlimit }
	// SYS_NOSYS = 145;  // { int nosys(void); }   { old setrlimit }
	// SYS_NOSYS = 146;  // { int nosys(void); }   { old killpg }
	SYS_SETSID = 147 // { int setsid(void); }
	// SYS_NOSYS = 148;  // { int nosys(void); }   { old setquota }
	// SYS_NOSYS = 149;  // { int nosys(void); }   { old qquota }
	// SYS_NOSYS = 150;  // { int nosys(void); }   { old getsockname }
	SYS_GETPGID     = 151 // { int getpgid(pid_t pid); }
	SYS_SETPRIVEXEC = 152 // { int setprivexec(int flag); }
	SYS_PREAD       = 153 // { user_ssize_t pread(int fd, user_addr_t buf, user_size_t nbyte, off_t offset); }
	SYS_PWRITE      = 154 // { user_ssize_t pwrite(int fd, user_addr_t buf, user_size_t nbyte, off_t offset); }
	SYS_NFSSVC      = 155 // { int nfssvc(int flag, caddr_t argp); }
	// SYS_NOSYS = 155;  // { int nosys(void); }
	// SYS_NOSYS = 156;  // { int nosys(void); }   { old getdirentries }
	SYS_STATFS  = 157 // { int statfs(char *path, struct statfs *buf); }
	SYS_FSTATFS = 158 // { int fstatfs(int fd, struct statfs *buf); }
	SYS_UNMOUNT = 159 // { int unmount(user_addr_t path, int flags); }
	// SYS_NOSYS = 160;  // { int nosys(void); }   { old async_daemon }
	SYS_GETFH = 161 // { int getfh(char *fname, fhandle_t *fhp); }
	// SYS_NOSYS = 161;  // { int nosys(void); }
	// SYS_NOSYS = 162;  // { int nosys(void); }   { old getdomainname }
	// SYS_NOSYS = 163;  // { int nosys(void); }   { old setdomainname }
	// SYS_NOSYS = 164;  // { int nosys(void); }
	SYS_QUOTACTL = 165 // { int quotactl(const char *path, int cmd, int uid, caddr_t arg); }
	// SYS_NOSYS = 166;  // { int nosys(void); }   { old exportfs }
	SYS_MOUNT = 167 // { int mount(char *type, char *path, int flags, caddr_t data); }
	// SYS_NOSYS = 168;  // { int nosys(void); }   { old ustat }
	SYS_CSOPS = 169 // { int csops(pid_t pid, uint32_t ops, user_addr_t useraddr, user_size_t usersize); }
	// SYS_NOSYS = 171;  // { int nosys(void); }   { old wait3 }
	// SYS_NOSYS = 172;  // { int nosys(void); }   { old rpause	}
	SYS_WAITID = 173 // { int waitid(idtype_t idtype, id_t id, siginfo_t *infop, int options); }
	// SYS_NOSYS = 174;  // { int nosys(void); }   { old getdents }
	// SYS_NOSYS = 175;  // { int nosys(void); }   { old gc_control }
	SYS_ADD_PROFIL = 176 // { int add_profil(short *bufbase, size_t bufsize, u_long pcoffset, u_int pcscale); }
	// SYS_NOSYS = 177;  // { int nosys(void); }
	// SYS_NOSYS = 178;  // { int nosys(void); }
	// SYS_NOSYS = 179;  // { int nosys(void); }
	SYS_KDEBUG_TRACE = 180 // { int kdebug_trace(int code, int arg1, int arg2, int arg3, int arg4, int arg5) NO_SYSCALL_STUB; }
	SYS_SETGID       = 181 // { int setgid(gid_t gid); }
	SYS_SETEGID      = 182 // { int setegid(gid_t egid); }
	SYS_SETEUID      = 183 // { int seteuid(uid_t euid); }
	SYS_SIGRETURN    = 184 // { int sigreturn(struct ucontext *uctx, int infostyle); }
	// SYS_NOSYS = 186;  // { int nosys(void); }
	// SYS_NOSYS = 187;  // { int nosys(void); }
	SYS_STAT      = 188 // { int stat(user_addr_t path, user_addr_t ub); }
	SYS_FSTAT     = 189 // { int fstat(int fd, user_addr_t ub); }
	SYS_LSTAT     = 190 // { int lstat(user_addr_t path, user_addr_t ub); }
	SYS_PATHCONF  = 191 // { int pathconf(char *path, int name); }
	SYS_FPATHCONF = 192 // { int fpathconf(int fd, int name); }
	// SYS_NOSYS = 193;  // { int nosys(void); }
	SYS_GETRLIMIT     = 194 // { int getrlimit(u_int which, struct rlimit *rlp); }
	SYS_SETRLIMIT     = 195 // { int setrlimit(u_int which, struct rlimit *rlp); }
	SYS_GETDIRENTRIES = 196 // { int getdirentries(int fd, char *buf, u_int count, long *basep); }
	SYS_MMAP          = 197 // { user_addr_t mmap(caddr_t addr, size_t len, int prot, int flags, int fd, off_t pos); }
	// SYS_NOSYS = 198;  // { int nosys(void); } 	{ __syscall }
	SYS_LSEEK     = 199 // { off_t lseek(int fd, off_t offset, int whence); }
	SYS_TRUNCATE  = 200 // { int truncate(char *path, off_t length); }
	SYS_FTRUNCATE = 201 // { int ftruncate(int fd, off_t length); }
	SYS___SYSCTL  = 202 // { int __sysctl(int *name, u_int namelen, void *old, size_t *oldlenp, void *new, size_t newlen); }
	SYS_MLOCK     = 203 // { int mlock(caddr_t addr, size_t len); }
	SYS_MUNLOCK   = 204 // { int munlock(caddr_t addr, size_t len); }
	SYS_UNDELETE  = 205 // { int undelete(user_addr_t path); }
	SYS_ATSOCKET  = 206 // { int ATsocket(int proto); }
	// SYS_NOSYS = 213;  // { int nosys(void); } 	{ Reserved for AppleTalk }
	// SYS_NOSYS = 206;  // { int nosys(void); }
	// SYS_NOSYS = 207;  // { int nosys(void); }
	// SYS_NOSYS = 208;  // { int nosys(void); }
	// SYS_NOSYS = 209;  // { int nosys(void); }
	// SYS_NOSYS = 210;  // { int nosys(void); }
	// SYS_NOSYS = 211;  // { int nosys(void); }
	// SYS_NOSYS = 212;  // { int nosys(void); }
	// SYS_NOSYS = 213;  // { int nosys(void); } 	{ Reserved for AppleTalk }
	SYS_KQUEUE_FROM_PORTSET_NP = 214 // { int kqueue_from_portset_np(int portset); }
	SYS_KQUEUE_PORTSET_NP      = 215 // { int kqueue_portset_np(int fd); }
	SYS_GETATTRLIST            = 220 // { int getattrlist(const char *path, struct attrlist *alist, void *attributeBuffer, size_t bufferSize, u_long options); }
	SYS_SETATTRLIST            = 221 // { int setattrlist(const char *path, struct attrlist *alist, void *attributeBuffer, size_t bufferSize, u_long options); }
	SYS_GETDIRENTRIESATTR      = 222 // { int getdirentriesattr(int fd, struct attrlist *alist, void *buffer, size_t buffersize, u_long *count, u_long *basep, u_long *newstate, u_long options); }
	SYS_EXCHANGEDATA           = 223 // { int exchangedata(const char *path1, const char *path2, u_long options); }
	// SYS_NOSYS = 224;  // { int nosys(void); } { was checkuseraccess }
	SYS_SEARCHFS = 225 // { int searchfs(const char *path, struct fssearchblock *searchblock, u_long *nummatches, u_long scriptcode, u_long options, struct searchstate *state); }
	SYS_DELETE   = 226 // { int delete(user_addr_t path) NO_SYSCALL_STUB; }       { private delete (Carbon semantics) }
	SYS_COPYFILE = 227 // { int copyfile(char *from, char *to, int mode, int flags) NO_SYSCALL_STUB; }
	// SYS_NOSYS = 228;  // { int nosys(void); }
	// SYS_NOSYS = 229;  // { int nosys(void); }
	SYS_POLL         = 230 // { int poll(struct pollfd *fds, u_int nfds, int timeout); }
	SYS_WATCHEVENT   = 231 // { int watchevent(struct eventreq *u_req, int u_eventmask); }
	SYS_WAITEVENT    = 232 // { int waitevent(struct eventreq *u_req, struct timeval *tv); }
	SYS_MODWATCH     = 233 // { int modwatch(struct eventreq *u_req, int u_eventmask); }
	SYS_GETXATTR     = 234 // { user_ssize_t getxattr(user_addr_t path, user_addr_t attrname, user_addr_t value, size_t size, uint32_t position, int options); }
	SYS_FGETXATTR    = 235 // { user_ssize_t fgetxattr(int fd, user_addr_t attrname, user_addr_t value, size_t size, uint32_t position, int options); }
	SYS_SETXATTR     = 236 // { int setxattr(user_addr_t path, user_addr_t attrname, user_addr_t value, size_t size, uint32_t position, int options); }
	SYS_FSETXATTR    = 237 // { int fsetxattr(int fd, user_addr_t attrname, user_addr_t value, size_t size, uint32_t position, int options); }
	SYS_REMOVEXATTR  = 238 // { int removexattr(user_addr_t path, user_addr_t attrname, int options); }
	SYS_FREMOVEXATTR = 239 // { int fremovexattr(int fd, user_addr_t attrname, int options); }
	SYS_LISTXATTR    = 240 // { user_ssize_t listxattr(user_addr_t path, user_addr_t namebuf, size_t bufsize, int options); }
	SYS_FLISTXATTR   = 241 // { user_ssize_t flistxattr(int fd, user_addr_t namebuf, size_t bufsize, int options); }
	SYS_FSCTL        = 242 // { int fsctl(const char *path, u_long cmd, caddr_t data, u_long options); }
	SYS_INITGROUPS   = 243 // { int initgroups(u_int gidsetsize, gid_t *gidset, int gmuid); }
	SYS_POSIX_SPAWN  = 244 // { int posix_spawn(pid_t *pid, const char *path, const struct _posix_spawn_args_desc *adesc, char **argv, char **envp); }
	// SYS_NOSYS = 245;  // { int nosys(void); }
	// SYS_NOSYS = 246;  // { int nosys(void); }
	SYS_NFSCLNT = 247 // { int nfsclnt(int flag, caddr_t argp); }
	// SYS_NOSYS = 247;  // { int nosys(void); }
	SYS_FHOPEN = 248 // { int fhopen(const struct fhandle *u_fhp, int flags); }
	// SYS_NOSYS = 248;  // { int nosys(void); }
	// SYS_NOSYS = 249;  // { int nosys(void); }
	SYS_MINHERIT = 250 // { int minherit(void *addr, size_t len, int inherit); }
	SYS_SEMSYS   = 251 // { int semsys(u_int which, int a2, int a3, int a4, int a5); }
	// SYS_NOSYS = 251;  // { int nosys(void); }
	SYS_MSGSYS = 252 // { int msgsys(u_int which, int a2, int a3, int a4, int a5); }
	// SYS_NOSYS = 252;  // { int nosys(void); }
	SYS_SHMSYS = 253 // { int shmsys(u_int which, int a2, int a3, int a4); }
	// SYS_NOSYS = 253;  // { int nosys(void); }
	SYS_SEMCTL = 254 // { int semctl(int semid, int semnum, int cmd, semun_t arg); }
	SYS_SEMGET = 255 // { int semget(key_t key, int	nsems, int semflg); }
	SYS_SEMOP  = 256 // { int semop(int semid, struct sembuf *sops, int nsops); }
	// SYS_NOSYS = 257;  // { int nosys(void); }
	// SYS_NOSYS = 254;  // { int nosys(void); }
	// SYS_NOSYS = 255;  // { int nosys(void); }
	// SYS_NOSYS = 256;  // { int nosys(void); }
	// SYS_NOSYS = 257;  // { int nosys(void); }
	SYS_MSGCTL = 258 // { int msgctl(int msqid, int cmd, struct	msqid_ds *buf); }
	SYS_MSGGET = 259 // { int msgget(key_t key, int msgflg); }
	SYS_MSGSND = 260 // { int msgsnd(int msqid, void *msgp, size_t msgsz, int msgflg); }
	SYS_MSGRCV = 261 // { user_ssize_t msgrcv(int msqid, void *msgp, size_t msgsz, long msgtyp, int msgflg); }
	// SYS_NOSYS = 258;  // { int nosys(void); }
	// SYS_NOSYS = 259;  // { int nosys(void); }
	// SYS_NOSYS = 260;  // { int nosys(void); }
	// SYS_NOSYS = 261;  // { int nosys(void); }
	SYS_SHMAT  = 262 // { user_addr_t shmat(int shmid, void *shmaddr, int shmflg); }
	SYS_SHMCTL = 263 // { int shmctl(int shmid, int cmd, struct shmid_ds *buf); }
	SYS_SHMDT  = 264 // { int shmdt(void *shmaddr); }
	SYS_SHMGET = 265 // { int shmget(key_t key, size_t size, int shmflg); }
	// SYS_NOSYS = 262;  // { int nosys(void); }
	// SYS_NOSYS = 263;  // { int nosys(void); }
	// SYS_NOSYS = 264;  // { int nosys(void); }
	// SYS_NOSYS = 265;  // { int nosys(void); }
	SYS_SHM_OPEN               = 266 // { int shm_open(const char *name, int oflag, int mode); }
	SYS_SHM_UNLINK             = 267 // { int shm_unlink(const char *name); }
	SYS_SEM_OPEN               = 268 // { user_addr_t sem_open(const char *name, int oflag, int mode, int value); }
	SYS_SEM_CLOSE              = 269 // { int sem_close(sem_t *sem); }
	SYS_SEM_UNLINK             = 270 // { int sem_unlink(const char *name); }
	SYS_SEM_WAIT               = 271 // { int sem_wait(sem_t *sem); }
	SYS_SEM_TRYWAIT            = 272 // { int sem_trywait(sem_t *sem); }
	SYS_SEM_POST               = 273 // { int sem_post(sem_t *sem); }
	SYS_SEM_GETVALUE           = 274 // { int sem_getvalue(sem_t *sem, int *sval); }
	SYS_SEM_INIT               = 275 // { int sem_init(sem_t *sem, int phsared, u_int value); }
	SYS_SEM_DESTROY            = 276 // { int sem_destroy(sem_t *sem); }
	SYS_OPEN_EXTENDED          = 277 // { int open_extended(user_addr_t path, int flags, uid_t uid, gid_t gid, int mode, user_addr_t xsecurity) NO_SYSCALL_STUB; }
	SYS_UMASK_EXTENDED         = 278 // { int umask_extended(int newmask, user_addr_t xsecurity) NO_SYSCALL_STUB; }
	SYS_STAT_EXTENDED          = 279 // { int stat_extended(user_addr_t path, user_addr_t ub, user_addr_t xsecurity, user_addr_t xsecurity_size) NO_SYSCALL_STUB; }
	SYS_LSTAT_EXTENDED         = 280 // { int lstat_extended(user_addr_t path, user_addr_t ub, user_addr_t xsecurity, user_addr_t xsecurity_size) NO_SYSCALL_STUB; }
	SYS_FSTAT_EXTENDED         = 281 // { int fstat_extended(int fd, user_addr_t ub, user_addr_t xsecurity, user_addr_t xsecurity_size) NO_SYSCALL_STUB; }
	SYS_CHMOD_EXTENDED         = 282 // { int chmod_extended(user_addr_t path, uid_t uid, gid_t gid, int mode, user_addr_t xsecurity) NO_SYSCALL_STUB; }
	SYS_FCHMOD_EXTENDED        = 283 // { int fchmod_extended(int fd, uid_t uid, gid_t gid, int mode, user_addr_t xsecurity) NO_SYSCALL_STUB; }
	SYS_ACCESS_EXTENDED        = 284 // { int access_extended(user_addr_t entries, size_t size, user_addr_t results, uid_t uid) NO_SYSCALL_STUB; }
	SYS_SETTID                 = 285 // { int settid(uid_t uid, gid_t gid) NO_SYSCALL_STUB; }
	SYS_GETTID                 = 286 // { int gettid(uid_t *uidp, gid_t *gidp) NO_SYSCALL_STUB; }
	SYS_SETSGROUPS             = 287 // { int setsgroups(int setlen, user_addr_t guidset) NO_SYSCALL_STUB; }
	SYS_GETSGROUPS             = 288 // { int getsgroups(user_addr_t setlen, user_addr_t guidset) NO_SYSCALL_STUB; }
	SYS_SETWGROUPS             = 289 // { int setwgroups(int setlen, user_addr_t guidset) NO_SYSCALL_STUB; }
	SYS_GETWGROUPS             = 290 // { int getwgroups(user_addr_t setlen, user_addr_t guidset) NO_SYSCALL_STUB; }
	SYS_MKFIFO_EXTENDED        = 291 // { int mkfifo_extended(user_addr_t path, uid_t uid, gid_t gid, int mode, user_addr_t xsecurity) NO_SYSCALL_STUB; }
	SYS_MKDIR_EXTENDED         = 292 // { int mkdir_extended(user_addr_t path, uid_t uid, gid_t gid, int mode, user_addr_t xsecurity) NO_SYSCALL_STUB; }
	SYS_IDENTITYSVC            = 293 // { int identitysvc(int opcode, user_addr_t message) NO_SYSCALL_STUB; }
	SYS_SHARED_REGION_CHECK_NP = 294 // { int shared_region_check_np(uint64_t *start_address) NO_SYSCALL_STUB; }
	SYS_SHARED_REGION_MAP_NP   = 295 // { int shared_region_map_np(int fd, uint32_t count, const struct shared_file_mapping_np *mappings) NO_SYSCALL_STUB; }
	// SYS_NOSYS = 296;  // { int nosys(void); } { old load_shared_file }
	// SYS_NOSYS = 297;  // { int nosys(void); } { old reset_shared_file }
	// SYS_NOSYS = 298;  // { int nosys(void); } { old new_system_shared_regions }
	// SYS_ENOSYS = 299;  // { int enosys(void); } { old shared_region_map_file_np }
	// SYS_ENOSYS = 300;  // { int enosys(void); } { old shared_region_make_private_np }
	SYS___PTHREAD_MUTEX_DESTROY  = 301 // { int __pthread_mutex_destroy(int mutexid); }
	SYS___PTHREAD_MUTEX_INIT     = 302 // { int __pthread_mutex_init(user_addr_t  mutex, user_addr_t attr); }
	SYS___PTHREAD_MUTEX_LOCK     = 303 // { int __pthread_mutex_lock(int mutexid); }
	SYS___PTHREAD_MUTEX_TRYLOCK  = 304 // { int __pthread_mutex_trylock(int mutexid); }
	SYS___PTHREAD_MUTEX_UNLOCK   = 305 // { int __pthread_mutex_unlock(int mutexid); }
	SYS___PTHREAD_COND_INIT      = 306 // { int __pthread_cond_init(user_addr_t cond, user_addr_t attr); }
	SYS___PTHREAD_COND_DESTROY   = 307 // { int __pthread_cond_destroy(int condid); }
	SYS___PTHREAD_COND_BROADCAST = 308 // { int __pthread_cond_broadcast(int condid); }
	SYS___PTHREAD_COND_SIGNAL    = 309 // { int __pthread_cond_signal(int condid); }
	SYS_GETSID                   = 310 // { int getsid(pid_t pid); }
	SYS_SETTID_WITH_PID          = 311 // { int settid_with_pid(pid_t pid, int assume) NO_SYSCALL_STUB; }
	SYS___PTHREAD_COND_TIMEDWAIT = 312 // { int __pthread_cond_timedwait(int condid, int mutexid, user_addr_t abstime); }
	SYS_AIO_FSYNC                = 313 // { int aio_fsync(int op, user_addr_t aiocbp); }
	SYS_AIO_RETURN               = 314 // { user_ssize_t aio_return(user_addr_t aiocbp); }
	SYS_AIO_SUSPEND              = 315 // { int aio_suspend(user_addr_t aiocblist, int nent, user_addr_t timeoutp); }
	SYS_AIO_CANCEL               = 316 // { int aio_cancel(int fd, user_addr_t aiocbp); }
	SYS_AIO_ERROR                = 317 // { int aio_error(user_addr_t aiocbp); }
	SYS_AIO_READ                 = 318 // { int aio_read(user_addr_t aiocbp); }
	SYS_AIO_WRITE                = 319 // { int aio_write(user_addr_t aiocbp); }
	SYS_LIO_LISTIO               = 320 // { int lio_listio(int mode, user_addr_t aiocblist, int nent, user_addr_t sigp); }
	SYS___PTHREAD_COND_WAIT      = 321 // { int __pthread_cond_wait(int condid, int mutexid); }
	SYS_IOPOLICYSYS              = 322 // { int iopolicysys(int cmd, void *arg) NO_SYSCALL_STUB; }
	// SYS_NOSYS = 323;  // { int nosys(void); }
	SYS_MLOCKALL   = 324 // { int mlockall(int how); }
	SYS_MUNLOCKALL = 325 // { int munlockall(int how); }
	// SYS_NOSYS = 326;  // { int nosys(void); }
	SYS_ISSETUGID              = 327 // { int issetugid(void); }
	SYS___PTHREAD_KILL         = 328 // { int __pthread_kill(int thread_port, int sig); }
	SYS___PTHREAD_SIGMASK      = 329 // { int __pthread_sigmask(int how, user_addr_t set, user_addr_t oset); }
	SYS___SIGWAIT              = 330 // { int __sigwait(user_addr_t set, user_addr_t sig); }
	SYS___DISABLE_THREADSIGNAL = 331 // { int __disable_threadsignal(int value); }
	SYS___PTHREAD_MARKCANCEL   = 332 // { int __pthread_markcancel(int thread_port); }
	SYS___PTHREAD_CANCELED     = 333 // { int __pthread_canceled(int  action); }
	SYS___SEMWAIT_SIGNAL       = 334 // { int __semwait_signal(int cond_sem, int mutex_sem, int timeout, int relative, time_t tv_sec, int32_t tv_nsec); }
	// SYS_NOSYS = 335;  // { int nosys(void); }   { old utrace }
	SYS_PROC_INFO = 336 // { int proc_info(int32_t callnum,int32_t pid,uint32_t flavor, uint64_t arg,user_addr_t buffer,int32_t buffersize) NO_SYSCALL_STUB; }
	SYS_SENDFILE  = 337 // { int sendfile(int fd, int s, off_t offset, off_t *nbytes, struct sf_hdtr *hdtr, int flags); }
	// SYS_NOSYS = 337;  // { int nosys(void); }
	SYS_STAT64           = 338 // { int stat64(user_addr_t path, user_addr_t ub); }
	SYS_FSTAT64          = 339 // { int fstat64(int fd, user_addr_t ub); }
	SYS_LSTAT64          = 340 // { int lstat64(user_addr_t path, user_addr_t ub); }
	SYS_STAT64_EXTENDED  = 341 // { int stat64_extended(user_addr_t path, user_addr_t ub, user_addr_t xsecurity, user_addr_t xsecurity_size) NO_SYSCALL_STUB; }
	SYS_LSTAT64_EXTENDED = 342 // { int lstat64_extended(user_addr_t path, user_addr_t ub, user_addr_t xsecurity, user_addr_t xsecurity_size) NO_SYSCALL_STUB; }
	SYS_FSTAT64_EXTENDED = 343 // { int fstat64_extended(int fd, user_addr_t ub, user_addr_t xsecurity, user_addr_t xsecurity_size) NO_SYSCALL_STUB; }
	SYS_GETDIRENTRIES64  = 344 // { user_ssize_t getdirentries64(int fd, void *buf, user_size_t bufsize, off_t *position) NO_SYSCALL_STUB; }
	SYS_STATFS64         = 345 // { int statfs64(char *path, struct statfs64 *buf); }
	SYS_FSTATFS64        = 346 // { int fstatfs64(int fd, struct statfs64 *buf); }
	SYS_GETFSSTAT64      = 347 // { int getfsstat64(user_addr_t buf, int bufsize, int flags); }
	SYS___PTHREAD_CHDIR  = 348 // { int __pthread_chdir(user_addr_t path); }
	SYS___PTHREAD_FCHDIR = 349 // { int __pthread_fchdir(int fd); }
	SYS_AUDIT            = 350 // { int audit(void *record, int length); }
	SYS_AUDITON          = 351 // { int auditon(int cmd, void *data, int length); }
	// SYS_NOSYS = 352;  // { int nosys(void); }
	SYS_GETAUID       = 353 // { int getauid(au_id_t *auid); }
	SYS_SETAUID       = 354 // { int setauid(au_id_t *auid); }
	SYS_GETAUDIT      = 355 // { int getaudit(struct auditinfo *auditinfo); }
	SYS_SETAUDIT      = 356 // { int setaudit(struct auditinfo *auditinfo); }
	SYS_GETAUDIT_ADDR = 357 // { int getaudit_addr(struct auditinfo_addr *auditinfo_addr, int length); }
	SYS_SETAUDIT_ADDR = 358 // { int setaudit_addr(struct auditinfo_addr *auditinfo_addr, int length); }
	SYS_AUDITCTL      = 359 // { int auditctl(char *path); }
	// SYS_NOSYS = 350;  // { int nosys(void); }
	// SYS_NOSYS = 351;  // { int nosys(void); }
	// SYS_NOSYS = 352;  // { int nosys(void); }
	// SYS_NOSYS = 353;  // { int nosys(void); }
	// SYS_NOSYS = 354;  // { int nosys(void); }
	// SYS_NOSYS = 355;  // { int nosys(void); }
	// SYS_NOSYS = 356;  // { int nosys(void); }
	// SYS_NOSYS = 357;  // { int nosys(void); }
	// SYS_NOSYS = 358;  // { int nosys(void); }
	// SYS_NOSYS = 359;  // { int nosys(void); }
	SYS_BSDTHREAD_CREATE    = 360 // { user_addr_t bsdthread_create(user_addr_t func, user_addr_t func_arg, user_addr_t stack, user_addr_t pthread, uint32_t flags) NO_SYSCALL_STUB; }
	SYS_BSDTHREAD_TERMINATE = 361 // { int bsdthread_terminate(user_addr_t stackaddr, size_t freesize, uint32_t port, uint32_t sem) NO_SYSCALL_STUB; }
	SYS_KQUEUE              = 362 // { int kqueue(void); }
	SYS_KEVENT              = 363 // { int kevent(int fd, const struct kevent *changelist, int nchanges, struct kevent *eventlist, int nevents, const struct timespec *timeout); }
	SYS_LCHOWN              = 364 // { int lchown(user_addr_t path, uid_t owner, gid_t group); }
	SYS_STACK_SNAPSHOT      = 365 // { int stack_snapshot(pid_t pid, user_addr_t tracebuf, uint32_t tracebuf_size, uint32_t options) NO_SYSCALL_STUB; }
	SYS_BSDTHREAD_REGISTER  = 366 // { int bsdthread_register(user_addr_t threadstart, user_addr_t wqthread, int pthsize) NO_SYSCALL_STUB; }
	SYS_WORKQ_OPEN          = 367 // { int workq_open(void) NO_SYSCALL_STUB; }
	SYS_WORKQ_OPS           = 368 // { int workq_ops(int options, user_addr_t item, int prio) NO_SYSCALL_STUB; }
	// SYS_NOSYS = 369;  // { int nosys(void); }
	// SYS_NOSYS = 370;  // { int nosys(void); }
	// SYS_NOSYS = 371;  // { int nosys(void); }
	// SYS_NOSYS = 372;  // { int nosys(void); }
	// SYS_NOSYS = 373;  // { int nosys(void); }
	// SYS_NOSYS = 374;  // { int nosys(void); }
	// SYS_NOSYS = 375;  // { int nosys(void); }
	// SYS_NOSYS = 376;  // { int nosys(void); }
	// SYS_NOSYS = 377;  // { int nosys(void); }
	// SYS_NOSYS = 378;  // { int nosys(void); }
	// SYS_NOSYS = 379;  // { int nosys(void); }
	SYS___MAC_EXECVE      = 380 // { int __mac_execve(char *fname, char **argp, char **envp, struct mac *mac_p); }
	SYS___MAC_SYSCALL     = 381 // { int __mac_syscall(char *policy, int call, user_addr_t arg); }
	SYS___MAC_GET_FILE    = 382 // { int __mac_get_file(char *path_p, struct mac *mac_p); }
	SYS___MAC_SET_FILE    = 383 // { int __mac_set_file(char *path_p, struct mac *mac_p); }
	SYS___MAC_GET_LINK    = 384 // { int __mac_get_link(char *path_p, struct mac *mac_p); }
	SYS___MAC_SET_LINK    = 385 // { int __mac_set_link(char *path_p, struct mac *mac_p); }
	SYS___MAC_GET_PROC    = 386 // { int __mac_get_proc(struct mac *mac_p); }
	SYS___MAC_SET_PROC    = 387 // { int __mac_set_proc(struct mac *mac_p); }
	SYS___MAC_GET_FD      = 388 // { int __mac_get_fd(int fd, struct mac *mac_p); }
	SYS___MAC_SET_FD      = 389 // { int __mac_set_fd(int fd, struct mac *mac_p); }
	SYS___MAC_GET_PID     = 390 // { int __mac_get_pid(pid_t pid, struct mac *mac_p); }
	SYS___MAC_GET_LCID    = 391 // { int __mac_get_lcid(pid_t lcid, struct mac *mac_p); }
	SYS___MAC_GET_LCTX    = 392 // { int __mac_get_lctx(struct mac *mac_p); }
	SYS___MAC_SET_LCTX    = 393 // { int __mac_set_lctx(struct mac *mac_p); }
	SYS_SETLCID           = 394 // { int setlcid(pid_t pid, pid_t lcid) NO_SYSCALL_STUB; }
	SYS_GETLCID           = 395 // { int getlcid(pid_t pid) NO_SYSCALL_STUB; }
	SYS_READ_NOCANCEL     = 396 // { user_ssize_t read_nocancel(int fd, user_addr_t cbuf, user_size_t nbyte) NO_SYSCALL_STUB; }
	SYS_WRITE_NOCANCEL    = 397 // { user_ssize_t write_nocancel(int fd, user_addr_t cbuf, user_size_t nbyte) NO_SYSCALL_STUB; }
	SYS_OPEN_NOCANCEL     = 398 // { int open_nocancel(user_addr_t path, int flags, int mode) NO_SYSCALL_STUB; }
	SYS_CLOSE_NOCANCEL    = 399 // { int close_nocancel(int fd) NO_SYSCALL_STUB; }
	SYS_WAIT4_NOCANCEL    = 400 // { int wait4_nocancel(int pid, user_addr_t status, int options, user_addr_t rusage) NO_SYSCALL_STUB; }
	SYS_RECVMSG_NOCANCEL  = 401 // { int recvmsg_nocancel(int s, struct msghdr *msg, int flags) NO_SYSCALL_STUB; }
	SYS_SENDMSG_NOCANCEL  = 402 // { int sendmsg_nocancel(int s, caddr_t msg, int flags) NO_SYSCALL_STUB; }
	SYS_RECVFROM_NOCANCEL = 403 // { int recvfrom_nocancel(int s, void *buf, size_t len, int flags, struct sockaddr *from, int *fromlenaddr) NO_SYSCALL_STUB; }
	SYS_ACCEPT_NOCANCEL   = 404 // { int accept_nocancel(int s, caddr_t name, socklen_t	*anamelen) NO_SYSCALL_STUB; }
	// SYS_NOSYS = 401;  // { int nosys(void); }
	// SYS_NOSYS = 402;  // { int nosys(void); }
	// SYS_NOSYS = 403;  // { int nosys(void); }
	// SYS_NOSYS = 404;  // { int nosys(void); }
	SYS_MSYNC_NOCANCEL   = 405 // { int msync_nocancel(caddr_t addr, size_t len, int flags) NO_SYSCALL_STUB; }
	SYS_FCNTL_NOCANCEL   = 406 // { int fcntl_nocancel(int fd, int cmd, long arg) NO_SYSCALL_STUB; }
	SYS_SELECT_NOCANCEL  = 407 // { int select_nocancel(int nd, u_int32_t *in, u_int32_t *ou, u_int32_t *ex, struct timeval *tv) NO_SYSCALL_STUB; }
	SYS_FSYNC_NOCANCEL   = 408 // { int fsync_nocancel(int fd) NO_SYSCALL_STUB; }
	SYS_CONNECT_NOCANCEL = 409 // { int connect_nocancel(int s, caddr_t name, socklen_t namelen) NO_SYSCALL_STUB; }
	// SYS_NOSYS = 409;  // { int nosys(void); }
	SYS_SIGSUSPEND_NOCANCEL = 410 // { int sigsuspend_nocancel(sigset_t mask) NO_SYSCALL_STUB; }
	SYS_READV_NOCANCEL      = 411 // { user_ssize_t readv_nocancel(int fd, struct iovec *iovp, u_int iovcnt) NO_SYSCALL_STUB; }
	SYS_WRITEV_NOCANCEL     = 412 // { user_ssize_t writev_nocancel(int fd, struct iovec *iovp, u_int iovcnt) NO_SYSCALL_STUB; }
	SYS_SENDTO_NOCANCEL     = 413 // { int sendto_nocancel(int s, caddr_t buf, size_t len, int flags, caddr_t to, socklen_t tolen) NO_SYSCALL_STUB; }
	// SYS_NOSYS = 413;  // { int nosys(void); }
	SYS_PREAD_NOCANCEL  = 414 // { user_ssize_t pread_nocancel(int fd, user_addr_t buf, user_size_t nbyte, off_t offset) NO_SYSCALL_STUB; }
	SYS_PWRITE_NOCANCEL = 415 // { user_ssize_t pwrite_nocancel(int fd, user_addr_t buf, user_size_t nbyte, off_t offset) NO_SYSCALL_STUB; }
	SYS_WAITID_NOCANCEL = 416 // { int waitid_nocancel(idtype_t idtype, id_t id, siginfo_t *infop, int options) NO_SYSCALL_STUB; }
	SYS_POLL_NOCANCEL   = 417 // { int poll_nocancel(struct pollfd *fds, u_int nfds, int timeout) NO_SYSCALL_STUB; }
	SYS_MSGSND_NOCANCEL = 418 // { int msgsnd_nocancel(int msqid, void *msgp, size_t msgsz, int msgflg) NO_SYSCALL_STUB; }
	SYS_MSGRCV_NOCANCEL = 419 // { user_ssize_t msgrcv_nocancel(int msqid, void *msgp, size_t msgsz, long msgtyp, int msgflg) NO_SYSCALL_STUB; }
	// SYS_NOSYS = 418;  // { int nosys(void); }
	// SYS_NOSYS = 419;  // { int nosys(void); }
	SYS_SEM_WAIT_NOCANCEL         = 420 // { int sem_wait_nocancel(sem_t *sem) NO_SYSCALL_STUB; }
	SYS_AIO_SUSPEND_NOCANCEL      = 421 // { int aio_suspend_nocancel(user_addr_t aiocblist, int nent, user_addr_t timeoutp) NO_SYSCALL_STUB; }
	SYS___SIGWAIT_NOCANCEL        = 422 // { int __sigwait_nocancel(user_addr_t set, user_addr_t sig) NO_SYSCALL_STUB; }
	SYS___SEMWAIT_SIGNAL_NOCANCEL = 423 // { int __semwait_signal_nocancel(int cond_sem, int mutex_sem, int timeout, int relative, time_t tv_sec, int32_t tv_nsec) NO_SYSCALL_STUB; }
	SYS___MAC_MOUNT               = 424 // { int __mac_mount(char *type, char *path, int flags, caddr_t data, struct mac *mac_p); }
	SYS___MAC_GET_MOUNT           = 425 // { int __mac_get_mount(char *path, struct mac *mac_p); }
	SYS___MAC_GETFSSTAT           = 426 // { int __mac_getfsstat(user_addr_t buf, int bufsize, user_addr_t mac, int macsize, int flags); }
)
