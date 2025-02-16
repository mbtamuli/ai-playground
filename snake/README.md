# Snake

A snake game that uses AI instead of a simple pathfinding algorithm.

## Code

This project uses 
- [Ebitengine v2](https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2)
- [Ebiten UI](https://pkg.go.dev/github.com/ebitenui/ebitenui)
- [Donburi](https://github.com/yohamta/donburi)

## Snake behaviors

The AI snake learns to seek out food through various configurable behaviors:
  - **Exploration Rate**: Determines how quickly the snake learns from its experiences. Lower values result in slower but more reliable learning.
  - **Risk Appetite/Discount Factor**: Dictates how much the snake values future rewards versus immediate rewards. Lower values make the snake prioritize immediate safety
  - **Learning Rate**: Controls how often the snake tries random moves instead of relying on learned behaviors. Lower values make the snake adhere to known successful strategies.

## Food Types

The game also features different types of food:
  - **Normal Food**
  - **Special Food**: Appears with a time limit, adding an extra challenge.
  - **Poison Food**: Reduces the snake's length by a fixed amount.

## Assets Required

The game requires the following assets:

### Textures
- Snake head (32x32px)
- Snake body segment (32x32px)
- Normal food sprite (32x32px)
- Special food sprite (32x32px)
- Poison food sprite (32x32px)
- Background tile (32x32px)

### UI Elements
- Button backgrounds (normal, hover, pressed states)
- Panel backgrounds for menu and AI settings
- Font files for text rendering

### Audio
- Background music
- Food collection sound
- Special food appearance jingle
- Game over sound
- Menu selection sounds
