package swarm

import "testing"

func TestBuildExhaustiveSwarmManifest(t *testing.T) {
	m := BuildExhaustiveSwarmManifest("forex-prop-firms")
	if m.TopicSlug != "forex-prop-firms" {
		t.Fatalf("unexpected slug: %s", m.TopicSlug)
	}
	if m.PerCallHardCapUSD > 0.15 {
		t.Fatalf("per-call hard cap exceeds $0.15 governor: %f", m.PerCallHardCapUSD)
	}
	if len(m.Substrates) < 5 {
		t.Fatalf("expected >= 5 substrates (Flash-Lite, Flash, Deep Research, Haiku, Local GPU), got %d", len(m.Substrates))
	}
	if len(m.SearchEngines) < 5 {
		t.Fatalf("expected >= 5 search engines (Exa, Tavily, Parallel, Brave, DuckDuckGo), got %d", len(m.SearchEngines))
	}
	if len(m.CrawlersAndRerankers) < 5 {
		t.Fatalf("expected >= 5 crawlers/rerankers (Firecrawl, Jina, msedge, Wayback, Cohere), got %d", len(m.CrawlersAndRerankers))
	}
	if len(m.FreePublicDatasets) < 7 {
		t.Fatalf("expected >= 7 free public datasets (BigQuery BLS/SEC/CFPB/Trends, FRED, BIS/ESMA, OpenAlex), got %d", len(m.FreePublicDatasets))
	}
}

func TestResolveTopicProfileAcrossFamilies(t *testing.T) {
	slugs := []string{
		"forex-prop-firms",
		"dropshipping",
		"ai-automation-agency",
		"airbnb-arbitrage",
		"youtube-channel",
	}
	for _, s := range slugs {
		p := ResolveTopicProfile(s)
		if p.FamilyCode == "" || p.ArchetypeCode == "" || len(p.OpenAlexQueries) == 0 {
			t.Fatalf("profile for %s incomplete: %+v", s, p)
		}
	}
}

func TestComputeDynamicBookSizingOrthogonalIndependence(t *testing.T) {
	// Young/emerging topic with sparse historical data (42 sources, 3 molds) BUT 5 high-quality Chapter 7 pivots:
	youngHighPivot := ComputeDynamicBookSizing("ai-automation-agency", "hot", 42, 3, 5)
	if youngHighPivot.ValidatedPivotsCount != 5 {
		t.Fatalf("expected young topic to preserve all 5 vetted pivots independently of topic age, got %d", youngHighPivot.ValidatedPivotsCount)
	}
	if youngHighPivot.StaticFrontMatterPages != 6 || youngHighPivot.StaticCoreBackMatterPages != 6 {
		t.Fatalf("expected static front/back matter bookends (6p each), got %d/%d", youngHighPivot.StaticFrontMatterPages, youngHighPivot.StaticCoreBackMatterPages)
	}
	if youngHighPivot.MinRealisticPages >= youngHighPivot.MaxRealisticPages {
		t.Fatalf("expected realistic page range, got [%d, %d]", youngHighPivot.MinRealisticPages, youngHighPivot.MaxRealisticPages)
	}

	// Older evergreen topic with deep historical data (180 sources, 5 molds) but only 2 vetted pivots:
	oldLowPivot := ComputeDynamicBookSizing("real-estate-wholesaling", "evergreen", 180, 5, 2)
	if oldLowPivot.ValidatedPivotsCount != 2 {
		t.Fatalf("expected older topic to preserve 2 vetted pivots without padding, got %d", oldLowPivot.ValidatedPivotsCount)
	}
}
