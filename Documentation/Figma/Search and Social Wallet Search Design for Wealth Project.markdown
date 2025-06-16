# Search and Social Wallet Search Design for Wealth Project

## Search

### Purpose
The **Search** page provides a global search functionality across Wealth Project’s features (tokens, users, articles, events, communities), enabling quick access to relevant content. It’s inspired by **TradingView**’s universal search and **GitHub**’s search bar, offering a fast, filterable experience.

**Goals**:
- Search tokens, users, articles, events, communities.
- Filter results by category (Tokens, Users, Articles).
- Display real-time suggestions (autocomplete).
- Integrate with side panel for search-related updates.

### Functionality
- **Search Bar**:
  - Input: Query (e.g., “SOL”, “@JohnDoeTrader”).
  - Autocomplete: Tokens (SOL, MATIC), users, articles.
  - Filter: Tokens, Users, Articles, Events, Communities.
- **Search Results**:
  - List: Tokens (SOL, $200), Users (@JohnDoeTrader), Articles (“How to Earn 20%”).
  - Details: Price, bio, excerpt, date.
- **Quick Actions**:
  - Tokens: Add to Watchlist, Trade.
  - Users: Follow, Message.
  - Articles: Read, Share.
- **Side Panel**:
  - Balance Updates: Search action (+0.1 WTH).
  - News: Search-related updates.
  - Social Alerts: Search trends (@WealthProject).

### Interactions
- **Search Bar**:
  - Type Query → Autocomplete suggestions.
  - Filter → Narrow results.
- **Search Results**:
  - Result Card → Modal (details, actions).
  - Quick Action → Perform (e.g., Trade → DeFi page).
- **Transitions**:
  - Query → Fade results (0.2s).
  - Result Card → Pop-up modal (0.2s).

### Descriptive Mockup
```
[Black Background #121212]
[Top Toolbar: Search Icon (Yellow), Search Input (400px), Filter (Category Dropdown)]
[Left Sidebar: Same as Main]
[Right Side Panel: 250px Wide, Black #1E222D]
  [Balance Updates: Card (+0.1 WTH, Search Action)]
  [News: Card (“Trending: SOL”)]
  [Social Alerts: Card (@WealthProject: “Search Update”)]
[Main Content: 750px Wide]
  [Header: "Search" (20px, White), Results: 50]
  [Search Bar: Input (Query: SOL), Autocomplete (SOL, MATIC), Filter]
  [Results: Cards (SOL: $200, @JohnDoeTrader, “How to Earn 20%”), 300x100px]
  [Quick Actions: Buttons (Add to Watchlist, Trade, Follow), 150x80px]
[Bottom Panel: Alerts (Search Completed), 750x80px]
[Mobile: Stacked Cards, Compact Search Bar]
```

## Social Wallet Search

### Purpose
The **Social Wallet Search** page allows users to find other users by their wallet address, fostering social connections and transparency. It’s inspired by **TradingView**’s user search and **GitHub**’s profile discovery, tailored for Web3.

**Goals**:
- Search users by wallet address (e.g., 0x123...abc).
- Display user profiles (username, bio, activity).
- Enable actions: Follow, Message, View Activity.
- Integrate with friends and feeds.

### Functionality
- **Search Bar**:
  - Input: Wallet address (full or partial).
  - Autocomplete: Matching users (@JohnDoeTrader).
- **Search Results**:
  - List: Users (e.g., @JohnDoeTrader, 0x123...abc).
  - Details: Username, bio, wallet (partial), activity (trades, votes).
- **Actions**:
  - Follow: Add to friends.
  - Message: Start chat.
  - View Activity: See user’s feed.
- **Side Panel**:
  - Balance Updates: User followed (+0.1 WTH).
  - News: Social updates.
  - Social Alerts: User discoveries (@WealthProject).

### Interactions
- **Search Bar**:
  - Type Address → Autocomplete.
- **Search Results**:
  - User Card → Profile modal.
  - Action → Perform (e.g., Follow → Friends page).
- **Transitions**:
  - Query → Fade results (0.2s).
  - User Card → Pop-up modal (0.2s).

### Descriptive Mockup
```
[Black Background #121212]
[Top Toolbar: Wallet Search Icon (Yellow), Search Input (400px)]
[Left Sidebar: Same as Main]
[Right Side Panel: 250px Wide, Black #1E222D]
  [Balance Updates: Card (+0.1 WTH, User Followed)]
  [News: Card (“New Social Feature”)]
  [Social Alerts: Card (@WealthProject: “Found User”)]
[Main Content: 750px Wide]
  [Header: "Social Wallet Search" (20px, White), Results: 10]
  [Search Bar: Input (Query: 0x123...abc), Autocomplete (@JohnDoeTrader)]
  [Results: Cards (@JohnDoeTrader, 0x123...abc, Bio), 300x100px]
  [Actions: Buttons (Follow, Message, View Activity), 150x80px]
[Bottom Panel: Alerts (User Found), 750x80px]
[Mobile: Stacked Cards, Compact Search Bar]
```

## Notes
- Search: TradingView-style universal search, GitHub-style filtering.
- Wallet Search: TradingView-style user discovery, GitHub-style profile cards.
- Enhances content access and social networking.