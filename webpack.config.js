const Webpack = require("webpack");
const Glob = require("glob");
const path = require("path");
const CopyWebpackPlugin = require("copy-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const ManifestPlugin = require("webpack-manifest-plugin");
const CleanUp = require('webpack-cleanup-plugin');
const TerserPlugin = require("terser-webpack-plugin");
const LiveReloadPlugin = require('webpack-livereload-plugin');

process.noDeprecation = true;

const configurator = {
  entries: function(){
    var entries = {
      application: [
        './node_modules/jquery-ujs/src/rails.js',
        './assets/css/application.scss',
      ],
    }
    
    Glob.sync("./assets/*/*.*").forEach((entry) => {
      if (entry === './assets/css/application.scss') {
        return
      }
      
      let key = entry.replace(/(\.\/assets\/(js|css|go)\/)|\.(js|s[ac]ss|go)/g, '')
      if(key.startsWith("_") || (/(js|s[ac]ss|go)$/i).test(entry) == false) {
        return
      }
      
      if( entries[key] == null) {
        entries[key] = [entry]
        return
      } 
      
      entries[key].push(entry)
    })

    return entries
  },

  plugins() {
    var plugins = [
      new CleanUp({exclude: ["robots.txt"],}),
      new Webpack.ProvidePlugin({$: "jquery",jQuery: "jquery", "window.$": "jquery", "window.jQuery": "jquery"}),
      new Webpack.ContextReplacementPlugin(/\.\/locale$/, 'empty-module', false, /js$/),
      new MiniCssExtractPlugin({filename: "[name].[contenthash].css"}),
      new CopyWebpackPlugin([{from: "./assets",to: ""}], {copyUnmodified: true,ignore: ["css/**", "js/**", "resources/**"] }),
      new Webpack.LoaderOptionsPlugin({minimize: true,debug: false}),
      new ManifestPlugin({fileName: "manifest.json"})
    ];

    return plugins
  },

  moduleOptions: function() {
    return {
      rules: [
        { test: /\.jsx?$/, loader: "babel-loader", exclude: /node_modules/, query: { presets: ['env', 'stage-2']}},
        { test: /\.s[ac]ss$/,
            use: [
              MiniCssExtractPlugin.loader,
              { loader: "css-loader", options: {sourceMap: true}},
              { loader: "sass-loader", options: {sourceMap: true}}
            ]
        },
        { test: /\.(woff|woff2|ttf|svg)(\?v=\d+\.\d+\.\d+)?$/, use: "url-loader?limit=200"},
        { test: /\.eot(\?v=\d+\.\d+\.\d+)?$/,use: "file-loader" },
        { test: require.resolve("jquery"), use: "expose-loader?jQuery!expose-loader?$"},
        { test: /\.go$/, use: "gopherjs-loader"}        
      ]
    }
  },

  buildConfig: function(){
    const env = process.env.NODE_ENV || "development";
    
    var config = {
      mode: env,
      entry: configurator.entries(),
      output: {filename: "[name].[hash].js", path: `${__dirname}/public/assets`},
      plugins: configurator.plugins(),
      module: configurator.moduleOptions(),
      resolve: {
        alias: {
            'jquery': path.join(__dirname, 'node_modules/jquery/dist/jquery'),
        }
      }
    }

    if( env === "development" ){
      config.plugins.push(new LiveReloadPlugin({appendScriptTag: true}))
      return config
    }

    const terser = new TerserPlugin({
      terserOptions: {
        compress: {},
        mangle: {
          keep_fnames: true
        },
        output: {
          comments: false,
        },
      },
      extractComments: false,
    })

    config.optimization = {
      minimizer: [terser]
    }

    return config
  }
}

module.exports = configurator.buildConfig()