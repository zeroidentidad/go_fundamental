module.exports = {
   entry: './index.js',
   output:{
       path: __dirname,
       filename: 'bundle.js'
    },
    resolve: {
        extensions: ['.js', '.jsx']
    },
    module:{
       rules:[
           {
               test: /\.jsx?$/,
               loader: 'babel-loader',
               exclude: /node_modules/,
               options: {
                   presets: ['@babel/env', '@babel/preset-react']
               }
           }
       ]
   },
   mode: 'none'
}