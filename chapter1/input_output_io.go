package main

// import "bytes"

// import (
// 	"bytes"
// 	"errors"
// 	"io"
// )

// import (
// 	"io"
// 	"testing"
// )

// import (
// 	"bytes"
// 	"fmt"
// 	"io"
// 	"strings"
// )

// ファイルがPNG形式かファイルの最初のマジックナンバーで確かめるコード
// func IsPNG(r io.Reader) (bool, error) {
// 	magicnum := []byte{137, 80, 78, 71}
// 	buf := make([]byte, len(magicnum))
// 	_, err := io.ReadAtLeast(r, buf, len(buf))
// 	if err != nil {
// 		return false, err
// 	}
// 	return bytes.Equal(magicnum, buf), nil
// }

// func main() {
// 	data := strings.NewReader(string([]byte{137, 80, 78, 71}))
// 	IsPNG, err := IsPNG(data)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println("Is PNG:", IsPNG)
// 	}
// }

// neverEndingはカスタムクラスでキャストした値を無限に供給するリーダー
// type neverEnding byte

// func (b neverEnding) Read(p []byte) (n int, err error) {
// 	for i := range p {
// 		p[i] = byte(b)
// 	}
// 	return len(p), nil
// }

// func TestIsPNG(t *testing.T) {
// 	n, want := int64(10), false
// 	r := io.LimitReader(neverEnding('x'), n)
// 	got, err := IsPNG(r)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if want != got {
// 		t.Error(want, "!=", got)
// 	}
// }

// MultiReaderはReadFullで読み進めたデータ分をもとの読み込まれたデータと読み込んだデータを連結することで元に戻す
// func NewPNG(r io.Reader) (io.Reader, error) {
// 	magicnum := []byte{137, 80, 78, 71}
// 	buf := make([]byte, len(magicnum))
// 	if _, err := io.ReadFull(r, buf); err != nil {
// 		return nil, err
// 	}
// 	if !bytes.Equal(magicnum, buf) {
// 		return nil, errors.New("PNG画像ではありません")
// 	}
// 	pngImg := io.MultiReader(bytes.NewReader(magicnum), r)
// 	return pngImg, nil
// }

// io.ReeReaderを使って、安全にチェック
// func TestUpperCount(t *testing.T) {
// 	str, want := "AbcD", 2
// 	var buf bytes.Buffer
// 	r := io.TeeReader(strings.NewReader(str), &buf)
// 	got, err := UpperCount(r)
// 	if err != nil { t.Fatal(err) }
// 	if got != want { t.Error(want, "!=", got) }
// 	if str != buf.String() { t.Error("読み込んだ文字列が一致しない") }
// }

// io.Pipeはstreamの役割を果たすpr(読み込み)pw(書き込み) encはpwに流れてきたmessageをjsonにエンコード
// func Post(m *Message) (rerr error) {
// 	pr, pw := io.Pipe()
// 	go func() {
// 		defer pw.Close()
// 		enc := json.NewEncoder(pw)
// 		err := enc.Encode(m)
// 		if err != nil { rerr = err }
// 	} ()
// 	const url = "http://example.com"
// 	const contentType = "application/json"
// 	_, err := http.Post(url, contentType, pr)
// 	if err != nil { return err }
// 	return nil
// }

// func main() {
// 	f, err := os.Create("sample.txt")
// 	if err != nil { /*エラー処理 */}
// 	h := sha256.New()
// 	w := io.MultiWriter(f, h)
// 	_, err = io.WriteString(w, "hello")
// 	if err != nil { /*エラー処理*/ }
// 	err = f.Close()
// 	if err != nil { /*エラー処理*/ }
// 	fmt.Printf("%x/n", h.Sum(nil))
// }
