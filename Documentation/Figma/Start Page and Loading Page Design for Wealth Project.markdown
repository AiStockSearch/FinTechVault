# Start Page and Loading Page Design for Wealth Project

## Start Page

### Purpose
The **Start Page** is the initial onboarding screen for new and returning users, welcoming them to Wealth Project and guiding them to key actions (sign up, log in, join airdrop, explore features). It’s designed to be engaging and informative, inspired by **TradingView**’s clean welcome screen and **GitHub**’s community-driven landing page, emphasizing Wealth Project’s UVP (unified Web3 interface, automation, security).

**Goals**:
- Introduce Wealth Project’s mission and features.
- Drive new user sign-ups and airdrop participation (100 WTH tokens).
- Provide quick login for returning users (NFT chip, OAuth).
- Showcase community activity (X/Telegram posts).
- Offer a visually appealing, trader-focused entry point.

### Functionality
- **Hero Section**:
  - Mission: “Simplify Web3 finance with security and automation.”
  - CTA: “Join Airdrop” (100 WTH), “Sign Up”, “Log In”.
- **Feature Highlights**:
  - Cards: Wallet, DeFi, Portfolio, Indices, Analytics, Tokenomics.
  - Short descriptions, icons, and links to feature pages.
- **Community Updates**:
  - Social feed: Recent X/Telegram posts (@WealthProject).
  - Airdrop countdown (ends July 12, 2025).
- **Login Options**:
  - NFT Chip: Scan for authentication.
  - OAuth: Google, X.
  - WalletConnect: MetaMask, Validium.
- **Quick Links**:
  - Blog, Events, Community, Support.

### Interactions
- **Hero CTAs**:
  - Join Airdrop → Airdrop page (external).
  - Sign Up → Profile page (sign-up form).
  - Log In → Modal (NFT chip, OAuth, WalletConnect).
- **Feature Cards**:
  - Click → Feature page (e.g., Wallet).
- **Community Updates**:
  - Social Post → Open X/Telegram.
- **Login Options**:
  - Scan NFT Chip → Verify (modal).
  - OAuth/WalletConnect → Authenticate, redirect to Main Screen.
- **Transitions**:
  - CTA Click → Fade to page/modal (0.2s).
  - Social Post → Open new tab.
  - Login → Slide to Main Screen (0.3s).

### Descriptive Mockup
```
[Black Background #121212]
[Top Toolbar: Logo (Yellow WP, 40x40px), Sign Up (Yellow Button), Log In (Blue Button)]
[Main Content: 1200x800px, Centered]
  [Hero: Background (Blockchain Image, Unsplash), H1: "Wealth Project" (32px, White)]
  [Mission: "Simplify Web3 Finance" (20px, White), 600x100px]
  [CTAs: Buttons (Join Airdrop, Sign Up, Log In), 200x50px, Yellow/Blue]
  [Feature Highlights: 6 Cards (Wallet, DeFi, Portfolio, Indices, Analytics, Tokenomics), 300x150px, Gray #2A2E39, Icons]
  [Community Updates: Cards (X Post: “Airdrop Live!”, Telegram: “Join Now”), 300x100px, Countdown: 30d]
  [Login Options: Buttons (NFT Chip, Google, X, WalletConnect), 150x80px]
  [Quick Links: Blog, Events, Community, Support, 14px, Blue #2962FF]
[Footer: Social Links (X, Telegram, Discord), Contact: support@wealthproject.io]
[Mobile: Stacked Cards, Smaller Hero, Dropdown for Login Options]
```

## Loading Page

### Purpose
The **Loading Page** is a transitional screen displayed during app initialization, wallet connection, or data fetching, ensuring a smooth user experience. It’s minimalistic, inspired by **TradingView**’s loading animations and **GitHub**’s progress indicators, reinforcing Wealth Project’s tech-forward branding.

**Goals**:
- Provide visual feedback during loading (e.g., wallet sync, data fetch).
- Maintain brand consistency with yellow accents.
- Keep users engaged with a subtle animation.

### Functionality
- **Loading Animation**:
  - Animated Wealth Project logo (yellow WP monogram).
  - Progress bar (yellow, 0–100%).
  - Status message: “Connecting Wallet…”, “Fetching Data…”.
- **Branding**:
  - Tagline: “Powering Your Web3 Future”.
- **Error Handling**:
  - Display error (e.g., “Wallet Connection Failed”) with retry button.

### Interactions
- **Animation**:
  - Auto-progress bar (2–5s, adjustable).
- **Error**:
  - Retry Button → Reattempt action.
  - Back → Return to Start Page.
- **Transitions**:
  - Complete → Fade to Main Screen (0.3s).
  - Error → Pop-up error modal (0.2s).

### Descriptive Mockup
```
[Black Background #121212]
[Main Content: 1200x800px, Centered]
  [Logo: Animated WP (Yellow, 100x100px), Pulsing Effect]
  [Progress Bar: Yellow #F4D03F, 400x20px, 0–100%]
  [Status: "Connecting Wallet…" (16px, White)]
  [Tagline: "Powering Your Web3 Future" (14px, Light Gray)]
  [Error (Optional): Card (“Connection Failed”, Retry Button, Back), 300x150px, Red #EF5350 Border]
[Mobile: Centered Logo, Smaller Bar (200x10px)]
```

## Notes
- Start Page: TradingView-style hero, GitHub-style community updates.
- Loading Page: Minimalistic, TradingView-inspired animation.
- Drives airdrop engagement and onboarding.