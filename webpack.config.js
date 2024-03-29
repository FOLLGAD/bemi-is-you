let path = require('path')

module.exports = {
	entry: './client/main.js',
	watch: true,
	mode: 'development',
	output: {
		path: path.resolve(__dirname, 'public', 'dist'),
		filename: 'bundle.js'
	}
}