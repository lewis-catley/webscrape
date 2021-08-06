package scraper

import (
	"bytes"
	"fmt"
	"net/url"
	"sync"

	"github.com/lewis-catley/webscrape/scraper/pkg/models"
	"github.com/lewis-catley/webscrape/scraper/pkg/page"
	"github.com/lewis-catley/webscrape/scraper/pkg/web"
)

type Scraper struct {
	ID          string
	urlsVisited map[string]interface{}
	baseURL     *url.URL
	mu          sync.Mutex
	wg          *sync.WaitGroup
	queue       chan *url.URL
	JobOut      chan models.URLSFound
	Finished    chan string
}

/**
What should happen in this pacakge
 - Save to mongoOnce a URL is finished
 - Only store the URL's Visited as we need this to validate where we have been
 - Once everything has been done call jobFinished
**/

const MAX_JOBS_PROCESSED = 10

func New(id string, baseURL *url.URL) *Scraper {
	return &Scraper{
		ID:          id,
		urlsVisited: map[string]interface{}{},
		baseURL:     baseURL,
		wg:          new(sync.WaitGroup),
		queue:       make(chan *url.URL, MAX_JOBS_PROCESSED),
		JobOut:      make(chan models.URLSFound),
		Finished:    make(chan string),
	}
}

func (s *Scraper) FindAllURLs() {
	s.wg.Add(1)
	go s.findURLs(s.baseURL)
	go s.processJobs()
	s.wg.Wait()
	fmt.Println("I finish processing", s.baseURL)
	s.Finished <- s.ID
}

func (s *Scraper) processJobs() {
	for {
		select {
		case nextJob := <-s.queue:
			fmt.Printf("Processing Job %s, base URL %s\n", nextJob.String(), s.baseURL.String())
			s.wg.Add(1)
			go s.findURLs(nextJob)
		case <-s.Finished:
			return
		}
	}
}

func (s *Scraper) findURLs(u *url.URL) {
	defer s.wg.Done()
	p, err := s.visitURL(u)
	if err != nil {
		// TODO: Need to mark an error on the url
		return
	}
	go p.GetLinks()

	for {
		select {
		case foundURL := <-p.URLChannel:
			if s.shouldVisitURL(foundURL) {
				s.queue <- foundURL
				fmt.Printf("Found URL %s, base URL %s\n", foundURL.String(), s.baseURL.String())
			}
		case <-p.QuitChannel:
			s.JobOut <- models.URLSFound{
				URL:       u.String(),
				URLSFound: p.Print(),
			}
			return
		}
	}
}

func (s *Scraper) shouldVisitURL(u *url.URL) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if u == nil {
		return false
	}

	// Don't visit a page that is not the same host
	if s.baseURL.Hostname() != u.Hostname() {
		return false
	}

	urlNoParams := web.GenerateURLNoParams(u)

	// Have we already visited the page
	if _, ok := s.urlsVisited[urlNoParams]; ok {
		return false
	}

	return true
}

func (s *Scraper) visitURL(url *url.URL) (*page.Page, error) {
	p := s.addPageAsVisited(url)
	body, err := web.GetURLContent(*url)
	if err != nil {
		return nil, err
	}
	nodes, err := web.ParseReaderHTMLNode(bytes.NewReader(*body))
	if err != nil {
		return nil, err
	}
	p.SetNodes(nodes)
	return p, nil
}

func (s *Scraper) addPageAsVisited(url *url.URL) *page.Page {
	s.mu.Lock()
	defer s.mu.Unlock()
	p := page.New(s.baseURL)
	s.urlsVisited[web.GenerateURLNoParams(url)] = nil
	return p
}
