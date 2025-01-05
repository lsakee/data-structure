# Hash 자료구조

## 해시(Hash)의 개념

- 가변 길이의 입력된 데이터를 고정된 길이의 데이터로 매핑하는 단방향 함수의 자료구조
- 키(key)를 해시 함수를 통해 해시값(hash value)으로 변환하고, 이 값을 인덱스로 사용하여 데이터를 저장/검색 가능

<img width="727" alt="스크린샷 2025-01-05 오후 5 47 21" src="https://github.com/user-attachments/assets/f487b607-ad82-4f3d-9111-489cd905cf17" />

## 해시의 주요 구성 요소

해시 함수: 키를 해시값으로 변환하는 함수

### 좋은 해시 함수의 조건

- 일관성: 같은 입력에 항상 같은 출력
- 균일성: 해시값이 고르게 분포
- 효율성: 계산이 빠름

## 해시 테이블

실제 데이터가 저장되는 배열
각 저장 공간을 버킷(bucket)이라고 함

### 해시 테이블 Deep Dive (밑에서 나올 개념 체이닝을 포함하고 있습니다.)

<img width="722" alt="스크린샷 2025-01-05 오후 5 52 50" src="https://github.com/user-attachments/assets/e5a930c7-bf9f-4372-915a-9172d1d70aed" />
예시) ESAK이란 가변의 데이터를 입력
계산 방법: 각 알파벳의 넘버  a,b,c,d,e....e=5

<img width="726" alt="스크린샷 2025-01-05 오후 5 57 40" src="https://github.com/user-attachments/assets/42a1fe15-ce11-449e-9c73-81422861975f" />

DESK,PEAK라는 가변의 데이터를 추가로 입력

DESK의 해시값은 3으로 계산이 되어 버킷에 저장이 됩니다

여기서 PEAK의 해시값도 3으로 계산이 되어 충돌이 발생하게 됩니다.

이러한 경우 체이닝을 통해 연결리스트로 데이터를 저장하여 해결합니다.

### 개방 주소법

<img width="732" alt="스크린샷 2025-01-05 오후 6 02 43" src="https://github.com/user-attachments/assets/1ae9cb2b-34f8-4b18-a9cf-bbb5da2af1c0" />
직관적으로 보았을 경우 여유공간이 있습니다.

만약 충돌이 발생한 다면 다른 빈 버킷으로 저장 합니다.

이러한 경우 포인터를 사용하지 않고 순차적으로 빈 버킷을 탐색하는 방식으로 진행됩니다.

근데 만약 해시값이 꽉차는 경우와 삽입에서 인접한 버킷을 차례로 검사하게 되므로 데이터의 군집화가 발생합니다.

검색시간과 삽입에 대한 성능도 저하됩니다.

또한 추가해시를 사용하지 않는경우 리사이징도 처리해주어야 합니다.

리사이징은 비용이 큰작업입니다.
<img width="742" alt="스크린샷 2025-01-05 오후 6 04 31" src="https://github.com/user-attachments/assets/049cf104-ed2e-4547-a781-b72657c1afa7" />

1. 일시적으로 기존 데이터를 유지하는 공간을 위한 메모리 사용
2. 데이터에 해시 작업 반복
3. 모든 데이터의 새로운 위치
4. 동기화 문제

##장점

검색 속도가 매우 빠름

평균적으로 O(1) 시간 복잡도
키를 알고 있다면 즉시 접근 가능

삽입/삭제가 효율적

다른 데이터의 이동이 필요 없음
평균적으로 O(1) 시간 복잡도

중복 처리가 용이

같은 키는 같은 해시값을 가짐
데이터 무결성 유지에 유용

## 단점

충돌(Collision) 발생

서로 다른 키가 같은 해시값을 가지는 현상
해결 방법:

체이닝: 같은 버킷에 연결 리스트로 저장
개방 주소법: 다른 빈 버킷을 찾아 저장

공간 효율성

충돌을 대비한 추가 공간 필요
일정 수준의 빈 공간 유지 필요

순서가 없음

데이터 간의 순서 관계를 알 수 없음
정렬된 데이터가 필요한 경우 부적합

## 메모리 효율성

로드 팩터(Load Factor)

해시 테이블의 공간 사용률
(저장된 데이터 수) / (전체 버킷 수)
일반적으로 0.7~0.8 유지 권장

## 메모리 사용

체이닝: 추가 포인터 공간 필요
개방 주소법: 더 많은 버킷 필요

## 예시

저는 코드를 직접 구현하기 보다는 안드로이드 내부 로직을 살펴보겠습니다.
내부에 데이터를 key,value 저장하는 SharedPreferences를 살펴 보겠습니다.

실제 데이터를 해시테이블에 저장하는 부분

코드를 부분적으로 보겠습니다.

 ```
// 삭제 처리
if (v == this || v == null) {
// 키가 있을 때만 삭제 진행
if (!mapToWriteToDisk.containsKey(k)) {
continue;
}
mapToWriteToDisk.remove(k);
}
// 추가/수정 처리
else {
// 키가 이미 존재하는 경우 값이 동일하면 작업 스킵합니다.
if (mapToWriteToDisk.containsKey(k)) {
Object existingValue = mapToWriteToDisk.get(k);

if (existingValue != null && existingValue.equals(v)) {
continue;
}
}
// 새 값을 해시테이블에 저장합니다.
mapToWriteToDisk.put(k, v);
}
changesMade = true;
// 리스너가 있으면 변경된 키 기록 -> 이부분을 통해 데이터가 변경된 경우 UI등 변경이 가능합니다.
if (hasListeners) {
keysModified.add(k);
} 

```

그럼 사용하고 있는 자바의 HashMap을 살펴 보겠습니다.

해시값을 계산하는 과정의 코드

```
static final int hash(Object key) {
int h;
return (key == null) ? 0 : (h = key.hashCode()) ^ (h >>> 16);
}

```

키의 hashCode()를 호출

상위 16비트를 하위 16비트와 XOR 연산

해시 충돌 가능성을 줄이는 효과

버킷 인덱스 계산

index = (n - 1) & hash

<img width="388" alt="스크린샷 2025-01-05 오후 6 39 54" src="https://github.com/user-attachments/assets/9d37e12f-3eb6-43d8-8cc9-c8a6670d433d" />

<img width="721" alt="스크린샷 2025-01-05 오후 6 40 10" src="https://github.com/user-attachments/assets/bd1e4ea2-723e-4b6d-ba2b-59ca66164fb1" />

    // 테이블이 비어있으면 생성
    if ((tab = table) == null || (n = tab.length) == 0)
        n = (tab = resize()).length;
    
    // 버킷이 비어있으면 새 노드 생성
    if ((p = tab[i = (n - 1) & hash]) == null)
        tab[i] = newNode(hash, key, value, null);
    else {
        Node<K,V> e; K k;
        // 첫 노드가 찾는 키와 같은 경우
        if (p.hash == hash && 
            ((k = p.key) == key || (key != null && key.equals(k))))
            e = p;
        // 트리 노드인 경우
        else if (p instanceof TreeNode)
            e = ((TreeNode<K,V>)p).putTreeVal(this, tab, hash, key, value);
        // 링크드 리스트인 경우
        else {
            for (int binCount = 0; ; ++binCount) {
                if ((e = p.next) == null) {
                    p.next = newNode(hash, key, value, null);
                    // 임계값 초과시 트리로 변환
                    if (binCount >= TREEIFY_THRESHOLD - 1)
                        treeifyBin(tab, hash);
                    break;
                }
                if (e.hash == hash &&
                    ((k = e.key) == key || (key != null && key.equals(k))))
                    break;
                p = e;
            }
        }
        // 기존 키가 있으면 값 갱신
        if (e != null) {
            V oldValue = e.value;
            if (!onlyIfAbsent || oldValue == null)
                e.value = value;
            return oldValue;
        }
    }
    ++modCount;
    // 크기 초과시 리사이징
    if (++size > threshold)
        resize();
    return null;

<img width="324" alt="스크린샷 2025-01-05 오후 6 44 37" src="https://github.com/user-attachments/assets/76689038-e782-4283-8408-87f5332d2c1d" />

임계값 도달 (노드 = 8)

TREEIFY_THRESHOLD (8) 도달
성능 저하 가능성 감지
링크드 리스트가 메모리를 더적게 사용하고 적은 데이터에 경우에는 순차검색도 충분히 빠름

성능 향상 log n / 최악의 경우에도 성능 보장
단점 메모리 사용 증가 