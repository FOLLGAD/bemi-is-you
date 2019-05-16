import * as PIXI from 'pixi.js'
import { DropShadowFilter } from '@pixi/filter-drop-shadow';

import Level from './level'

const TILE_SIZE = 128 // 64*64 pixels represents a tile

// Events
const MOVE = 0,
	DEATH = 1,
	SPAWN = 2
// Transform: death+spawn

export default class Game {
	constructor() {
		PIXI.settings.SCALE_MODE = PIXI.SCALE_MODES.NEAREST
		this.app = new PIXI.Application({
			width: TILE_SIZE * 18,
			height: TILE_SIZE * 12,
			resolution: window.devicePixelRatio || 1,
			backgroundColor: 0xeeeeee,
		})
		this.playerNum = null
	}
	start() {
		document.querySelector("#canv").appendChild(this.app.view)

		this.chars = new Map()

		const background = PIXI.Texture.from('../textures/background.png')
		let bgElement = new PIXI.Sprite.from(background)
		bgElement.width = this.app.screen.width
		bgElement.height = this.app.screen.height
		this.app.stage.addChild(bgElement)

		const container = new PIXI.Container()
		this.container = container
		container.x = 0
		container.y = 0

		// Add a dropshadow effect
		this.container.filters = [new DropShadowFilter({ distance: 5, color: 0x0, alpha: 0.3, blur: 50, pixelSize: 1 / 16 })]

		this.app.stage.addChild(container)

		this.level.objects.forEach(ob => {
			this.setChar(ob)
		})
		this.container.sortChildren()

		this.app.renderer.resize(TILE_SIZE * this.level.width, TILE_SIZE * this.level.height);
	}
	setChar(char) {
		const texture = PIXI.Texture.from('../textures/' + this.getTexture(char))
		const pixiChar = new PIXI.Sprite(texture)

		pixiChar.width = pixiChar.height = TILE_SIZE

		if (["wall"].indexOf(char.item) != -1) { // Full-tile blocks should be on bottom
			pixiChar.zIndex = 0
		} else if (char.kind == 0) {
			pixiChar.zIndex = 5
		} else { // Text should be on top
			pixiChar.zIndex = 10
		}

		pixiChar.x = TILE_SIZE * char.pos.x
		pixiChar.y = TILE_SIZE * char.pos.y
		this.container.addChild(pixiChar)
		this.chars.set(char.id, pixiChar)

		// pixiChar.filters = [new DropShadowFilter({distance:0})]
	}
	updateChar(charObject) {
		let char = this.chars.get(charObject.id)

		if (char) {
			char.x = TILE_SIZE * charObject.pos.x
			char.y = TILE_SIZE * charObject.pos.y
		}
	}
	getTexture(object) {
		if (object.item == "1" || object.item == "2") {
			if (this.playerNum == Number(object.item)) {
				return "1.png"
			} else {
				return "2.png"
			}
		} else if (object.kind == 0) return object.item + "_block.png"
		else return object.item + ".png"
	}
	removeChar(characterId) {
		let char = this.chars.get(characterId)
		if (!char) {
			console.warn("Tried to remove non-existing character", characterId)
			return
		}
		char.destroy() // Destroy the PIXI instance to remove it from the canvas
		this.chars.delete(characterId)
	}
	setLevel(data) {
		console.log(data)
		this.level = new Level(data)
	}
	listen(elem, func) {
		elem.addEventListener('keydown', func)
	}
	deltaTick(tick) {
		if (tick != null) {
			tick.forEach(change => {
				switch (change.event) {
					case MOVE: // Move
						let char = this.level.objects.find(ob => ob.id == change.id)
						char.pos.x += change.pos.x
						char.pos.y += change.pos.y
						this.updateChar(char)
						break
					case DEATH:
						this.level.objects = this.level.objects.filter(d => d.id != change.id)
						this.removeChar(change.id)
						break
					case SPAWN:
						let newChar = { pos: change.pos, id: change.id, item: change.item, kind: change.kind }
						this.level.objects.push(newChar)
						this.setChar(newChar)
						break
				}
			})
		} else {
			console.warn("Tick was null!")
		}
		this.container.sortChildren()
	}
}