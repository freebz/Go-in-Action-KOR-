// 예제 4.6  원소의 포인터 요소에 접근하기

// 다섯 개의 원소를 가지는 정수 포인터 배열을 선언한다.
// 인덱스 0과 1을 정수 포인터로 초기화한다.
array := [5]*int{0: new(int), 1: new(int)}

// 인덱스 0과 1에 값을 대입한다.
*array[0] = 10
*array[1] = 20
