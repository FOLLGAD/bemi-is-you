import * as PIXI from 'pixi.js'

window.onload = main

function main() {
	const app = new PIXI.Application({
		width: 1024, height: 1024, resolution: window.devicePixelRatio || 1, forceCanvas: true, backgroundColor: 0xeeeeee,
	})

	document.body.appendChild(app.view)

	const background = PIXI.Texture.from('../textures/background.png')
	// background.width = app.view.width
	// background.height = app.view.height
	// Create the background 
	let yah = new PIXI.Sprite.from(background);
	yah.width = app.view.width
	yah.height = app.view.height
	app.stage.addChild(yah);

	const container = new PIXI.Container()

	app.stage.addChild(container)

	const char = new PIXI.Graphics()
	char.beginFill(0x0011ee)
	char.drawRect(0, 0, 200, 200);
	// char.anchor.set(0.5)
	char.x = 20
	char.y = 20
	container.addChild(char)

	// Move container to the center
	container.x = app.screen.width / 2
	container.y = app.screen.height / 2

	// Center char sprite in local container coordinates
	container.pivot.x = container.width / 2
	container.pivot.y = container.height / 2
}