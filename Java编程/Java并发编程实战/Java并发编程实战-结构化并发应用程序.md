---
title: Java并发编程实战二(结构化并发应用程序)
date: 2017-07-31 11:59
---

# Java并发编程实战二-结构化并发应用程序
## 任务执行
* 大多数并发程序围绕"任务执行"构造，通常将任务分解为多个抽象且离散的工作单元，从而简化程序组织结构，提供自然的事物边界和优化错误恢复过程，以及提供一种自然的并行工作结构来提高并发性.
## 在线程中执行任务:
* 当围绕任务执行构建应用程序时,为(利用并发）提高程序的吞吐量和响应性，应首先确定清晰的任务边界和明确的任务执行策略;
    * 单个线程中串行的执行任务，效率低，需等待；
    * 显式的为任务创建线程，子任务的执行从主线程分离，提高响应性,可并行，但需要保证子任务处理代码必须线程安全。同时为每个任务分配一个线程存在缺陷(尤其线程数较多时):
        * 线程生命周期（创建／销毁)的开销非常高
        * 资源消耗,若线程大于可用处理器的数量，则控制的线程可能会占用大量内存资源，给垃圾回收器带来压力，且大量线程在竞争CPU资源时也存在性能开销
        * 稳定性:在可创建线程的数量上存在一个限制(可能与平台,jvm启动参数,Thread构造函数中请求栈的大小，及底层OS对线程的限制)，若破坏这个限制可能抛出OutOfMemoryError异常，且从这种错误恢复存在危险,应尽量避免这种错误；

## Executor框架:
* 线程池简化了线程的管理工作，并且java.util.concurrent提供了一种灵活的线程池实现作为Executor框架的一部分。在Java类库中任务执行的的主要抽象为Executor
```
public interface Executor{
    void execute(Runnable command);
}
```
* Executor接口为灵活而强大的异步任务执行框架提供了基础，可以支持多种不同类型的任务执行策略。提供一种标准的方式将任务提交过程与执行过程解耦开来，并用Runnable来表示任务.此外Executor还支持对生命周期的支持，以及统计信息收集，应用程序管理机制和性能监视等机制.
* Executor基于生产者-消费者模式，提交任务的操作相当于生产者(生成待完成的工作单元)，执行任务的线程相当于消费者(执行生成的工作单元)，通过Executor的不同实现(重写execute()方法)可以实现不同的执行策略；
* 执行策略为一种资源管理工具:定义了任务执行的相关信息:在什么线程中执行，执行顺序(FIFO,LIFO，优先级等)，可并发执行的数量,等待队列的线程数量，系统过载时应选择拒绝哪一个任务，如何通知应用程序有任务被拒绝，在执行任务前后需要哪些准备工作等；
> 最佳策略取决于可用计算资源以及对服务质量的需求，通过限制并发任务的数量，可以确保应用程序不会由于资源耗尽而失败,或者由于在稀缺资源上发生竞争而严重影响性能。通过将任务的提交与执行策略分离有助于在部署阶段选择与可用硬件资源最匹配的执行策略.
* 线程池: 指管理一组同构工作线程的资源池，线程池与工作队列密切相关，其中工作队列保存了所有等待执行的线程，工作者线程从工作队列中获取一个线程，执行任务，然后返回线程池并等待下一个任务。
* 可通过调用Executors的静态工厂方法来创建一个线程池:
    * newFixedThreadPool:创建一个固定长度的线程池，每当提交一个任务创建一个线程，直到达到线程池的最大数量，此时线程规模不再变化(如某个线程发生异常结束，则补充一个新的线程)
    * newCachedThreadPool：创建一个可缓存线程池，如果线程池的当前规模超过了处理需求时，则将回收空闲的线程.当需求增加时，可添加新的线程，线程池的规模不受限制
    * newSingleThreadExecutor：是一个单线程的Executor,创建单个工作者线程来执行任务,如此线程异常结束,则创建另一个线程来替代，可确保依照任务在队列中的顺序来串行执行(单线程的Executor还提供了大量的内部同步机制，确保任务执行的任何内存写入操作对于后续任务可见，)
    * newScheduledThreadPool：创建一个固定长度的线程池，且以延迟或定时的方式来执行任务，类似Timer
> 在线程池中执行任务比为每个任务分配一个线程更具优势:通过重用现有线程，可减少在处理多个请求时线程创建和销毁过程的开销，减少了创建线程的请求延时，同时避免的过多线程的高负载，此外应注意适当调整线程池的大小。
* Executor的生命周期:为了解决执行服务的生命周期问题,Executor扩展了ExecutorService接口，添加了一些用于生命周期管理的方法(此外包括用于任务提交的便利方法):
```
public interface ExecutorService extends Executor{
    void shutdown();  //执行平缓关闭(执行完现有线程，不再接收新的线程)
    List<Runnable> shutdownNow();　　//立即关闭
    boolean isShutdown();
    boolean isTerminated();
    boolean awaitTermination(long timeout,TimeUnit unit);  //等待终止
    .......
}
```
>ExecutorService包含运行/关闭/已终止三种状态，关闭后提交的任务由"拒绝执行处理器"处理，会抛弃任务，使得execute方法抛出一个未受检查的RejectedExecutionException
* 延迟／周期任务:相比Timer，ScheduledThreadPoolExcutor更加完善:
    * Timer在执行所有定时任务时只会创建一个线程(若一个任务执行时间过长会影响其他任务的精确定时)，Timer线程并不捕获异常，因此在TimerTask抛出未受检查异常时定时线程将终止，为完成的任务不会再执行，出现"线程泄露"；
    * 可以利用DelayQueue(实现了BlockingQueue,并为ScheduledThreadPoolExecutor提供调度功能)实现自己的调度服务:DelayQueue管理着一组Delay对象，每个Delay对象都有一个相应的延迟时间(只有某个元素逾期后才能从DelayQueue中执行take操作)，从DelayQueue中返回的对象将根据它们的延迟时间进行排序；

## 找出可利用的并发性:
* 使用Executor必须将任务表述为一个Runnable,有时任务边界并非显而易见，可能在单个任务中仍存在可并行的部分，在选择任务边界时，应权衡各种条件，发掘潜在的可并行策略
* Executor使用Runnable作为任务抽象，不能够返回一个值或抛出一个异常，然而在实际中可能存在延迟计算的情况,此时Callable是一种更好的任务抽象：它认为主入口点(即call)将返回一个值或抛出一个异常，同时在Executor中包含一些辅助方法可以将其他类型的任务封装为一个Callable(如Runnable,java.security.PrivilegedAction)
* Runnable和Callable描述的均为抽象的计算任务(通常有明确的生命周期)，Future表示一个任务的生命周期，并提供相应的方法来判断相应的任务是否完成或取消，以及获取任务的结果或取消任务，Future.get()若任务完成则立即返回或将EXception封装为ExecutionException抛出,若任务没有完成则阻塞等待直到任务完成，如果任务取消则抛出CancellationException.
* 可以通过多个方法获得Future:ExecutorService中的所有submit()方法会返回一个Future（因此将一个Runnabel或Callable向Executor提交时将得到一个相应的Future）,或者显式的为某个Runnable或Callable实例化一个FutureTask(其实现了Runnable,可以提交给Executor执行，或直接调用其run()方法)
> ExecutorService实现可以改写AbstractExecutorService中的newTaskFor()方法，根据已提交的Runnable或Callable来控制Futured的实例话过程(默认仅创建一个新的FutureTask)
> 将Runnable或Callable提交到Executor或设置Future结果时均包含安全发布，确保线程安全
* 在异构任务并行化中存在的局限:
 * 需要在相似的任务间找出细粒度的并行性，同时需要权衡任务的分配(各自合理的运行时间，较低的任务协调开销)
 * CompletionService:相比通过Future.get()轮询任务状态，CompletionService将Executor和BlockingQueue的功能融合，其可以负责执行Callable任务，使用类似队列的take/pull等方法获得已完成的结果(这些结果会在完成时封装为Future)
 * ExecutorCompletionService实现了CompletionService并将计算部分委托给Executor,实现:在构造函数中床架一个BockingQueue保存计算完成结果，计算完成时调用Futuretask的done方法，当提交任务时，首相将任务包装为一个QueueingFuture(为FutureTask的子类)，然后改写子类的done方法将结果放入BlockingQueue
 * Future可以设定时间限制，当任务在指定时间没有完成时，get()将抛出TimeoutException，此时应利用Future取消任务


## 取消与关闭:
125



