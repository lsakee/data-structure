# HEAP

## Heap이란?

  In computer science, a heap is a tree-based data structure that satisfies the heap property
    - [wiki pedia](https://en.wikipedia.org/wiki/Heap_(data_structure))

heap이란 우선순위 큐(Priority Queue)를 구현하기 위한 트리 기반 데이터 구조로, 완전 이진 트리의 형태를 가지며 힙 속성(heap property) 을 만족합니다.
삽입과 삭제의 시간 복잡도가 O(logN)입니다.

## Heap Property 종류
- max heap property: 부모 노드의 키 값은 자식 노드의 키 값보다 크거나 같다
- min heap property: 부모 노드의 키 값은 자식 노드의 키 값보다 작거나 같다

## 장단점

### 장점
- 빠른 삽입 및 삭제
  - 시간 복잡도: 삽입과 삭제 연산은 O(log N) 시간에 처리됩니다.
  - 트리 구조에서 부모-자식 관계를 활용하여 효율적으로 삽입 및 삭제가 가능합니다.

- 최소값 또는 최대값을 빠르게 찾을 수 있음
  - **최소 힙(Min-Heap)**에서는 루트 노드에 가장 작은 값이 위치하고, **최대 힙(Max-Heap)**에서는 루트 노드에 가장 큰 값이 위치하여, O(1) 시간에 최소값 또는 최대값을 빠르게 찾을 수 있습니다.

- 완전 이진 트리로 구현 가능
  - 힙은 완전 이진 트리(Complete Binary Tree) 구조를 가지며, 배열로 구현할 수 있습니다.
  - 배열을 사용하면 포인터가 필요 없고, 트리의 부모-자식 노드를 간단한 수학적 계산으로 찾을 수 있어 메모리 효율적입니다.

- 효율적인 우선순위 큐 구현
  - **우선순위 큐(Priority Queue)**를 구현할 때 유용하며, 다익스트라 알고리즘, 힙 정렬(Heap Sort), 허프만 코딩(Huffman coding) 등 다양한 알고리즘에 활용됩니다.
  - 트랜잭션을 처리하거나 우선순위가 높은 항목을 빠르게 꺼내는 데 유리합니다.

- 배열 기반으로 간단한 구현
  - 힙은 배열을 사용하여 구현할 수 있어 메모리 오버헤드가 적고, 구현이 간단합니다.
  - 포인터 기반 트리 구조에 비해 메모리 사용이 효율적입니다.

### 단점
- 검색이 느림
  - 힙은 특정 값을 찾는 데 O(N) 시간이 걸립니다.
  - 정렬된 자료구조가 아니기 때문에 힙 내에서 임의의 값을 찾는 데 시간이 많이 걸립니다.

- 정렬이 불완전함
  - 힙은 정렬된 자료구조가 아니므로 전체 데이터를 정렬하는 데 O(N log N)의 시간이 걸립니다.
  - 힙은 최소값 또는 최대값을 빠르게 찾을 수 있지만, 전체 데이터를 정렬하려면 추가적인 시간이 필요합니다.

- 메모리 공간이 비효율적일 수 있음
  - 배열 기반 힙은 완전 이진 트리로 구현되므로 빈 공간이 발생할 수 있습니다.
 - 트리가 깊어지면 일부 자식 노드 공간이 낭비될 수 있고, 이를 해결하려면 추가적인 메모리 공간이 필요합니다.

- 동적 크기 변경 시 추가 작업 필요
  - 힙은 동적 크기 조정이 자동으로 이루어지지 않으며, 삽입이나 삭제가 있을 때마다 재구성 작업이 필요합니다.
  - 이로 인해 성능 저하가 발생할 수 있으며, 특히 삽입 및 삭제가 빈번한 경우 비효율적일 수 있습니다.

- 균형 잡힌 이진 트리의 특성을 보장하지 않음
  - 힙은 완전 이진 트리이지만, 균형 이진 트리는 아닙니다.
  - 일부 연산에서는 불균형으로 인해 성능이 떨어질 수 있습니다. (예: 레드-블랙 트리, AVL 트리와 비교)

## 과정(insert and delete)
![max heap insert](./maxheap-insertion.png)
![max heap delete](./maxheap-delete.png)

```javascript
/// min heap
let minheap = [];

function insert(heap, num){
    heap.push(num);
    let ind = heap.length;
    while(ind>1){
        if(heap[Math.floor(ind/2)-1]>heap[ind-1]){
                const temp = heap[ind-1];
                heap[ind-1] = heap[Math.floor(ind/2)-1];
                heap[Math.floor(ind/2)-1] = temp;
                ind = Math.floor(ind/2);
        }
        else{
            break;
        }
    }
    return heap;
}

function del(heap){
    heap[0] = heap[heap.length-1];
    heap.pop();
    const len = heap.length;
    let ind = 1;
    while(ind*2<=len){
        if(heap[ind-1]>heap[ind*2-1] && (heap[2*ind]===undefined ||heap[ind*2-1] < heap[ind*2])){
            const temp = heap[ind*2-1];
            heap[ind*2-1] = heap[ind-1];
            heap[ind-1] = temp;
            ind = ind*2
        }
        else if(heap[ind-1]>heap[ind*2]){
            const temp = heap[ind*2];
            heap[ind*2] = heap[ind-1];
            heap[ind-1] = temp;
            ind = ind*2+1
        }
        else
            break;
    }
    return heap
}
```

```javascript
/// min heap
class Node{
    constructor(value){
        this.left = null;
        this.right = null;
        this.parent = null;
        this.value = value;
    }
}

class MinHeap{
    #root;
    #length;

    constructor(){
        this.#root = null;
        this.#length = 0;
    }

    get root(){
        return this.#root;
    }

    get length(){
        return this.#length;
    }

    get top(){
        return this.root.value;
    }

    static from(node){
        if(!(node instanceof Node)){
            throw Error(`${node} is not Node Type.`);
        }
        const minHeap = new MinHeap();
        minHeap.#root = node;
        minHeap.#length++;
        return minHeap;
    }

    /**
     * 
     * @param {Node} parentNode 
     * @param {Node} childNode 
     */
    #swapValue(parentNode, childNode){
        const value = childNode.value;
        childNode.value = parentNode.value;
        parentNode.value = value;
    }

    add(value){
        if(this.length === 0){
            const newNode = new Node(value);
            this.#root = newNode;
        }
        else{
            const newNode = new Node(value);
            const len = (this.length + 1).toString(2);

            let current = this.root;

            for(let i = 1; i < len.length - 1; i++){
                current = (len[i] === '0' ? current.left : current.right);
            }

            len[len.length - 1] === '0' ? current.left = newNode : current.right = newNode;
            newNode.parent = current;
            current = newNode;

            let parent = current.parent;

            while(current.parent !== null && current.value < parent.value){
                this.#swapValue(parent, current);
                current = parent;
                parent = current.parent;
            }
        }
        this.#length++;
    }

    remove(){
        if(this.root === null){
            throw Error('Heap is empty.');
        }
        const value = this.root.value;
        const len = (this.length).toString(2);

        let current = this.root;

        for(let i = 1; i < len.length; i++){
            current = (len[i] === '0' ? current.left : current.right);
        }

        this.#root.value = current.value;

        if(current !== this.root){
            if(len[len.length - 1] === '0'){
                current.parent.left = null;
            }
            else{
                current.parent.right = null;
            }
        }
        this.#length--;

        current = this.root;

        while(current.left || current.right){
            if(current.left && current.right){
                if(current.left.value > current.right.value){
                    if(current.value > current.right.value){
                        this.#swapValue(current, current.right);
                        current = current.right;
                    }
                    else{
                        break;
                    }
                }
                else{
                    if(current.value > current.left.value){
                        this.#swapValue(current, current.left);
                        current = current.left;
                    }
                    else{
                        break;
                    }
                }
            }
            else{
                if(current.value > current.left.value){
                    this.#swapValue(current, current.left);
                    current = current.left;
                }
                else{
                    break;
                }
            }
        }
        return value;
    }
}
```

```go
package main

import "fmt"

type Heap[T comparable] struct {
	elements []T
	Compare  func(x, y T) bool
}

func (h *Heap[T]) Insert(value T) {
	if h.Length() == 0 {
		h.elements = make([]T, 0)
	}

	h.elements = append(h.elements, value)

	index := h.Length() - 1
	for !h.shiftUp(index) {
		switch index % 2 {
		case 0:
			index = (index - 2) / 2
		case 1:
			index = (index - 1) / 2
		}
	}
}

func (h *Heap[T]) Delete() (top T) {
	if h.Length() == 0 {
		return
	}

	root := h.elements[0]

	h.swap(0, h.Length()-1)
	h.elements = h.elements[:h.Length()-1]

	index := 0
	for next, done := h.shiftDown(index); !done; next, done = h.shiftDown(next) {
	}

	return root
}

func (h Heap[T]) Length() int {
	return len(h.elements)
}

func (h Heap[T]) Top() (top T) {
	top = h.elements[0]
	return
}

func (h Heap[T]) Print() {
	fmt.Printf("%v\n", h.elements)
}

func (h *Heap[T]) shiftUp(index int) (done bool) {
	if index == 0 {
		return true
	}

	current := h.elements[index]
	var parentIndex int

	if index%2 == 0 {
		parentIndex = (index - 2) / 2
	} else {
		parentIndex = (index - 1) / 2
	}

	parent := h.elements[parentIndex]

	if h.Compare(current, parent) {
		h.swap(index, parentIndex)
		return
	}

	return true
}

func (h *Heap[T]) shiftDown(index int) (nextIndex int, done bool) {
	leftChildIndex := 2*index + 1
	rightChildIndex := 2*index + 2

	if leftChildIndex > h.Length()-1 {
		return -1, true
	} else if rightChildIndex > h.Length()-1 {
		if h.Compare(h.elements[leftChildIndex], h.elements[index]) {
			h.swap(leftChildIndex, index)
			return leftChildIndex, false
		}

		return -1, true
	}

	if h.Compare(h.elements[leftChildIndex], h.elements[rightChildIndex]) {
		if h.Compare(h.elements[leftChildIndex], h.elements[index]) {
			h.swap(leftChildIndex, index)
			return leftChildIndex, false
		}

		return -1, true
	} else {
		if h.Compare(h.elements[rightChildIndex], h.elements[index]) {
			h.swap(rightChildIndex, index)
			return rightChildIndex, false
		}
		return -1, true
	}
}

func (h *Heap[T]) swap(x, y int) {
	temp := h.elements[x]
	h.elements[x] = h.elements[y]
	h.elements[y] = temp
}

func main() {
	maxHeap := &Heap[int]{
		elements: make([]int, 0),
		Compare: func(x, y int) bool {
			return x > y
		},
	}

	maxHeap.Insert(7)
	maxHeap.Print() // [7]
	maxHeap.Insert(8)
	maxHeap.Print() // [8 7]
	maxHeap.Insert(1)
	maxHeap.Print() // [8 7 1]
	maxHeap.Insert(2)
	maxHeap.Print() // [8 7 1 2]
	maxHeap.Insert(3)
	maxHeap.Print() // [8 7 1 2 3]

	fmt.Println(maxHeap.Delete())
	maxHeap.Print()
	fmt.Println(maxHeap.Delete())
	maxHeap.Print()
}
```