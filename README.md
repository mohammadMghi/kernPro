# kernPro
KernPro is a tool that assists you in interacting with the kernel and logging components of the operating system.


## What is process in linux:
Every running app in linux need to run a process to work .
A process consists of the program code, data, stack, and other resources required for its execution. It is managed and controlled by the Linux kernel, which provides the necessary services and resources to facilitate the execution of processes.

## types of process

User Processes: These processes are initiated and controlled by users. They can include applications, command-line programs, and other software running on the system.

System Processes: These processes are started and managed by the operating system itself. They perform essential system tasks, such as memory management, process scheduling, and handling hardware devices.

Parent and Child Processes: In Linux, a process can spawn child processes. The process that initiates the creation of a child process is called the parent process. Parent processes can have multiple child processes, forming a hierarchical relationship.

Daemon Processes: Daemon processes are background processes that run continuously, usually providing specific services or functionality. They often start during system boot and operate independently of user interaction. Examples include web servers (e.g., Apache), database servers (e.g., MySQL), and network services (e.g., SSH).

Zombie Processes: Zombie processes are terminated processes that have completed their execution but still have an entry in the process table. They exist to allow the parent process to retrieve information about the child's termination status. Zombie processes are usually cleaned up by the parent process using the wait system call.

Orphan Processes: Orphan processes are child processes whose parent process has terminated or ended unexpectedly. These processes are adopted by the init process (with PID 1), which becomes their new parent process. The init process ensures that orphan processes are properly handled and do not become zombie processes.


## How to use?
First of all clone the project then run project by bellow command :

```
$ go run main.go -type=cpu // For get cpu utilization
```

For print logs of process use below command:
```
$ go run main.go -type=process
```

For print logs of heavy cpu process use bellow comand:
```
$ go run main.go -type=cpu_heavy
```

For print logs of process state with pid :
```
$ go run main.go -type=process_state -pid=x
```

For print logs of child process use :
```
$ go run main.go -type=process_child -pid=x
```











ب
