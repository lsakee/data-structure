##Hash
- [ ]  해시
    - 임의의 크기를 가진 데이터를 고정된 크기의 값 으로 변환하는 함수
    - 해시테이블
        - 키-값 쌍을 저장하는 자료구조
        - 해시 함수를 사용하여 키를 특정 인덱스로 매핑
- [ ]  해시 테이블 시간복잡도
    - 삽입(Insertion)
        - 최선: O(1) - 해시충돌이 없는 경우
        - 평균: O(1)
        - 최악: O(n) - 모든 키가 같은 해시값을 가져 하나의 버킷에 모든 데이터가 체이닝된 경우
    - 삭제(Deletion)
        - 최선: O(1) - 해시충돌이 없는 경우
        - 평균: O(1)
        - 최악: O(n) - 모든 키가 같은 해시값을 가져 하나의 버킷에 모든 데이터가 체이닝된 경우
    - 검색(Search)
        - 최선: O(1) - 해시충돌이 없는 경우
        - 평균: O(1)
        - 최악: O(n) - 모든 키가 같은 해시값을 가져 하나의 버킷에 모든 데이터가 체이닝된 경우
- [ ]  해시 충돌
    - 어떤 상황에 발생할 수 있는지?
        - 서로 다른 key를 해시 함수에 넣었는데, 우연히 같은 해시 값이 나올 때. 그래서 해시 테이블의 같은 index에 value가 저장될 때
        
        ```swift
        // String을 키로 사용하는 경우
        var hashTable = [String: String]()
        
        // 두 개의 다른 키
        "ABC" -> hash() -> 12345
        "XYZ" -> hash() -> 12345  // 우연히 같은 해시값이 나오는 경우 -> 같은 index에 저장
        [ 0: [("ABC", 3), ("XYZ", 4)], 1: nil...] // 이런식일듯?
        ```
        
    
    - 해결방법
        - 체이닝
            - 연결리스트 형태로 데이터를 저장
            - 검색 할 때 키를 비교하여 정확한 값을 찾아냄.
            
            ```swift
            struct Person: Hashable {
                let name: String
                let age: Int
                
                // Equatable
                static func == (lhs: Person, rhs: Person) -> Bool {
                    return lhs.name == rhs.name && lhs.age == rhs.age
                }
            }
            
            let person1 = Person(name: "Kim", age: 20)  // 해시값: 12345
            let person2 = Person(name: "Lee", age: 30)  // 해시값: 12345
            
            // 둘 다 동일한 인덱스에 저장
            dict[person1] = "학생"
            dict[person2] = "교사"
            
            인덱스 5: (person1, "학생") -> (person2, "교사")
            
            // person1에 해당하는 값을 찾으려고 할 때
            let searchPerson = Person(name: "Kim", age: 20)
            
            // 1. 해시값으로 인덱스 찾기
            인덱스 = hash(searchPerson) % 테이블크기  // 인덱스 5
            
            // 2. 인덱스 5의 연결 리스트 탐색
            person1 == searchPerson 비교
              // name: "Kim" == "Kim"
              // age: 20 == 20
              // true 반환, "학생" 찾음!
            ```
            
        
        - 개방 주소법
            - 충돌 발생시 다른 공간에 저장
                - linear probing - 순차 탐색
                - quadratic probing - 제곱수만큼 건너뛰며 탐색
                - double hashing - 다른 해시 함수로 새 위치 계산
- [ ]  hash table의 실제 구조
    - 버킷 - 해시 테이블의 실제 저장소 역할
        
        ```swift
        // 해시 테이블의 구조 (단순화)
        [
            0: nil,                                         // 비어있는 버킷
            1: [("apple", "사과")],                         // 하나의 데이터가 있는 버킷
            2: [("banana", "바나나"), ("grape", "포도")],    // 충돌로 인해 두 개의 데이터가 있는 버킷
            3: nil,
            4: [("orange", "오렌지")]
        ]
        ```
        

- [ ]  Swift에서는 Dictionary 타입이 hash table 자료구조 기반의 타입이다.

```swift
// MARK: Dictionary

// key는 Hashable이라는 프로토콜을 채택 
@frozen public struct Dictionary<Key, Value> where Key : Hashable

// MARK: Hashable
// Hashable이라는 프로토콜은 Equatable이라는 프로토콜을 채택
public protocol Hashable : Equatable {
    var hashValue: Int { get }
    func hash(into hasher: inout Hasher)
}

// MARK: Equatable
// 두 객체가 동일한 객체인가?를 비교할 수 있는 타입이 채택하는 프로토콜
public protocol Equatable {
    static func == (lhs: Self, rhs: Self) -> Bool
}

```
