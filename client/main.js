import Game from './game';

window.onload = main

const TILE_SIZE = 16 // 16x16 pixels represents a tile

function main() {

	let game = new Game()

	game.listen(window) // Initiate listeners on the 'window' object
	game.start()
}