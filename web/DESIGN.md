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

- Radii: `--radius-lg` 32px (banner, map), `--radius` 20px (cards, panels), `--radius-sm` 12px
  (photos, bars), `--pill` for buttons, nav and chips.
- Content width 1200px; side padding `--space` (16–32px); sections 56–104px apart.
- One soft shadow, `--shadow`, on cards, panels, nav and map.

## Components

- **Header**: sticky, blurred cream; leaf logo and wordmark "Wildroot"; nav in a white pill; lime
  "Explore" button (hidden on phones).
- **Banner**: forest rounded panel; spinning 3D globe behind it with a lime dot per sighting
  (`assets/js/scenes/globe.js`); a tilted glassy card with the latest photographed sighting.
- **Numbers**: four stat tiles (the first forest-filled), a sparkline of sightings per year and
  bar lists for top animals and countries.
- **Map**: Leaflet, clustered, forest markers with lime outline, framed in a paper border.
- **Timeline**: years down the left (sticky), a line with lime dots, one card per sighting with a
  4:3 photo, date in soil brown, name, scientific name, place, count, observers and weather.
- **Error page**: a big forest status code and a forest button home.

## Avoid

- Text in lime on cream or paper.
- Inline styles built from GBIF data.
