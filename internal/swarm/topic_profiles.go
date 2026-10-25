package swarm

import (
	"strings"
)

// DynamicBookSizingSpec governs natural range-based book length: Front Matter (~6p) and core Back Matter
// (~6p including the 2-page Research Machine spread, Glossary, Contract, Results Typical & Colophon) are
// fairly static structural bookends, while every interior section (Ch 3/5 empirical history, Ch 4 molds,
// Ch 7 pivots, Ch 8 exit/honest-test) and the Bulk Source List scale INDEPENDENTLY based strictly on what
// cleared the quality & context bar—never assuming topic age or data volume correlates with pivot count or mold count,
// and never stretching (padding) or compacting (squeezing) to hit a fixed page number.
type DynamicBookSizingSpec struct {
	TopicSlug                    string  `yaml:"topic_slug" json:"topic_slug"`
	TopicAgeNote                 string  `yaml:"topic_age_note" json:"topic_age_note"`
	EstimatedTopicAgeYears       float64 `yaml:"estimated_topic_age_years" json:"estimated_topic_age_years"`
	ValidatedSourceCount         int     `yaml:"validated_source_count" json:"validated_source_count"`
	PopulatedMoldsCount          int     `yaml:"populated_molds_count" json:"populated_molds_count"`
	ValidatedPivotsCount         int     `yaml:"validated_pivots_count" json:"validated_pivots_count"`
	StaticFrontMatterPages       int     `yaml:"static_front_matter_pages" json:"static_front_matter_pages"`           // Fairly static ~6 pages
	StaticCoreBackMatterPages    int     `yaml:"static_core_back_matter_pages" json:"static_core_back_matter_pages"`   // Fairly static ~6 pages (2p Research Machine Spread + Glossary + Contract + Colophon + Index)
	NaturalBibliographyPageRange string  `yaml:"natural_bibliography_page_range" json:"natural_bibliography_page_range"` // Scales independently with harvested source count
	NaturalInteriorPageRange     string  `yaml:"natural_interior_page_range" json:"natural_interior_page_range"`         // Sum of independent chapter yields (no forced correlation)
	NaturalTotalBookPageRange    string  `yaml:"natural_total_book_page_range" json:"natural_total_book_page_range"`     // Realistic natural range (never forced to a single integer)
	MinRealisticPages            int     `yaml:"min_realistic_pages" json:"min_realistic_pages"`
	MaxRealisticPages            int     `yaml:"max_realistic_pages" json:"max_realistic_pages"`
	OrthogonalIndependenceRule   string  `yaml:"orthogonal_independence_rule" json:"orthogonal_independence_rule"`
}

// ComputeDynamicBookSizing calculates the realistic natural page range by summing the independent
// evidentiary yields of each section (a young topic with sparse historical data may still have 5+ high-quality
// Chapter 7 pivots, while an older topic may have 2 pivots). We use what reality and the quality/context bar produce.
func ComputeDynamicBookSizing(topicSlug string, velocityTier string, validatedSources int, populatedMolds int, validatedPivots int) DynamicBookSizingSpec {
	slug := strings.ToLower(strings.TrimSpace(topicSlug))

	ageNote := "Multi-year empirical history"
	ageYears := 6.5
	if velocityTier == "hot" || strings.Contains(slug, "ai-") || strings.Contains(slug, "tiktok") || strings.Contains(slug, "faceless") || strings.Contains(slug, "prompt") || strings.Contains(slug, "claw") {
		ageNote = "Recent/emerging topic (<2.5 yrs); historical cohort sections stay concise while pivots/mechanisms use every high-quality match found"
		ageYears = 1.8
	} else if velocityTier == "evergreen" || strings.Contains(slug, "real-estate") || strings.Contains(slug, "mlm") || strings.Contains(slug, "wholesal") || strings.Contains(slug, "franchise") {
		ageNote = "Multi-decade statutory & cohort history"
		ageYears = 15.0
	}

	// Each section's page range is computed independently from what passed quality & context vetting:
	// - Phase 1 (Ch 1-2): ~10-12 pages
	// - Ch 3, 5, 6 (Math, Industry & Verdict): ~10-18 pages depending on historical/regulatory dataset depth
	// - Ch 4 (Molds): 2 pages per vetted composite mold (populatedMolds * 2) + 2 framing pages
	// - Ch 7 (Pivots): ~1 page per vetted pivot that passes quality/context + 2 framing pages (independent of topic age!)
	// - Ch 8 (Honest Test & Exit): ~5-7 pages
	// - Appendix A Bulk Bibliography: ~1-4 pages based on validatedSources
	minBib, maxBib := 1, 2
	if validatedSources >= 120 {
		minBib, maxBib = 2, 4
	}

	minCh356, maxCh356 := 12, 16
	if ageYears < 2.5 && validatedSources < 70 {
		minCh356, maxCh356 = 9, 13
	} else if ageYears >= 12.0 && validatedSources >= 150 {
		minCh356, maxCh356 = 14, 20
	}

	ch4Pages := (populatedMolds * 2) + 2
	ch7Pages := validatedPivots + 2

	minInterior := 10 + minCh356 + ch4Pages + ch7Pages + 5
	maxInterior := 12 + maxCh356 + ch4Pages + ch7Pages + 7

	staticFront := 6
	staticCoreBack := 6 // 2p Dual-Page Research Machine Spread + 1p Glossary + 1p Contract + 1p Results Typical/Colophon + 1p Index

	minTotal := staticFront + minInterior + minBib + staticCoreBack
	maxTotal := staticFront + maxInterior + maxBib + staticCoreBack

	return DynamicBookSizingSpec{
		TopicSlug:                    slug,
		TopicAgeNote:                 ageNote,
		EstimatedTopicAgeYears:       ageYears,
		ValidatedSourceCount:         validatedSources,
		PopulatedMoldsCount:          populatedMolds,
		ValidatedPivotsCount:         validatedPivots,
		StaticFrontMatterPages:       staticFront,
		StaticCoreBackMatterPages:    staticCoreBack,
		NaturalBibliographyPageRange: "1–4 pages (scales directly with validated source count)",
		NaturalInteriorPageRange:     "Independent sum of Ch 1–2 + Ch 3/5/6 data depth + Ch 4 vetted molds + Ch 7 vetted pivots + Ch 8 exit protocol",
		NaturalTotalBookPageRange:    "40–84 pages realistic Series envelope (topic-specific natural range computed from independent section yields)",
		MinRealisticPages:            minTotal,
		MaxRealisticPages:            maxTotal,
		OrthogonalIndependenceRule:   "Section lengths are NOT assumed to be correlated: a young topic with sparse historical data may still yield many strong Chapter 7 pivots, and an older topic may yield fewer pivots. We include everything that clears the quality and context bar—never stretching or compacting to hit a fixed number.",
	}
}

// TopicFamilyProfile defines the dynamic, family-specific research configuration
// so the factory is equally proficient across all 54 topics in the 9 Y-IT categories (F1..F9).
type TopicFamilyProfile struct {
	TopicSlug              string   `yaml:"topic_slug" json:"topic_slug"`
	TopicLabel             string   `yaml:"topic_label" json:"topic_label"`
	FamilyCode             string   `yaml:"family_code" json:"family_code"`
	FamilyName             string   `yaml:"family_name" json:"family_name"`
	ArchetypeCode          string   `yaml:"archetype_code" json:"archetype_code"`
	VelocityTier           string   `yaml:"velocity_tier" json:"velocity_tier"`
	UnitEconomicKillEq     string   `yaml:"unit_economic_kill_eq" json:"unit_economic_kill_eq"`
	PrimaryAuthorities     []string `yaml:"primary_authorities" json:"primary_authorities"`
	FederalRegisterTerms   []string `yaml:"federal_register_terms" json:"federal_register_terms"`
	DirectTier1URLs        []struct {
		Label     string `yaml:"label" json:"label"`
		URL       string `yaml:"url" json:"url"`
		Publisher string `yaml:"publisher" json:"publisher"`
		Vector    string `yaml:"vector" json:"vector"`
	} `yaml:"direct_tier1_urls" json:"direct_tier1_urls"`
	OpenAlexQueries        []string `yaml:"openalex_queries" json:"openalex_queries"`
	CrossrefQueries        []string `yaml:"crossref_queries" json:"crossref_queries"`
	ArxivQueries           []string `yaml:"arxiv_queries" json:"arxiv_queries"`
	WaybackTargets         []string `yaml:"wayback_targets" json:"wayback_targets"`
	PractitionerHNQueries  []string `yaml:"practitioner_hn_queries" json:"practitioner_hn_queries"`
	OpenSourceToolQueries  []string `yaml:"open_source_tool_queries" json:"open_source_tool_queries"`
	LikelyDominantMolds    []string `yaml:"likely_dominant_molds" json:"likely_dominant_molds"`
	FamilyHonestTestModule string   `yaml:"family_honest_test_module" json:"family_honest_test_module"`
	FamilyExitModule       string   `yaml:"family_exit_module" json:"family_exit_module"`
	AffiliateEcosystemFocus []string `yaml:"affiliate_ecosystem_focus" json:"affiliate_ecosystem_focus"`
}

// ResolveTopicProfile dynamically constructs a full multi-engine research profile for ANY of the 54 Y-IT topics
// across all 9 families (F1..F9), ensuring equal research depth and proficiency on every topic.
func ResolveTopicProfile(topicSlug string) TopicFamilyProfile {
	slug := strings.ToLower(strings.TrimSpace(topicSlug))

	// Base universal authorities and queries shared across all side-hustle/wealth topics
	base := TopicFamilyProfile{
		TopicSlug:  slug,
		TopicLabel: formatSlugTitle(slug),
	}

	switch {
	case strings.Contains(slug, "forex") || strings.Contains(slug, "prop") || strings.Contains(slug, "crypto") || strings.Contains(slug, "option") || strings.Contains(slug, "trading"):
		base.FamilyCode = "F4"
		base.FamilyName = "Trading & Investing"
		base.ArchetypeCode = "ARCH-1: SPECULATION_BBOOK"
		base.VelocityTier = "classic"
		base.UnitEconomicKillEq = "Asymmetric Drawdown Barrier (5% daily vs 8-10% target) + Spread/Slippage Drag + 98.4% B-Book Counterparty Routing -> 91-98% Attrition & Reset-Fee Harvesting"
		base.PrimaryAuthorities = []string{"U.S. CFTC (RFED & Enforcement Dockets)", "SEC / FINRA", "BIS Triennial Central Bank Survey", "ESMA & UK FCA Mandatory CFD/FX Disclosures", "U.S. BLS OEWS", "Federal Reserve SHED"}
		base.FederalRegisterTerms = []string{
			"retail foreign exchange",
			"commodity futures trading commission forex",
			"deceptive earnings claims business opportunity",
			"leveraged retail derivatives",
			"proprietary trading risk controls",
			"consumer financial protection foreign exchange",
		}
		base.DirectTier1URLs = []struct {
			Label     string `yaml:"label" json:"label"`
			URL       string `yaml:"url" json:"url"`
			Publisher string `yaml:"publisher" json:"publisher"`
			Vector    string `yaml:"vector" json:"vector"`
		}{
			{"BIS Triennial Central Bank Survey of FX & OTC Derivatives ($7.5T/day)", "https://www.bis.org/statistics/rpfx22_fx.htm", "Bank for International Settlements (BIS)", "P2-Q1/P2-Q6"},
			{"Federal Reserve Report on the Economic Well-Being of U.S. Households (SHED)", "https://www.federalreserve.gov/publications/report-economic-well-being-us-households.htm", "Board of Governors of the Federal Reserve System", "P2-Q12/P3-Q13"},
			{"SEC Investor.gov High-Yield & Forex Trading Risk Bulletin", "https://www.investor.gov/introduction-investing/investing-basics/glossary/foreign-currency-trading-forex", "U.S. Securities and Exchange Commission (SEC)", "P2-Q5/P3-Q15"},
		}
		base.OpenAlexQueries = []string{
			"retail foreign exchange trading profitability",
			"day trading individual investor performance transaction costs",
			"contract for difference retail investor losses leverage",
			"foreign exchange market microstructure retail order flow",
			"overconfidence retail currency traders gambler ruin",
			"algorithmic trading backtest overfitting walk forward",
			"business intelligence data analyst wage return career",
			"financial influencer affiliate marketing disclosure consumer",
			"household emergency savings liquid financial buffer survey",
			"kelly criterion drawdown stopping time position sizing",
		}
		base.CrossrefQueries = []string{
			"retail forex trading profitability spread",
			"proprietary trading drawdown evaluation risk",
			"retail CFD leverage regulatory intervention",
			"forex broker conflict of interest market making",
			"foreign exchange news announcement spread slippage",
			"walk forward optimization trading strategy robustness",
		}
		base.ArxivQueries = []string{
			"all:foreign+exchange+retail+trading",
			"all:drawdown+ruin+probability+trading",
			"all:backtest+overfitting+sharpe",
			"all:limit+order+book+slippage+execution",
		}
		base.WaybackTargets = []string{"ftmo.com", "cftc.gov/LearnAndProtect", "forex.com"}
		base.PractitionerHNQueries = []string{"prop trading firm forex", "retail forex broker", "python backtesting trading strategy", "SQL python data analyst transition"}
		base.OpenSourceToolQueries = []string{"forex backtesting python", "kelly criterion position sizing python", "walk forward analysis trading python"}
		base.LikelyDominantMolds = []string{"E2", "H4", "H5", "C1", "C3"}
		base.FamilyHonestTestModule = "F4 Paper First, Position Limits & The Index Baseline"
		base.FamilyExitModule = "Close Leveraged Positions, Cancel Chart/Signal/VPS Subscriptions, Audit 120-Day Chargeback Window"
		base.AffiliateEcosystemFocus = []string{"Prop Challenge CPA Referral Codes (20-25%)", "Offshore Broker IB Lot Rebates ($6-$15/lot)", "TradingView/Indicator Recurring SaaS Kickbacks", "Ethical Pivot: SQL/Python Financial BI & QuantConnect"}

	case strings.Contains(slug, "dropship") || strings.Contains(slug, "fba") || strings.Contains(slug, "tiktok-shop") || strings.Contains(slug, "print-on-demand") || strings.Contains(slug, "shopify") || strings.Contains(slug, "etsy"):
		base.FamilyCode = "F1"
		base.FamilyName = "Ecommerce"
		base.ArchetypeCode = "ARCH-2: ECOMMERCE_AUCTION"
		base.VelocityTier = "classic"
		base.UnitEconomicKillEq = "Net Margin = AOV - COGS - Platform Referral/Fulfillment Take (15-45%) - Paid Ad CAC (Meta/TikTok CPM Auction) - Returns/Chargeback Reserves"
		base.PrimaryAuthorities = []string{"SEC 10-K Filings (Amazon, Shopify, Etsy, Meta)", "FTC Mail/Internet Order Merchandise Rule & Business Opportunity Dockets", "U.S. CBP De-Minimis Tariff Filings", "U.S. BLS OEWS", "Federal Reserve SHED"}
		base.FederalRegisterTerms = []string{
			"mail or telephone order merchandise rule",
			"deceptive earnings claims ecommerce dropshipping",
			"de minimis customs tariff e-commerce",
			"online marketplace seller fee disclosure",
			"consumer review fairness act fake reviews",
			"payment processor merchant reserve chargeback",
		}
		base.DirectTier1URLs = []struct {
			Label     string `yaml:"label" json:"label"`
			URL       string `yaml:"url" json:"url"`
			Publisher string `yaml:"publisher" json:"publisher"`
			Vector    string `yaml:"vector" json:"vector"`
		}{
			{"Federal Reserve Report on the Economic Well-Being of U.S. Households (SHED)", "https://www.federalreserve.gov/publications/report-economic-well-being-us-households.htm", "Board of Governors of the Federal Reserve System", "P2-Q12/P3-Q13"},
			{"FTC Business Opportunity & E-Commerce Enforcement Guidance", "https://www.ftc.gov/business-guidance/resources/selling-internet-prompt-delivery-rules", "Federal Trade Commission (FTC)", "P2-Q1/P2-Q5"},
		}
		base.OpenAlexQueries = []string{
			"e-commerce dropshipping survival rate profitability",
			"online marketplace seller competition platform fees",
			"social media advertising customer acquisition cost auction",
			"third party seller amazon fba inventory risk",
			"supply chain lead time chargeback retail ecommerce",
			"e-commerce operations supply chain analyst wage career",
		}
		base.CrossrefQueries = []string{
			"dropshipping supply chain retail margin",
			"platform marketplace fee seller survival",
			"digital advertising auction customer acquisition cost",
			"online retail product return rate profitability",
		}
		base.ArxivQueries = []string{"all:sponsored+search+auction+bidding+margin", "all:ecommerce+recommendation+long+tail+seller"}
		base.WaybackTargets = []string{"shopify.com/pricing", "sellercentral.amazon.com", "autods.com"}
		base.PractitionerHNQueries = []string{"dropshipping margins ad spend", "amazon fba seller fees", "shopify app churn", "supply chain data analyst"}
		base.OpenSourceToolQueries = []string{"ecommerce unit economics calculator", "inventory forecasting python", "marketing mix modeling open source"}
		base.LikelyDominantMolds = []string{"E1", "E3", "H1", "H2", "C2"}
		base.FamilyHonestTestModule = "F1 Sell 10 Units By Hand Before Building a Store or Buying Ads · The Unit Economics Sheet"
		base.FamilyExitModule = "Freeze Ad Campaigns, Liquidate Inventory at Landed Cost, Downgrade Store/App Stack, Reconcile Processor Reserves"
		base.AffiliateEcosystemFocus = []string{"Store-Builder Trial Bounty ($58-$150 CPA)", "AutoDS/Helium10/JungleScout Recurring SaaS (20-30%)", "3PL & Ad-Spy Tool Referrals", "Ethical Pivot: Supply-Chain / E-Commerce BI & Paid-Media Analytics"}

	case strings.Contains(slug, "agency") || strings.Contains(slug, "smma") || strings.Contains(slug, "automation") || strings.Contains(slug, "seo") || strings.Contains(slug, "freelance") || strings.Contains(slug, "virtual-assistant") || strings.Contains(slug, "coaching"):
		base.FamilyCode = "F2"
		base.FamilyName = "Services & Agencies"
		base.ArchetypeCode = "ARCH-3: AGENCY_CHURN"
		base.VelocityTier = "hot"
		base.UnitEconomicKillEq = "0.2-0.5% Cold-Outreach Booking Rate + 60-90 Day SMB Client Churn + Unbilled Scope Creep = 70%+ of Hours Spent Replacing Churned Retainers"
		base.PrimaryAuthorities = []string{"FTC Telemarketing Sales Rule & Business Opportunity Filings", "FCC TCPA Cold-Calling/SMS Regulations", "U.S. BLS Professional & Technical Services OEWS", "SBA Small Business Survival Cohorts"}
		base.FederalRegisterTerms = []string{
			"telemarketing sales rule business to business",
			"telephone consumer protection act robocall lead generation",
			"deceptive business opportunity coaching seminar",
			"independent contractor classification wage",
		}
		base.OpenAlexQueries = []string{
			"small business digital marketing agency client churn",
			"B2B cold outreach conversion rate email fatigue",
			"freelance platform gig economy effective hourly wage",
			"marketing automation software adoption small enterprise",
		}
		base.CrossrefQueries = []string{
			"B2B service retainer churn customer lifetime value",
			"freelance labor marketplace price competition",
		}
		base.ArxivQueries = []string{"all:large+language+model+workflow+automation+reliability"}
		base.WaybackTargets = []string{"gohighlevel.com", "make.com/en/pricing", "upwork.com"}
		base.PractitionerHNQueries = []string{"AI automation agency churn", "SMMA cold email deliverability", "GoHighLevel affiliate", "B2B consulting retainer"}
		base.OpenSourceToolQueries = []string{"n8n workflow automation", "crm open source self hosted", "cold email deliverability dmarc"}
		base.LikelyDominantMolds = []string{"E2", "E4", "H3", "H4", "C1"}
		base.FamilyHonestTestModule = "F2 One Paid Client First (Zero SaaS Stack Before First Signed Invoice) · Rate × Hours × Utilization Sheet"
		base.FamilyExitModule = "Cancel Multi-Seat SaaS/Lead-Gen Subscriptions, Close Out Subcontractor Agreements, Transition Surviving Clients to Fixed-Scope Handover"
		base.AffiliateEcosystemFocus = []string{"GoHighLevel 40% Recurring Affiliate Pyramid", "Make/Zapier/Instantly SaaS Referrals", "$997-$5,000 Skool Mastermind Upsells", "Ethical Pivot: Direct W-2/Contract RevOps & Solutions Engineering"}

	case strings.Contains(slug, "airbnb") || strings.Contains(slug, "real-estate") || strings.Contains(slug, "wholesal"):
		base.FamilyCode = "F7"
		base.FamilyName = "Real Estate & Property"
		base.ArchetypeCode = "ARCH-4: ASSET_LEASE_LEVERAGE"
		base.VelocityTier = "evergreen"
		base.UnitEconomicKillEq = "Fixed Master-Lease / Debt Obligation + Furnishing Capex + Cleaning/Platform Fees vs. Seasonal Occupancy Troughs & Municipal Short-Term Rental Bans"
		base.PrimaryAuthorities = []string{"Municipal STR Enforcement Registries", "HUD / RESPA & State Real Estate Licensing Boards", "FTC Real Estate Seminar Enforcement (Zurixx / Nudge LLC)", "SBA & Commercial Lease Default Data"}
		base.FederalRegisterTerms = []string{"real estate settlement procedures act", "short term rental housing regulation", "deceptive real estate investment seminar"}
		base.OpenAlexQueries = []string{"short term rental airbnb regulation profitability", "real estate wholesaling contract assignment regulation", "residential lease arbitrage occupancy risk"}
		base.CrossrefQueries = []string{"airbnb occupancy rate municipal regulation", "real estate seminar consumer loss"}
		base.ArxivQueries = []string{"all:short+term+rental+housing+market+pricing"}
		base.WaybackTargets = []string{"airbnb.com/help", "airdna.co", "biggerpockets.com"}
		base.PractitionerHNQueries = []string{"airbnb arbitrage lease", "short term rental ban", "real estate wholesaling cold calling"}
		base.OpenSourceToolQueries = []string{"real estate cash flow calculator", "mortgage amortization python"}
		base.LikelyDominantMolds = []string{"E1", "H1", "H5", "C3", "C4"}
		base.FamilyHonestTestModule = "F7 Run Three Deals on Paper (Underwrite 20% Vacancy & Regulatory Ban Risk Before Signing Any Lease)"
		base.FamilyExitModule = "Negotiate Master-Lease Surrender/Sublease, Liquidate Furnishings, Cancel Dynamic-Pricing SaaS"
		base.AffiliateEcosystemFocus = []string{"$10k-$40k Real Estate Guru Seminar Financing Kickbacks", "Skip-Tracing & Dialer SaaS Affiliates", "AirDNA/PriceLabs Referrals", "Ethical Pivot: Licensed Property Management, Commercial Appraisal & REIT/Index Baseline"}

	default:
		// F3 Content / F6 Influencer / F8 Flipping / F9 Emerging Tech universal fallback profile
		base.FamilyCode = "F3/F9"
		base.FamilyName = "Content, Digital Products & Emerging Tech"
		base.ArchetypeCode = "ARCH-5: ALGORITHMIC_ATTENTION"
		base.VelocityTier = "hot"
		base.UnitEconomicKillEq = "Creator Power-Law Distribution (Top 1% Capture >85% of Pool) + Algorithmic Demonetization / API Cost Spikes + Unpaid Production Hours -> <$5/hr Effective Wage"
		base.PrimaryAuthorities = []string{"U.S. Copyright Office AI Rulings", "FTC Endorsement & Testimonial Guides", "Platform Creator Fund & RPM Filings", "U.S. BLS OEWS", "Federal Reserve SHED"}
		base.FederalRegisterTerms = []string{"copyright registration artificial intelligence", "guides concerning use of endorsements and testimonials", "deceptive earnings claims online business"}
		base.OpenAlexQueries = []string{"creator economy power law income distribution", "youtube monetization algorithmic demonetization", "micro saas churn customer acquisition cost"}
		base.CrossrefQueries = []string{"creator economy platform monetization inequality", "digital product marketplace saturation"}
		base.ArxivQueries = []string{"all:creator+economy+recommender+system+fairness"}
		base.WaybackTargets = []string{"support.google.com/youtube", "gumroad.com", "substack.com"}
		base.PractitionerHNQueries = []string{"micro saas churn", "youtube automation demonetized", "kdp low content amazon ban"}
		base.OpenSourceToolQueries = []string{"static site generator newsletter", "video editing automation ffmpeg"}
		base.LikelyDominantMolds = []string{"E2", "E3", "H3", "H5", "C3"}
		base.FamilyHonestTestModule = "F3 Publish 20 Pieces Before Buying Gear / F9 Ten Paying Customers Before Writing Code"
		base.FamilyExitModule = "Cancel Generative AI/SaaS Subscriptions, Export Owned Email List, Repurpose Portfolio into B2B Production Retainers"
		base.AffiliateEcosystemFocus = []string{"AI Video/Voice Wrapper SaaS Affiliates (30% recurring)", "Course Hosting & Community Platform Referrals", "Ethical Pivot: B2B Technical Writing, Product Engineering & Owned-List Consulting"}
	}

	return base
}

func formatSlugTitle(slug string) string {
	parts := strings.Split(slug, "-")
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, " ")
}
