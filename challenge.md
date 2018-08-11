FizzBuzzのHTTPサーバをGoで書いたもの。
FizzBuzzは[Wiki参照](https://ja.wikipedia.org/wiki/Fizz_Buzz)
　　　
# 仕様

HTTPで `http://localhost:8080/fizzbuzz/1` のようにリクエストを送ったら、HTTPステータスコード200で結果を返す。
もし数字じゃない値を送ったら、ステータスコード400でinvalidのメッセージを返す。
細かい仕様は適当で良い。

## 例
```
/fizzbuzz/9 → "9"が返る
/fizzbuzz/10 → "Buzz" が返る
/fizzbuzz/15 → "Fizz Buzz" が返る
/fizzbuzz/hoge → "invalid number" が返る
```
