package ja3_test

import (
	"testing"

	"github.com/airdb/go-demo/ja3"
)

func TestJa3(t *testing.T) {
	ja3List := []ja3.Browser{
		{Ja3: "771,4865-4866-4867-49196-49195-52393-49200-49199-52392-49188-49187-49162-49161-49192-49191-49172-49171-157-156-61-60-53-47-49160-49170-10,0-23-65281-10-11-16-5-13-18-51-45-43-21,29-23-24-25,0", UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36 Edg/92.0.902.67"},
	}
	for _, j := range ja3List {
		ja3.Req(j)
	}
}
