---
author: wonhee
pubDatetime: 2025-01-05
modDatetime:
title: hash
description: hash, hash table, hash function
---

## Hash

### 사전적 의미

- 어원: 해시 함수가 입력 데이터를 섞어서 출력을 만들어 내는 과정에 대한 비유로 일상 용어인 hash 를 사용 (to chop up or make a mess out of something)
- 컴퓨터 과학에서의 정의: 임의 크기의 데이터를 고정 크기의 데이터로 변환하는 함수와 그 결과 값

## Hash Function

### 기본 특성

- 고정 출력 크기: 입력 크기에 상관없이 고정된 크기의 출력을 생성
- 결정론적 (Deterministic): 동일한 입력에 대해 항상 동일한 출력
- 균일성 (Uniformity): 출력 값이 가능한 범위에서 균등하게 분포해야 함
- 효율성: 빠르게 계산 가능하고 적은 리소스를 사용해야 함

### 비 암호화 해시 함수

- MurmurHash: 빠른 해싱 (SHA 대비 4배 빠름), 해시테이블 구현에 사용
- CRC32: 오류 검출에 사용

## 암호화 해시 함수

### 요구사항

- 충돌 저항성 (Collision Resistance): 서로 다른 입력에 대해 동일한 출력이 나오는 경우를 최소화해야 함 - 불가능에 가까워야 함
- 전위 저항성 (Preimage Resistance): 출력 값으로 입력 값을 찾는 것이 어려워야 함 - 주어진 해시값 h에 대해 h = H(m)을 만족하는 메시지 m을 찾기 어려워야 함
- 제2 전위 저항성 (2nd Preimage Resistance): 주어진 입력값 m1에 대해 h = H(m1)을 만족하는 다른 메시지 m2를 찾기 어려워야 함

### 주요 알고리즘

- MD5: 보안 취약점 존재해서 사용 권장하지 않음
- SHA 계열: 블록체인, Git, SSL/TLS
- bcrypt, scrypt: 비밀번호 해싱에 사용

## Hash Table

### 기본 구조

- 해시 함수: key를 해시값으로 변환 -> 모듈러 연산 후 배열의 인덱스로 사용
- 버킷: 해시값이 동일한 데이터를 저장하는 공간 - 배열의 각 요소
- 적재율 (Load Factor): 저장된 데이터 수 / 버킷 수 - 버킷 수 대비 저장된 데이터의 비율
- 리사이징: 적재율이 일정 수준을 넘으면 버킷 수 조절 - 보통 0.7 ~ 0.8을 기준으로 함

## Hash Collision

- 해시 충돌: 서로 다른 입력에 대해 동일한 해시값이 나오는 경우

## Solution

### Chaining

- Linked List: 버킷을 연결 리스트로 구현 - 충돌 발생 시 연결 리스트에 추가, 최악의 경우 O(n)
- Red-Black Tree: 더 나은 성능을 위해 RB Tree로 구현 - 항상 O(logn)의 탐색/삽입/삭제 보장

### Open Addressing

- Linear Probing: Probing = 탐사. 충돌 발생 시 선형적으로 +1씩 인덱스를 옮겨가며 빈 버킷 탐색
- Quadratic Probing: 충돌 발생 시 +1, +4, +9, +16, ... 순으로 탐색
- Double Hashing: 충돌 발생 시 두 번째 해시 함수를 사용해 다음 인덱스를 계산

### Chaining vs Open Addressing

- 캐시 효율성: Open Addressing이 Chaining보다 더 효율적
- Chaining은 다음 노드를 찾기 위해 포인터로 여기저기 메모리 여기저기 분포, 지역성이 떨어짐
- Open Addressing은 연속된 메모리 공간 사용으로 캐시 효율성이 높음
- 대부분의 언어는 Chaining을 사용: 구현이 단순하고 직관적, 삭제 연산이 더 간단, Load factor가 높아도 안정적인 성능 보장
- Open Addressing은 삭제 연산이 복잡하고 성능이 떨어질 수 있음, Python에서는 메모리 효율성을 위해 Open Addressing 사용
