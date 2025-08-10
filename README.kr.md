네, `rzh` 프로그램의 문서를 기술 문서 형식에 맞춰 한국어로 번역해 드리겠습니다.

-----

# rzh

`rzh`는 Zsh 히스토리를 관리하는 프로그램입니다.

## 목차

  - [주요 기능](#주요-기능)
  - [주의사항](#주의사항)
  - [사용법](#사용법)
  - [설치](#설치)
  - [라이선스](#라이선스)

## 주요 기능

  - Zsh 히스토리에 명령어 추가
  - Zsh 히스토리에서 명령어 삭제
  - Zsh 히스토리 중복 항목 제거
  - Zsh 히스토리 백업

## 주의사항

이 프로그램은 홈 디렉터리(예: `~/.zsh_history`)에 직접 위치한 `.zsh_history` 파일을 대상으로 작동합니다. 만약 파일이 다른 위치에 있다면, 프로그램이 올바르게 작동하지 않습니다.

## 사용법

### 명령어 추가

```shell
rzh add "<command1>" ["<command2>"...] [--backup/-b] [--force/-f]
```

하나 이상의 새 명령어를 `.zsh_history` 파일에 추가합니다.

  - `--backup` 플래그는 추가 작업을 수행하기 전에 히스토리 파일의 전체 백업을 생성합니다.
  - `--force` 플래그는 확인 메시지를 건너뛰고 즉시 명령어를 추가합니다.

<!-- end list -->

```shell
# "ls -la" 명령어 하나를 히스토리에 추가
> rzh add "ls -la"

# 여러 명령어를 한 번에 추가
> rzh add "go build" "go test" "git push"

# 백업 생성 후 여러 명령어 추가
> rzh add "npm install" "npm run dev" --backup
```

### 명령어 제거

```shell
rzh remove "<command1>" ["<command2>"...] [--prefix/-p] [--backup/-b] [--force/-f]
```

지정한 검색어 중 하나라도 포함하는 모든 명령어를 `~/.zsh_history` 파일에서 제거합니다. 하나 이상의 검색어를 제공할 수 있습니다.

기본적으로, 삭제될 모든 일치하는 라인을 보여주고 확인을 요청합니다. 검색은 대소문자를 구분합니다.

  - `--backup` 플래그는 제거 작업을 수행하기 전에 히스토리 파일의 전체 백업을 생성합니다.
  - `--force` 플래그는 확인 프롬프트를 건너뛰고 일치하는 라인을 즉시 삭제합니다.

### 히스토리 중복 제거

```shell
rzh dedup [--force/-f]
```

중복된 명령어 항목을 제거하여 `.zsh_history` 파일을 정리합니다.

각 중복된 명령어에 대해 타임스탬프 기준 **가장 최근의 항목**을 남기고, 오래된 다른 모든 항목을 제거합니다.

  - `--force` 플래그는 확인 프롬프트를 건너뜁니다.

<!-- end list -->

```shell
# 확인 프롬프트와 함께 히스토리 중복 제거
> rzh dedup

# 확인 없이 히스토리 중복 제거
> rzh dedup --force
```

### 백업 관리

```shell
rzh backup [create | list | remove | restore] [--force/-f]
```

`~/rzh/backup/`에 저장된 히스토리 파일의 백업을 관리합니다. 기본 동작은 `create`입니다.

  - `create`: 타임스탬프 형식의 파일명으로 새 백업을 생성합니다.

  - `list`: 사용 가능한 모든 백업을 보여줍니다.

  - `remove`: 백업을 선택하고 삭제할 수 있는 대화형 인터페이스를 실행합니다.

  - `restore`: 사용 가능한 백업 목록을 번호와 함께 보여주고, 복원할 백업을 선택하라는 메시지를 표시합니다. 선택된 백업의 내용이 현재 `.zsh_history` 파일을 완전히 덮어씁니다.

  - `--force` 플래그는 `remove` 및 `restore` 시 대화형 프롬프트를 건너뛰는 데 사용됩니다.

<!-- end list -->

```shell
# 새 백업 생성 (기본 동작)
> rzh backup
> rzh backup create

# 모든 백업 목록 보기
> rzh backup list

# 대화형으로 복원할 백업 선택
> rzh backup restore
```

## 설치

1.  `rzh`의 [GitHub 릴리스 페이지](https://github.com/chebread/rzh/releases)를 방문합니다.
2.  사용 중인 운영체제와 아키텍처에 맞는 파일을 다운로드합니다.
3.  다운로드한 파일의 압축을 해제합니다.
4.  `rzh` 실행 파일을 실행합니다.
5.  더 쉽게 접근하려면 `rzh` 실행 파일을 시스템의 PATH 환경 변수에 추가하는 것을 고려하세요.

## 라이선스

MIT LICENSE &copy; 2025 Cha Haneum