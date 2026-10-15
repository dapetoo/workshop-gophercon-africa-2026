# Design

Inspired by the Agroverde sustainable-farming landing page on Dribbble: calm and natural, deep
greens with a fresh lime accent on cream, big rounded panels and pill-shaped buttons.

## Colours ("Farmhouse green")

Custom properties in `assets/css/app.css` `:root`:

| Property | Hex | Use |
|----------|-----|-----|
| `--forest` | `#1E3A2B` | headings, banner, primary dark surfaces, map markers |
| `--forest-deep` | `#142A1E` | footer |
| `--moss` | `#2F5A41` | banner glow, the 3D globe |
| `--lime` | `#B8E04A` | buttons, accents, timeline dots, focus rings, globe dots |
| `--lime-soft` | `#E4F3BC` | eyebrow chips, chart bars, empty photo slots |
| `--cream` | `#F6F3EA` | page background |
| `--paper` | `#FFFDF7` | cards and panels |
| `--soil` | `#6B4F3A` | dates, weather icons |
| `--ink` | `#16241B` | body text |
| `--muted` | `#535C55` | secondary text (passes 4.5:1 on cream and paper) |

Lime is never used for text on light backgrounds, only on forest.

## Lettering

Plus Jakarta Sans (Google Fonts), 400–800. Headings 800 with tight letter-spacing; small
uppercase labels at 700 with 0.08em tracking.

## Shape and spacing

- Radii: `--radius-lg` 32px (banner, map), `--radius` 22px (cards, panels), `--radius-sm` 14px
  (photos, bars), `--pill` for buttons, nav and chips.
- Content width 1240px (banner 1360px); side padding `--space` (16–40px); sections 64–112px apart.
- `--shadow` on cards, panels, nav and map; `--shadow-lift` on hover and the numbers strip.

## Components

- **Header**: sticky, blurred cream; leaf logo and wordmark "Wildroot"; nav (Animals, Map,
  Timeline) in a white pill; lime "Explore" button (hidden on phones).
- **Banner**: a full-width wildlife photo (a big cat when there is one) under a forest fade from
  the left and bottom; huge cream headline with "real place" in lime; lime and outline pill
  buttons; a glassy "latest" chip top right with a pulsing lime dot.
- **Numbers strip**: a white rounded bar overlapping the banner's bottom edge, four big figures
  split by thin lines.
- **Meet the regulars**: photo tiles of the five most-seen animals, the first twice the size;
  lime count chip, name over a dark fade, a round arrow that turns lime on hover.
- **Globe panel**: forest rounded panel; three.js globe with cream land, moss sea, lime dots per
  sighting and a lime rim glow (`assets/js/scenes/globe.js`); top countries as bars beside it.
- **Trends**: lime columns for the last twelve years, the latest in forest; bars for top animals.
- **Map**: Leaflet, clustered, forest markers with lime outline, framed in a paper border.
- **Timeline**: years down the left (sticky) on a line with lime dots; cards three across with a
  4:3 photo, a date chip on it, name, Latin name, place, observer and a cream weather box. Nine
  per year, the rest behind a "Show N more" pill.
- **Closing banner**: forest panel, a photo fading in from the right, "Seen something wild?".
- **Footer**: deep forest, brand blurb and two link columns with lime headings.
- **Error page**: a big forest status code and a forest button home.

Photos from iNaturalist are requested at "medium" (cards) or "large" (banners), not the originals.

## Avoid

- Text in lime on cream or paper.
- Inline styles built from GBIF data.
