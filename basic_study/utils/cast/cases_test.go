package cast

import (
	"testing"
)

func TestToSnake(t *testing.T) {
	cases := map[string]string{
		"":                                  "",
		"A":                                 "a",
		"a":                                 "a",
		"AA":                                "aa",
		"AdjustThis":                        "adjust_this",
		"adjustThis":                        "adjust_this",
		"JSONData":                          "json_data",
		"JSONdata":                          "json_data",
		"dataHTML":                          "data_html",
		"dataHELLOWorld":                    "data_hello_world",
		"hello":                             "hello",
		"Gateway123TODO":                    "gateway123_todo",
		"AAAbbb":                            "aa_abbb",
		"userID":                            "user_id",
		"IDUserID":                          "id_user_id",
		"appIDUserID":                       "app_id_user_id",
		"QueryRowxContextJSONHTTPIDDataURL": "query_rowx_context_json_http_id_data_url",
	}
	for k, v := range cases {
		a := ToSnake(k)
		if a != v {
			t.Errorf("%s != %s", a, v)
		}
	}
}

func BenchmarkSnake(b *testing.B) {
	const c = "AdjustThisWord"
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		ToSnake(c)
	}
}
