# About postkr

[인터넷 우체국 오픈API][2]를 위한 Go언어 패키지 입니다.

현재, 우편번호, 종추적조회, EMS 중 우편번호 기능 만 구현되어 있습니다.

# Documentation

## Prerequisites

[Install Go][1]

## Installation

    $ go get github.com/suapapa/go_postkr

> 참고: 내부 인코딩변환(cp949 tofrom utf8)을 위해 Donovan Jimenez의
[iconv-go][5] 패키지를 사용하며 위 명령으로 함께 설치됩니다.

## General Documentation

[godoc.org][4] 링크에서 온라인 문서를 보거나,
다음 명령으로 로컬에서 문서를 볼 수 있습니다.

    go doc github.com/suapapa/go_postkr


## Example

postkr을 사용하려면, 인터넷 우체국의 [오픈API사용신청][3]을 통해 키를
발급 받아야 합니다. 발급받은 키를 아래 코드의 `yourKey` 변수에 할당하세요.

    package main

    import (
            "fmt"
            "github.com/suapapa/go_postkr"
    )

    func main() {
            // 발급받은 오픈API 키로 바꾸세요
            yourKey := "abcedfghijklmnopqrstuvwxyz1234"

            s := postkr.NewService(yourKey)
            l, err := s.SerchZipCode("내곡동")
            if err != nil {
                    fmt.Println(err)
                    return
            }

            for _, p := range l {
                    fmt.Printf("(%s) %s\n", p.Code, p.Address)
            }
    }

실행 결과

    $ go run postkr.go
    (137180) 서울 서초구 내곡동
    (412260) 경기 고양시 덕양구 내곡동
    (210160) 강원 강릉시 내곡동
    (210701) 강원 강릉시 내곡동 관동대
    ...

# Author

Homin Lee &lt;homin.lee@suapapa.net&gt;

# Copyright & License

Copyright (c) 2012, Homin Lee.
All rights reserved.
Use of this source code is governed by a BSD-style license that can be
found in the LICENSE file.

[1]: http://golang.org/doc/install
[2]: http://biz.epost.go.kr/eportal/custom/custom_9.jsp?subGubun=sub_3&subGubun_1=cum_17&gubun=m07
[3]: http://biz.epost.go.kr/eportal/custom/custom_11.jsp?subGubun=sub_3&subGubun_1=cum_19&gubun=m07
[4]: http://godoc.org/github.com/suapapa/go_postkr
[5]: https://github.com/djimenez/iconv-go
