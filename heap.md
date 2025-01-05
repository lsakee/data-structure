- [ ]  최댓값, 최솟값을 빠르게 찾아내기 위한 자료구조
    - 완전 이진트리를 기본으로 함(값 자체의 대소관계는 의미 X)
    - 완전 이진트리?
        - 마지막 레벨(리프 노드)을 제외한 모든 레벨이 완전히 채워져 있음
        - 마지막 레벨(리프 노드)은 왼쪽부터 차례로 채워져 있음
        - 노드 값의 크기나 순서는 상관없음

- [ ]  최소힙
    - 루트 노드가 최솟값
    - 부모 노드의 값이 항상 자식 노드의 값보다 같거나 작음
- [ ]  최대힙
    - 루트 노드가 최댓값
    - 부모 노드의 값이 항상 자식 노드의 값보다 같거나 큼
- [ ]  시간 복잡도
    - 삽입 O(log n)
    - 삭제 O(log n)
    - 조회 O(1)

- [ ]  HEAP 구현 코드 (swift)

```swift
import Foundation

struct HEAP<T> {
    private var elements: [T]
    
    private let priorityFunction: (T, T) -> Bool
    
    //MARK: 힙 초기화 메서드
    // - Parameters:
    /// - elements: 초기 배열 (기본값 빈 배열)
    /// - priorityFunction: 정렬 기준 클로저 (최소힙: <, 최대힙: >)
    init (elements: [T] = [], priorityFunction: @escaping (T, T) -> Bool) {
        self.elements = []
        self.priorityFunction = priorityFunction
        buildHeap()
    }
    
    // MARK: - 힙이 비어있는지
    var isEmpty: Bool {
        return elements.isEmpty
    }
    
    // MARK: - 힙의 원소 갯수
    var count: Int {
        return elements.count
    }
    
    // MARK: 힙의 최상단 요소 확인(최솟값 or 최댓값) - 제거하지 않음
    func peek() -> T? {
        return elements.first
    }
    
    // MARK: 힙에 원소추가
    mutating func insert(_ element: T) {
        elements.append(element)
        siftUp(from: elements.count - 1)  // 마지막에 추가된 요소를 올바른 위치로 이동
    }
    
    // MARK: 힙의 최상단 원소 제거 후 반환
    mutating func remove() -> T? {
        guard !isEmpty else { return nil }
        
        elements.swapAt(0, count - 1) // 루트원소와 마지막 원소 교환
        let element = elements.removeLast() // 마지막 원소(원래 루트노드의 원소) 제거
        if !isEmpty {
            siftDown(from: 0)
        }
        
        return element
    }
    
    //MARK: - 초기 배열을 힙으로 구성하는 메서드
    mutating func buildHeap() {
        // 마지막 비단말 노드부터 루트까지 siftDown 수행
        for i in stride(from: elements.count/2 - 1, through: 0, by: -1) {
            siftDown(from: i)
        }
    }
    
    // MARK: - 특정 노드를 위로 이동시키는 메서드 (삽입 시 사용)
    private mutating func siftUp(from index: Int) {
        var child = index
        var parent = parentIndex(of: child)
        
        // 부모 노드와 비교하며 우선순위가 높은 경우 위치 교환
        while child > 0 && priorityFunction(elements[child], elements[parent]) {
            elements.swapAt(child, parent)
            child = parent
            parent = parentIndex(of: child)
        }
    }
    
    // MARK: - 특정 노드를 아래로 이동시킨다. (삭제, 힙 구현시 사용)
    private mutating func siftDown(from index: Int) {
        var parent = index
        
        while true {
            let leftChild = leftChildIndex(of: parent)
            let rightChild = rightChildIndex(of: parent)
            var candidate = parent
            
            // 왼쪽 자식과 비교, 왼쪽 자식이 더 작으면 교환 후보로 선택
            if leftChild < count && priorityFunction(elements[leftChild], elements[candidate]) {
                candidate = leftChild
            }
            
            // 오른쪽 자식과 비교, 오른쪽 자식이 더 작으면 교환 후보로 선택
            if rightChild < count && priorityFunction(elements[rightChild], elements[candidate]) {
                candidate = rightChild
            }
            
            // 더 이상 교환이 필요없으면 종료
            if candidate == parent {
                return
            }
            
            // 교환 후보와 교환
            elements.swapAt(parent, candidate)
            parent = candidate
        }
    }
    
    // MARK: - 부모 노드의 인덱스를 반환하는 메서드
    private func parentIndex(of index: Int) -> Int {
        return (index - 1) / 2
    }
    
    // MARK: - 왼쪽 자식 노드의 인덱스를 반환하는 메서드
    private func leftChildIndex(of index: Int) -> Int {
        return index * 2 + 1
    }
    
    // MARK: - 오른쪽 자신 노드의 인덱스를 반환하는 메서드
    private func rightChildIndex(of index: Int) -> Int {
        return index * 2 + 2
    }
}

var minHeap = HEAP<Int> { $0 < $1 } // 최소힙

minHeap.insert(5)
minHeap.insert(3)
minHeap.insert(7)
minHeap.insert(1)

while let min = minHeap.remove() {
    print(min)  // 1, 3, 5, 7 순서로 출력
}

print("=======================")

var maxHeap = HEAP<Int> { $0 > $1 } // 최대힙

maxHeap.insert(5)
maxHeap.insert(3)
maxHeap.insert(7)
maxHeap.insert(1)

while let max = maxHeap.remove() {
    print(max)
}

```
