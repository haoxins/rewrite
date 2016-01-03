package rewrite

import "net/http/httptest"
import "net/http"
import "testing"

type testFixture struct {
	from string
	to   string
}

type testCase struct {
	pattern  string
	to       string
	fixtures []testFixture
}

func TestRewrite(t *testing.T) {
	tests := []testCase{
		testCase{
			pattern: "/a",
			to:      "/b",
			fixtures: []testFixture{
				testFixture{from: "/a", to: "/b"},
			},
		},
		testCase{
			pattern: "/a/(.*)",
			to:      "/bb",
			fixtures: []testFixture{
				testFixture{from: "/a", to: "/a"},
				testFixture{from: "/a/", to: "/bb"},
				testFixture{from: "/a/a", to: "/bb"},
				testFixture{from: "/a/b/c", to: "/bb"},
			},
		},
	}

	for _, test := range tests {
		t.Logf("Test - pattern: %s, to: %s", test.pattern, test.to)

		for _, fixture := range test.fixtures {
			req, err := http.NewRequest("GET", fixture.from, nil)
			if err != nil {
				t.Fatalf("Fixture %v - create HTTP request error: %v", fixture, err)
			}

			h := NewHandler(map[string]string{
				test.pattern: test.to,
			})

			t.Logf("From: %s", req.URL.Path)
			if req.URL.Path != fixture.from {
				t.Errorf("Invalid test fixture: %s", fixture.from)
			}

			res := httptest.NewRecorder()
			h.ServeHTTP(res, req)

			t.Logf("Rewrited: %s", req.URL.Path)
			if req.URL.Path != fixture.to {
				t.Errorf("Test failed - pattern: %s, to: %s. Fixture from %s to %s",
					test.pattern, test.to, fixture.from, fixture.to)
			}

			t.Log(res.Header().Get(headerField))
		}
	}
}
