## Trie

## 구조적 특징

- 트라이 (/ˈtraɪ/, /ˈtriː/) 는 prefix tree, digital tree 로도 불리는 검색을 위한 트리형태의 자료구조
- 루트에서 특정 노드까지의 경로가 해당 키를 표현
- 모든 자식 노드는 부모 노드의 접두사를 공유함
- 주로 접두사를 공유하는 문자열을 압축하여 저장 - Patricia tree
- 시간 복잡도는 검색/삽입/삭제 모두 문자열 길이 L 의 경우 O(L)

### 노드

```ts
class TrieNode {
  children: Map<string, TrieNode>;
  isTerminal: boolean;
  // value?: any; // 필요 시 추가 데이터 저장

  constructor() {
    this.children = new Map();
    this.isTerminal = false;
  }
}
```

- 단어의 끝을 나타내는 표시 (`isTerminal`)
- 각 노드는 여러 자식 노드를 가질 수 있음
- 한 노드에서 자식 노드로 이동할 때 매핑된 key string을 거쳐감

## 동작

### 조회

```ts
search(word: string): boolean {
  let current = this.root;

  for (const char of word) {
    if (!current.children.has(char)) {
      return false;
    }
    current = current.children.get(char)!;
  }

  return current.isTerminal;
}
```

### 삽입

```ts
insert(word: string): void {
  let current = this.root;

  for (const char of word) {
    if (!current.children.has(char)) {
      current.children.set(char, new TrieNode());
    }

    current = current.children.get(char)!;
  }

  current.isTerminal = true;
}
```

### 삭제

```ts
delete(word: string): boolean {
  return this.deleteHelper(this.root, word, 0);
}

// return 값으로 재귀 호출 중 삭제 여부 판별
private deleteHelper(
  current: TrieNode,
  word: string,
  index: number
): boolean {
  if (index === word.length) {
    // 해당 경로가 존재하지만 저장된 단어가 아닌 경우
    // ex. trees 가 저장되어있고 tree 삭제 시도 시
    if (!current.isTerminal) {
      return false;
    }

	// 한 단어의 끝에 해당하는 표시 제거 - 사실 상 해당 단어를 삭제한 것
    current.isTerminal = false;
    // 자식이 없는 경우 아예 노드를 삭제 가능 - a
    return current.children.size === 0;
  }

  const char = word[index];
  if (!current.children.has(char)) {
    return false;
  }

  const shouldDeleteCurrentNode = this.deleteHelper(
    current.children.get(char),
    word,
    index + 1
  );

  if (shouldDeleteCurrentNode) {
    current.children.delete(char); // a 에 의해 자식 노드 삭제
    // 더 이상 자식이 없고, 다른 단어의 끝이 아닌 경우 현재 노드 삭제 가능
    return current.children.size === 0 && !current.isTerminal;
  }

  return false;
}
```

## 주요 변형

### Patricia tree

- Practical Algorithm To Retrieve Information Coded In Alphanumeric
- 압축된 trie의 한 형태
- 단일 자식만 갖는 노드들을 합축해서 공간 최적화
- ex. "team", "test" 를 저장한 경우:

```
일반 Trie:         Patricia Tree:
    root              root
     |                 |
     t                 t
     |                / \
     e              ea  es
    / \              |   |
   a   e            m    t
   |   |
   m   s
       |
       t
```

- 검색 시간은 O(k), 여기서 k는 접두사의 수

## 활용

### 라우팅

1. IP 라우팅
2. DNS 라우팅
3. 브라우져 URL 라우팅

### 블록체인

Merkel Patricia Trie (이더리움)

- Patricia Tree 에 Merkle Tree 특성을 추가한 것
- 주소 문자열을 경로로 관리하고 리프 노드에 해당 계정의 상태를 관리
- 모든 노드가 해시로 참조됨
  - 데이터 무결성 보장
  - 변경 감지 가능
- 부분 업데이트 효율성
  - 전체 트리를 다시 계산할 필요 없음
  - 상태가 변경되면 해당 경로의 해시만 재계산
- 상태 검증
  - 특정 계정의 잔액이나 상태가 올바른 지 검증
  - 전체 블록체인의 데이터가 없이도 부분 검증 가능
- 효울적인 상태 업데이트
  - 기존 작액: A 계정 10ETH
  - 트랜젝션: A -> B 2ETH 전송
  - 변경되는 것: A, B 계정 경로의 해시들만 재계산
  - A 계정 경로 = A 계정 주소가 0xabcd... 인 경우

```
Root Hash
   |
[Branch Node] ('a'에 해당하는 슬롯)
   |
[Extension Node] ('bc'와 같은 공통 경로)
   |
[Branch Node] ('d'에 해당하는 슬롯)
   |
[Leaf Node] (나머지 주소 + 계정 상태)
```
