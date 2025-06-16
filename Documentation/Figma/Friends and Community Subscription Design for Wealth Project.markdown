# Friends and Community Subscription Design for Wealth Project

## Friends

### Purpose
The **Friends** page enables users to connect with other Wealth Project users, view their activities, and collaborate on trading strategies, fostering a social network. It’s inspired by **TradingView**’s social features and **GitHub**’s follower system, enhancing community engagement.

**Goals**:
- Manage friend connections (add, remove, follow).
- Display friends’ activities (trades, publications, DAO votes).
- Enable private messaging and group chats.
- Integrate with feeds and side panel updates.

### Functionality
- **Friend List**:
  - List: Friends (e.g., @AliceTrader, @BobCrypto).
  - Details: Username, bio, wallet address (partial, e.g., 0x123...abc).
  - Status: Online, offline, last active.
- **Friend Actions**:
  - Add Friend: Search by username/wallet.
  - Remove/Follow: Manage connections.
  - Message: Private chat, group chat (e.g., “DeFi Group”).
- **Activity Feed**:
  - Friends’ actions: Trades, publications, DAO votes.
  - Filter: Type (Trades, Publications), Date.
- **Side Panel**:
  - Balance Updates: Friend added (+0.1 WTH).
  - News: Social updates.
  - Social Alerts: Friend activities (@WealthProject).

### Interactions
- **Friend List**:
  - Friend Card → Profile modal (view, message).
  - Search → Autocomplete (username/wallet).
- **Friend Actions**:
  - Add Friend → Modal (search, confirm).
  - Message → Chat window (pop-up).
- **Activity Feed**:
  - Filter → Update feed.
  - Item → Modal (details, like).
- **Transitions**:
  - Friend Card → Pop-up modal (0.2s).
  - Message → Slide chat (0.3s).

### Descriptive Mockup
```
[Black Background #121212]
[Top Toolbar: Friends Icon (Yellow), Search (Input, 200px), Add Friend (Yellow Button)]
[Left Sidebar: Same as Main]
[Right Side Panel: 250px Wide, Black #1E222D]
  [Balance Updates: Card (+0.1 WTH, Friend Added)]
  [News: Card (“New Social Feature”)]
  [Social Alerts: Card (@WealthProject: “New Friend”)]
[Main Content: 750px Wide]
  [Header: "Friends" (20px, White), Total Friends: 25]
  [Friend List: Cards (@AliceTrader, Online, 0x123...abc), 300x100px, Gray #2A2E39]
  [Actions: Buttons (Add Friend, Message, Follow), 150x80px]
  [Activity Feed: Cards (“Alice Published Article”, 2025-06-11), 600x150px]
  [Chat (On Click): Window (Messages, Send), 400x300px]
[Bottom Panel: Alerts (New Message), 750x80px]
[Mobile: Stacked Cards, Compact Chat]
```

## Community Subscription

### Purpose
The **Community Subscription** page allows users to join Wealth Project communities (e.g., DeFi Traders, NFT Collectors), subscribe to premium groups, and engage with exclusive content. It’s inspired by **TradingView**’s community features and **GitHub**’s organization pages, boosting social interaction.

**Goals**:
- Display available communities with details (members, focus).
- Enable subscription (free or paid, e.g., $10/month).
- Provide access to exclusive content (articles, analytics).
- Integrate with side panel for community updates.

### Functionality
- **Community List**:
  - List: Communities (e.g., “DeFi Traders”, 500 members).
  - Details: Focus (DeFi), members, subscription (free, $10/month).
  - Filter: Free, Paid, Category (Trading, NFT).
- **Subscription Form**:
  - Select community, payment (WTH, USDT).
  - Confirmation: Access granted (email, QR-code).
- **Exclusive Content**:
  - Articles: Premium posts (e.g., “DeFi Strategy”).
  - Analytics: Community-specific Julia simulations.
- **Side Panel**:
  - Balance Updates: Subscribed (+0.1 WTH).
  - News: Community updates.
  - Social Alerts: Community posts (@WealthProject).

### Interactions
- **Community List**:
  - Community Card → Modal (details, subscribe).
  - Filter → Update list.
- **Subscription Form**:
  - Submit → Confirmation modal.
- **Exclusive Content**:
  - Article → Modal (read, share).
- **Transitions**:
  - Community Card → Pop-up modal (0.2s).
  - Submit → Fade confirmation (0.2s).

### Descriptive Mockup
```
[Black Background #121212]
[Top Toolbar: Community Icon (Yellow), Filter (Category Dropdown)]
[Left Sidebar: Same as Main]
[Right Side Panel: 250px Wide, Black #1E222D]
  [Balance Updates: Card (+0.1 WTH, Subscribed)]
  [News: Card (“New DeFi Community”)]
  [Social Alerts: Card (@WealthProject: “Join Group”)]
[Main Content: 750px Wide]
  [Header: "Community Subscription" (20px, White), Communities: 15]
  [Community List: Cards (“DeFi Traders”, 500 Members, $10/month), 300x150px]
  [Subscription Form: Dropdown (Community), Payment (WTH), Submit]
  [Content: Cards (“DeFi Strategy”, Premium Article), 300x150px]
[Bottom Panel: Alerts (Subscribed to Community), 750x80px]
[Mobile: Stacked Cards, Compact Form]
```

## Notes
- Friends: TradingView-style social connections, GitHub-style activity feed.
- Community: TradingView-style groups, GitHub-style subscription model.
- Enhances social networking and exclusive access.