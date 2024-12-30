## 특징

- 완전 이진 트리 구조
  - leaf node를 제외한 모든 레벨이 왼쪽부터 순차적으로 채워짐
  - 전체 노드 수 n = 2^0 + 2^1 + ... + 2^h = 2^(h+1) - 1
  - 높이 h = log2(n+1) - 1 (꽉찬 경우)
- 최대 힙(Max Heap): 부모 노드가 항상 자식 노드보다 크거나 같음
- 최소 힙(Min Heap): 부모 노드가 항상 자식 노드보다 작거나 같음
- 형제 노드간에는 대소 관계 없음

## 동작

### 삽입

1. 마지막 인덱스로 삽입
2. 부모 노드와 비교하며 위치 교환
3. 최대 교환 횟수 = 트리의 높이 = O(logN) (N: 노드 개수) (정확히는 밑이 2인 로그)

### 삭제

1. 루트 노드 삭제
2. 마지막 노드를 루트로 이동
3. 자식 노드와 비교하며 위치 교환
4. 최대 교환 횟수 = 트리의 높이 = O(logN) (N: 노드 개수) (정확히는 밑이 2인 로그)

### 조회 (최대값, 최소값)

1. 최대 힙: 루트 노드 - O(1)

## 구현

```ts
abstract class MinHeap<T> {
  abstract push(value: T): void;
  abstract pop(): T | undefined;
  abstract peek(): T | undefined;
  abstract size(): number;
  abstract isEmpty(): boolean;
}

class MinHeapImpl<T> implements MinHeap<T> {
  private heap: T[];

  constructor() {
    this.heap = [];
  }

  private getParentIndex(index: number): number {
    return Math.floor((index - 1) / 2);
  }

  private getLeftChildIndex(index: number): number {
    return 2 * index + 1;
  }

  private getRightChildIndex(index: number): number {
    return 2 * index + 2;
  }

  private swap(index1: number, index2: number): void {
    [this.heap[index1], this.heap[index2]] = [this.heap[index2], this.heap[index1]];
  }

  private bubbleUp(index: number): void {
    while (index > 0) {
      const parentIndex = this.getParentIndex(index);

      if (this.heap[parentIndex] <= this.heap[index]) break;
      this.swap(index, parentIndex);
      index = parentIndex;
    }
  }

  private bubbleDown(index: number): void {
    while (true) {
      let smallestIndex = index;
      const leftChild = this.getLeftChildIndex(index);
      const rightChild = this.getRightChildIndex(index);

      if (leftChild < this.heap.length && this.heap[leftChild] < this.heap[smallestIndex]) {
        smallestIndex = leftChild;
      }

      if (rightChild < this.heap.length && this.heap[rightChild] < this.heap[smallestIndex]) {
        smallestIndex = rightChild;
      }

      if (smallestIndex === index) break;

      this.swap(index, smallestIndex);
      index = smallestIndex;
    }
  }

  public push(value: T): void {
    this.heap.push(value);
    this.bubbleUp(this.heap.length - 1);
  }

  public pop(): T | undefined {
    if (this.heap.length === 0) return undefined;
    if (this.heap.length === 1) return this.heap.pop();

    const min = this.heap[0];
    this.heap[0] = this.heap.pop()!;
    this.bubbleDown(0);

    return min;
  }

  public peek(): T | undefined {
    return this.heap[0];
  }

  public size(): number {
    return this.heap.length;
  }

  public isEmpty(): boolean {
    return this.heap.length === 0;
  }
}
```

- 완전 이진 트리 라서 배열로 구현 가능 - index 계산
- RB Tree 는 색, 회전 연산같은 요구사항이 있어서 노드 기반(연결리스트)으로 구현해야함

## 실제 사용

- 최대, 최소 힙을 이용한 우선순위 큐
- 스케쥴링에 사용
  - 실제 OS의 스케줄링 구현에는 RB Tree 나 다른 자료구조 사용 (범위 검색이 등을 위해)
  - starvation 방지는 자료구조보다는 스케쥴링 정책으로 해결 ex. aging
  - 리눅스 RT 스케쥴러, JAVA ThreadPoolExecutor, Redis의 타이머 이벤트 등에 실제로 사용
- 다익스트라
  - 최소 비용으로 경로를 찾기 위해 사용
  - 네트워크 라우팅, GPS 등
