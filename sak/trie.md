## Trie
- 문자열을 저장하고 효율적인 탐색하기 위한 트리 형태의 자료구조
- 각 노드가 문자를 저장 및 경로로 단어를 표현
- 루트는 비어있다.
- 단어 끝에 마커 존재

## 특징
- 공통 접두어를 가진 문자열들은 같은 경로 공유
- 문자열 검색시 문자열 길이만큼만 탐색하면 됨
- 노드는 해당 문자와 자식 노드들의 정보를 가짐

## 장점
- 빠른 문자열 검색 (문자열 길이에만 비례)
- 접두사 검색에 매우 효율적
- 자동완성, 사전 검색 등에 최적화
- 정렬된 상태로 문자열 보관 가능

## 단점
- 각 노드마다 자식 노드 정보를 위한 큰 메모리 필요
- 해시테이블보다 구현이 복잡
- 단일 문자열 검색만 하는 경우 공간 낭비
- 삭제 연산이 복잡함

## 메모리 관점
- 각 노드는 알파벳 크기만큼의 포인터 배열 필요
- 문자열들의 공통 접두어가 많을수록 메모리 효율 증가
- 최악의 경우 문자열 개수 * 길이 * 알파벳 크기만큼 공간 필요
- 실제 저장할 문자열보다 더 많은 공간 차지 가능

## 성능
- 검색: O(M) (M은 찾는 문자열 길이)
- 삽입: O(M)
- 삭제: O(M)
- 메모리: O(ALPHABET_SIZE * M * N)

(N: 문자열 개수, M: 평균 문자열 길이)

## 예시 <apache에 Trie 라이브러리 >

### 단순 사용
```
package io.delightlabs.trie

import org.apache.commons.collections4.trie.PatriciaTrie

fun main() {
    val trie = PatriciaTrie<Int>()

    trie["leeSak"] = 1
    trie["leelee"] = 2

    println(trie.prefixMap("lee"))
    
    // res = [leesak,1],[leelee,2]
}
```

### 검색 부분
```
    @Override
    public V get(final Object k) {
        final TrieEntry<K, V> entry = getEntry(k);
        return entry != null ? entry.getValue() : null;
    }
    TrieEntry<K,V> getEntry(final Object k) {
        final K key = castKey(k);
        if (key == null) {
            return null;
        }
        
        // 이부분에서 비트 단위로 키를 비교? -> 성능 최적화 부분
        final int lengthInBits = lengthInBits(key);
        final TrieEntry<K,V> entry = getNearestEntryForKey(key, lengthInBits);
        // 가까운 노드와 이부분에서 실제 키 일치하는 지 확인 -> 불필요 비교 줄임 -> 성능 개선 추가로 정확하게
        return !entry.isEmpty() && compareKeys(key, entry.key) ? entry : null;
        
    }
    
    TrieEntry<K, V> getNearestEntryForKey(final K key, final int lengthInBits) {
        TrieEntry<K, V> current = root.left;
        TrieEntry<K, V> path = root;
        while(true) {
            if (current.bitIndex <= path.bitIndex) {
                return current;
            }

            path = current;
            // 비트에 따라 0 left , 1 right
            if (!isBitSet(key, current.bitIndex, lengthInBits)) {
                current = current.left;
            } else {
                current = current.right;
            }
        }
    }
```
문자열이 아닌 비트 단위로 키를 비교 하고 있음 => 시간복잡도는 O(bit길이)

getEntry() 부분을 보면 노드가 비어있으면 바로 리턴

ex) 검색 "sak" 인데 근접 노드 sao 발견시 false로 반환

=> 정학성 향상

### 삽입 부분
코드는 기본적으로 별거없고 비트 인덱스로 삽입하는 로직 + 기존키 중복 같은 확인 코드

### 삭제 부분
```
  @Override
    public V remove(final Object k) {
        if (k == null) {
            return null;
        }

        final K key = castKey(k);
        final int lengthInBits = lengthInBits(key);
       
        TrieEntry<K, V> current = root.left;
        TrieEntry<K, V> path = root;
        while (true) {
            if (current.bitIndex <= path.bitIndex) {
                if (!current.isEmpty() && compareKeys(key, current.key)) {
                    return removeEntry(current);
                }
                return null;
            }

            path = current;

            if (!isBitSet(key, current.bitIndex, lengthInBits)) {
                current = current.left;
            } else {
                current = current.right;
            }
        }
    }
```
#### 순서
1. 루트의 leftChild 선 검색
2. 키 일치 시 삭제
3. 비트 비교후 좌,우 이동 선택

### 성능 최적화 부분
```
   final int bitIndex(final K key, final K foundKey) {
        return keyAnalyzer.bitIndex(key, 0, lengthInBits(key), foundKey, 0, lengthInBits(foundKey));
    }
    
    int bitIndex = getKeyAnalyzer().bitIndex(key1, 0, lengthInBits(key1),
                                        key2, 0, lengthInBits(key2));

    boolean isSet = isBitSet(key1, 16, lengthInBits(key1));
```
키 비교시 전체 문자열이 아닌 비트 단위로 비교 

근데 문자열이랑 비트 단위로 비교를 통한 성능 최적화를 예시를 통해 보자

예를 들어  sihyun sak jamin

그러면 여기서 "sak" 을 검색해보자

일단 현재 그래프는 sihyun - leftChild(jaemin) - rightChild(sak)
1. "sak" 검색 
2. sihyun - s일치  - i 비교 a -> right child -> "sak 발견" -> s , a, k 3번 비교
3. "sak 발견" 비교연산 s-s , a vs i , 문자 일치 3번 = 총 5회 비교
만약에 비트로 계산 하는 경우 (비트는 맘대로 찍었습니다.)

0101010(s) -비트 i,h,y,u,n

0101010(s) -비트 a,k

0111111(j) -비트 a,m,i,n

트리 구조 first 비트 -> leftChild(jaemin) - 두번째 단어 비트 9 - left(sak) -right(sihyun)

1. s(첫번째 비트 비교 ) 오른쪽 1회
2. a(비교 9번째) 2회
3. "sak 발견" 총 2회비교
### 메모리 최적화

지연초기화 , 스레드 간 가시성 보장 등

### 캐시 지역성 최적화
```
    /**
     * Returns the nearest entry for a given key.  This is useful
     * for finding knowing if a given key exists (and finding the value
     * for it), or for inserting the key.
     *
     * The actual get implementation. This is very similar to
     * selectR but with the exception that it might return the
     * root Entry even if it's empty.
     */
    TrieEntry<K, V> getNearestEntryForKey(final K key, final int lengthInBits) {
        TrieEntry<K, V> current = root.left;
        TrieEntry<K, V> path = root;
        while(true) {
            if (current.bitIndex <= path.bitIndex) {
                return current;
            }

            path = current;
            if (!isBitSet(key, current.bitIndex, lengthInBits)) {
                current = current.left;
            } else {
                current = current.right;
            }
        }
    }
```
메모리 인접 노드를 연속적으로 접근 자식 노드들이 캐시에 있음
Ex)  

        1

    2       3

        s       d
3번 접근시 s,d 로드

예측 가능한 다음 위치로 캐시 히트 상승
