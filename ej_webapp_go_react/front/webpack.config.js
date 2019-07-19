module.exports = {
   entry: './index.js',
   output:{
       path: __dirname,
       filename: 'bundle.js'
   },
   module:{
       rules:[
           {
               test: /\.jsx?$/,
               loader: 'babel-loader',
               exclude: /node_modules/,
               options: {
                   presets: ['@babel/preset-env']
               }
           }
       ]
   },
   mode: 'none'
}