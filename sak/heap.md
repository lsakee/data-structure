# Heap 자료구조

## Heap의 개념

- Heap은 완전 이진 트리(Complete Binary Tree) 기반의 자료구조

## Heap의 활용

1. 우선순위 큐 구현
2. 힙 정렬
3. 다익스트라 알고리즘
4. 중앙값 찾기 (MedianFinder)
5. 작업 스케줄링


## 메모리 효율성

- 완전 이진 트리 특성으로 인해 메모리 낭비가 적음
- 배열 기반 구현으로 포인터 오버헤드 없음
- 동적 크기 조절 가능 (MutableList 사용)

## 주의사항

1. 힙 속성 유지
    - 삽입/삭제 시 항상 힙 속성을 만족하도록 유지
    - heapify 과정의 정확한 구현 필요
2. 인덱스 계산
    - 배열 기반 구현에서 인덱스 계산 시 주의
    - 범위 검사 필수

# 안드로이드 ThreadPoolExecutor

(코루틴 대체
tmi - 안쓰는 이유는 크게 Os 레벨 스레드 관리,선점형 , 고정크기 스택메모리 할당, 동시 생성 스레드수 제한 등)
개념: 스레드 풀 관리
작업 처리 방식: 코어 스레드가 먼저 작업 처리 - 작업 많아지면 추가 스레드 - 유휴 스레드는 유지시간 확인 후 종료
방식: 우선순위 높은 작업 먼저 처리, 큐 비었으면 스레드 자동 대기, 새로운 작업 추가시 우선순위 재정렬

안드로이드 내장 함수 로직에 구현 되어 있는 큐를 보겠습니다.

```
  public static <E extends Comparable> PriorityQueue<E> newPriorityQueue() {
    return new PriorityQueue<E>();
  }
	
	public static <E extends Comparable> PriorityQueue<E> newPriorityQueue(
      Iterable<? extends E> elements) {
    if (elements instanceof Collection) {
      return new PriorityQueue<E>((Collection<? extends E>) elements);
    }
    PriorityQueue<E> queue = new PriorityQueue<E>();
    Iterables.addAll(queue, elements);
    return queue;
  }
```

```
public static <E extends Comparable> PriorityBlockingQueue<E> newPriorityBlockingQueue() {
    return new PriorityBlockingQueue<E>();
}

public static <E extends Comparable> PriorityBlockingQueue<E> newPriorityBlockingQueue(
    Iterable<? extends E> elements) {
    if (elements instanceof Collection) {
        return new PriorityBlockingQueue<E>((Collection<? extends E>) elements);
    }
    PriorityBlockingQueue<E> queue = new PriorityBlockingQueue<E>();
    Iterables.addAll(queue, elements);
    return queue;
}
```

```
// 위로 올리며 정렬 (add/offer 시 호출)
private static <T> void siftUpComparable(int k, T x, Object[] es) {
    Comparable<? super T> key = (Comparable<? super T>) x;
    while (k > 0) {
        int parent = (k - 1) >>> 1;  // 부모 노드 찾기
        Object e = es[parent];
        if (key.compareTo((T) e) >= 0)  // 부모보다 크거나 같으면 중단
            break;
        es[k] = e;                   // 부모를 아래로 내림
        k = parent;
    }
    es[k] = key;
}
```

```
private static <T> void siftDownComparable(int k, T x, Object[] es, int n) {
    Comparable<? super T> key = (Comparable<? super T>)x;
    int half = n >>> 1;
    while (k < half) {
        int child = (k << 1) + 1;    // 왼쪽 자식
        Object c = es[child];
        int right = child + 1;       // 오른쪽 자식
        if (right < n && 
            ((Comparable<? super T>) c).compareTo((T) es[right]) > 0)
            c = es[child = right];   // 더 작은 자식 선택
        if (key.compareTo((T) c) <= 0) 
            break;
        es[k] = c;                   // 선택된 자식을 위로 올림
        k = child;
    }
    es[k] = key;
}
```

```
public boolean offer(E e) {
    if (e == null)
        throw new NullPointerException();
    final ReentrantLock lock = this.lock;
    lock.lock();
    try {
        // ... 크기 조정 로직 ...
        if (comparator == null)
            siftUpComparable(n, e, es);      // 자연 순서 사용
        else
            siftUpUsingComparator(n, e, es, cmp);  // 커스텀 비교자 사용
        size = n + 1;
        notEmpty.signal();
    } finally {
        lock.unlock();
    }
    return true;
}
```