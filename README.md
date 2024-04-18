# 경희대학교 
## 풀스텍 네트워크 서비스 
### 이성원 교수님

이 레포지 토리는 Nodejs에서 제공하는 "net" 모듈과  socket.io 를 비교하기 위해서 만든 레포지 토리 입니다.

기존의 python 코드를 js로 변경하였습니다.
1. branch main - README
2. branch nodeJs - nodejs 로 구현한 TCP -socket code
3. branch python - python 으로 구현한  TCP -socket code
4. branch golang - golang 으로 구현한 TCP -socket code
#### NodeJs net에 대한 추가적인 설명 참조

공식문서 : https://nodejs.org/api/net.html

Node.js의 기본 내장 모듈인 net 모듈을 사용하여 소켓 통신을 구현할 수 있습니다. net 모듈은 TCP 소켓을 다루는 데 사용됩니다. 이 모듈을 사용하면 소켓을 통해 데이터를 읽고 쓸 수 있습니다. 이것은 전통적인 소켓 프로그래밍의 방법입니다.

반면에 socket.io는 Node.js를 위한 실시간 웹 소켓 통신 라이브러리입니다. 이는 웹 소켓을 사용하여 클라이언트와 서버 간의 양방향 통신을 구현하는 것을 목적으로 합니다. socket.io는 웹 소켓을 사용하되, 웹 소켓을 지원하지 않는 브라우저에 대해서는 다양한 폴백 옵션을 제공하여 실시간 통신을 가능하게 합니다. 또한, socket.io는 이벤트 기반의 통신을 간편하게 처리할 수 있는 API를 제공하므로, 개발자가 더 쉽게 웹 소켓을 사용할 수 있습니다.

간단하게 말하면, net 모듈은 소켓 프로그래밍에 사용되는 기본적인 TCP 소켓을 다루는 데에 사용되고, socket.io는 실시간 웹 소켓 통신을 구현하는 데에 사용됩니다. socket.io는 더 높은 수준의 추상화를 제공하고, 웹 애플리케이션에서 실시간 통신을 구현하는 데에 더 적합합니다.



