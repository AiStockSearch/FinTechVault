# Blog and Economic Calendar Design for Wealth Project

## Blog for News

### Purpose
The **Blog for News** page is a dedicated section for Wealth Project updates, crypto market insights, and community news, fostering engagement and education. It’s inspired by **TradingView**’s news feed and **GitHub**’s blog, providing a structured, searchable content hub for Web3 enthusiasts, traders, and beginners.

**Goals**:
- Share platform updates (e.g., new features, airdrop progress).
- Provide market insights (e.g., “BTC ETF Impact”).
- Encourage community interaction (comments, shares).
- Integrate with side panel for real-time updates.

### Functionality
- **Blog Posts**:
  - List: Title, excerpt, author, date, tags (e.g., DeFi, Airdrop).
  - Categories: Platform Updates, Market Insights, Tutorials.
  - Search: Filter by title, tag, date.
- **Post Details**:
  - Full content (Markdown-rendered), images, links.
  - Comments: Add, view, like (tied to user profile).
  - Share: X, Telegram, Discord, QR-code.
- **Analytics**:
  - Views, shares, engagement (e.g., 500 views, 3.5%).
- **Side Panel**:
  - Balance Updates: New blog post interaction (+0.1 WTH).
  - News: Related market updates.
  - Social Alerts: Blog shares (@WealthProject).

### Interactions
- **Blog Posts**:
  - Post Card → Blog post details page.
  - Search → Update list.
  - Category Filter → Narrow posts.
- **Post Details**:
  - Comment → Submit (modal).
  - Share → Copy link, open social app.
  - QR-code → Download.
- **Analytics**:
  - Toggle timeframe (1W, 1M).
- **Transitions**:
  - Post Card → Slide to details (0.3s).
  - Comment → Pop-up modal (0.2s).

### Descriptive Mockup
```
[Black Background #121212]
[Top Toolbar: Blog Icon (Yellow), Search (Input, 200px), Filter (Category Dropdown)]
[Left Sidebar: Same as Main]
[Right Side Panel: 250px Wide, Black #1E222D]
  [Balance Updates: Card (+0.1 WTH, Blog Comment)]
  [News: Card (“BTC ETF Impact”)]
  [Social Alerts: Card (@WealthProject: “New Blog Post”)]
[Main Content: 750px Wide]
  [Header: "Blog" (20px, White), Total Posts: 50]
  [Post List: Cards (“BTC ETF Impact”, Excerpt, 2025-06-11, Tags: DeFi), 300x150px, Gray #2A2E39]
  [Search: Input (Filter by Title/Tag), 200x50px]
  [Analytics: Bar Chart (Views: 500, Shares: 50), 600x200px]
  [Post Details (On Click): Title, Content (Markdown), Comments (Add, Like), Share (QR-code, X)]
[Bottom Panel: Alerts (New Post Published), 750x80px]
[Mobile: Stacked Cards, Compact Search]
```

## Economic Calendar

### Purpose
The **Economic Calendar** page displays upcoming crypto and financial events (e.g., token launches, ETF approvals, Fed meetings), helping traders plan strategies. It’s inspired by **TradingView**’s calendar and **GitHub**’s event timeline, offering a data-rich, filterable interface.

**Goals**:
- Show events with dates, times, and impacts (high, medium, low).
- Allow filtering by category (crypto, macro, regulatory).
- Integrate with watchlists for event-based triggers.
- Provide event details and registration links.

### Functionality
- **Calendar View**:
  - Monthly grid or list: Events (e.g., “SOL Staking Launch”, 2025-06-15).
  - Impact: High (red), Medium (yellow), Low (green).
  - Filter: Crypto, Macro, Regulatory, Date Range.
- **Event Details**:
  - Title, date, time, description, impact.
  - Link to register (e.g., webinar, meetup).
  - Add to Watchlist: Trigger for event (e.g., price alert).
- **Side Panel**:
  - Balance Updates: Event-related action (+0.1 WTH).
  - News: Related updates (“ETH ETF Vote”).
  - Social Alerts: Event announcements (@WealthProject).

### Interactions
- **Calendar View**:
  - Filter → Update events.
  - Event Card → Modal (details, register).
- **Event Details**:
  - Register → Open external link.
  - Add to Watchlist → Set trigger (modal).
- **Transitions**:
  - Filter Apply → Fade events (0.2s).
  - Event Card → Pop-up modal (0.2s).

### Descriptive Mockup
```
[Black Background #121212]
[Top Toolbar: Calendar Icon (Yellow), Filter (Category Dropdown), Date Picker]
[Left Sidebar: Same as Main]
[Right Side Panel: 250px Wide, Black #1E222D]
  [Balance Updates: Card (+0.1 WTH, Event Action)]
  [News: Card (“ETH ETF Vote”)]
  [Social Alerts: Card (@WealthProject: “Event Soon”)]
[Main Content: 750px Wide]
  [Header: "Economic Calendar" (20px, White), Events: 25]
  [Calendar View: Grid (June 2025, Cards: “SOL Staking”, High Impact, Red), 600x400px]
  [List View: Table (Date, Event, Impact), 600x300px, Gray #2A2E39]
  [Details (On Click): Card (Title, Description, Register Link, Add to Watchlist), 400x300px]
[Bottom Panel: Alerts (Event Added to Watchlist), 750x80px]
[Mobile: List View Only, Compact Cards]
```

## Notes
- Blog: TradingView-style feed, GitHub-style structured posts.
- Calendar: TradingView-style grid, filterable for trader utility.
- Enhances community engagement and trading planning.