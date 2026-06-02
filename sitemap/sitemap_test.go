package sitemap

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestEnsureOutputDir(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(dir)

	out, err := ensureOutputDir("generated")
	if err != nil {
		t.Fatal(err)
	}
	want := filepath.Join(dir, "generated")
	if out != want {
		t.Fatalf("got %q, want %q", out, want)
	}
	if _, err := os.Stat(want); err != nil {
		t.Fatal(err)
	}
}

func TestGetLastModifiedOrNow(t *testing.T) {
	fixed := time.Date(2024, 3, 15, 12, 0, 0, 0, time.UTC)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Last-Modified", fixed.Format(time.RFC1123))
	}))
	defer srv.Close()

	got, err := getLastModifiedOrNow(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	if got != "2024-03-15" {
		t.Fatalf("got %q, want 2024-03-15", got)
	}
}

func TestSitemapSave(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(dir)

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Last-Modified", "Mon, 15 Mar 2024 12:00:00 GMT")
	}))
	defer srv.Close()

	s, err := NewSitemap().Path("out")
	if err != nil {
		t.Fatal(err)
	}
	if err := s.AddURL(srv.URL); err != nil {
		t.Fatal(err)
	}
	if err := s.Save(); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "out", "sitemap.xml"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)
	if !strings.Contains(content, srv.URL) {
		t.Fatalf("expected sitemap to contain %q, got:\n%s", srv.URL, content)
	}
	if !strings.Contains(content, "2024-03-15") {
		t.Fatalf("expected lastmod in sitemap, got:\n%s", content)
	}
}

func TestVideoSitemapSave(t *testing.T) {
	dir := t.TempDir()
	t.Chdir(dir)

	v, err := NewVideoSitemap().Path("out")
	if err != nil {
		t.Fatal(err)
	}
	if err := v.AddVideoURL(VideoURL{
		Loc: "https://www.example.com/videos/video1.html",
		Videos: []Video{{
			ThumbnailLoc: "https://www.example.com/thumb.png",
			Title:        "title",
			Description:  "desc",
			ContentLoc:   "https://www.example.com/video.mp4",
			PlayerLoc:    "https://www.example.com/player",
		}},
	}); err != nil {
		t.Fatal(err)
	}
	if err := v.Save(); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(filepath.Join(dir, "out", "sitemap_video.xml"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)
	if !strings.Contains(content, "video:thumbnail_loc") {
		t.Fatalf("expected video namespace elements, got:\n%s", content)
	}
}
