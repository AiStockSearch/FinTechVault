# Event Registration and Feeds Design for Wealth Project

## Event Registration

### Purpose
The **Event Registration** page allows users to sign up for Wealth Project events (meetups, webinars, conferences like Consensus), fostering community participation. It’s inspired by **TradingView**’s event pages and **GitHub**’s community event listings, providing a streamlined registration process.

**Goals**:
- Display upcoming events with details (date, location, type).
- Enable registration with user profile data.
- Integrate with Economic Calendar and side panel updates.
- Promote airdrop and community engagement.

### Functionality
- **Event List**:
  - Upcoming events: “Web3 Meetup” (2025-06-20, NYC, Webinar).
  - Details: Date, time, location, description, speaker.
  - Filter: Type (Meetup, Webinar, Conference), Location.
- **Registration Form**:
  - Fields: Name, email (auto-filled from profile), event selection.
  - Option: Join airdrop during registration.
  - Confirmation: Email ticket (QR-code).
- **Side Panel**:
  - Balance Updates: Event registered (+0.1 WTH).
  - News: Event-related updates.
  - Social Alerts: Event promotions (@WealthProject).

### Interactions
- **Event List**:
  - Event Card → Modal (details, register).
  - Filter → Update list.
- **Registration Form**:
  - Submit → Confirmation modal (QR-code).
  - Join Airdrop → Checkbox, redirect to airdrop.
- **Transitions**:
  - Event Card → Pop-up modal (0.2s).
  - Submit → Fade confirmation (0.2s).

### Descriptive Mockup
```
[Black Background #121212]
[Top Toolbar: Events Icon (Yellow), Filter (Type Dropdown)]
[Left Sidebar: Same as Main]
[Right Side Panel: 250px Wide, Black #1E222D]
  [Balance Updates: Card (+0.1 WTH, Event Registered)]
  [News: Card (“Consensus 2025 Announced”)]
  [Social Alerts: Card (@WealthProject: “Join Meetup”)]
[Main Content: 750px Wide]
  [Header: "Event Registration" (20px, White), Events: 10]
  [Event List: Cards (“Web3 Meetup”, 2025-06-20, NYC), 300x150px, Gray #2A2E39]
  [Registration Form: Inputs (Name, Email), Dropdown (Event), Checkbox (Join Airdrop), Submit]
  [Details (On Click): Card (Description, Speaker, Register), 400x300px]
[Bottom Panel: Alerts (Registered for Meetup), 750x80px]
[Mobile: Stacked Cards, Compact Form]
```

## Reading Feeds

### Purpose
The **Reading Feeds** page aggregates social and news feeds, including friends’ activities, platform updates, and market news, creating a community-driven experience. It’s inspired by **TradingView**’s social feed and **GitHub**’s activity timeline, enhancing user engagement.

**Goals**:
- Display feeds: Friends’ actions, Wealth Project posts, market news.
- Allow interaction: Like, comment, share.
- Integrate with side panel for real-time updates.
- Support airdrop promotion.

### Functionality
- **Feed Types**:
  - Friends: Actions (e.g., “John published article”, “Alice voted DAO”).
  - Platform: Wealth Project posts (e.g., “Airdrop Update”).
  - News: Market updates (e.g., “BTC ETF Approved”).
  - Filter: Friends, Platform, News, Date.
- **Feed Items**:
  - Content: Text, images, links.
  - Actions: Like, Comment, Share (X, Telegram).
  - Timestamp, author (e.g., @JohnDoeTrader).
- **Side Panel**:
  - Balance Updates: Feed interaction (+0.1 WTH).
  - News: Feed-related updates.
  - Social Alerts: Feed shares (@WealthProject).

### Interactions
- **Feed Types**:
  - Filter → Update feed.
  - Feed Item → Modal (full content, comments).
- **Feed Items**:
  - Like → Increment counter.
  - Comment → Submit (modal).
  - Share → Copy link, open social app.
- **Transitions**:
  - Filter Apply → Fade feed (0.2s).
  - Feed Item → Pop-up modal (0.2s).

### Descriptive Mockup
```
[Black Background #121212]
[Top Toolbar: Feeds Icon (Yellow), Filter (Type Dropdown), Search]
[Left Sidebar: Same as Main]
[Right Side Panel: 250px Wide, Black #1E222D]
  [Balance Updates: Card (+0.1 WTH, Feed Like)]
  [News: Card (“BTC ETF Approved”)]
  [Social Alerts: Card (@WealthProject: “Airdrop Feed”)]
[Main Content: 750px Wide]
  [Header: "Reading Feeds" (20px, White), Total Items: 100]
  [Feed Types: Tabs (Friends, Platform, News), Blue Active]
  [Feed Items: Cards (“John Published Article”, 2025-06-11, Like, Comment), 600x150px]
  [Details (On Click): Card (Content, Comments, Share), 600x400px]
[Bottom Panel: Alerts (New Feed Item), 750x80px]
[Mobile: Stacked Cards, Dropdown for Tabs]
```

## Notes
- Event Registration: TradingView-style event cards, GitHub-style registration.
- Feeds: TradingView-style social feed, GitHub-style activity timeline.
- Enhances community and event participation.