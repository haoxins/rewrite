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
		testCase{
			pattern: "/r/(.*)",
			to:      `/r/v1/$1`,
			fixtures: []testFixture{
				testFixture{from: "/a", to: "/a"},
				testFixture{from: "/r", to: "/r"},
				testFixture{from: "/r/a", to: "/r/v1/a"},
				testFixture{from: "/r/a/b", to: "/r/v1/a/b"},
			},
		},
		testCase{
			pattern: "/r/(.*)/a/(.*)",
			to:      `/r/v1/$1/a/$2`,
			fixtures: []testFixture{
				testFixture{from: "/r/1/2", to: "/r/1/2"},
				testFixture{from: "/r/1/a/2", to: "/r/v1/1/a/2"},
				testFixture{from: "/r/1/a/2/3", to: "/r/v1/1/a/2/3"},
			},
		},
		testCase{
			pattern: "/r/(.*)/a/(.*)",
			to:      `/r/v1/$2/a/$1`,
			fixtures: []testFixture{
				testFixture{from: "/r/1/a/2", to: "/r/v1/2/a/1"},
				testFixture{from: "/r/1/a/2/3", to: "/r/v1/2/3/a/1"},
			},
		},
		testCase{
			pattern: "/from/:one/to/:two",
			to:      "/from/:two/to/:one",
			fixtures: []testFixture{
				testFixture{from: "/from/123/to/456", to: "/from/456/to/123"},
				testFixture{from: "/from/abc/to/def", to: "/from/def/to/abc"},
			},
		},
		testCase{
			pattern: "/from/:one/to/:two",
			to:      "/:one/:two/:three/:two/:one",
			fixtures: []testFixture{
				testFixture{from: "/from/123/to/456", to: "/123/456/:three/456/123"},
				testFixture{from: "/from/abc/to/def", to: "/abc/def/:three/def/abc"},
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
				t.Errorf("Test failed \n pattern: %s, to: %s, \n fixture: %s to %s, \n result: %s",
					test.pattern, test.to, fixture.from, fixture.to, req.URL.Path)
			}

			t.Logf("Flag: %s", res.Header().Get(headerField))
		}
	}
}
